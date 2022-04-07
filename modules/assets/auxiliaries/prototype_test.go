// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package auxiliaries

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

func Test_Auxiliaries_Prototype(t *testing.T) {
	require.Equal(t, baseHelpers.NewAuxiliaries(), Prototype())
}
