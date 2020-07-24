package order

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type queryRequest struct {
	OrderID types.ID
}

var _ utilities.QueryRequest = (*queryRequest)(nil)

func (QueryRequest queryRequest) FromCLI(cliCommand utilities.CLICommand, _ context.CLIContext) utilities.QueryRequest {
	return newQueryRequest(base.NewID(cliCommand.ReadString(constants.OrderID)))
}

func (QueryRequest queryRequest) FromMap(vars map[string]string) utilities.QueryRequest {
	return newQueryRequest(base.NewID(vars[constants.OrderID.GetName()]))
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

func newQueryRequest(orderID types.ID) utilities.QueryRequest {
	return queryRequest{OrderID: orderID}
}
