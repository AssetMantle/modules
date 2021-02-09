/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package applications

import (
	"encoding/json"
	"io"

	"github.com/tendermint/tendermint/libs/log"

	tendermintABCITypes "github.com/tendermint/tendermint/abci/types"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	tendermintTypes "github.com/tendermint/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"
)

type Application interface {
	tendermintABCITypes.Application

	LoadHeight(int64) error
	ExportApplicationStateAndValidators(bool, []string) (json.RawMessage, []tendermintTypes.GenesisValidator, error)

	Initialize(applicationName string, codec *codec.Codec, enabledProposals []wasm.ProposalType, moduleAccountPermissions map[string][]string, tokenReceiveAllowedModules map[string]bool,
		logger log.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string, baseAppOptions ...func(*baseapp.BaseApp)) Application
}
