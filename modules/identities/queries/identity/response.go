package identity

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type queryResponse struct {
	Identities mappers.InterIdentities `json:"identities" valid:"required~required field identities missing"`
}

var _ utilities.QueryResponse = (*queryResponse)(nil)

func queryResponsePrototype() utilities.QueryResponse {
	return queryResponse{}
}

func newQueryResponse(identities mappers.InterIdentities) utilities.QueryResponse {
	return queryResponse{Identities: identities}
}
