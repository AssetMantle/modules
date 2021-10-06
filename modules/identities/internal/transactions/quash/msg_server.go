package quash

import (
	"context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/schema/mappables" //nolint:typecheck
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type msgServer struct {
	transactionKeeper
}

var _ MsgServer = msgServer{}

func (msgServer msgServer) Quash(goCtx context.Context, message *Message) (*TransactionResponse, error) {
	ctx := sdkTypes.UnwrapSDKContext(goCtx)
	if auxiliaryResponse := msgServer.transactionKeeper.verifyAuxiliary.GetKeeper().Help(ctx, verify.NewAuxiliaryRequest(message.From.AsSDKTypesAccAddress(), &message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	identities := msgServer.transactionKeeper.mapper.NewCollection(ctx).Fetch(key.FromID(&message.IdentityID))

	identity := identities.Get(key.FromID(&message.IdentityID))
	if identity == nil {
		return nil, errors.EntityNotFound
	}

	metaProperties, Error := supplement.GetMetaPropertiesFromResponse(msgServer.transactionKeeper.supplementAuxiliary.GetKeeper().Help(ctx, supplement.NewAuxiliaryRequest(identity.(mappables.InterIdentity).GetExpiry())))
	if Error != nil {
		return nil, Error
	}

	expiryHeightMetaFact := metaProperties.Get(base.NewID(properties.Expiry))
	if expiryHeightMetaFact == nil {
		return nil, errors.EntityNotFound
	}

	expiryHeight, Error := expiryHeightMetaFact.GetMetaFact().GetData().AsHeight()
	if Error != nil {
		return nil, Error
	}

	if expiryHeight.Compare(base.NewHeight(ctx.BlockHeight())) > 0 {
		return nil, errors.NotAuthorized
	}

	identities.Remove(identity)

	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}
