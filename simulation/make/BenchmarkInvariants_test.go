// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

import (
	"fmt"
	"os"
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/AssetMantle/modules/schema/applications/base"
)

func BenchmarkInvariants(b *testing.B) {
	config, db, _, logger, _, closeFn, err := setupRun(b, "leveldb-app-invariant-bench", "Simulation")
	defer closeFn()

	require.NoError(b, err, "simulation setup failed")

	config.AllInvariants = false

	prototype := base.NewSimulationApplication(applicationName, moduleBasicManager, wasm.EnableAllProposals, moduleAccountPermissions, tokenReceiveAllowedModules).(*base.SimulationApplication)
	simulationApplication := prototype.InitializeSimulationApplication(logger, db, nil, true, simapp.FlagPeriodValue, map[int64]bool{}, prototype.GetDefaultNodeHome(), interBlockCacheOpt()).(*base.SimulationApplication)

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

	ctx := simulationApplication.GetBaseApp().NewContext(true, abci.Header{Height: simulationApplication.GetBaseApp().LastBlockHeight() + 1})

	// 3. Benchmark each invariant separately
	//
	// NOTE: We use the crisis keeper as it has all the invariants registered with
	// their respective metadata which makes it useful for testing/benchmarking.
	for _, cr := range simulationApplication.CrisisKeeper.Routes() {
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
