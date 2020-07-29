package exchanges

import (
	"github.com/persistenceOne/persistenceSDK/modules/exchanges/auxiliaries/swap"
	"github.com/persistenceOne/persistenceSDK/modules/exchanges/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/exchanges/mapper"
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
	[]helpers.Auxiliary{swap.Auxiliary},
	[]helpers.Query{},
	[]helpers.Transaction{},
)
