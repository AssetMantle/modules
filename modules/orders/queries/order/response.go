package order

import (
	"github.com/persistenceOne/persistenceSDK/types"
)

type queryResponse struct {
	Orders types.Orders
}

var _ types.QueryResponse = (*queryResponse)(nil)

func queryResponsePrototype() types.QueryResponse {
	return queryResponse{}
}

func newQueryResponse(orders types.Orders) types.QueryResponse {
	return queryResponse{Orders: orders}
}
