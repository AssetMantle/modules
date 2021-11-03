package keeper

import (
	"context"
	"fmt"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/metas/types"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

func (q Querier) GetMeta(c context.Context, req *types.QueryMetaRequest) (*types.QueryMetaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var meta types.Meta
	ctx := sdk.UnwrapSDKContext(c)

	id := base.NewID(req.MetaID)
	metaID := types.MetaIDFromInterface(&id)

	meta, err := q.Keeper.GetMeta(ctx, metaID)
	if err != nil {
		return &types.QueryMetaResponse{Success: false, Error: fmt.Sprintf("meta %s not found", req.MetaID), Value: types.Meta{}}, status.Errorf(codes.NotFound, "meta %s not found", req.MetaID)
	}

	return &types.QueryMetaResponse{Success: true, Error: "", Value: meta}, nil
}

// Params queries the staking parameters
func (q Querier) GetParameters(c context.Context, _ *types.QueryParametersRequest) (*types.QueryParametersResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := q.Keeper.GetParameters(ctx)

	return &types.QueryParametersResponse{Parameters: params}, nil
}
