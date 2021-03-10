/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package applications

import (
	"io"

	serverTypes "github.com/cosmos/cosmos-sdk/server/types"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/baseapp"
	tendermintDB "github.com/tendermint/tm-db"
)

type Application interface {
	serverTypes.Application
	LoadHeight(int64) error
	ExportApplicationStateAndValidators(bool, []string) (serverTypes.ExportedApp, error)

	Initialize(applicationName string, encodingConfig EncodingConfig, enabledProposals []wasm.ProposalType, moduleAccountPermissions map[string][]string, tokenReceiveAllowedModules map[string]bool,
		logger log.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string, applicationOptions serverTypes.AppOptions, baseAppOptions ...func(*baseapp.BaseApp)) Application
}
