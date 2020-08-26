/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package maintainers

import (
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/maintain"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/super"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/queries/maintainer"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/transactions/deputize"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Module = base.NewModule(
	mapper.ModuleName,
	mapper.DefaultParamspace,
	mapper.QueryRoute,
	mapper.TransactionRoute,
	genesis.State,
	mapper.Mapper,
	[]helpers.Auxiliary{maintain.Auxiliary, super.Auxiliary},
	[]helpers.Query{maintainer.Query},
	[]helpers.Transaction{deputize.Transaction},
)
