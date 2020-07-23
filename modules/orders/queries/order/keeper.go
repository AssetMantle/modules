package order

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type queryKeeper struct {
	mapper utilities.Mapper
}

var _ utilities.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) Enquire(context sdkTypes.Context, queryRequest utilities.QueryRequest) utilities.QueryResponse {
	return newQueryResponse(mapper.NewOrders(queryKeeper.mapper, context).Fetch(queryRequestFromInterface(queryRequest).OrderID))
}

func initializeQueryKeeper(mapper utilities.Mapper) utilities.QueryKeeper {
	return queryKeeper{mapper: mapper}
}
