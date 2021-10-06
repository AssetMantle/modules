package deputize

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/deputize"
)

type msgServer struct {
	transactionKeeper
}

func (msgServer msgServer) Deputize(goCtx context.Context, message *Message) (*TransactionResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	if auxiliaryResponse := msgServer.transactionKeeper.verifyAuxiliary.GetKeeper().Help(ctx, verify.NewAuxiliaryRequest(message.From.AsSDKTypesAccAddress(), &message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	if auxiliaryResponse := msgServer.transactionKeeper.deputizeAuxiliary.GetKeeper().Help(ctx, deputize.NewAuxiliaryRequest(&message.FromID, &message.ToID, &message.ClassificationID, &message.MaintainedProperties, message.AddMaintainer, message.RemoveMaintainer, message.MutateMaintainer)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}

var _ MsgServer = msgServer{}
