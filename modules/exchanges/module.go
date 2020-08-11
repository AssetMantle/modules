/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package exchanges

import (
	"github.com/persistenceOne/persistenceSDK/modules/exchanges/auxiliaries/custody"
	"github.com/persistenceOne/persistenceSDK/modules/exchanges/auxiliaries/reverse"
	"github.com/persistenceOne/persistenceSDK/modules/exchanges/auxiliaries/swap"
	"github.com/persistenceOne/persistenceSDK/modules/exchanges/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/exchanges/mapper"
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
	[]helpers.Auxiliary{swap.Auxiliary, custody.Auxiliary, reverse.Auxiliary},
	[]helpers.Query{},
	[]helpers.Transaction{},
)
