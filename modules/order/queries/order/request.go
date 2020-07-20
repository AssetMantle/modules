package order

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

type queryRequest struct {
	OrderID types.ID
}

var _ types.QueryRequest = (*queryRequest)(nil)

func (QueryRequest queryRequest) FromCLI(cliCommand types.CLICommand, _ context.CLIContext) types.QueryRequest {
	return newQueryRequest(types.NewID(cliCommand.ReadString(constants.OrderID)))
}

func (QueryRequest queryRequest) FromMap(vars map[string]string) types.QueryRequest {
	return newQueryRequest(types.NewID(vars[constants.OrderID.GetName()]))
}

func queryRequestPrototype() types.QueryRequest {
	return queryRequest{}
}

func queryRequestFromInterface(QueryRequest types.QueryRequest) queryRequest {
	switch value := QueryRequest.(type) {
	case queryRequest:
		return value
	default:
		return queryRequest{}
	}
}

func newQueryRequest(orderID types.ID) types.QueryRequest {
	return queryRequest{OrderID: orderID}
}
