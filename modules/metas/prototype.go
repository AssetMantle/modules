/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package metas

import (
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/queries"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/simulator"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/transactions"
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
