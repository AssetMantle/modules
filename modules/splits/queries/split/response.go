package split

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type queryResponse struct {
	Splits mappers.Splits `json:"splits" valid:"required~required field splits missing"`
}

var _ utilities.QueryResponse = (*queryResponse)(nil)

func queryResponsePrototype() utilities.QueryResponse {
	return queryResponse{}
}

func newQueryResponse(splits mappers.Splits) utilities.QueryResponse {
	return queryResponse{Splits: splits}
}
