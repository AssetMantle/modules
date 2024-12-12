// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identities

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries"
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
