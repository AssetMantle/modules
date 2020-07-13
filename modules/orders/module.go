package orders

import (
	"github.com/persistenceOne/persistenceSDK/modules/orders/constants"
	"github.com/persistenceOne/persistenceSDK/modules/orders/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/orders/queries/asset"
	"github.com/persistenceOne/persistenceSDK/modules/orders/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/orders/transactions/mint"
	"github.com/persistenceOne/persistenceSDK/modules/orders/transactions/mutate"
	"github.com/persistenceOne/persistenceSDK/types"
)

var Module = types.NewModule(
	constants.ModuleName,
	constants.StoreKey,
	constants.DefaultParamspace,
	constants.QuerierRoute,
	constants.TransactionRoute,
	genesis.GenesisState,
	mapper.Mapper,
	[]types.Query{asset.Query},
	[]types.Transaction{burn.Transaction, mint.Transaction, mutate.Transaction},
)
