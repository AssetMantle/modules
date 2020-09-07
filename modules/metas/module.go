/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package metas

import (
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/queries"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/transactions"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Module = base.NewModule(
	mapper.ModuleName,
	mapper.DefaultParamspace,
	mapper.QueryRoute,
	mapper.TransactionRoute,
	mapper.Mapper,
	genesis.State,
	parameters.Parameters,
	auxiliaries.Auxiliaries,
	queries.Queries,
	transactions.Transactions,
)
