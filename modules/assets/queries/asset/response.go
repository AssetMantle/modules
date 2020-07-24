package asset

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
)

type queryResponse struct {
	Assets mappers.InterNFTs `json:"assets" valid:"required~required field assets missing"`
}

var _ helpers.QueryResponse = (*queryResponse)(nil)

func queryResponsePrototype() helpers.QueryResponse {
	return queryResponse{}
}

func newQueryResponse(assets mappers.InterNFTs) helpers.QueryResponse {
	return queryResponse{Assets: assets}
}
