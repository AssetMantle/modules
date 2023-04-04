// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package splits

import (
	"testing"

	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/x/classifications/auxiliaries"
	"github.com/AssetMantle/modules/x/splits/internal/block"
	"github.com/AssetMantle/modules/x/splits/internal/genesis"
	"github.com/AssetMantle/modules/x/splits/internal/invariants"
	"github.com/AssetMantle/modules/x/splits/internal/mapper"
	"github.com/AssetMantle/modules/x/splits/internal/module"
	"github.com/AssetMantle/modules/x/splits/internal/parameters"
	"github.com/AssetMantle/modules/x/splits/internal/queries"
	"github.com/AssetMantle/modules/x/splits/internal/simulator"
	"github.com/AssetMantle/modules/x/splits/internal/transactions"
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
