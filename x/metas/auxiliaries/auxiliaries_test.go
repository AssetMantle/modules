// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package auxiliaries

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().GetAuxiliary("supplement").GetName(), baseHelpers.NewAuxiliaries(
		scrub.Auxiliary,
		supplement.Auxiliary,
	).GetAuxiliary("supplement").GetName())
}
