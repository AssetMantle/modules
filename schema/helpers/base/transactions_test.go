/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/stretchr/testify/require"
	"testing"
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
