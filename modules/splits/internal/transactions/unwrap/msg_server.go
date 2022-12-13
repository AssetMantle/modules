package unwrap

import (
	"context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/utilities"
)

type msgServer struct {
	transactionKeeper
}

func (msgServer msgServer) Unwrap(goCtx context.Context, message *Message) (*TransactionResponse, error) {
	ctx := sdkTypes.UnwrapSDKContext(goCtx)
	if auxiliaryResponse := msgServer.transactionKeeper.verifyAuxiliary.GetKeeper().Help(ctx, verify.NewAuxiliaryRequest(message.From.AsSDKTypesAccAddress(), &message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	splits := msgServer.transactionKeeper.mapper.NewCollection(ctx)
	if _, Error := utilities.SubtractSplits(splits, &message.FromID, &message.OwnableID, sdkTypes.NewDecFromInt(message.Value)); Error != nil {
		return nil, Error
	}

	if Error := msgServer.transactionKeeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, module.Name, message.From.AsSDKTypesAccAddress(), sdkTypes.NewCoins(sdkTypes.NewCoin(message.OwnableID.String(), message.Value))); Error != nil {
		return nil, Error
	}

	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}

var _ MsgServer = msgServer{}
