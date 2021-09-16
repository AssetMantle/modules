package renumerate

import (
	"context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/maintain"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/renumerate"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type msgServer struct {
	transactionKeeper
}

var _ MsgServer = msgServer{}

func (msgServer msgServer) Renumerate(goCtx context.Context, message *Message) (*TransactionResponse, error) {
	ctx := sdkTypes.UnwrapSDKContext(goCtx)
	if auxiliaryResponse := msgServer.transactionKeeper.verifyAuxiliary.GetKeeper().Help(ctx, verify.NewAuxiliaryRequest(message.From.AsSDKTypesAccAddress(), &message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	assets := msgServer.transactionKeeper.mapper.NewCollection(ctx).Fetch(key.FromID(&message.AssetID))

	asset := assets.Get(key.FromID(&message.AssetID))
	if asset == nil {
		return nil, errors.EntityNotFound
	}

	if auxiliaryResponse := msgServer.transactionKeeper.maintainAuxiliary.GetKeeper().Help(ctx, maintain.NewAuxiliaryRequest(asset.(mappables.InterNFT).GetClassificationID(), &message.FromID, base.NewProperties(base.NewProperty(base.NewID(properties.Value), nil)))); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	metaProperties, Error := supplement.GetMetaPropertiesFromResponse(msgServer.transactionKeeper.supplementAuxiliary.GetKeeper().Help(ctx, supplement.NewAuxiliaryRequest(asset.(mappables.InterNFT).GetValue())))
	if Error != nil {
		return nil, Error
	}

	if valueMetaProperty := metaProperties.Get(base.NewID(properties.Value)); valueMetaProperty != nil {
		if value, Error := valueMetaProperty.GetMetaFact().GetData().AsDec(); Error != nil {
			return nil, Error
		} else if auxiliaryResponse := msgServer.transactionKeeper.renumerateAuxiliary.GetKeeper().Help(ctx, renumerate.NewAuxiliaryRequest(message.FromID, message.AssetID, value)); !auxiliaryResponse.IsSuccessful() {
			return nil, auxiliaryResponse.GetError()
		}
	} else {
		return nil, errors.MetaDataError
	}

	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}
