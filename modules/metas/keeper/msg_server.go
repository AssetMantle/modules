package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	metaTypes "github.com/persistenceOne/persistenceSDK/modules/metas/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) metaTypes.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ metaTypes.MsgServer = msgServer{}

func (k msgServer) Reveal(goCtx context.Context, msg *metaTypes.MsgReveal) (*metaTypes.MsgRevealResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var data types.Data
	err := k.cdc.UnpackAny(msg.MetaFact.Data, &data)
	if err != nil {
		return nil, err
	}
	metaID := metaTypes.GenerateMetaID(data)
	k.Keeper.SetMeta(ctx, metaTypes.NewMeta(metaID, data))
	fmt.Println("META ID:")
	fmt.Println(metaID.String())

	return &metaTypes.MsgRevealResponse{}, nil
}
