/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package order

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type queryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.QueryKeeper = (*queryKeeper)(nil)
var _ QueryServer = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) LegacyEnquire(context sdkTypes.Context, queryRequest helpers.QueryRequest) helpers.QueryResponse {
	return newQueryResponse(queryKeeper.mapper.NewCollection(context).Fetch(key.FromID(queryRequestFromInterface(queryRequest).OrderID)), nil)
}

func (queryKeeper queryKeeper) Enquire(context context.Context, queryRequest *QueryRequest) (*QueryResponse, error) {
	response := newQueryResponse(queryKeeper.mapper.NewCollection(sdkTypes.UnwrapSDKContext(context)).Fetch(key.FromID(queryRequest.OrderID)), nil)
	return &response, response.GetError()
}

func (queryKeeper queryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	queryKeeper.mapper = mapper
	return queryKeeper
}

func (queryKeeper queryKeeper) RegisterGRPCGatewayRoute(clientContext client.Context, serveMux *runtime.ServeMux) {
	err := RegisterQueryHandlerClient(context.Background(), serveMux, NewQueryClient(clientContext))
	if err != nil {
		panic(err)
	}
}

func (queryKeeper queryKeeper) RegisterService(cfg module.Configurator) {
	RegisterQueryServer(cfg.QueryServer(), queryKeeper)
}

func keeperPrototype() helpers.QueryKeeper {
	return queryKeeper{}
}
