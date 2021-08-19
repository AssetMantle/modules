/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package ownable

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/utilities"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type queryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.QueryKeeper = (*queryKeeper)(nil)
var _ QueryServer = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) LegacyEnquire(context sdkTypes.Context, queryRequest helpers.QueryRequest) helpers.QueryResponse {
	return newQueryResponse(utilities.GetOwnableTotalSplitsValue(queryKeeper.mapper.NewCollection(context), queryRequestFromInterface(queryRequest).OwnableID), nil)
}

func (queryKeeper queryKeeper) Enquire(context context.Context, queryRequest *QueryRequest) (*QueryResponse, error) {
	response := newQueryResponse(utilities.GetOwnableTotalSplitsValue(queryKeeper.mapper.NewCollection(sdkTypes.UnwrapSDKContext(context)), queryRequest.OwnableID), nil)
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
