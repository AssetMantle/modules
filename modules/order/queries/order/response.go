package order

import "github.com/persistenceOne/persistenceSDK/types"

type queryResponse struct {
	Orders types.InterNFTs
}

var _ types.QueryResponse = (*queryResponse)(nil)

func queryResponsePrototype() types.QueryResponse {
	return queryResponse{}
}

func newQueryResponse(orders types.InterNFTs) types.QueryResponse {
	return queryResponse{Orders: orders}
}
