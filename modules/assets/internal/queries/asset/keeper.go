// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package asset

import (
	"context"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/schema/helpers"
)

type queryKeeper struct {
	mapper helpers.Mapper
}

func (queryKeeper queryKeeper) Asset(ctx context.Context, request *QueryRequest) (*QueryResponse, error) {
	return queryKeeper.Enquire(sdkTypes.UnwrapSDKContext(ctx), request).(*QueryResponse), nil
}

var _ helpers.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) Enquire(context sdkTypes.Context, queryRequest helpers.QueryRequest) helpers.QueryResponse {
	assetID, err := baseIDs.ReadAssetID(queryRequestFromInterface(queryRequest).AssetID)
	if err != nil {
		panic("Incorrect AssetID")
	}
	return newQueryResponse(queryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(assetID)), nil)
}

func (queryKeeper queryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	queryKeeper.mapper = mapper
	return queryKeeper
}

func keeperPrototype() helpers.QueryKeeper {
	return queryKeeper{}
}
