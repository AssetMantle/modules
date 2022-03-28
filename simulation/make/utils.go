package make

import (
	"os"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"
)

type testingLogging interface {
	Error(...interface{})
	Fatal(...interface{})
}

func setupRun(t testingLogging, dirPrefix, dbName string) (config simulation.Config, db db.DB, dir string, logger log.Logger, skip bool, closeFn func(), err error) {
	closeFn = func() {}

	config, db, dir, logger, skip, err = simapp.SetupSimulation(dirPrefix, dbName)
	if err != nil {
		return
	}

	closeFn = func() {
		var gotError bool

		if db != nil {
			err = db.Close()
			if err != nil {
				gotError = true

				t.Error(err)
			}
		}

		err = os.RemoveAll(dir)
		if err != nil {
			gotError = true

			t.Error(err)
		}

		if gotError {
			t.Fatal()
		}
	}

	return
}
