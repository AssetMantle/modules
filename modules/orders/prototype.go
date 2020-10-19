/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package orders

import (
	"github.com/persistenceOne/persistenceSDK/modules/orders/auxiliaries"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/genesis"
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
		simulator.Prototype,
		parameters.Prototype,
		genesis.Prototype,
		auxiliaries.Prototype,
		queries.Prototype,
		transactions.Prototype,
	)
}
