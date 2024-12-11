// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maintainers

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries"
	"github.com/AssetMantle/modules/x/maintainers/block"
	"github.com/AssetMantle/modules/x/maintainers/constants"
	"github.com/AssetMantle/modules/x/maintainers/genesis"
	"github.com/AssetMantle/modules/x/maintainers/invariants"
	"github.com/AssetMantle/modules/x/maintainers/mapper"
	"github.com/AssetMantle/modules/x/maintainers/parameters"
	"github.com/AssetMantle/modules/x/maintainers/queries"
	"github.com/AssetMantle/modules/x/maintainers/simulator"
	"github.com/AssetMantle/modules/x/maintainers/transactions"
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
		migrations.Prototype,
		parameters.Prototype,
		queries.Prototype,
		simulator.Prototype,
		transactions.Prototype,
	).Name())
}
