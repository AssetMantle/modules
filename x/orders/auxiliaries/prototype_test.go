// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package auxiliaries

import (
	"testing"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/stretchr/testify/require"
)

func Test_Auxiliary_Prototype(t *testing.T) {
	require.Equal(t, baseHelpers.NewAuxiliaries(), Prototype())
}
