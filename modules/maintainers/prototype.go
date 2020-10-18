/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package maintainers

import (
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/queries"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/simulator"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/transactions"
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
