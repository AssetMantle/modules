// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
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
	return newQueryResponse(queryKeeper.mapper.NewCollection(context).FetchRecord(queryRequest.Key)), nil
}

func (queryKeeper queryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	queryKeeper.mapper = mapper
	helpers.PanicOnUninitializedKeeperFields(queryKeeper)
	return queryKeeper
}

func keeperPrototype() helpers.QueryKeeper {
	return queryKeeper{}
}
