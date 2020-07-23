package split

import (
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type queryResponse struct {
	Splits schema.Splits `json:"splits" valid:"required~Enter the Splits"`
}

var _ utility.QueryResponse = (*queryResponse)(nil)

func queryResponsePrototype() utility.QueryResponse {
	return queryResponse{}
}

func newQueryResponse(splits schema.Splits) utility.QueryResponse {
	return queryResponse{Splits: splits}
}
