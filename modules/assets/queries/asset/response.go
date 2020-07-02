package asset

import "github.com/persistenceOne/persistenceSDK/types"

type queryResponse struct {
	Assets types.InterNFTs
}

var _ types.QueryResponse = (*queryResponse)(nil)

func queryResponsePrototype() types.QueryResponse {
	return queryResponse{}
}

func NewQueryResponse(assets types.InterNFTs) types.QueryResponse {
	return queryResponse{Assets: assets}
}
