/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package metas

import (
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/block"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/mapper"
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
