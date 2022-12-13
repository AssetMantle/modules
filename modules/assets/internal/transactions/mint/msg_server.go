package mint

import (
	"context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	propertiesConstants "github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/conform"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/mint"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type msgServer struct {
	transactionKeeper
}

var _ MsgServer = msgServer{}

func (msgServer msgServer) Mint(goCtx context.Context, message *Message) (*TransactionResponse, error) {
	ctx := sdkTypes.UnwrapSDKContext(goCtx)
	if auxiliaryResponse := msgServer.transactionKeeper.verifyAuxiliary.GetKeeper().Help(ctx, verify.NewAuxiliaryRequest(message.From.AsSDKTypesAccAddress(), &message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	immutableMetaProperties, Error := scrub.GetPropertiesFromResponse(msgServer.transactionKeeper.scrubAuxiliary.GetKeeper().Help(ctx, scrub.NewAuxiliaryRequest(message.ImmutableMetaProperties.GetList()...)))
	if Error != nil {
		return nil, Error
	}

	immutableProperties := base.NewProperties(append(immutableMetaProperties.GetList(), message.ImmutableProperties.GetList()...)...)

	assetID := key.NewAssetID(&message.ClassificationID, immutableProperties)

	assets := msgServer.transactionKeeper.mapper.NewCollection(ctx).Fetch(key.FromID(assetID))
	if assets.Get(key.FromID(assetID)) != nil {
		return nil, errors.EntityAlreadyExists
	}

	mutableMetaProperties, Error := scrub.GetPropertiesFromResponse(msgServer.transactionKeeper.scrubAuxiliary.GetKeeper().Help(ctx, scrub.NewAuxiliaryRequest(message.MutableMetaProperties.GetList()...)))
	if Error != nil {
		return nil, Error
	}

	mutableProperties := base.NewProperties(append(mutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...)

	if auxiliaryResponse := msgServer.transactionKeeper.conformAuxiliary.GetKeeper().Help(ctx, conform.NewAuxiliaryRequest(&message.ClassificationID, immutableProperties, mutableProperties)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	split := sdkTypes.SmallestDec()

	if metaProperties := base.NewMetaProperties(append(message.ImmutableMetaProperties.GetList(), message.MutableMetaProperties.GetList()...)...); metaProperties.Get(base.NewID(propertiesConstants.Lock)) != nil {
		if split, Error = metaProperties.Get(base.NewID(propertiesConstants.Lock)).GetMetaFact().GetData().AsDec(); Error != nil {
			return nil, errors.MetaDataError
		}
	}

	if auxiliaryResponse := msgServer.transactionKeeper.mintAuxiliary.GetKeeper().Help(ctx, mint.NewAuxiliaryRequest(message.ToID, assetID, split)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	assets.Add(mappable.NewAsset(assetID, immutableProperties, mutableProperties))

	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}
