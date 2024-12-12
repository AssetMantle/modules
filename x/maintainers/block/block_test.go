// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package block

import (
	"context"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	"testing"

	tendermintDB "github.com/cometbft/cometbft-db"
	abciTypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/maintainers/mapper"
	"github.com/AssetMantle/modules/x/maintainers/parameters"
)

func CreateTestInput(t *testing.T) context.Context {
	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, storeTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	return sdkTypes.WrapSDKContext(context)
}

func Test_Block_Methods(t *testing.T) {
	block := Prototype()
	block.Initialize(mapper.Prototype(), parameters.Prototype(), []helpers.Auxiliary{})
	context := CreateTestInput(t)
	block.Begin(context, abciTypes.RequestBeginBlock{})
	block.End(context, abciTypes.RequestEndBlock{})
}
