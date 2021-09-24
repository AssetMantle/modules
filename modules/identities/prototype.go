/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package identities

import (
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/block"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/queries"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/simulator"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions"
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
