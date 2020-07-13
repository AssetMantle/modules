package identities

import (
	"github.com/persistenceOne/persistenceSDK/modules/identities/constants"
	"github.com/persistenceOne/persistenceSDK/modules/identities/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/identities/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/identities/queries/identity"
	"github.com/persistenceOne/persistenceSDK/modules/identities/transactions/issue"
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
	[]types.Query{identity.Query},
	[]types.Transaction{issue.Transaction},
)
