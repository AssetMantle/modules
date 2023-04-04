// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package assets

import (
	"testing"

	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/assets/internal/block"
	"github.com/AssetMantle/modules/modules/assets/internal/genesis"
	"github.com/AssetMantle/modules/modules/assets/internal/invariants"
	"github.com/AssetMantle/modules/modules/assets/internal/mapper"
	"github.com/AssetMantle/modules/modules/assets/internal/module"
	"github.com/AssetMantle/modules/modules/assets/internal/parameters"
	"github.com/AssetMantle/modules/modules/assets/internal/queries"
	"github.com/AssetMantle/modules/modules/assets/internal/simulator"
	"github.com/AssetMantle/modules/modules/assets/internal/transactions"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries"
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
