// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package auxiliaries

import (
	"testing"

	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("supplement").GetName(), baseHelpers.NewAuxiliaries(
		scrub.Auxiliary,
		supplement.Auxiliary,
	).Get("supplement").GetName())
}
