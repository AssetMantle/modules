/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package applications

import (
	"encoding/json"
	"io"

	"github.com/cosmos/cosmos-sdk/baseapp"
	tendermintABCITypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintTypes "github.com/tendermint/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"
)

type Application interface {
	tendermintABCITypes.Application

	LoadHeight(int64) error
	ExportApplicationStateAndValidators(bool, []string) (json.RawMessage, []tendermintTypes.GenesisValidator, error)

	Initialize(log.Logger, tendermintDB.DB, io.Writer, bool, uint, map[int64]bool, string, ...func(*baseapp.BaseApp)) Application
}
