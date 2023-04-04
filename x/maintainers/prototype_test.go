// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maintainers

import (
	"testing"

	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/x/classifications/auxiliaries"
	"github.com/AssetMantle/modules/x/maintainers/internal/block"
	"github.com/AssetMantle/modules/x/maintainers/internal/genesis"
	"github.com/AssetMantle/modules/x/maintainers/internal/invariants"
	"github.com/AssetMantle/modules/x/maintainers/internal/mapper"
	"github.com/AssetMantle/modules/x/maintainers/internal/module"
	"github.com/AssetMantle/modules/x/maintainers/internal/parameters"
	"github.com/AssetMantle/modules/x/maintainers/internal/queries"
	"github.com/AssetMantle/modules/x/maintainers/internal/simulator"
	"github.com/AssetMantle/modules/x/maintainers/internal/transactions"
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
