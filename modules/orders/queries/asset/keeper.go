package asset

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types"
)

type queryKeeper struct {
	mapper types.Mapper
}

var _ types.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) Query(context sdkTypes.Context, queryRequest types.QueryRequest) types.QueryResponse {
	return nil //newQueryResponse(mapper.NewAssets(queryKeeper.mapper, context).Fetch(queryRequestFromInterface(queryRequest).AssetID))
}

func initializeQueryKeeper(mapper types.Mapper) types.QueryKeeper {
	return queryKeeper{mapper: mapper}
}
