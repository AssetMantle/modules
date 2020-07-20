package order

import (
	"github.com/persistenceOne/persistenceSDK/modules/order/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/order/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/order/queries/order"
	"github.com/persistenceOne/persistenceSDK/modules/order/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/order/transactions/mint"
	"github.com/persistenceOne/persistenceSDK/modules/order/transactions/mutate"
	"github.com/persistenceOne/persistenceSDK/types"
)

var Module = types.NewModule(
	mapper.ModuleName,
	mapper.StoreKey,
	mapper.DefaultParamspace,
	mapper.QueryRoute,
	mapper.TransactionRoute,
	genesis.GenesisState,
	mapper.Mapper,
	[]types.Query{order.Query},
	[]types.Transaction{burn.Transaction, mint.Transaction, mutate.Transaction},
)
