package rest

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"net/http"
)

func PostProcessResponse(responseWriter http.ResponseWriter, context client.Context, response interface{}) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(responseWriter).Encode(response)
}

func CheckBadRequestError(responseWriter http.ResponseWriter, err error) bool {
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		_, _ = responseWriter.Write([]byte(err.Error()))
		return true
	}
	return false
}

func ParseQueryHeightOrReturnBadRequest(responseWriter http.ResponseWriter, clientContext client.Context, request *http.Request) (client.Context, bool) {
	// TODO correct
	return clientContext, true
}

func CheckInternalServerError(responseWriter http.ResponseWriter, err error) bool {
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		_, _ = responseWriter.Write([]byte(err.Error()))
		return true
	}
	return false
}

func WriteErrorResponse(responseWriter http.ResponseWriter, statusCode int, message string) {
	responseWriter.WriteHeader(statusCode)
	_, _ = responseWriter.Write([]byte(message))
}
