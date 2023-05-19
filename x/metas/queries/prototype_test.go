// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/metas/queries/meta"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().GetQuery("metas").GetName(), baseHelpers.NewQueries(
		meta.Query,
	).GetQuery("metas").GetName())
}
