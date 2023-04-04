// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package classifications

import (
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries"
	"github.com/AssetMantle/modules/modules/classifications/internal/block"
	"github.com/AssetMantle/modules/modules/classifications/internal/genesis"
	"github.com/AssetMantle/modules/modules/classifications/internal/invariants"
	"github.com/AssetMantle/modules/modules/classifications/internal/mapper"
	"github.com/AssetMantle/modules/modules/classifications/internal/module"
	"github.com/AssetMantle/modules/modules/classifications/internal/parameters"
	"github.com/AssetMantle/modules/modules/classifications/internal/queries"
	"github.com/AssetMantle/modules/modules/classifications/internal/simulator"
	"github.com/AssetMantle/modules/modules/classifications/internal/transactions"
	"github.com/AssetMantle/schema/x/helpers"
	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"
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
