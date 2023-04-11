// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package auxiliaries

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/define"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, baseHelpers.NewAuxiliaries(
		conform.Auxiliary,
		define.Auxiliary,
	).Get(""), Prototype().Get(""))
}
