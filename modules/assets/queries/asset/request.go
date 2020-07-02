package asset

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/persistenceOne/persistenceSDK/modules/assets/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
)

type queryRequest struct {
	AssetID types.ID
}

var _ types.QueryRequest = (*queryRequest)(nil)

func (QueryRequest queryRequest) FromCLI(cliCommand types.CLICommand, cliContext context.CLIContext) types.QueryRequest {
	return NewQueryRequest(mapper.ReadAssetID(cliCommand.ReadString(constants.AssetID)))
}

func (QueryRequest queryRequest) FromMap(vars map[string]string) types.QueryRequest {
	return NewQueryRequest(mapper.ReadAssetID(vars[constants.AssetID.GetName()]))
}

func queryRequestPrototype() types.QueryRequest {
	return &queryRequest{}
}

func queryRequestFromInterface(QueryRequest types.QueryRequest) queryRequest {
	switch value := QueryRequest.(type) {
	case queryRequest:
		return value
	default:
		return queryRequest{}
	}
}

func NewQueryRequest(assetID types.ID) types.QueryRequest {
	return &queryRequest{AssetID: assetID}
}
