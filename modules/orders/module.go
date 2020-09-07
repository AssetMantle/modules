/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package orders

import (
	"github.com/persistenceOne/persistenceSDK/modules/orders/auxiliaries"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/queries"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions"
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
