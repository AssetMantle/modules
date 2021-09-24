/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package orders

import (
	"github.com/persistenceOne/persistenceSDK/modules/orders/auxiliaries"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/block"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/queries"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/simulator"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func Prototype() helpers.Module {
	return base.NewModule(
		module.Name,
		auxiliaries.Prototype,
		genesis.Prototype,
		mapper.Prototype,
		parameters.Prototype,
		queries.Prototype,
		simulator.Prototype,
		transactions.Prototype,
		block.Prototype,
	)
}
