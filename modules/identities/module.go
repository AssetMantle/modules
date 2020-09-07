/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package identities

import (
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/queries"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Module = base.NewModule(
	mapper.ModuleName,
	mapper.DefaultParamspace,
	mapper.QueryRoute,
	mapper.TransactionRoute,
	mapper.Mapper,
	genesis.Genesis,
	parameters.Parameters,
	auxiliaries.Auxiliaries,
	queries.Queries,
	transactions.Transactions,
)
