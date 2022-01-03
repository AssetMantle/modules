package keeper

import (
	"context"

	"github.com/persistenceOne/persistenceSDK/modules/liquidStaking/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) AddChain(goCtx context.Context, msg *types.MsgAddChain) (*types.MsgAddChainResponse, error) {

	return &types.MsgAddChainResponse{}, nil
}
