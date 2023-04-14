// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package classifications

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries"
	"github.com/AssetMantle/modules/x/classifications/internal/block"
	"github.com/AssetMantle/modules/x/classifications/internal/genesis"
	"github.com/AssetMantle/modules/x/classifications/internal/invariants"
	"github.com/AssetMantle/modules/x/classifications/internal/mapper"
	"github.com/AssetMantle/modules/x/classifications/internal/module"
	"github.com/AssetMantle/modules/x/classifications/internal/parameters"
	"github.com/AssetMantle/modules/x/classifications/internal/queries"
	"github.com/AssetMantle/modules/x/classifications/internal/simulator"
	"github.com/AssetMantle/modules/x/classifications/internal/transactions"
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
