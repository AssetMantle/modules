// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"github.com/AssetMantle/modules/utilities/rest"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	sdkModuleTypes "github.com/cosmos/cosmos-sdk/types/module"

	abciTypes "github.com/cometbft/cometbft/abci/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/AssetMantle/modules/helpers"
)

type query struct {
	name                 string
	cliCommand           helpers.CLICommand
	moduleName           string
	queryKeeper          helpers.QueryKeeper
	requestPrototype     func() helpers.QueryRequest
	responsePrototype    func() helpers.QueryResponse
	keeperPrototype      func() helpers.QueryKeeper
	serviceRegistrar     func(grpc.ServiceRegistrar, helpers.QueryKeeper)
	grpcGatewayRegistrar func(client.Context, *runtime.ServeMux) error
	queryHandler         func(client.Context, helpers.QueryRequest) (helpers.QueryResponse, error)
}

var _ helpers.Query = (*query)(nil)

func (query query) RegisterService(configurator sdkModuleTypes.Configurator) {
	if query.queryKeeper == nil {
		panic(fmt.Errorf("query keeper for query %s not initialized", query.name))
	}
	query.serviceRegistrar(configurator.QueryServer(), query.queryKeeper)
}
func (query query) RegisterGRPCGatewayRoute(context client.Context, serveMux *runtime.ServeMux) {
	if err := query.grpcGatewayRegistrar(context, serveMux); err != nil {
		panic(err)
	}
}
func (query query) GetName() string { return query.name }
func (query query) Command() *cobra.Command {
	runE := func(command *cobra.Command, args []string) error {
		clientContext, err := client.GetClientTxContext(command)
		if err != nil {
			return err
		}
		queryRequest, err := query.requestPrototype().FromCLI(query.cliCommand, clientContext)
		if err != nil {
			return err
		}

		queryResponse, err := query.queryHandler(clientContext, queryRequest)
		if err != nil {
			return err
		}

		return clientContext.PrintProto(queryResponse)
	}

	return query.cliCommand.CreateCommand(runE)
}
func (query query) HandleQuery(context context.Context, requestQuery abciTypes.RequestQuery) ([]byte, error) {
	request, err := query.requestPrototype().Decode(requestQuery.Data)
	if err != nil {
		return nil, err
	}

	result, err := query.queryKeeper.Enquire(context, request)
	if err != nil {
		return nil, err
	}
	return result.Encode()
}

func (query query) RESTQueryHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		responseWriter.Header().Set("Content-Type", "application/json")

		clientContext, ok := rest.ParseQueryHeightOrReturnBadRequest(responseWriter, context, httpRequest)
		if !ok {
			return
		}
		queryRequest, err := query.requestPrototype().FromHTTPRequest(httpRequest)
		if err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		if err := queryRequest.Validate(); err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
		}

		response, err := query.queryHandler(clientContext, queryRequest)
		if err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(responseWriter, clientContext, response)
	}
}
func (query query) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaryKeepers ...interface{}) helpers.Query {
	query.queryKeeper = query.keeperPrototype().Initialize(mapper, parameterManager, auxiliaryKeepers).(helpers.QueryKeeper)
	return query
}

func NewQuery(name string, short string, long string, moduleName string, requestPrototype func() helpers.QueryRequest, responsePrototype func() helpers.QueryResponse, keeperPrototype func() helpers.QueryKeeper, serviceRegistrar func(grpc.ServiceRegistrar, helpers.QueryKeeper), grpcGatewayRegistrar func(client.Context, *runtime.ServeMux) error, queryHandler func(client.Context, helpers.QueryRequest) (helpers.QueryResponse, error), flagList ...helpers.CLIFlag) helpers.Query {
	return query{
		name:                 name,
		cliCommand:           NewCLICommand(name, short, long, flagList),
		moduleName:           moduleName,
		requestPrototype:     requestPrototype,
		responsePrototype:    responsePrototype,
		keeperPrototype:      keeperPrototype,
		serviceRegistrar:     serviceRegistrar,
		grpcGatewayRegistrar: grpcGatewayRegistrar,
		queryHandler:         queryHandler,
	}
}
