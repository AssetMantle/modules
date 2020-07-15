package identity

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/identities/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
)

type queryKeeper struct {
	mapper types.Mapper
}

var _ types.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) Query(context sdkTypes.Context, queryRequest types.QueryRequest) types.QueryResponse {
	return newQueryResponse(mapper.NewIdentities(queryKeeper.mapper, context).Fetch(queryRequestFromInterface(queryRequest).IdentityID))
}

func initializeQueryKeeper(mapper types.Mapper) types.QueryKeeper {
	return queryKeeper{mapper: mapper}
}
