// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package meta

import (
	"context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/metas/module/key"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

type queryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.QueryKeeper = (*queryKeeper)(nil)
var _ QueryServer = &queryKeeper{}

func (queryKeeper queryKeeper) Meta(ctx context.Context, request *QueryRequest) (*QueryResponse, error) {
	sdkCtx := sdkTypes.UnwrapSDKContext(ctx)
	response := queryKeeper.Enquire(sdkCtx, request)
	//TODO: QueryResponse already contains error, no need to add separately
	return response.(*QueryResponse), nil
}

func (queryKeeper queryKeeper) mustEmbedUnimplementedQueryServer() {
	// TODO implement me
	panic("implement me")
}

func (queryKeeper queryKeeper) Enquire(context sdkTypes.Context, queryRequest helpers.QueryRequest) helpers.QueryResponse {
	return newQueryResponse(queryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(baseIDs.NewDataID(baseIDs.NewStringID(queryRequestFromInterface(queryRequest).DataID.Type.IdString), baseIDs.NewHashID(queryRequestFromInterface(queryRequest).DataID.HashId.IdBytes)))), nil)
}

func (queryKeeper queryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	queryKeeper.mapper = mapper
	return queryKeeper
}

func keeperPrototype() helpers.QueryKeeper {
	return queryKeeper{}
}
