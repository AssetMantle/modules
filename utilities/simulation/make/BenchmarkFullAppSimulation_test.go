/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package make

import (
	"os"
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/persistenceOne/persistenceSDK/schema/applications/base"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// Profile with:
// /usr/local/go/bin/go test -benchmem -run=^$ github.com/persistence/persistenceSDK/simapp -bench ^BenchmarkFullAppSimulation$ -Commit=true -cpuprofile cpu.out
func BenchmarkFullAppSimulation(b *testing.B) {
	config, db, dir, logger, _, err := simapp.SetupSimulation("goleveldb-app-sim", "Simulation")
	if err != nil {
		b.Fatalf("simulation setup failed: %s", err.Error())
	}

	defer func() {
		Error := db.Close()
		err = os.RemoveAll(dir)
		if err != nil {
			b.Fatal(err)
		} else if Error != nil {
			b.Fatal(Error)
		}
	}()

	app := base.NewSimApp().Initialize(base.ApplicationName, base.MakeCodec(), wasm.EnableAllProposals, base.ModuleAccountPermissions, base.AllowedReceivingModuleAccounts, logger, db, nil, true, simapp.FlagPeriodValue, map[int64]bool{}, base.DefaultNodeHome, interBlockCacheOpt()).(base.SimulationApplication)

	// run randomized simulation
	_, simParams, simErr := simulation.SimulateFromSeed(
		b, os.Stdout, app.GetBaseApp(), simapp.AppStateFn(app.Codec(), app.SimulationManager()),
		simapp.SimulationOperations(app, app.Codec(), config),
		app.ModuleAccountAddrs(), config,
	)

	// export state and simParams before the simulation error is checked
	if err = simapp.CheckExportSimulation(app, config, simParams); err != nil {
		b.Fatal(err)
	}

	if simErr != nil {
		b.Fatal(simErr)
	}

	if config.Commit {
		simapp.PrintStats(db)
	}
}
