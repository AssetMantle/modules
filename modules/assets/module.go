/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package assets

import (
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/auxiliaries"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/queries"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Module = base.NewModule(
	mapper.ModuleName,
	mapper.DefaultParamspace,
	mapper.QueryRoute,
	mapper.TransactionRoute,
	mapper.Mapper,
	genesis.Prototype,
	parameters.Prototype,
	auxiliaries.Auxiliaries,
	queries.Queries,
	transactions.Transactions,
)
