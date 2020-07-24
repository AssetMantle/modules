package asset

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type queryResponse struct {
	Assets mappers.InterNFTs `json:"assets" valid:"required~required field assets missing"`
}

var _ utilities.QueryResponse = (*queryResponse)(nil)

func queryResponsePrototype() utilities.QueryResponse {
	return queryResponse{}
}

func newQueryResponse(assets mappers.InterNFTs) utilities.QueryResponse {
	return queryResponse{Assets: assets}
}
