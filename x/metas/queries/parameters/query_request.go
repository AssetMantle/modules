package parameters

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
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

func (queryRequest *QueryRequest) Encode() ([]byte, error) {
	return base.CodecPrototype().MarshalJSON(queryRequest)
}

func (queryRequest *QueryRequest) Decode(bytes []byte) (helpers.QueryRequest, error) {
	if err := base.CodecPrototype().UnmarshalJSON(bytes, queryRequest); err != nil {
		return nil, err
	}

	return queryRequest, nil
}
func requestPrototype() helpers.QueryRequest {
	return &QueryRequest{}
}
func queryRequestFromInterface(request helpers.QueryRequest) *QueryRequest {
	switch value := request.(type) {
	case *QueryRequest:
		return value
	default:
		return &QueryRequest{}
	}
}
