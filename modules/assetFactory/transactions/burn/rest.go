package burn

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/persistenceOne/persistenceSDK/types"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

type request struct {
	baseReq rest.BaseReq
	assetID string
}

func RestRequestHandler(cliContext context.CLIContext) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		var request request
		if !rest.ReadRESTReq(responseWriter, httpRequest, cliContext.Codec, &request) {
			return
		}

		request.baseReq = request.baseReq.Sanitize()
		if !request.baseReq.ValidateBasic(responseWriter) {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "")
			return
		}

		_, Error := govalidator.ValidateStruct(request)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}

		from, Error := sdkTypes.AccAddressFromBech32(request.baseReq.From)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}

		message := Message{
			from:    from,
			assetID: types.BaseID{IDString: request.assetID},
		}
		client.WriteGenerateStdTxResponse(responseWriter, cliContext, request.baseReq, []sdkTypes.Msg{message})
	}
}
