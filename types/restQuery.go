package types

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type RESTQuery interface {
	CreateQuery(context.CLIContext, func(map[string]string) []byte) http.HandlerFunc
}

type restQuery struct {
	Route string
	Path  string
}

var _ RESTQuery = (*restQuery)(nil)

func (restQuery restQuery) CreateQuery(cliContext context.CLIContext, makeQueryBytes func(map[string]string) []byte) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		responseWriter.Header().Set("Content-Type", "application/json")
		cliContext, ok := rest.ParseQueryHeightOrReturnBadRequest(responseWriter, cliContext, httpRequest)
		if !ok {
			return
		}

		vars := mux.Vars(httpRequest)
		bytes := makeQueryBytes(vars)
		response, height, err := cliContext.QueryWithData(strings.Join([]string{"", "custom", restQuery.Route, restQuery.Path}, "/"), bytes)
		if err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		cliContext = cliContext.WithHeight(height)
		rest.PostProcessResponse(responseWriter, cliContext, response)
	}
}
func NewRESTQuery(route string, path string) RESTQuery {
	return &restQuery{
		Route: route,
		Path:  path,
	}
}
