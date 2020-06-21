package types

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client"
	"net/http"
)

type RESTRequest interface {
	RequestHandler(func() Request) func(context.CLIContext) http.HandlerFunc
}

type restRequest struct {
}

var _ RESTRequest = (*restRequest)(nil)

func (restRequest restRequest) RequestHandler(requestPrototype func() Request) func(cliContext context.CLIContext) http.HandlerFunc {
	return func(cliContext context.CLIContext) http.HandlerFunc {
		return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
			request := requestPrototype()
			if !rest.ReadRESTReq(responseWriter, httpRequest, cliContext.Codec, &request) {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "")
				return
			}

			baseReq := request.GetBaseReq()
			msg := request.MakeMsg()

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
}
func NewRESTRequest() RESTRequest {
	return &restRequest{}
}
