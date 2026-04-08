// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package testutil

import (
	"testing"

	storeTypes "cosmossdk.io/store/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTestContext(t *testing.T) {
	key := storeTypes.NewKVStoreKey("test")
	ctx := NewTestContext(t, key)

	assert.Equal(t, ChainID, ctx.ChainID())

	// verify store is accessible and functional
	store := ctx.KVStore(key)
	require.NotNil(t, store)
	store.Set([]byte("hello"), []byte("world"))
	assert.Equal(t, []byte("world"), store.Get([]byte("hello")))
}

func TestNewTestContext_multipleKeys(t *testing.T) {
	key1 := storeTypes.NewKVStoreKey("store1")
	key2 := storeTypes.NewKVStoreKey("store2")
	ctx := NewTestContext(t, key1, key2)

	store1 := ctx.KVStore(key1)
	store2 := ctx.KVStore(key2)

	store1.Set([]byte("k"), []byte("v1"))
	store2.Set([]byte("k"), []byte("v2"))

	assert.Equal(t, []byte("v1"), store1.Get([]byte("k")))
	assert.Equal(t, []byte("v2"), store2.Get([]byte("k")))
}

func TestNewTestContextWithBankAuth(t *testing.T) {
	moduleStoreKey := storeTypes.NewKVStoreKey("testModule")
	ctx, keepers, genesisAddr := NewTestContextWithBankAuth(t, "testModule", moduleStoreKey)

	assert.NotNil(t, ctx)
	assert.NotNil(t, keepers.AuthKeeper)
	assert.NotNil(t, keepers.BankKeeper)
	assert.NotEmpty(t, genesisAddr)

	// verify genesis address has coins
	balance := keepers.BankKeeper.GetBalance(ctx, genesisAddr, Denom)
	assert.Equal(t, int64(GenesisSupply), balance.Amount.Int64())
}
