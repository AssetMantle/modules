// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package metas

import (
	"testing"

	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/x/classifications/auxiliaries"
	"github.com/AssetMantle/modules/x/metas/internal/block"
	"github.com/AssetMantle/modules/x/metas/internal/genesis"
	"github.com/AssetMantle/modules/x/metas/internal/invariants"
	"github.com/AssetMantle/modules/x/metas/internal/mapper"
	"github.com/AssetMantle/modules/x/metas/internal/module"
	"github.com/AssetMantle/modules/x/metas/internal/parameters"
	"github.com/AssetMantle/modules/x/metas/internal/queries"
	"github.com/AssetMantle/modules/x/metas/internal/simulator"
	"github.com/AssetMantle/modules/x/metas/internal/transactions"
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
