// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package classifications

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries"
	"github.com/AssetMantle/modules/x/classifications/block"
	"github.com/AssetMantle/modules/x/classifications/constants"
	"github.com/AssetMantle/modules/x/classifications/genesis"
	"github.com/AssetMantle/modules/x/classifications/invariants"
	"github.com/AssetMantle/modules/x/classifications/mapper"
	"github.com/AssetMantle/modules/x/classifications/parameters"
	"github.com/AssetMantle/modules/x/classifications/queries"
	"github.com/AssetMantle/modules/x/classifications/simulator"
	"github.com/AssetMantle/modules/x/classifications/transactions"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Name(), baseHelpers.NewModule(
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
	).Name())
}
