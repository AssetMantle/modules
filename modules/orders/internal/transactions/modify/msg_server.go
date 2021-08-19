package modify

import (
	"context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/conform"
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

func (msgServer msgServer) Modify(goCtx context.Context, msg *Message) (*TransactionResponse, error) {
	message := messageFromInterface(msg)
	ctx := sdkTypes.UnwrapSDKContext(goCtx)
	if auxiliaryResponse := msgServer.transactionKeeper.verifyAuxiliary.GetKeeper().Help(ctx, verify.NewAuxiliaryRequest(message.From.AsSDKTypesAccAddress(), message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	orders := msgServer.transactionKeeper.mapper.NewCollection(ctx).Fetch(key.FromID(message.OrderID))

	order := orders.Get(key.FromID(message.OrderID))
	if order == nil {
		return nil, errors.EntityNotFound
	}

	metaProperties, Error := supplement.GetMetaPropertiesFromResponse(msgServer.transactionKeeper.supplementAuxiliary.GetKeeper().Help(ctx, supplement.NewAuxiliaryRequest(order.(mappables.Order).GetMakerOwnableSplit())))
	if Error != nil {
		return nil, Error
	}

	transferMakerOwnableSplit := sdkTypes.ZeroDec()

	if makerOwnableSplitProperty := metaProperties.Get(base.NewID(properties.MakerOwnableSplit)); makerOwnableSplitProperty != nil {
		oldMakerOwnableSplit, Error := makerOwnableSplitProperty.GetMetaFact().GetData().AsDec()
		if Error != nil {
			return nil, Error
		}

		transferMakerOwnableSplit = message.MakerOwnableSplit.Sub(oldMakerOwnableSplit)
	} else {
		return nil, errors.MetaDataError
	}

	if transferMakerOwnableSplit.LT(sdkTypes.ZeroDec()) {
		if auxiliaryResponse := msgServer.transactionKeeper.transferAuxiliary.GetKeeper().Help(ctx, transfer.NewAuxiliaryRequest(base.NewID(module.Name), message.FromID, order.(mappables.Order).GetMakerOwnableID(), transferMakerOwnableSplit.Abs())); !auxiliaryResponse.IsSuccessful() {
			return nil, auxiliaryResponse.GetError()
		}
	} else if transferMakerOwnableSplit.GT(sdkTypes.ZeroDec()) {
		if auxiliaryResponse := msgServer.transactionKeeper.transferAuxiliary.GetKeeper().Help(ctx, transfer.NewAuxiliaryRequest(message.FromID, base.NewID(module.Name), order.(mappables.Order).GetMakerOwnableID(), transferMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
			return nil, auxiliaryResponse.GetError()
		}
	}

	mutableMetaProperties := message.MutableMetaProperties.Add(base.NewMetaProperty(base.NewID(properties.MakerOwnableSplit), base.NewMetaFact(base.NewDecData(message.MakerOwnableSplit))))
	mutableMetaProperties = mutableMetaProperties.Add(base.NewMetaProperty(base.NewID(properties.Expiry), base.NewMetaFact(base.NewHeightData(base.NewHeight(message.ExpiresIn.Get()+ctx.BlockHeight())))))

	scrubbedMutableMetaProperties, Error := scrub.GetPropertiesFromResponse(msgServer.transactionKeeper.scrubAuxiliary.GetKeeper().Help(ctx, scrub.NewAuxiliaryRequest(mutableMetaProperties.GetList()...)))
	if Error != nil {
		return nil, Error
	}

	updatedMutables := order.(mappables.Order).GetMutableProperties().Mutate(append(scrubbedMutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...)

	if auxiliaryResponse := msgServer.transactionKeeper.conformAuxiliary.GetKeeper().Help(ctx, conform.NewAuxiliaryRequest(order.(mappables.Order).GetClassificationID(), order.(mappables.Order).GetImmutableProperties(), updatedMutables)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	orders.Remove(order)
	orders.Add(mappable.NewOrder(
		key.NewOrderID(
			order.(mappables.Order).GetClassificationID(),
			order.(mappables.Order).GetMakerOwnableID(),
			order.(mappables.Order).GetTakerOwnableID(),
			base.NewID(message.TakerOwnableSplit.QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(message.MakerOwnableSplit).String()),
			base.NewID(order.(mappables.Order).GetCreation().GetMetaFact().GetData().String()),
			order.(mappables.Order).GetMakerID(), order.(mappables.Order).GetImmutableProperties(),
		),
		order.(mappables.Order).GetImmutableProperties(),
		updatedMutables),
	)
	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}

var _ MsgServer = msgServer{}
