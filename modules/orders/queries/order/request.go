package order

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type queryRequest struct {
	OrderID schema.ID
}

var _ utility.QueryRequest = (*queryRequest)(nil)

func (QueryRequest queryRequest) FromCLI(cliCommand utility.CLICommand, _ context.CLIContext) utility.QueryRequest {
	return newQueryRequest(schema.NewID(cliCommand.ReadString(constants.OrderID)))
}

func (QueryRequest queryRequest) FromMap(vars map[string]string) utility.QueryRequest {
	return newQueryRequest(schema.NewID(vars[constants.OrderID.GetName()]))
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

func newQueryRequest(orderID schema.ID) utility.QueryRequest {
	return queryRequest{OrderID: orderID}
}
