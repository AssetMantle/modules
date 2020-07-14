package identities

import (
	"github.com/persistenceOne/persistenceSDK/modules/identities/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/identities/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/identities/queries/identity"
	"github.com/persistenceOne/persistenceSDK/modules/identities/transactions/issue"
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
	[]types.Query{identity.Query},
	[]types.Transaction{issue.Transaction},
)
