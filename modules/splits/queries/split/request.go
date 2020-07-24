package split

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type queryRequest struct {
	SplitID types.ID `json:"splitid" valid:"required~required field splitid missing"`
}

var _ utilities.QueryRequest = (*queryRequest)(nil)

func (QueryRequest queryRequest) FromCLI(cliCommand utilities.CLICommand, _ context.CLIContext) utilities.QueryRequest {
	return newQueryRequest(base.NewID(cliCommand.ReadString(constants.SplitID)))
}

func (QueryRequest queryRequest) FromMap(vars map[string]string) utilities.QueryRequest {
	return newQueryRequest(base.NewID(vars[constants.SplitID.GetName()]))
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

func newQueryRequest(splitID types.ID) utilities.QueryRequest {
	return queryRequest{SplitID: splitID}
}
