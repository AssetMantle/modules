// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supply

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/splits/utilities"
)

type queryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) Enquire(context context.Context, queryRequest helpers.QueryRequest) (helpers.QueryResponse, error) {
	return queryKeeper.Handle(context, queryRequest.(*QueryRequest))
}
func (queryKeeper queryKeeper) Handle(context context.Context, queryRequest *QueryRequest) (*QueryResponse, error) {
	if err := queryRequest.Validate(); err != nil {
		return nil, err
	}
	return newQueryResponse(utilities.GetTotalSupply(queryKeeper.mapper.NewCollection(context), queryRequest.AssetID)), nil
}

func (queryKeeper queryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	queryKeeper.mapper = mapper
	return queryKeeper
}

func keeperPrototype() helpers.QueryKeeper {
	return queryKeeper{}
}
