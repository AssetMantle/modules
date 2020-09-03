/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package orders

import (
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/queries/order"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/cancel"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/define"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/make"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/take"
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
	[]helpers.Query{order.Query},
	[]helpers.Transaction{define.Transaction, cancel.Transaction, make.Transaction, take.Transaction},
)
