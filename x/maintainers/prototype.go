// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maintainers

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries"
	"github.com/AssetMantle/modules/x/maintainers/internal/block"
	"github.com/AssetMantle/modules/x/maintainers/internal/genesis"
	"github.com/AssetMantle/modules/x/maintainers/internal/invariants"
	"github.com/AssetMantle/modules/x/maintainers/internal/mapper"
	"github.com/AssetMantle/modules/x/maintainers/internal/module"
	"github.com/AssetMantle/modules/x/maintainers/internal/parameters"
	"github.com/AssetMantle/modules/x/maintainers/internal/queries"
	"github.com/AssetMantle/modules/x/maintainers/internal/simulator"
	"github.com/AssetMantle/modules/x/maintainers/internal/transactions"
)

func Prototype() helpers.Module {
	return baseHelpers.NewModule(
		module.Name,
		module.ConsensusVersion,
		auxiliaries.Prototype,
		block.Prototype,
		genesis.Prototype,
		invariants.Prototype,
		mapper.Prototype,
		parameters.Prototype,
		queries.Prototype,
		simulator.Prototype,
		transactions.Prototype,
	)
}
