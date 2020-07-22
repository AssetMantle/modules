package identities

import (
	"github.com/persistenceOne/persistenceSDK/modules/identities/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/identities/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/identities/queries/identity"
	"github.com/persistenceOne/persistenceSDK/modules/identities/transactions/issue"
	"github.com/persistenceOne/persistenceSDK/modules/identities/transactions/provision"
	"github.com/persistenceOne/persistenceSDK/modules/identities/transactions/unprovision"
	"github.com/persistenceOne/persistenceSDK/types/utility"
	"github.com/persistenceOne/persistenceSDK/types/utility/base"
)

var Module = base.NewModule(
	mapper.ModuleName,
	mapper.StoreKey,
	mapper.DefaultParamspace,
	mapper.QueryRoute,
	mapper.TransactionRoute,
	genesis.GenesisState,
	mapper.Mapper,
	[]utility.Auxiliary{},
	[]utility.Query{identity.Query},
	[]utility.Transaction{issue.Transaction, provision.Transaction, unprovision.Transaction},
)
