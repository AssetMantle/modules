// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package applications

import (
	"io"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
)

type SimulationApplication interface {
	Application
	simapp.App

	GetBaseApp() *baseapp.BaseApp
	GetKey(storeKey string) *sdk.KVStoreKey
	GetTKey(storeKey string) *sdk.TransientStoreKey
	GetSubspace(moduleName string) params.Subspace
	GetModuleAccountPermissions() map[string][]string
	GetBlackListedAddresses() map[string]bool
	ModuleManager() *module.Manager

	CheckBalance(*testing.T, sdk.AccAddress, sdk.Coins)

	AddTestAddresses(sdk.Context, int, sdk.Int) []sdk.AccAddress

	Setup(bool) SimulationApplication
	SetupWithGenesisAccounts([]exported.GenesisAccount) SimulationApplication
	NewTestApplication(bool) (SimulationApplication, sdk.Context)
	InitializeSimulationApplication(logger log.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string, baseAppOptions ...func(*baseapp.BaseApp)) SimulationApplication
}
