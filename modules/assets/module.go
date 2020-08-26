/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package assets

import (
	"github.com/persistenceOne/persistenceSDK/modules/assets/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/assets/queries/asset"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/define"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/mint"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/mutate"
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
	[]helpers.Auxiliary{},
	[]helpers.Query{asset.Query},
	[]helpers.Transaction{burn.Transaction, define.Transaction, mint.Transaction, mutate.Transaction},
)
