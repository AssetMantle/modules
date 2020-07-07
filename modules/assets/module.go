package assets

import (
	"github.com/persistenceOne/persistenceSDK/modules/assets/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/assets/queries/asset"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/mint"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/mutate"
	"github.com/persistenceOne/persistenceSDK/types"
)

var Module = types.NewModule(
	constants.ModuleName,
	constants.StoreKey,
	constants.DefaultParamspace,
	constants.QuerierRoute,
	constants.TransactionRoute,
	genesis.NewGenesisState(),
	[]types.Query{asset.Query},
	[]types.Transaction{burn.Transaction, mint.Transaction, mutate.Transaction},
	registerCodec,
	mapper.NewMapper,
)
