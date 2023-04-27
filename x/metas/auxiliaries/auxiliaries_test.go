// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package auxiliaries

import (
	"testing"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/stretchr/testify/require"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("supplement").GetName(), baseHelpers.NewAuxiliaries(
		scrub.Auxiliary,
		supplement.Auxiliary,
	).Get("supplement").GetName())
}
