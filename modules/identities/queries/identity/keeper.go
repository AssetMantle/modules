package identity

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/identities/mapper"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type queryKeeper struct {
	mapper utility.Mapper
}

var _ utility.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) Enquire(context sdkTypes.Context, queryRequest utility.QueryRequest) utility.QueryResponse {
	return newQueryResponse(mapper.NewIdentities(queryKeeper.mapper, context).Fetch(queryRequestFromInterface(queryRequest).IdentityID))
}

func initializeQueryKeeper(mapper utility.Mapper) utility.QueryKeeper {
	return queryKeeper{mapper: mapper}
}
