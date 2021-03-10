/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"bytes"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptoCodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/persistenceOne/persistenceSDK/schema/applications"
	tendermintProto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
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

	context := sdkTypes.NewContext(commitMultiStore, tendermintProto.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	return context, storeKey, paramsTransientStoreKeys
}

func MakeCodec() *codec.LegacyAmino {
	var Codec = codec.NewLegacyAmino()

	schema.RegisterCodec(Codec)
	sdkTypes.RegisterLegacyAminoCodec(Codec)
	cryptoCodec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vestingTypes.RegisterLegacyAminoCodec(Codec)

	return Codec
}

// MakeEncodingConfig creates an EncodingConfig for an amino based test configuration.
func MakeEncodingConfig() applications.EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)

	return applications.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}
}

// key struct, implements helpers.Key
type testKey struct {
	ID string
}

var _ helpers.Key = (*testKey)(nil)

func (t testKey) GenerateStoreKeyBytes() []byte {
	return append([]byte{0x11}, []byte(t.ID)...)
}

func (t testKey) RegisterCodec(codec *codec.LegacyAmino) {
	codec.RegisterConcrete(testKey{}, "test/testKey", nil)
}

func (t testKey) IsPartial() bool {
	return t.ID != ""
}

func (t testKey) Matches(key helpers.Key) bool {
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

func (t testMappable) RegisterCodec(c *codec.LegacyAmino) {
	c.RegisterConcrete(testMappable{}, "test/testMappable", nil)
}

func NewMappable(id string, value string) helpers.Mappable {
	return testMappable{ID: id, Value: value}
}

func MappablePrototype() helpers.Mappable {
	return testMappable{}
}
