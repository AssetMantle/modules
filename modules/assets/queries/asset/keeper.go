package asset

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type queryKeeper struct {
	mapper utility.Mapper `json:"mapper" valid:"required~Enter the Mapper"`
}

var _ utility.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) Enquire(context sdkTypes.Context, queryRequest utility.QueryRequest) utility.QueryResponse {
	return newQueryResponse(mapper.NewAssets(queryKeeper.mapper, context).Fetch(queryRequestFromInterface(queryRequest).AssetID))
}

func initializeQueryKeeper(mapper utility.Mapper) utility.QueryKeeper {
	return queryKeeper{mapper: mapper}
}
