/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package block

import (
	"testing"

	cryptoCodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	tendermintProto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
)

func CreateTestInput(t *testing.T) sdkTypes.Context {
	var Codec = codec.NewLegacyAmino()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterLegacyAminoCodec(Codec)
	cryptoCodec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vestingTypes.RegisterLegacyAminoCodec(Codec)
	Codec.Seal()
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

	return context
}

func Test_Block_Methods(t *testing.T) {
	block := Prototype()
	block.Initialize(mapper.Prototype(), parameters.Prototype(), []helpers.Auxiliary{})
	context := CreateTestInput(t)
	block.Begin(context, abciTypes.RequestBeginBlock{})
	block.End(context, abciTypes.RequestEndBlock{})
}
