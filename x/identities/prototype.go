// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identities

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/identities/auxiliaries"
	"github.com/AssetMantle/modules/x/identities/block"
	"github.com/AssetMantle/modules/x/identities/constants"
	"github.com/AssetMantle/modules/x/identities/genesis"
	"github.com/AssetMantle/modules/x/identities/invariants"
	"github.com/AssetMantle/modules/x/identities/mapper"
	"github.com/AssetMantle/modules/x/identities/parameters"
	"github.com/AssetMantle/modules/x/identities/queries"
	"github.com/AssetMantle/modules/x/identities/simulator"
	"github.com/AssetMantle/modules/x/identities/transactions"
)

func Prototype() helpers.Module {
	return baseHelpers.NewModule(
		constants.ModuleName,
		constants.ModuleConsensusVersion,
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
