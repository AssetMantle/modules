/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package make

import (
	"fmt"
	"os"
	"testing"

	"github.com/persistenceOne/persistenceSDK/schema/applications/base"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func TestAppSimulationAfterImport(t *testing.T) {
	config, db, dir, logger, skip, err := simapp.SetupSimulation("leveldb-app-sim", "Simulation")
	if skip {
		t.Skip("skipping application simulation after import")
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
	stopEarly, simParams, simErr := simulation.SimulateFromSeed(
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

	if stopEarly {
		fmt.Printf("can't export or import a zero-validator genesis, exiting test...")
		return
	}

	fmt.Printf("exporting genesis...\n")

	appState, _, err := app.ExportAppStateAndValidators(true, []string{})
	require.NoError(t, err)

	fmt.Printf("importing genesis...\n")

	_, newDB, newDir, _, _, err := simapp.SetupSimulation("leveldb-app-sim-2", "Simulation-2")
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		Error := newDB.Close()
		require.Nil(t, Error)
		require.NoError(t, os.RemoveAll(newDir))
	}()

	newApp := base.NewSimApp().Initialize(base.ApplicationName, base.MakeCodec(), wasm.EnableAllProposals, base.ModuleAccountPermissions, base.AllowedReceivingModuleAccounts, logger, newDB, nil, true, simapp.FlagPeriodValue, map[int64]bool{}, base.DefaultNodeHome, fauxMerkleModeOpt).(base.SimulationApplication)
	require.Equal(t, "SimulationApplication", newApp.Name())

	newApp.InitChain(abci.RequestInitChain{
		AppStateBytes: appState,
	})

	_, _, err = simulation.SimulateFromSeed(
		t, os.Stdout, newApp.GetBaseApp(), simapp.AppStateFn(app.Codec(), app.SimulationManager()),
		simapp.SimulationOperations(newApp, newApp.Codec(), config),
		newApp.ModuleAccountAddrs(), config,
	)
	require.NoError(t, err)
}
