/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/spf13/cobra"
	abciTypes "github.com/tendermint/tendermint/abci/types"
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
func (query query) Command(codec *codec.Codec) *cobra.Command {
	runE := func(command *cobra.Command, args []string) error {
		cliContext := context.NewCLIContext().WithCodec(codec)

		queryRequest := query.requestPrototype().FromCLI(query.cliCommand, cliContext)
		responseBytes, _, err := query.query(queryRequest, cliContext)

		if err != nil {
			return err
		}

		response, err := query.responsePrototype().Decode(responseBytes)
		if err != nil {
			return err
		}

		return cliContext.PrintOutput(response)
	}

	return query.cliCommand.CreateCommand(runE)
}
func (query query) HandleMessage(context sdkTypes.Context, requestQuery abciTypes.RequestQuery) ([]byte, error) {
	request, err := query.requestPrototype().Decode(requestQuery.Data)
	if err != nil {
		return nil, err
	}

	return query.queryKeeper.Enquire(context, request).Encode()
}

func (query query) RESTQueryHandler(outerCliContext context.CLIContext) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		responseWriter.Header().Set("Content-Type", "application/json")

		cliContext, ok := rest.ParseQueryHeightOrReturnBadRequest(responseWriter, outerCliContext, httpRequest)
		if !ok {
			return
		}

		queryRequest := query.requestPrototype().FromMap(mux.Vars(httpRequest))

		response, height, err := query.query(queryRequest, cliContext)
		if err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, err.Error())
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

func (query query) query(queryRequest helpers.QueryRequest, cliContext context.CLIContext) ([]byte, int64, error) {
	bytes, err := queryRequest.Encode()
	if err != nil {
		return nil, 0, err
	}

	return cliContext.QueryWithData("custom"+"/"+query.moduleName+"/"+query.name, bytes)
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
