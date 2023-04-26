// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package assets

import (
	"testing"

	"github.com/AssetMantle/modules/x/assets/block"
	"github.com/AssetMantle/modules/x/assets/genesis"
	"github.com/AssetMantle/modules/x/assets/invariants"
	"github.com/AssetMantle/modules/x/assets/mapper"
	"github.com/AssetMantle/modules/x/assets/module"
	"github.com/AssetMantle/modules/x/assets/parameters"
	"github.com/AssetMantle/modules/x/assets/queries"
	"github.com/AssetMantle/modules/x/assets/simulator"
	"github.com/AssetMantle/modules/x/assets/transactions"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries"
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
