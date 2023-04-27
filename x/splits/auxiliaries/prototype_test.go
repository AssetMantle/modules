// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package auxiliaries

import (
	"testing"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/burn"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/mint"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/renumerate"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
	"github.com/stretchr/testify/require"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("transfer").GetName(), baseHelpers.NewAuxiliaries(
		burn.Auxiliary,
		mint.Auxiliary,
		renumerate.Auxiliary,
		transfer.Auxiliary,
	).Get("transfer").GetName())
}
