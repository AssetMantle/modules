/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package splits

import (
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/burn"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/mint"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/transfer"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/queries/split"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/transactions/send"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/transactions/unwrap"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/transactions/wrap"
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
	[]helpers.Auxiliary{mint.Auxiliary, burn.Auxiliary, transfer.Auxiliary},
	[]helpers.Query{split.Query},
	[]helpers.Transaction{send.Transaction, unwrap.Transaction, wrap.Transaction},
)
