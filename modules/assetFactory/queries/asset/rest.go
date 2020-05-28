package asset

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"net/http"
	"strings"
)

func RestQueryHandler(cliContext context.CLIContext) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		responseWriter.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(httpRequest)
		id := types.BaseID{BaseString: vars["id"]}

		cliContext, ok := rest.ParseQueryHeightOrReturnBadRequest(responseWriter, cliContext, httpRequest)
		if !ok {
			return
		}
		bytes := packageCodec.MustMarshalJSON(query{id: id})
		response, height, err := cliContext.QueryWithData(strings.Join([]string{"", "custom", constants.QuerierRoute, constants.AssetQuery}, "/"), bytes)
		if err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, err.Error())
			return
		}

		cliContext = cliContext.WithHeight(height)
		rest.PostProcessResponse(responseWriter, cliContext, response)
	}
}
