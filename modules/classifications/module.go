package classifications

import (
	"github.com/persistenceOne/persistenceSDK/modules/classifications/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/queries/classification"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/transactions/create"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/transactions/delete"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/transactions/update"
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
	[]helpers.Auxiliary{},
	[]helpers.Query{classification.Query},
	[]helpers.Transaction{create.Transaction, update.Transaction, delete.Transaction},
)
