// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"testing"

	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/schema/helpers"
)

func SetupTest(t *testing.T) (sdkTypes.Context, *sdkTypes.KVStoreKey, *sdkTypes.TransientStoreKey) {
	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	return context, storeKey, paramsTransientStoreKeys
}

// key struct, implements helpers.Key
type testKey struct {
	ID string
}

var _ helpers.Key = (*testKey)(nil)

func (t testKey) String() string {
	return t.ID
}

func (t testKey) GenerateStoreKeyBytes() []byte {
	return append([]byte{0x11}, []byte(t.ID)...)
}

func (t testKey) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	legacyAmino.RegisterConcrete(testKey{}, "test/testKey", nil)
}

func (t testKey) IsPartial() bool {
	return t.ID != ""
}

func (t testKey) Equals(key helpers.Key) bool {
	return bytes.Equal([]byte(t.ID), []byte(key.(testKey).ID))
}

func NewKey(id string) helpers.Key {
	return testKey{ID: id}
}

func KeyPrototype() helpers.Key {
	return testKey{}
}

// mappable struct, implements helpers.Mappable
type testMappable struct {
	ID    string
	Value string
}

var _ helpers.Mappable = (*testMappable)(nil)

func (t testMappable) GetKey() helpers.Key {
	return NewKey(t.ID)
}

func (t testMappable) RegisterLegacyAminoCodec(c *codec.LegacyAmino) {
	c.RegisterConcrete(testMappable{}, "test/testMappable", nil)
}

func NewMappable(id string, value string) helpers.Mappable {
	return testMappable{ID: id, Value: value}
}

func MappablePrototype() helpers.Mappable {
	return testMappable{}
}
