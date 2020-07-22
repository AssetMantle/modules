package order

import (
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type queryResponse struct {
	Orders schema.Orders
}

var _ utility.QueryResponse = (*queryResponse)(nil)

func queryResponsePrototype() utility.QueryResponse {
	return queryResponse{}
}

func newQueryResponse(orders schema.Orders) utility.QueryResponse {
	return queryResponse{Orders: orders}
}
