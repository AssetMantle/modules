package revoke

import (
	"context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/revoke"
)

type msgServer struct {
	transactionKeeper
}

var _ MsgServer = msgServer{}

func (msgServer msgServer) Revoke(goCtx context.Context, message *Message) (*TransactionResponse, error) {
	ctx := sdkTypes.UnwrapSDKContext(goCtx)
	if auxiliaryResponse := msgServer.transactionKeeper.verifyAuxiliary.GetKeeper().Help(ctx, verify.NewAuxiliaryRequest(message.From.AsSDKTypesAccAddress(), &message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	if auxiliaryResponse := msgServer.transactionKeeper.revokeAuxiliary.GetKeeper().Help(ctx, revoke.NewAuxiliaryRequest(&message.FromID, &message.ToID, &message.ClassificationID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}
