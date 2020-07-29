package splits

import (
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/mint"
	"github.com/persistenceOne/persistenceSDK/modules/splits/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/splits/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/splits/queries/split"
	"github.com/persistenceOne/persistenceSDK/modules/splits/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/splits/transactions/send"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Module = base.NewModule(
	mapper.ModuleName,
	mapper.DefaultParamspace,
	mapper.QueryRoute,
	mapper.TransactionRoute,
	genesis.GenesisState,
	mapper.Mapper,
	[]helpers.Auxiliary{mint.Auxiliary},
	[]helpers.Query{split.Query},
	[]helpers.Transaction{burn.Transaction, send.Transaction},
)
