// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	sdkModuleTypes "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gogo/protobuf/grpc"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"golang.org/x/net/context"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
)

type query struct {
	name                 string
	cliCommand           helpers.CLICommand
	moduleName           string
	queryKeeper          helpers.QueryKeeper
	requestPrototype     func() helpers.QueryRequest
	responsePrototype    func() helpers.QueryResponse
	keeperPrototype      func() helpers.QueryKeeper
	serviceRegistrar     func(grpc.Server, helpers.QueryKeeper)
	grpcGatewayRegistrar func(client.Context, *runtime.ServeMux) error
}

var _ helpers.Query = (*query)(nil)

func (query query) RegisterService(configurator sdkModuleTypes.Configurator) {
	if query.queryKeeper == nil {
		panic(errorConstants.UninitializedUsage.Wrapf("query keeper for query %s not initialized", query.name))
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

		responseBytes, _, err := query.query(queryRequest, clientContext)
		if err != nil {
			return err
		}

		response, err := query.responsePrototype().Decode(responseBytes)
		if err != nil {
			return err
		}

		return clientContext.PrintProto(response)
	}

	return query.cliCommand.CreateCommand(runE)
}
func (query query) HandleQuery(context context.Context, requestQuery abciTypes.RequestQuery) ([]byte, error) {
	request, err := query.requestPrototype().Decode(requestQuery.Data)
	if err != nil {
		return nil, err
	}

	return query.queryKeeper.Enquire(context, request).Encode()
}

func (query query) RESTQueryHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		responseWriter.Header().Set("Content-Type", "application/json")

		clientContext, ok := rest.ParseQueryHeightOrReturnBadRequest(responseWriter, context, httpRequest)
		if !ok {
			return
		}

		queryRequest, err := query.requestPrototype().FromMap(mux.Vars(httpRequest))
		if err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		response, height, err := query.query(queryRequest, clientContext)
		if err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		clientContext = clientContext.WithHeight(height)
		rest.PostProcessResponse(responseWriter, clientContext, response)
	}
}
func (query query) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaryKeepers ...interface{}) helpers.Query {
	query.queryKeeper = query.keeperPrototype().Initialize(mapper, parameterManager, auxiliaryKeepers).(helpers.QueryKeeper)
	return query
}

func (query query) query(queryRequest helpers.QueryRequest, context client.Context) ([]byte, int64, error) {
	bytes, err := queryRequest.Encode()
	if err != nil {
		return nil, 0, err
	}

	return context.QueryWithData("custom"+"/"+query.moduleName+"/"+query.name, bytes)
}

func NewQuery(name string, short string, long string, moduleName string, requestPrototype func() helpers.QueryRequest, responsePrototype func() helpers.QueryResponse, keeperPrototype func() helpers.QueryKeeper, serviceRegistrar func(grpc.Server, helpers.QueryKeeper), grpcGatewayRegistrar func(client.Context, *runtime.ServeMux) error, flagList ...helpers.CLIFlag) helpers.Query {
	return query{
		name:                 name,
		cliCommand:           NewCLICommand(name, short, long, flagList),
		moduleName:           moduleName,
		requestPrototype:     requestPrototype,
		responsePrototype:    responsePrototype,
		keeperPrototype:      keeperPrototype,
		serviceRegistrar:     serviceRegistrar,
		grpcGatewayRegistrar: grpcGatewayRegistrar,
	}
}
