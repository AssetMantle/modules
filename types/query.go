package types

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"net/http"
	"strings"
)

type Query interface {
	GetModuleName() string
	GetName() string
	Command(*codec.Codec) *cobra.Command
	HandleMessage(sdkTypes.Context, QueryKeeper, abciTypes.RequestQuery) ([]byte, error)
	RESTQueryHandler(context.CLIContext) http.HandlerFunc
	RegisterCodec(*codec.Codec)
}

type query struct {
	ModuleName             string
	Name                   string
	CLICommand             CLICommand
	PackageCodec           *codec.Codec
	Codec                  func(*codec.Codec)
	QueryRequestPrototype  func() QueryRequest
	QueryResponsePrototype func() QueryResponse
}

var _ Query = (*query)(nil)

func (query query) GetModuleName() string { return query.ModuleName }

func (query query) GetName() string { return query.Name }

func (query query) Command(codec *codec.Codec) *cobra.Command {
	runE := func(command *cobra.Command, args []string) error {
		cliContext := context.NewCLIContext().WithCodec(codec)

		queryRequest := query.QueryRequestPrototype().FromCLI(query.CLICommand, cliContext)
		bytes := codec.MustMarshalJSON(queryRequest)
		responseBytes, _, Error := cliContext.QueryWithData(strings.Join([]string{"", "custom", query.ModuleName, query.Name}, "/"), bytes)
		if Error != nil {
			return Error
		}
		response := query.QueryResponsePrototype()
		if Error := codec.UnmarshalJSON(responseBytes, &response); Error != nil {
			return Error
		}
		return cliContext.PrintOutput(response)
	}
	return query.CLICommand.CreateCommand(runE)
}
func (query query) HandleMessage(context sdkTypes.Context, queryKeeper QueryKeeper, requestQuery abciTypes.RequestQuery) ([]byte, error) {
	queryRequest := query.QueryRequestPrototype()
	Error := query.PackageCodec.UnmarshalJSON(requestQuery.Data, &queryRequest)
	if Error != nil {
		return nil, Error
	}
	return queryKeeper.Query(context, queryRequest)
}

func (query query) RESTQueryHandler(cliContext context.CLIContext) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		responseWriter.Header().Set("Content-Type", "application/json")
		cliContext, ok := rest.ParseQueryHeightOrReturnBadRequest(responseWriter, cliContext, httpRequest)
		if !ok {
			return
		}

		bytes, Error := query.PackageCodec.MarshalBinaryBare(query.QueryRequestPrototype().FromMap(mux.Vars(httpRequest)))
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, Error.Error())
			return
		}
		response, height, Error := cliContext.QueryWithData(strings.Join([]string{"", "custom", query.ModuleName, query.Name}, "/"), bytes)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, Error.Error())
			return
		}

		cliContext = cliContext.WithHeight(height)
		rest.PostProcessResponse(responseWriter, cliContext, response)
	}
}
func (query query) RegisterCodec(codec *codec.Codec) {
	query.Codec(codec)
}
func NewQuery(module string, name string, short string, long string, queryRequestPrototype func() QueryRequest, queryResponsePrototype func() QueryResponse, registerCodec func(*codec.Codec), flagList []CLIFlag) Query {
	return &query{
		ModuleName:             module,
		Name:                   name,
		CLICommand:             NewCLICommand(name, short, long, flagList),
		Codec:                  registerCodec,
		QueryRequestPrototype:  queryRequestPrototype,
		QueryResponsePrototype: queryResponsePrototype,
	}
}
