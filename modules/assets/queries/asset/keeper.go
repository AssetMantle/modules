package asset

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type queryKeeper struct {
	mapper utilities.Mapper `json:"mapper" valid:"required~required field mapper missing"`
}

var _ utilities.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) Enquire(context sdkTypes.Context, queryRequest utilities.QueryRequest) utilities.QueryResponse {
	return newQueryResponse(mapper.NewAssets(queryKeeper.mapper, context).Fetch(queryRequestFromInterface(queryRequest).AssetID))
}

func initializeQueryKeeper(mapper utilities.Mapper) utilities.QueryKeeper {
	return queryKeeper{mapper: mapper}
}
