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

	prototype := base.NewSimulationApplication(applicationName, moduleBasicManager, wasm.EnableAllProposals, moduleAccountPermissions, tokenReceiveAllowedModules)
	simulationApplication := prototype.Initialize(logger, db, nil, true, simapp.FlagPeriodValue, map[int64]bool{}, prototype.GetDefaultNodeHome(), fauxMerkleModeOpt).(*base.SimulationApplication)
	require.Equal(t, "SimulationApplication", simulationApplication.Name())

	// Run randomized simulation
	_, simParams, simErr := simulation.SimulateFromSeed(
		t, os.Stdout, simulationApplication.GetBaseApp(), simapp.AppStateFn(simulationApplication.Codec(), simulationApplication.SimulationManager()),
		simapp.SimulationOperations(simulationApplication, simulationApplication.Codec(), config),
		simulationApplication.ModuleAccountAddrs(), config,
	)

	// export state and simParams before the simulation error is checked
	err = simapp.CheckExportSimulation(simulationApplication, config, simParams)
	require.NoError(t, err)
	require.NoError(t, simErr)

	if config.Commit {
		simapp.PrintStats(db)
	}

	fmt.Printf("exporting genesis...\n")

	appState, _, err := simulationApplication.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err)

	fmt.Printf("importing genesis...\n")

	_, newDB, newDir, _, _, err := simapp.SetupSimulation("leveldb-app-sim-2", "Simulation-2")
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		Error := newDB.Close()
		require.Nil(t, Error)
		require.NoError(t, os.RemoveAll(newDir))
	}()

	newSimulationApplication := prototype.Initialize(logger, db, nil, true, simapp.FlagPeriodValue, map[int64]bool{}, prototype.GetDefaultNodeHome()).(*base.SimulationApplication)
	require.Equal(t, "SimulationApplication", newSimulationApplication.Name())

	var genesisState simapp.GenesisState
	err = simulationApplication.Codec().UnmarshalJSON(appState, &genesisState)
	require.NoError(t, err)

	ctxA := simulationApplication.GetBaseApp().NewContext(true, abci.Header{Height: simulationApplication.GetBaseApp().LastBlockHeight()})
	ctxB := newSimulationApplication.GetBaseApp().NewContext(true, abci.Header{Height: simulationApplication.GetBaseApp().LastBlockHeight()})
	newSimulationApplication.ModuleManager().InitGenesis(ctxB, genesisState)

	fmt.Printf("comparing stores...\n")

	storeKeysPrefixes := []StoreKeysPrefixes{
		{simulationApplication.GetKey(baseapp.MainStoreKey), newSimulationApplication.GetKey(baseapp.MainStoreKey), [][]byte{}},
		{simulationApplication.GetKey(auth.StoreKey), newSimulationApplication.GetKey(auth.StoreKey), [][]byte{}},
		{simulationApplication.GetKey(staking.StoreKey), newSimulationApplication.GetKey(staking.StoreKey),
			[][]byte{
				staking.UnbondingQueueKey, staking.RedelegationQueueKey, staking.ValidatorQueueKey,
			}}, // ordering may change but it doesn't matter
		{simulationApplication.GetKey(slashing.StoreKey), newSimulationApplication.GetKey(slashing.StoreKey), [][]byte{}},
		{simulationApplication.GetKey(mint.StoreKey), newSimulationApplication.GetKey(mint.StoreKey), [][]byte{}},
		{simulationApplication.GetKey(distribution.StoreKey), newSimulationApplication.GetKey(distribution.StoreKey), [][]byte{}},
		{simulationApplication.GetKey(supply.StoreKey), newSimulationApplication.GetKey(supply.StoreKey), [][]byte{}},
		{simulationApplication.GetKey(params.StoreKey), newSimulationApplication.GetKey(params.StoreKey), [][]byte{}},
		{simulationApplication.GetKey(gov.StoreKey), newSimulationApplication.GetKey(gov.StoreKey), [][]byte{}},
		{simulationApplication.GetKey(wasm.StoreKey), newSimulationApplication.GetKey(wasm.StoreKey), [][]byte{}},
		{simulationApplication.GetKey(assets.Prototype().Name()), newSimulationApplication.GetKey(assets.Prototype().Name()), [][]byte{}},
		{simulationApplication.GetKey(classifications.Prototype().Name()), newSimulationApplication.GetKey(classifications.Prototype().Name()), [][]byte{}},
		{simulationApplication.GetKey(identities.Prototype().Name()), newSimulationApplication.GetKey(identities.Prototype().Name()), [][]byte{}},
		{simulationApplication.GetKey(maintainers.Prototype().Name()), newSimulationApplication.GetKey(maintainers.Prototype().Name()), [][]byte{}},
		{simulationApplication.GetKey(metas.Prototype().Name()), newSimulationApplication.GetKey(metas.Prototype().Name()), [][]byte{}},
		{simulationApplication.GetKey(orders.Prototype().Name()), newSimulationApplication.GetKey(orders.Prototype().Name()), [][]byte{}},
		{simulationApplication.GetKey(splits.Prototype().Name()), newSimulationApplication.GetKey(splits.Prototype().Name()), [][]byte{}},
	}

	for _, skp := range storeKeysPrefixes {
		storeA := ctxA.KVStore(skp.A)
		storeB := ctxB.KVStore(skp.B)

		failedKVAs, failedKVBs := sdk.DiffKVStores(storeA, storeB, skp.Prefixes)
		require.Equal(t, len(failedKVAs), len(failedKVBs), "unequal sets of key-values to compare")

		fmt.Printf("compared %d key/value pairs between %s and %s\n", len(failedKVAs), skp.A, skp.B)
		require.Equal(t, len(failedKVAs), 0, simapp.GetSimulationLog(skp.A.Name(), simulationApplication.SimulationManager().StoreDecoders, simulationApplication.Codec(), failedKVAs, failedKVBs))
	}
}
