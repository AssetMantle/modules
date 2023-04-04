// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package auxiliaries

import (
	"testing"

	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/define"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, baseHelpers.NewAuxiliaries(
		conform.Auxiliary,
		define.Auxiliary,
	).Get(""), Prototype().Get(""))
}
