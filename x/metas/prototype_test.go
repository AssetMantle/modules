// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package metas

import (
	"testing"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries"
	"github.com/AssetMantle/modules/x/metas/block"
	"github.com/AssetMantle/modules/x/metas/genesis"
	"github.com/AssetMantle/modules/x/metas/invariants"
	"github.com/AssetMantle/modules/x/metas/mapper"
	"github.com/AssetMantle/modules/x/metas/module"
	"github.com/AssetMantle/modules/x/metas/parameters"
	"github.com/AssetMantle/modules/x/metas/queries"
	"github.com/AssetMantle/modules/x/metas/simulator"
	"github.com/AssetMantle/modules/x/metas/transactions"
	"github.com/stretchr/testify/require"
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
