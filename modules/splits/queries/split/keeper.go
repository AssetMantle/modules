package split

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/splits/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type queryKeeper struct {
	mapper utilities.Mapper `json:"mapper" valid:"required~required field mapper missing"`
}

var _ utilities.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) Enquire(context sdkTypes.Context, queryRequest utilities.QueryRequest) utilities.QueryResponse {
	return newQueryResponse(mapper.NewSplits(queryKeeper.mapper, context).Fetch(queryRequestFromInterface(queryRequest).SplitID))
}

func initializeQueryKeeper(mapper utilities.Mapper) utilities.QueryKeeper {
	return queryKeeper{mapper: mapper}
}
