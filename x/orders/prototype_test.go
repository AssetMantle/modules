// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package orders

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries"
	"github.com/AssetMantle/modules/x/orders/internal/block"
	"github.com/AssetMantle/modules/x/orders/internal/genesis"
	"github.com/AssetMantle/modules/x/orders/internal/invariants"
	"github.com/AssetMantle/modules/x/orders/internal/mapper"
	"github.com/AssetMantle/modules/x/orders/internal/module"
	"github.com/AssetMantle/modules/x/orders/internal/parameters"
	"github.com/AssetMantle/modules/x/orders/internal/queries"
	"github.com/AssetMantle/modules/x/orders/internal/simulator"
	"github.com/AssetMantle/modules/x/orders/internal/transactions"
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
