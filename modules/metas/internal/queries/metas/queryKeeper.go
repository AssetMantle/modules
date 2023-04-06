// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package metas

import (
	"context"
	"github.com/AssetMantle/modules/schema/ids/base"

	"github.com/AssetMantle/modules/modules/metas/internal/key"
	"github.com/AssetMantle/modules/schema/helpers"
)

type queryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) Enquire(context context.Context, queryRequest helpers.QueryRequest) helpers.QueryResponse {
	queryResponse, _ := queryKeeper.Handle(context, queryRequestFromInterface(queryRequest))
	return queryResponse
}
func (queryKeeper queryKeeper) Handle(context context.Context, queryRequest *QueryRequest) (*QueryResponse, error) {
	return newQueryResponse(queryKeeper.mapper.NewCollection(context).FetchPaginated(key.NewKey(base.PrototypeDataID()), queryRequest.Pagination), nil), nil
}

func (queryKeeper queryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	queryKeeper.mapper = mapper
	return queryKeeper
}

func keeperPrototype() helpers.QueryKeeper {
	return queryKeeper{}
}
