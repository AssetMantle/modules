// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/helpers"
)

func TestNewTransactions(t *testing.T) {
	Transactions1 := NewTransactions()
	Transactions2 := NewTransactions(transaction{})

	// GetName
	require.Equal(t, nil, Transactions1.Get(""))
	require.Equal(t, "", Transactions2.Get("").GetName())

	// GetList
	require.Equal(t, []helpers.Transaction(nil), Transactions1.GetList())
	require.Equal(t, []helpers.Transaction{transaction{}}, Transactions2.GetList())
}
