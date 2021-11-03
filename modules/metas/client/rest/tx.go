package rest

import (
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/modules/metas/types"
)

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc(
		"/persistenceSDK/metas/meta/reveal",
		newPostRevealHandlerFn(clientCtx),
	).Methods("POST")
}

type (
	RevealRequest struct {
		BaseReq   rest.BaseReq `json:"baseReq" yaml:"baseReq"`
		DataType  string       `json:"dataTypes" yaml:"dataTypes"`
		DataValue string       `json:"dataValue" yaml:"dataValue"`
	}
)

func newPostRevealHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RevealRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		data := base.NewStringData(req.DataValue)
		fromAddr, err := sdkTypes.AccAddressFromBech32(req.BaseReq.From)
		if rest.CheckBadRequestError(w, err) {
			return
		}
		msg := types.NewMsgReveal(fromAddr, data)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
