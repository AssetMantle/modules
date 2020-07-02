package asset

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
)

type queryKeeper struct {
	mapper mapper.Mapper
}

var _ types.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) Query(context sdkTypes.Context, queryRequest types.QueryRequest) types.QueryResponse {
	return NewQueryResponse(mapper.NewAssets(queryKeeper.mapper, context).Fetch(queryRequestFromInterface(queryRequest).AssetID))
}

func NewQueryKeeper(mapper mapper.Mapper) types.QueryKeeper {
	return queryKeeper{mapper: mapper}
}
