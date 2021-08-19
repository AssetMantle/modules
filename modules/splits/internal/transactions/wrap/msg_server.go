package wrap

import (
	"context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/utilities"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type msgServer struct {
	transactionKeeper
}

func (msgServer msgServer) Wrap(goCtx context.Context, msg *Message) (*TransactionResponse, error) {
	message := messageFromInterface(msg)
	ctx := sdkTypes.UnwrapSDKContext(goCtx)
	if auxiliaryResponse := msgServer.transactionKeeper.verifyAuxiliary.GetKeeper().Help(ctx, verify.NewAuxiliaryRequest(message.From.AsSDKTypesAccAddress(), message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	if Error := msgServer.transactionKeeper.bankKeeper.SendCoinsFromAccountToModule(ctx, message.From.AsSDKTypesAccAddress(), module.Name, message.Coins); Error != nil {
		return nil, Error
	}

	for _, coin := range message.Coins {
		if _, Error := utilities.AddSplits(msgServer.transactionKeeper.mapper.NewCollection(ctx), message.FromID, base.NewID(coin.Denom), sdkTypes.NewDecFromInt(coin.Amount)); Error != nil {
			return nil, Error
		}
	}

	return &TransactionResponse{}, nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}

var _ MsgServer = msgServer{}
