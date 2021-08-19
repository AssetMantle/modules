package unprovision

import (
	"context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
)

type msgServer struct {
	transactionKeeper
}

var _ MsgServer = msgServer{}

func (msgServer msgServer) Unprovision(goCtx context.Context, msg *Message) (*TransactionResponse, error) {
	message := messageFromInterface(msg)
	ctx := sdkTypes.UnwrapSDKContext(goCtx)
	identityID := message.IdentityID
	identities := msgServer.transactionKeeper.mapper.NewCollection(ctx).Fetch(key.FromID(identityID))

	identity := identities.Get(key.FromID(identityID))
	if identity == nil {
		return nil, errors.EntityNotFound
	}

	if !identity.(mappables.InterIdentity).IsProvisioned(message.From.AsSDKTypesAccAddress()) {
		return nil, errors.NotAuthorized
	}

	if !identity.(mappables.InterIdentity).IsProvisioned(message.To.AsSDKTypesAccAddress()) {
		return nil, errors.EntityNotFound
	}

	identities.Mutate(identity.(mappables.InterIdentity).UnprovisionAddress(message.To.AsSDKTypesAccAddress()))

	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}
