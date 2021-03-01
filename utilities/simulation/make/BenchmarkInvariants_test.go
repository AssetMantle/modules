/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package make

import (
	"fmt"
	"os"
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/persistenceOne/persistenceSDK/schema/applications/base"
	abci "github.com/tendermint/tendermint/abci/types"
)

func BenchmarkInvariants(b *testing.B) {
	config, db, dir, logger, _, err := simapp.SetupSimulation("leveldb-app-invariant-bench", "Simulation")
	if err != nil {
		b.Fatalf("simulation setup failed: %s", err.Error())
	}

	config.AllInvariants = false

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

	ctx := app.GetBaseApp().NewContext(true, abci.Header{Height: app.GetBaseApp().LastBlockHeight() + 1})

	// 3. Benchmark each invariant separately
	//
	// NOTE: We use the crisis keeper as it has all the invariants registered with
	// their respective metadata which makes it useful for testing/benchmarking.
	for _, cr := range app.CrisisKeeper.Routes() {
		cr := cr
		b.Run(fmt.Sprintf("%s/%s", cr.ModuleName, cr.Route), func(b *testing.B) {
			if res, stop := cr.Invar(ctx); stop {
				b.Fatalf(
					"broken invariant at block %d of %d\n%s",
					ctx.BlockHeight()-1, config.NumBlocks, res,
				)
			}
		})
	}
}
