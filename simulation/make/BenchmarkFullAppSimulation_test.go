// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

import (
	"os"
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/applications/base"
)

// Profile with:
// /usr/local/go/bin/go test -benchmem -run=^$ github.com/AssetMantle/modules/simapp -bench ^BenchmarkFullAppSimulation$ -Commit=true -cpuprofile cpu.out
func BenchmarkFullAppSimulation(b *testing.B) {
	config, db, _, logger, _, closeFn, err := setupRun(b, "goleveldb-app-sim", "Simulation")
	defer closeFn()

	require.NoError(b, err, "simulation setup failed")

	prototype := base.NewSimulationApplication(applicationName, moduleBasicManager, wasm.EnableAllProposals, moduleAccountPermissions, tokenReceiveAllowedModules)
	simulationApplication := prototype.Initialize(logger, db, nil, true, simapp.FlagPeriodValue, map[int64]bool{}, prototype.GetDefaultNodeHome(), interBlockCacheOpt()).(*base.SimulationApplication)

	// run randomized simulation
	_, simParams, simErr := simulation.SimulateFromSeed(
		b, os.Stdout, simulationApplication.GetBaseApp(), simapp.AppStateFn(simulationApplication.Codec(), simulationApplication.SimulationManager()),
		simapp.SimulationOperations(simulationApplication, simulationApplication.Codec(), config),
		simulationApplication.ModuleAccountAddrs(), config,
	)

	// export state and simParams before the simulation error is checked
	if err = simapp.CheckExportSimulation(simulationApplication, config, simParams); err != nil {
		b.Fatal(err)
	}

	if simErr != nil {
		b.Fatal(simErr)
	}

	if config.Commit {
		simapp.PrintStats(db)
	}
}
