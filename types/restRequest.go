package types

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client"
	"net/http"
)

type RESTRequest interface {
	CreateRequest(context.CLIContext, func(Request) (rest.BaseReq, sdkTypes.Msg)) http.HandlerFunc
}
type Request interface{}

type restRequest struct {
	Request Request
}

var _ RESTRequest = (*restRequest)(nil)

func (restRequest restRequest) CreateRequest(cliContext context.CLIContext, makeBaseReqAndMsg func(Request) (rest.BaseReq, sdkTypes.Msg)) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {

		if !rest.ReadRESTReq(responseWriter, httpRequest, cliContext.Codec, &restRequest.Request) {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "")
			return
		}

		baseReq, msg := makeBaseReqAndMsg(restRequest.Request)

		baseReq = baseReq.Sanitize()
		if !baseReq.ValidateBasic(responseWriter) {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "")
			return
		}

		Error := msg.ValidateBasic()
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}
		client.WriteGenerateStdTxResponse(responseWriter, cliContext, baseReq, []sdkTypes.Msg{msg})
	}
}
func NewRESTRequest(request interface{}) RESTRequest { return &restRequest{Request: request} }
