package split

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type queryRequest struct {
	SplitID schema.ID `json:"splitId" valid:"required~Enter the SplitID"`
}

var _ utility.QueryRequest = (*queryRequest)(nil)

func (QueryRequest queryRequest) FromCLI(cliCommand utility.CLICommand, _ context.CLIContext) utility.QueryRequest {
	return newQueryRequest(schema.NewID(cliCommand.ReadString(constants.SplitID)))
}

func (QueryRequest queryRequest) FromMap(vars map[string]string) utility.QueryRequest {
	return newQueryRequest(schema.NewID(vars[constants.SplitID.GetName()]))
}

func queryRequestPrototype() utility.QueryRequest {
	return queryRequest{}
}

func queryRequestFromInterface(QueryRequest utility.QueryRequest) queryRequest {
	switch value := QueryRequest.(type) {
	case queryRequest:
		return value
	default:
		return queryRequest{}
	}
}

func newQueryRequest(splitID schema.ID) utility.QueryRequest {
	return queryRequest{SplitID: splitID}
}
