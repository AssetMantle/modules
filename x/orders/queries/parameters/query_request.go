package parameters

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/cosmos/cosmos-sdk/client"
	"net/http"
)

var _ helpers.QueryRequest = (*QueryRequest)(nil)

func (queryRequest *QueryRequest) Validate() error {
	return nil
}

func (queryRequest *QueryRequest) FromCLI(_ helpers.CLICommand, _ client.Context) (helpers.QueryRequest, error) {
	return requestPrototype(), nil
}

func (queryRequest *QueryRequest) FromHTTPRequest(_ *http.Request) (helpers.QueryRequest, error) {
	return requestPrototype(), nil
}

func requestPrototype() helpers.QueryRequest {
	return &QueryRequest{}
}
