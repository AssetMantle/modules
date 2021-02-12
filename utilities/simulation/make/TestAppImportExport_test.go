/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package make

import (
	"fmt"
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/store"

	"github.com/persistenceOne/persistenceSDK/schema/applications/base"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/persistenceOne/persistenceSDK/modules/assets"
	"github.com/persistenceOne/persistenceSDK/modules/classifications"
	"github.com/persistenceOne/persistenceSDK/modules/identities"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers"
	"github.com/persistenceOne/persistenceSDK/modules/metas"
	"github.com/persistenceOne/persistenceSDK/modules/orders"
	"github.com/persistenceOne/persistenceSDK/modules/splits"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

// Get flags every time the simulator is run
func init() {
	simapp.GetSimulatorFlags()
}

type StoreKeysPrefixes struct {
	A        sdk.StoreKey
	B        sdk.StoreKey
	Prefixes [][]byte
}

// fauxMerkleModeOpt returns a BaseApp option to use a dbStoreAdapter instead of
// an IAVLStore for faster simulation speed.
func fauxMerkleModeOpt(baseApplication *baseapp.BaseApp) {
	baseApplication.SetFauxMerkleMode()
}

func interBlockCacheOpt() func(*baseapp.BaseApp) {
	return baseapp.SetInterBlockCache(store.NewCommitKVStoreCacheManager())
}

func TestAppImportExport(t *testing.T) {
	config, db, dir, logger, skip, err := simapp.SetupSimulation("leveldb-app-sim", "Simulation")
	if skip {
		t.Skip("skipping application import/export simulation")
	}
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		Error := db.Close()
		require.Nil(t, Error)
		require.NoError(t, os.RemoveAll(dir))
	}()

	app := base.NewSimApp().Initialize(base.ApplicationName, base.MakeCodec(), wasm.EnableAllProposals, base.ModuleAccountPermissions, base.AllowedReceivingModuleAccounts, logger, db, nil, true, simapp.FlagPeriodValue, map[int64]bool{}, base.DefaultNodeHome, fauxMerkleModeOpt).(base.SimulationApplication)
	require.Equal(t, "SimulationApplication", app.Name())

	// Run randomized simulation
	_, simParams, simErr := simulation.SimulateFromSeed(
		t, os.Stdout, app.GetBaseApp(), simapp.AppStateFn(app.Codec(), app.SimulationManager()),
		simapp.SimulationOperations(app, app.Codec(), config),
		app.ModuleAccountAddrs(), config,
	)

	// export state and simParams before the simulation error is checked
	err = simapp.CheckExportSimulation(app, config, simParams)
	require.NoError(t, err)
	require.NoError(t, simErr)

	if config.Commit {
		simapp.PrintStats(db)
	}

	fmt.Printf("exporting genesis...\n")

	appState, _, err := app.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err)

	fmt.Printf("importing genesis...\n")

	_, newDB, newDir, _, _, err := simapp.SetupSimulation("leveldb-app-sim-2", "Simulation-2")
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		Error := newDB.Close()
		require.Nil(t, Error)
		require.NoError(t, os.RemoveAll(newDir))
	}()

	newApp := base.NewSimApp().Initialize(base.ApplicationName, base.MakeCodec(), wasm.EnableAllProposals, base.ModuleAccountPermissions, base.AllowedReceivingModuleAccounts, logger, db, nil, true, simapp.FlagPeriodValue, map[int64]bool{}, base.DefaultNodeHome).(base.SimulationApplication)
	require.Equal(t, "SimulationApplication", newApp.Name())

	var genesisState simapp.GenesisState
	err = app.Codec().UnmarshalJSON(appState, &genesisState)
	require.NoError(t, err)

	ctxA := app.GetBaseApp().NewContext(true, abci.Header{Height: app.GetBaseApp().LastBlockHeight()})
	ctxB := newApp.GetBaseApp().NewContext(true, abci.Header{Height: app.GetBaseApp().LastBlockHeight()})
	newApp.ModuleManager().InitGenesis(ctxB, genesisState)

	fmt.Printf("comparing stores...\n")

	storeKeysPrefixes := []StoreKeysPrefixes{
		{app.GetKey(baseapp.MainStoreKey), newApp.GetKey(baseapp.MainStoreKey), [][]byte{}},
		{app.GetKey(auth.StoreKey), newApp.GetKey(auth.StoreKey), [][]byte{}},
		{app.GetKey(staking.StoreKey), newApp.GetKey(staking.StoreKey),
			[][]byte{
				staking.UnbondingQueueKey, staking.RedelegationQueueKey, staking.ValidatorQueueKey,
			}}, // ordering may change but it doesn't matter
		{app.GetKey(slashing.StoreKey), newApp.GetKey(slashing.StoreKey), [][]byte{}},
		{app.GetKey(mint.StoreKey), newApp.GetKey(mint.StoreKey), [][]byte{}},
		{app.GetKey(distribution.StoreKey), newApp.GetKey(distribution.StoreKey), [][]byte{}},
		{app.GetKey(supply.StoreKey), newApp.GetKey(supply.StoreKey), [][]byte{}},
		{app.GetKey(params.StoreKey), newApp.GetKey(params.StoreKey), [][]byte{}},
		{app.GetKey(gov.StoreKey), newApp.GetKey(gov.StoreKey), [][]byte{}},
		{app.GetKey(wasm.StoreKey), newApp.GetKey(wasm.StoreKey), [][]byte{}},
		{app.GetKey(assets.Prototype().Name()), newApp.GetKey(assets.Prototype().Name()), [][]byte{}},
		{app.GetKey(classifications.Prototype().Name()), newApp.GetKey(classifications.Prototype().Name()), [][]byte{}},
		{app.GetKey(identities.Prototype().Name()), newApp.GetKey(identities.Prototype().Name()), [][]byte{}},
		{app.GetKey(maintainers.Prototype().Name()), newApp.GetKey(maintainers.Prototype().Name()), [][]byte{}},
		{app.GetKey(metas.Prototype().Name()), newApp.GetKey(metas.Prototype().Name()), [][]byte{}},
		{app.GetKey(orders.Prototype().Name()), newApp.GetKey(orders.Prototype().Name()), [][]byte{}},
		{app.GetKey(splits.Prototype().Name()), newApp.GetKey(splits.Prototype().Name()), [][]byte{}},
	}

	for _, skp := range storeKeysPrefixes {
		storeA := ctxA.KVStore(skp.A)
		storeB := ctxB.KVStore(skp.B)

		failedKVAs, failedKVBs := sdk.DiffKVStores(storeA, storeB, skp.Prefixes)
		require.Equal(t, len(failedKVAs), len(failedKVBs), "unequal sets of key-values to compare")

		fmt.Printf("compared %d key/value pairs between %s and %s\n", len(failedKVAs), skp.A, skp.B)
		require.Equal(t, len(failedKVAs), 0, simapp.GetSimulationLog(skp.A.Name(), app.SimulationManager().StoreDecoders, app.Codec(), failedKVAs, failedKVBs))
	}
}
