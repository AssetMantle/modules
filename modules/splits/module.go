package splits

import (
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/mint"
	"github.com/persistenceOne/persistenceSDK/modules/splits/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/splits/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/splits/queries/split"
	"github.com/persistenceOne/persistenceSDK/modules/splits/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/splits/transactions/send"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
	"github.com/persistenceOne/persistenceSDK/schema/utilities/base"
)

var Module = base.NewModule(
	mapper.ModuleName,
	mapper.StoreKey,
	mapper.DefaultParamspace,
	mapper.QueryRoute,
	mapper.TransactionRoute,
	genesis.GenesisState,
	mapper.Mapper,
	[]utilities.Auxiliary{mint.Auxiliary},
	[]utilities.Query{split.Query},
	[]utilities.Transaction{burn.Transaction, send.Transaction},
)
