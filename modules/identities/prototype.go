// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identities

import (
	"github.com/AssetMantle/modules/modules/identities/auxiliaries"
	"github.com/AssetMantle/modules/modules/identities/internal/block"
	"github.com/AssetMantle/modules/modules/identities/internal/genesis"
	"github.com/AssetMantle/modules/modules/identities/internal/mapper"
	"github.com/AssetMantle/modules/modules/identities/internal/module"
	"github.com/AssetMantle/modules/modules/identities/internal/parameters"
	"github.com/AssetMantle/modules/modules/identities/internal/queries"
	"github.com/AssetMantle/modules/modules/identities/internal/simulator"
	"github.com/AssetMantle/modules/modules/identities/internal/transactions"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

func Prototype() helpers.Module {
	return baseHelpers.NewModule(
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
