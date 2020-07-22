package orders

import (
	"github.com/persistenceOne/persistenceSDK/modules/orders/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/orders/queries/order"
	"github.com/persistenceOne/persistenceSDK/modules/orders/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/orders/transactions/mint"
	"github.com/persistenceOne/persistenceSDK/modules/orders/transactions/mutate"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

var Module = utility.NewModule(
	mapper.ModuleName,
	mapper.StoreKey,
	mapper.DefaultParamspace,
	mapper.QueryRoute,
	mapper.TransactionRoute,
	genesis.GenesisState,
	mapper.Mapper,
	[]utility.Auxiliary{},
	[]utility.Query{order.Query},
	[]utility.Transaction{burn.Transaction, mint.Transaction, mutate.Transaction},
)
