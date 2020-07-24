package identity

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
)

type queryResponse struct {
	Identities mappers.InterIdentities `json:"identities" valid:"required~required field identities missing"`
}

var _ helpers.QueryResponse = (*queryResponse)(nil)

func queryResponsePrototype() helpers.QueryResponse {
	return queryResponse{}
}

func newQueryResponse(identities mappers.InterIdentities) helpers.QueryResponse {
	return queryResponse{Identities: identities}
}
