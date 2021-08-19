package send

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/utilities"
)

type msgServer struct {
	transactionKeeper
}

func (msgServer msgServer) Send(goCtx context.Context, msg *Message) (*TransactionResponse, error) {
	message := messageFromInterface(msg)
	ctx := types.UnwrapSDKContext(goCtx)
	if auxiliaryResponse := msgServer.transactionKeeper.verifyAuxiliary.GetKeeper().Help(ctx, verify.NewAuxiliaryRequest(message.From.AsSDKTypesAccAddress(), message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	splits := msgServer.transactionKeeper.mapper.NewCollection(ctx)

	if _, Error := utilities.SubtractSplits(splits, message.FromID, message.OwnableID, message.Value); Error != nil {
		return nil, Error
	}

	if _, Error := utilities.AddSplits(splits, message.ToID, message.OwnableID, message.Value); Error != nil {
		return nil, Error
	}

	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}

var _ MsgServer = msgServer{}
