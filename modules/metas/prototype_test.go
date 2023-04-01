// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package metas

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/classifications/auxiliaries"
	"github.com/AssetMantle/modules/modules/metas/internal/block"
	"github.com/AssetMantle/modules/modules/metas/internal/genesis"
	"github.com/AssetMantle/modules/modules/metas/internal/invariants"
	"github.com/AssetMantle/modules/modules/metas/internal/mapper"
	"github.com/AssetMantle/modules/modules/metas/internal/module"
	"github.com/AssetMantle/modules/modules/metas/internal/parameters"
	"github.com/AssetMantle/modules/modules/metas/internal/queries"
	"github.com/AssetMantle/modules/modules/metas/internal/simulator"
	"github.com/AssetMantle/modules/modules/metas/internal/transactions"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
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
