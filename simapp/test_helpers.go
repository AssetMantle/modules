package simapp

import (
	"github.com/cosmos/cosmos-sdk/simapp"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/codec"
)

// Setup initializes a new SimulationApplication. A Nop logger is set in SimulationApplication.
func Setup(isCheckTx bool) *SimulationApplication {
	db := dbm.NewMemDB()
	application := NewSimApp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, 0)

	if !isCheckTx {
		// init chain must be called to stop deliverState from being nil
		genesisState := simapp.NewDefaultGenesisState()

		stateBytes, err := codec.MarshalJSONIndent(application.Codec(), genesisState)
		if err != nil {
			panic(err)
		}

		// Initialize the chain
		application.InitChain(
			abci.RequestInitChain{
				Validators:    []abci.ValidatorUpdate{},
				AppStateBytes: stateBytes,
			},
		)
	}

	return application
}
