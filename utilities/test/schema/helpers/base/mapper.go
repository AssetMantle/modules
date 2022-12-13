/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"bytes"
	cryptoCodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmProto "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"
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
	Error := commitMultiStore.LoadLatestVersion()
	require.Nil(t, Error)

	context := sdkTypes.NewContext(commitMultiStore, tmProto.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	return context, storeKey, paramsTransientStoreKeys
}

func MakeCodec() *codec.LegacyAmino {
	var Codec = codec.NewLegacyAmino()

	schema.RegisterLegacyAminoCodec(Codec)
	sdkTypes.RegisterLegacyAminoCodec(Codec)
	cryptoCodec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vestingTypes.RegisterLegacyAminoCodec(Codec)

	return Codec
}

// key struct, implements helpers.Key
type testKey struct {
	ID string
}

var _ helpers.Key = (*testKey)(nil)

func (t testKey) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	panic("implement me")
}

func (t testKey) GetStructReference() codec.ProtoMarshaler {
	panic("implement me")
}

func (t testKey) GenerateStoreKeyBytes() []byte {
	return append([]byte{0x11}, []byte(t.ID)...)
}

func (t testKey) RegisterCodec(codec *codec.LegacyAmino) {
	codec.RegisterConcrete(testKey{}, "test/testKey", nil)
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

func (t testMappable) RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codec.RegisterConcrete(testMappable{}, "test/testMappable", nil)
}

func (t testMappable) Size() int {
	panic("implement me")
}

func (t testMappable) MarshalTo(data []byte) (int, error) {
	panic("implement me")
}

func (t testMappable) Unmarshal(dAtA []byte) error {
	panic("implement me")
}

func (t testMappable) Reset() {
	panic("implement me")
}

func (t testMappable) String() string {
	panic("implement me")
}

func (t testMappable) ProtoMessage() {
	panic("implement me")
}

func (t testMappable) Marshal() ([]byte, error) {
	panic("implement me")
}

func (t testMappable) MarshalToSizedBuffer(i []byte) (int, error) {
	panic("implement me")
}

func (t testMappable) GetStructReference() codec.ProtoMarshaler {
	panic("implement me")
}

func (t testMappable) GetKey() helpers.Key {
	return NewKey(t.ID)
}

func NewMappable(id string, value string) helpers.Mappable {
	return testMappable{ID: id, Value: value}
}

func MappablePrototype() helpers.Mappable {
	return testMappable{}
}
