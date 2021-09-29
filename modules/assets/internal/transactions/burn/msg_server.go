package burn

import (
	"context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/burn"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type msgServer struct {
	transactionKeeper
}

var _ MsgServer = msgServer{}

func (msgServer msgServer) Burn(goCtx context.Context, message *Message) (*TransactionResponse, error) {
	ctx := sdkTypes.UnwrapSDKContext(goCtx)
	if auxiliaryResponse := msgServer.transactionKeeper.verifyAuxiliary.GetKeeper().Help(ctx, verify.NewAuxiliaryRequest(message.From.AsSDKTypesAccAddress(), &message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	assets := msgServer.transactionKeeper.mapper.NewCollection(ctx).Fetch(key.FromID(&message.AssetID))

	asset := assets.Get(key.FromID(&message.AssetID))
	if asset == nil {
		return nil, errors.EntityNotFound
	}

	metaProperties, Error := supplement.GetMetaPropertiesFromResponse(msgServer.transactionKeeper.supplementAuxiliary.GetKeeper().Help(ctx, supplement.NewAuxiliaryRequest(asset.(mappables.InterNFT).GetBurn())))
	if Error != nil {
		return nil, Error
	}

	burnHeightMetaFact := metaProperties.Get(base.NewID(properties.Burn))
	if burnHeightMetaFact == nil {
		return nil, errors.EntityNotFound
	}

	burnHeight, Error := burnHeightMetaFact.GetMetaFact().GetData().AsHeight()
	if Error != nil {
		return nil, Error
	}

	if burnHeight.Compare(base.NewHeight(ctx.BlockHeight())) > 0 {
		return nil, errors.NotAuthorized
	}

	if auxiliaryResponse := msgServer.transactionKeeper.burnAuxiliary.GetKeeper().Help(ctx, burn.NewAuxiliaryRequest(message.FromID, message.AssetID, sdkTypes.SmallestDec())); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	assets.Remove(asset)

	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}
