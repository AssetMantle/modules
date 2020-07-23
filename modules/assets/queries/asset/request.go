package asset

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type queryRequest struct {
	AssetID schema.ID `json:"assetID" valid:"required~Enter the AssetID"`
}

var _ utility.QueryRequest = (*queryRequest)(nil)

func (QueryRequest queryRequest) FromCLI(cliCommand utility.CLICommand, _ context.CLIContext) utility.QueryRequest {
	return newQueryRequest(schema.NewID(cliCommand.ReadString(constants.AssetID)))
}

func (QueryRequest queryRequest) FromMap(vars map[string]string) utility.QueryRequest {
	return newQueryRequest(schema.NewID(vars[constants.AssetID.GetName()]))
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

func newQueryRequest(assetID schema.ID) utility.QueryRequest {
	return queryRequest{AssetID: assetID}
}
