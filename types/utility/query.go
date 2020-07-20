package utility

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
	GetRoute() string
	Command(*codec.Codec) *cobra.Command
	HandleMessage(sdkTypes.Context, abciTypes.RequestQuery) ([]byte, error)
	RESTQueryHandler(context.CLIContext) http.HandlerFunc
	RegisterCodec(*codec.Codec)
	InitializeKeeper(Mapper)
	query(QueryRequest, context.CLIContext) ([]byte, int64, error)
}

type query struct {
	moduleName             string
	name                   string
	route                  string
	queryKeeper            QueryKeeper
	cliCommand             CLICommand
	packageCodec           *codec.Codec
	registerCodec          func(*codec.Codec)
	initializeKeeper       func(Mapper) QueryKeeper
	queryRequestPrototype  func() QueryRequest
	queryResponsePrototype func() QueryResponse
}

var _ Query = (*query)(nil)

func (query query) GetModuleName() string { return query.moduleName }
func (query query) GetName() string       { return query.name }
func (query query) GetRoute() string      { return query.route }
func (query query) Command(codec *codec.Codec) *cobra.Command {
	runE := func(command *cobra.Command, args []string) error {
		cliContext := context.NewCLIContext().WithCodec(codec)

		queryRequest := query.queryRequestPrototype().FromCLI(query.cliCommand, cliContext)
		responseBytes, _, Error := query.query(queryRequest, cliContext)
		if Error != nil {
			return Error
		}
		response := query.queryResponsePrototype()
		if Error := query.packageCodec.UnmarshalJSON(responseBytes, &response); Error != nil {
			return Error
		}
		return cliContext.PrintOutput(response)
	}
	return query.cliCommand.CreateCommand(runE)
}
func (query query) HandleMessage(context sdkTypes.Context, requestQuery abciTypes.RequestQuery) ([]byte, error) {
	queryRequest := query.queryRequestPrototype()
	Error := query.packageCodec.UnmarshalJSON(requestQuery.Data, &queryRequest)
	if Error != nil {
		return nil, Error
	}
	return query.packageCodec.MarshalJSON(query.queryKeeper.Query(context, queryRequest))
}

func (query query) RESTQueryHandler(cliContext context.CLIContext) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		responseWriter.Header().Set("Content-Type", "application/json")
		cliContext, ok := rest.ParseQueryHeightOrReturnBadRequest(responseWriter, cliContext, httpRequest)
		if !ok {
			return
		}

		queryRequest := query.queryRequestPrototype().FromMap(mux.Vars(httpRequest))
		response, height, Error := query.query(queryRequest, cliContext)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, Error.Error())
			return
		}
		cliContext = cliContext.WithHeight(height)
		rest.PostProcessResponse(responseWriter, cliContext, response)
	}
}
func (query query) RegisterCodec(codec *codec.Codec) {
	query.registerCodec(codec)
}

func (query *query) InitializeKeeper(mapper Mapper) {
	query.queryKeeper = query.initializeKeeper(mapper)
}

func (query query) query(queryRequest QueryRequest, cliContext context.CLIContext) ([]byte, int64, error) {
	bytes, Error := query.packageCodec.MarshalJSON(queryRequest)
	if Error != nil {
		return nil, 0, Error
	}
	return cliContext.QueryWithData(strings.Join([]string{"", "custom", query.moduleName, query.name}, "/"), bytes)
}

func NewQuery(module string, name string, route string, short string, long string, packageCodec *codec.Codec, registerCodec func(*codec.Codec), initializeKeeper func(Mapper) QueryKeeper, queryRequestPrototype func() QueryRequest, queryResponsePrototype func() QueryResponse, flagList []CLIFlag) Query {
	return &query{
		moduleName:             module,
		name:                   name,
		route:                  route,
		cliCommand:             NewCLICommand(name, short, long, flagList),
		packageCodec:           packageCodec,
		registerCodec:          registerCodec,
		initializeKeeper:       initializeKeeper,
		queryRequestPrototype:  queryRequestPrototype,
		queryResponsePrototype: queryResponsePrototype,
	}
}
