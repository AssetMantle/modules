package clientUtils

import (
	"github.com/cosmos/cosmos-sdk/client"
	"net/http"
)

func ParseQueryHeightOrReturnBadRequest(responseWriter http.ResponseWriter, clientContext client.Context, request *http.Request) (client.Context, bool) {
	// TODO correct
	return clientContext, true
}
