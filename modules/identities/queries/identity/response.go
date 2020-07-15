package identity

import "github.com/persistenceOne/persistenceSDK/types"

type queryResponse struct {
	Identities types.InterIdentities
}

var _ types.QueryResponse = (*queryResponse)(nil)

func queryResponsePrototype() types.QueryResponse {
	return queryResponse{}
}

func newQueryResponse(identities types.InterIdentities) types.QueryResponse {
	return queryResponse{Identities: identities}
}
