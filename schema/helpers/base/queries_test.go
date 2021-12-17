/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/schema/helpers"
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
