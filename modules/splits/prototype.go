/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package splits

import (
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/block"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/queries"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/simulator"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/transactions"
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
