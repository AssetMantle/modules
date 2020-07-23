package identity

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type queryRequest struct {
	IdentityID types.ID `json:"identityId" valid:"required~Enter the IdentityID"`
}

var _ utilities.QueryRequest = (*queryRequest)(nil)

func (QueryRequest queryRequest) FromCLI(cliCommand utilities.CLICommand, _ context.CLIContext) utilities.QueryRequest {
	return newQueryRequest(base.NewID(cliCommand.ReadString(constants.IdentityID)))
}

func (QueryRequest queryRequest) FromMap(vars map[string]string) utilities.QueryRequest {
	return newQueryRequest(base.NewID(vars[constants.IdentityID.GetName()]))
}

func queryRequestPrototype() utilities.QueryRequest {
	return queryRequest{}
}

func queryRequestFromInterface(QueryRequest utilities.QueryRequest) queryRequest {
	switch value := QueryRequest.(type) {
	case queryRequest:
		return value
	default:
		return queryRequest{}
	}
}

func newQueryRequest(identityID types.ID) utilities.QueryRequest {
	return queryRequest{IdentityID: identityID}
}
