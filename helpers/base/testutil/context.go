// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package testutil

import (
	"testing"

	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"cosmossdk.io/store"
	storeMetrics "cosmossdk.io/store/metrics"
	storeTypes "cosmossdk.io/store/types"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	cosmosDB "github.com/cosmos/cosmos-db"
	addressCodec "github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	authKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
)

// NewTestContext creates a test SDK context with the given store keys mounted.
// Returns the context and the first store key (convenience for single-key tests).
// Uses NoOp metrics and nil db per store for proper SDK v0.50 prefixing.
func NewTestContext(t testing.TB, storeKeys ...*storeTypes.KVStoreKey) sdkTypes.Context {
	t.Helper()

	memDB := cosmosDB.NewMemDB()
	cms := store.NewCommitMultiStore(memDB, log.NewNopLogger(), storeMetrics.NewNoOpMetrics())

	for _, key := range storeKeys {
		cms.MountStoreWithDB(key, storeTypes.StoreTypeIAVL, nil)
	}

	require.NoError(t, cms.LoadLatestVersion())

	return sdkTypes.NewContext(cms, protoTendermintTypes.Header{
		ChainID: ChainID,
	}, false, log.NewNopLogger())
}

// BankAuthKeepers holds auth and bank keeper instances for tests that need them.
type BankAuthKeepers struct {
	AuthKeeper authKeeper.AccountKeeper
	BankKeeper bankKeeper.BaseKeeper
	AuthStoreKey *storeTypes.KVStoreKey
	BankStoreKey *storeTypes.KVStoreKey
}

// NewTestContextWithBankAuth creates a test context with auth and bank keepers,
// mints GenesisSupply coins to the TestMinterModuleName, and sends them to the
// returned genesis address. This is the pattern used by transaction keeper tests
// that need to transfer coins.
func NewTestContextWithBankAuth(t testing.TB, moduleName string, moduleStoreKey *storeTypes.KVStoreKey) (sdkTypes.Context, BankAuthKeepers, sdkTypes.AccAddress) {
	t.Helper()

	authStoreKey := storeTypes.NewKVStoreKey(authTypes.StoreKey)
	bankStoreKey := storeTypes.NewKVStoreKey(bankTypes.StoreKey)

	memDB := cosmosDB.NewMemDB()
	cms := store.NewCommitMultiStore(memDB, log.NewNopLogger(), storeMetrics.NewNoOpMetrics())
	cms.MountStoreWithDB(moduleStoreKey, storeTypes.StoreTypeIAVL, nil)
	cms.MountStoreWithDB(authStoreKey, storeTypes.StoreTypeIAVL, nil)
	cms.MountStoreWithDB(bankStoreKey, storeTypes.StoreTypeIAVL, nil)
	require.NoError(t, cms.LoadLatestVersion())

	ctx := sdkTypes.NewContext(cms, protoTendermintTypes.Header{ChainID: ChainID}, false, log.NewNopLogger())

	codec := baseHelpers.TestCodec()
	moduleAccountPermissions := map[string][]string{
		TestMinterModuleName: {authTypes.Minter},
		moduleName:           nil,
	}

	ak := authKeeper.NewAccountKeeper(
		codec,
		runtime.NewKVStoreService(authStoreKey),
		authTypes.ProtoBaseAccount,
		moduleAccountPermissions,
		addressCodec.NewBech32Codec(sdkTypes.GetConfig().GetBech32AccountAddrPrefix()),
		sdkTypes.GetConfig().GetBech32AccountAddrPrefix(),
		authTypes.NewModuleAddress(govTypes.ModuleName).String(),
	)

	blacklistedAddresses := map[string]bool{
		authTypes.NewModuleAddress(TestMinterModuleName).String(): false,
		authTypes.NewModuleAddress(moduleName).String():           false,
	}

	bk := bankKeeper.NewBaseKeeper(
		codec,
		runtime.NewKVStoreService(bankStoreKey),
		ak,
		blacklistedAddresses,
		authTypes.NewModuleAddress(govTypes.ModuleName).String(),
		log.NewNopLogger(),
	)

	coinSupply := sdkTypes.NewCoins(sdkTypes.NewCoin(Denom, math.NewInt(GenesisSupply)))
	require.NoError(t, bk.MintCoins(ctx, TestMinterModuleName, coinSupply))

	genesisAddress := TestAddress()
	require.NoError(t, bk.SendCoinsFromModuleToAccount(ctx, TestMinterModuleName, genesisAddress, coinSupply))

	keepers := BankAuthKeepers{
		AuthKeeper:   ak,
		BankKeeper:   bk,
		AuthStoreKey: authStoreKey,
		BankStoreKey: bankStoreKey,
	}

	return ctx, keepers, genesisAddress
}
