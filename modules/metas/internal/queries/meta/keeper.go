/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package meta

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/spf13/cobra"
)

type queryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) LegacyEnquire(context sdkTypes.Context, queryRequest helpers.QueryRequest) helpers.QueryResponse {
	queryResponse := newQueryResponse(queryKeeper.mapper.NewCollection(context).Fetch(key.FromID(base.NewID(queryRequestFromInterface(queryRequest).MetaID.IdString))), nil)
	return &queryResponse
}

func (queryKeeper queryKeeper) GetQueryClient(ctx client.Context) QueryClient {
	return NewQueryClient(ctx)
}

func (queryKeeper queryKeeper) Enquire(clientctx client.Context, ctx sdkTypes.Context, request helpers.QueryRequest) (helpers.QueryResponse, error) {
	queryRequest := request.(QueryRequest)
	queryCli := queryKeeper.GetQueryClient(clientctx)
	response, Error := queryCli.Get(sdkTypes.WrapSDKContext(ctx), &queryRequest)
	//response, Error := queryKeeper.Get(sdkTypes.WrapSDKContext(ctx), &queryRequest)
	return response, Error
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

func (queryKeeper queryKeeper) QueryInKeeper(ctx sdkTypes.Context, req helpers.QueryRequest) (json.RawMessage, error) {
	request := req.(QueryRequest)
	goCtx := sdkTypes.WrapSDKContext(ctx)
	queryServer := NewQueryServerImpl(queryKeeper)
	queryRes, Error := queryServer.Get(goCtx, &request)
	if Error == (errors.New("yes")) {
		panic(Error)
	}
	return json.Marshal(queryRes)
}

func (queryKeeper queryKeeper) RegisterService(cfg module.Configurator) {
	RegisterQueryServer(cfg.QueryServer(), NewQueryServerImpl(queryKeeper))
}

func keeperPrototype() helpers.QueryKeeper {
	return queryKeeper{}
}

func queryInKeeper(mapper helpers.Mapper, command *cobra.Command, clientCtx client.Context, req helpers.QueryRequest) (helpers.QueryResponse, error) {
	queryClient := NewQueryClient(clientCtx)
	meta, Error := command.Flags().GetString("metaID")
	if Error != nil {
		panic(Error)
	}
	newMetaID := base.NewID(meta)
	params := NewQueryGet(newMetaID)
	return queryClient.Get(command.Context(), params)
}

func NewQueryGet(metaID types.ID) *QueryRequest {
	return &QueryRequest{
		MetaID: *base.NewID(metaID.String()),
	}
}
