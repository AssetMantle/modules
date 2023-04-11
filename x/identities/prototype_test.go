// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identities

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries"
	"github.com/AssetMantle/modules/x/identities/internal/block"
	"github.com/AssetMantle/modules/x/identities/internal/genesis"
	"github.com/AssetMantle/modules/x/identities/internal/invariants"
	"github.com/AssetMantle/modules/x/identities/internal/mapper"
	"github.com/AssetMantle/modules/x/identities/internal/module"
	"github.com/AssetMantle/modules/x/identities/internal/parameters"
	"github.com/AssetMantle/modules/x/identities/internal/queries"
	"github.com/AssetMantle/modules/x/identities/internal/simulator"
	"github.com/AssetMantle/modules/x/identities/internal/transactions"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Name(), baseHelpers.NewModule(
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
	).Name())
}
