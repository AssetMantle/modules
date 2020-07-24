package asset

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type queryRequest struct {
	AssetID types.ID `json:"assetid" valid:"required field assetid missing"`
}

var _ utilities.QueryRequest = (*queryRequest)(nil)

func (QueryRequest queryRequest) FromCLI(cliCommand utilities.CLICommand, _ context.CLIContext) utilities.QueryRequest {
	return newQueryRequest(base.NewID(cliCommand.ReadString(constants.AssetID)))
}

func (QueryRequest queryRequest) FromMap(vars map[string]string) utilities.QueryRequest {
	return newQueryRequest(base.NewID(vars[constants.AssetID.GetName()]))
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

func newQueryRequest(assetID types.ID) utilities.QueryRequest {
	return queryRequest{AssetID: assetID}
}
