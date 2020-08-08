/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package splits

import (
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/burn"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/mint"
	"github.com/persistenceOne/persistenceSDK/modules/splits/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/splits/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/splits/queries/split"
	"github.com/persistenceOne/persistenceSDK/modules/splits/transactions/send"
	"github.com/persistenceOne/persistenceSDK/modules/splits/transactions/unwrap"
	"github.com/persistenceOne/persistenceSDK/modules/splits/transactions/wrap"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Module = base.NewModule(
	mapper.ModuleName,
	mapper.DefaultParamspace,
	mapper.QueryRoute,
	mapper.TransactionRoute,
	genesis.GenesisState,
	mapper.Mapper,
	[]helpers.Auxiliary{mint.Auxiliary, burn.Auxiliary},
	[]helpers.Query{split.Query},
	[]helpers.Transaction{send.Transaction, unwrap.Transaction, wrap.Transaction},
)
