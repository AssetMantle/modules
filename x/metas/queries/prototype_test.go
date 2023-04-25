// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"testing"

	"github.com/AssetMantle/modules/x/metas/queries/meta"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("metas").GetName(), baseHelpers.NewQueries(
		meta.Query,
	).Get("metas").GetName())
}
