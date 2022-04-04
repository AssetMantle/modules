// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/helpers"
)

func TestQueries(t *testing.T) {
	Queries1 := NewQueries(query{})
	Queries2 := NewQueries()

	// Get
	require.Equal(t, "", Queries1.Get("").GetName())
	require.Equal(t, nil, Queries2.Get(""))

	// GetList
	require.Equal(t, []helpers.Query{query{}}, Queries1.GetList())
	require.Equal(t, []helpers.Query(nil), Queries2.GetList())

}
