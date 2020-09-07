/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package splits

import (
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/queries"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/transactions"
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
