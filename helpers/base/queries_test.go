// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
)

func TestQueries(t *testing.T) {
	Queries1 := NewQueries(query{})
	Queries2 := NewQueries()

	// GetProperty
	require.Equal(t, "", Queries1.GetQuery("").GetServicePath())
	require.Equal(t, nil, Queries2.GetQuery(""))

	// GetAuxiliary
	require.Equal(t, []helpers.Query{query{}}, Queries1.Get())
	require.Equal(t, []helpers.Query(nil), Queries2.Get())

}
