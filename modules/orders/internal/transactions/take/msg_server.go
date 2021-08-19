package take

import (
	"context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/transfer"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type msgServer struct {
	transactionKeeper
}

func (msgServer msgServer) Take(goCtx context.Context, msg *Message) (*TransactionResponse, error) {
	message := messageFromInterface(msg)
	ctx := sdkTypes.UnwrapSDKContext(goCtx)
	if auxiliaryResponse := msgServer.transactionKeeper.verifyAuxiliary.GetKeeper().Help(ctx, verify.NewAuxiliaryRequest(message.From.AsSDKTypesAccAddress(), message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, errors.EntityNotFound
	}

	orderID := message.OrderID
	orders := msgServer.transactionKeeper.mapper.NewCollection(ctx).Fetch(key.FromID(orderID))
	order := orders.Get(key.FromID(orderID))

	if order == nil {
		return nil, errors.EntityNotFound
	}

	metaProperties, Error := supplement.GetMetaPropertiesFromResponse(msgServer.transactionKeeper.supplementAuxiliary.GetKeeper().Help(ctx, supplement.NewAuxiliaryRequest(order.(mappables.Order).GetTakerID(), order.(mappables.Order).GetMakerOwnableSplit())))
	if Error != nil {
		return nil, Error
	}

	if takerIDProperty := metaProperties.Get(base.NewID(properties.TakerID)); takerIDProperty != nil {
		takerID, Error := takerIDProperty.GetMetaFact().GetData().AsID()
		if Error != nil {
			return nil, errors.MetaDataError
		} else if takerID.Compare(base.NewID("")) != 0 && takerID.Compare(message.FromID) != 0 {
			return nil, errors.NotAuthorized
		}
	}

	exchangeRate, Error := order.(mappables.Order).GetExchangeRate().GetMetaFact().GetData().AsDec()
	if Error != nil {
		return nil, Error
	}

	makerOwnableSplitProperty := metaProperties.Get(base.NewID(properties.MakerOwnableSplit))
	if makerOwnableSplitProperty == nil {
		return nil, errors.MetaDataError
	}

	makerOwnableSplit, Error := makerOwnableSplitProperty.GetMetaFact().GetData().AsDec()
	if Error != nil {
		return nil, errors.MetaDataError
	}

	makerReceiveTakerOwnableSplit := makerOwnableSplit.MulTruncate(exchangeRate).MulTruncate(sdkTypes.SmallestDec())
	takerReceiveMakerOwnableSplit := message.TakerOwnableSplit.QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(exchangeRate)

	switch updatedMakerOwnableSplit := makerOwnableSplit.Sub(takerReceiveMakerOwnableSplit); {
	case updatedMakerOwnableSplit.Equal(sdkTypes.ZeroDec()):
		if message.TakerOwnableSplit.LT(makerReceiveTakerOwnableSplit) {
			return nil, errors.InsufficientBalance
		}

		orders.Remove(order)
	case updatedMakerOwnableSplit.LT(sdkTypes.ZeroDec()):
		if message.TakerOwnableSplit.LT(makerReceiveTakerOwnableSplit) {
			return nil, errors.InsufficientBalance
		}

		takerReceiveMakerOwnableSplit = makerOwnableSplit

		orders.Remove(order)
	default:
		makerReceiveTakerOwnableSplit = message.TakerOwnableSplit
		mutableProperties, Error := scrub.GetPropertiesFromResponse(msgServer.transactionKeeper.scrubAuxiliary.GetKeeper().Help(ctx, scrub.NewAuxiliaryRequest(base.NewMetaProperty(base.NewID(properties.MakerOwnableSplit), base.NewMetaFact(base.NewDecData(updatedMakerOwnableSplit))))))

		if Error != nil {
			return nil, Error
		}

		order = mappable.NewOrder(orderID, order.(mappables.Order).GetImmutableProperties(), order.(mappables.Order).GetImmutableProperties().Mutate(mutableProperties.GetList()...))
		orders.Mutate(order)
	}

	if auxiliaryResponse := msgServer.transactionKeeper.transferAuxiliary.GetKeeper().Help(ctx, transfer.NewAuxiliaryRequest(message.FromID, order.(mappables.Order).GetMakerID(), order.(mappables.Order).GetTakerOwnableID(), makerReceiveTakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	if auxiliaryResponse := msgServer.transactionKeeper.transferAuxiliary.GetKeeper().Help(ctx, transfer.NewAuxiliaryRequest(base.NewID(module.Name), message.FromID, order.(mappables.Order).GetMakerOwnableID(), takerReceiveMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}
	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}

var _ MsgServer = msgServer{}
