package asset

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type queryRequest struct {
	AssetID types.ID `json:"assetID" valid:"required field assetID missing"`
}

var _ helpers.QueryRequest = (*queryRequest)(nil)

func (QueryRequest queryRequest) FromCLI(cliCommand helpers.CLICommand, _ context.CLIContext) helpers.QueryRequest {
	return newQueryRequest(base.NewID(cliCommand.ReadString(constants.AssetID)))
}

func (QueryRequest queryRequest) FromMap(vars map[string]string) helpers.QueryRequest {
	return newQueryRequest(base.NewID(vars[constants.AssetID.GetName()]))
}

func queryRequestPrototype() helpers.QueryRequest {
	return queryRequest{}
}

func queryRequestFromInterface(QueryRequest helpers.QueryRequest) queryRequest {
	switch value := QueryRequest.(type) {
	case queryRequest:
		return value
	default:
		return queryRequest{}
	}
}

func newQueryRequest(assetID types.ID) helpers.QueryRequest {
	return queryRequest{AssetID: assetID}
}
