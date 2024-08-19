// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"github.com/AssetMantle/modules/utilities/rest"
	"github.com/cosmos/cosmos-sdk/client"
	sdkModuleTypes "github.com/cosmos/cosmos-sdk/types/module"
	"net/http"
	"strings"

	abciTypes "github.com/cometbft/cometbft/abci/types"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/AssetMantle/modules/helpers"
)

type query struct {
	serviceName       string
	cliCommand        helpers.CLICommand
	queryKeeper       helpers.QueryKeeper
	requestPrototype  func() helpers.QueryRequest
	responsePrototype func() helpers.QueryResponse
	keeperPrototype   func() helpers.QueryKeeper
	serviceRegistrar  func(grpc.ServiceRegistrar, helpers.QueryKeeper)
}

var _ helpers.Query = (*query)(nil)

func (query query) RegisterService(configurator sdkModuleTypes.Configurator) {
	if query.queryKeeper == nil {
		panic(fmt.Errorf("query keeper for query %s not initialized", query.serviceName))
	}
	query.serviceRegistrar(configurator.QueryServer(), query.queryKeeper)
}
func (query query) GetServicePath() string {
	splits := strings.Split(query.serviceName, ".")
	return "/" + splits[3] + "/" + splits[5]
}
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

		queryResponse := query.responsePrototype()
		if err := clientContext.Invoke(context.Background(), query.serviceName+"/Handle", queryRequest, queryResponse); err != nil {
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

func (query query) RESTQueryHandler(Context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		responseWriter.Header().Set("Content-Type", "application/json")

		clientContext, ok := rest.ParseQueryHeightOrReturnBadRequest(responseWriter, Context, httpRequest)
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

		queryResponse := query.responsePrototype()
		if err := clientContext.Invoke(context.Background(), query.serviceName+"/Handle", queryRequest, queryResponse); err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(responseWriter, clientContext, queryResponse)
	}
}
func (query query) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaryKeepers ...interface{}) helpers.Query {
	query.queryKeeper = query.keeperPrototype().Initialize(mapper, parameterManager, auxiliaryKeepers).(helpers.QueryKeeper)
	return query
}

func NewQuery(serviceName string, short string, long string, requestPrototype func() helpers.QueryRequest, responsePrototype func() helpers.QueryResponse, keeperPrototype func() helpers.QueryKeeper, serviceRegistrar func(grpc.ServiceRegistrar, helpers.QueryKeeper), flagList ...helpers.CLIFlag) helpers.Query {
	splits := strings.Split(serviceName, ".")
	return query{
		serviceName:       serviceName,
		cliCommand:        NewCLICommand(splits[5], short, long, flagList),
		requestPrototype:  requestPrototype,
		responsePrototype: responsePrototype,
		keeperPrototype:   keeperPrototype,
		serviceRegistrar:  serviceRegistrar,
	}
}
