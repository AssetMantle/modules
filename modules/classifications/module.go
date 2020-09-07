/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package classifications

import (
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/queries"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/transactions"

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
