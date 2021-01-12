/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package applications

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/baseapp"
	tendermintABCITypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintTypes "github.com/tendermint/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"
	"io"
)

type Application interface {
	LoadHeight(int64) error
	ExportApplicationStateAndValidators(bool, []string) (json.RawMessage, []tendermintTypes.GenesisValidator, error)
	tendermintABCITypes.Application
}

type NewApplication func(logger log.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string, baseAppOptions ...func(*baseapp.BaseApp)) Application
