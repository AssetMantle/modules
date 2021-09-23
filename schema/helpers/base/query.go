/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypesModule "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/spf13/cobra"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"net/http"
)

type query struct {
	name              string
	cliCommand        helpers.CLICommand
	moduleName        string
	queryKeeper       helpers.QueryKeeper
	requestPrototype  func() helpers.QueryRequest
	responsePrototype func() helpers.QueryResponse
	keeperPrototype   func() helpers.QueryKeeper
}

var _ helpers.Query = (*query)(nil)

func (query query) GetName() string { return query.name }

//func (query query) GetCommand() *cobra.Command {
//	cmd := query.cliCommand.CreateQueryCommand()
//	cmd.RunE = func(cmd *cobra.Command, args []string) error {
//		clientCtx, err := client.GetClientQueryContext(cmd)
//		if err != nil {
//			panic(err)
//		}
//		queryRequest := query.requestPrototype().FromCLI(query.cliCommand, clientCtx)
//		response, Error := query.queryKeeper.QueryInKeeper(cmd, clientCtx, queryRequest)
//		if Error != nil {
//			return Error
//		}
//
//		return clientCtx.PrintProto(response)
//	}
//	return cmd
//}
func (query query) Command() *cobra.Command {
	runE := func(command *cobra.Command, args []string) error {
		clientContext, err := client.GetClientQueryContext(command)
		if err != nil {
			return err
		}
		queryRequest := query.requestPrototype().FromCLI(query.cliCommand, clientContext)
		responseBytes, _, Error := query.query(queryRequest, clientContext)
		if Error != nil {
			return Error
		}
		response, Error := query.responsePrototype().LegacyAminoDecode(responseBytes)

		if Error != nil {
			return Error
		}

		return clientContext.PrintProto(response)
	}

	return query.cliCommand.CreateCommand(runE)
}
func (query query) HandleMessageByLegacyAmino(context sdkTypes.Context, legacyAmino *codec.LegacyAmino, requestQuery abciTypes.RequestQuery) ([]byte, error) {
	request, Error := query.requestPrototype().LegacyAminoDecode(requestQuery.Data)
	if Error != nil {
		return nil, Error
	}

	return query.queryKeeper.LegacyEnquire(context, request).LegacyAminoEncode()
}

func (query query) HandleMessage(context sdkTypes.Context, requestQuery abciTypes.RequestQuery) ([]byte, error) {
	request, Error := query.requestPrototype().LegacyAminoDecode(requestQuery.Data)
	if Error != nil {
		return nil, Error
	}

	return query.queryKeeper.QueryInKeeper(context, request).LegacyAminoEncode()

}

//func (query query) HandleMessageByProto(context sdkTypes.Context, requestQuery abciTypes.RequestQuery) ([]byte, error) {
//	request, Error := query.requestPrototype().Decode(requestQuery.Data)
//	if Error != nil {
//		return nil, Error
//	}
//	response, Error := query.queryKeeper.Enquire(context, request)
//	if Error != nil {
//		return nil, Error
//	}
//	return response.Encode()
//
//}

func (query query) RESTQueryHandler(cliContext client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		responseWriter.Header().Set("Content-Type", "application/json")
		cliContext, ok := rest.ParseQueryHeightOrReturnBadRequest(responseWriter, cliContext, httpRequest)

		if !ok {
			return
		}

		queryRequest := query.requestPrototype().FromMap(mux.Vars(httpRequest))
		response, height, Error := query.query(queryRequest, cliContext)

		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, Error.Error())
			return
		}

		cliContext = cliContext.WithHeight(height)
		rest.PostProcessResponse(responseWriter, cliContext, response)
	}
}
func (query query) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaryKeepers ...interface{}) helpers.Query {
	query.queryKeeper = query.keeperPrototype().Initialize(mapper, parameters, auxiliaryKeepers).(helpers.QueryKeeper)
	return query
}

func (query query) RegisterGRPCGatewayRoute(clientContext client.Context, serveMux *runtime.ServeMux) {
	query.keeperPrototype().RegisterGRPCGatewayRoute(clientContext, serveMux)
}

func (query query) query(queryRequest helpers.QueryRequest, cliContext client.Context) ([]byte, int64, error) {
	bytes, Error := queryRequest.Encode(cliContext.JSONMarshaler)
	if Error != nil {
		return nil, 0, Error
	}

	return cliContext.QueryWithData("custom"+"/"+query.moduleName+"/"+query.name, bytes)
}

func (query query) RegisterService(configurator sdkTypesModule.Configurator) {
	query.keeperPrototype().RegisterService(configurator)
}

func NewQuery(name string, short string, long string, moduleName string, requestPrototype func() helpers.QueryRequest, responsePrototype func() helpers.QueryResponse, keeperPrototype func() helpers.QueryKeeper, flagList ...helpers.CLIFlag) helpers.Query {
	return query{
		name:              name,
		cliCommand:        NewCLICommand(name, short, long, flagList),
		moduleName:        moduleName,
		requestPrototype:  requestPrototype,
		responsePrototype: responsePrototype,
		keeperPrototype:   keeperPrototype,
	}
}
