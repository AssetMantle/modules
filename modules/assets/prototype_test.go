// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package assets

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/assets/internal/block"
	"github.com/AssetMantle/modules/modules/assets/internal/genesis"
	"github.com/AssetMantle/modules/modules/assets/internal/mapper"
	"github.com/AssetMantle/modules/modules/assets/internal/module"
	"github.com/AssetMantle/modules/modules/assets/internal/parameters"
	"github.com/AssetMantle/modules/modules/assets/internal/queries"
	"github.com/AssetMantle/modules/modules/assets/internal/simulator"
	"github.com/AssetMantle/modules/modules/assets/internal/transactions"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries"
	"github.com/AssetMantle/modules/schema/helpers/base"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Name(), base.NewModule(
		module.Name,
		auxiliaries.Prototype,
		genesis.Prototype,
		mapper.Prototype,
		parameters.Prototype,
		queries.Prototype,
		simulator.Prototype,
		transactions.Prototype,
		block.Prototype,
	).Name())

}
