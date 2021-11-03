package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/persistenceOne/persistenceSDK/modules/metas/types"
	"net/http"
)

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// Get a single validator info
	r.HandleFunc(
		"/persistenceSDK/metas/meta/{metaID}",
		validatorHandlerFn(clientCtx),
	).Methods("GET")

	// Get the current staking parameter values
	r.HandleFunc(
		"/persistenceSDK/metas/parameters",
		paramsHandlerFn(clientCtx),
	).Methods("GET")
}

// HTTP request handler to query the validator information from a given validator address
func validatorHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		clientCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		params := types.QueryMetaRequest{MetaID: vars["metaID"]}
		bz, err := clientCtx.LegacyAmino.MarshalJSON(params)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		res, height, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryMeta), bz)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}

// HTTP request handler to query the staking params values
func paramsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		res, height, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryParameters), nil)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}
