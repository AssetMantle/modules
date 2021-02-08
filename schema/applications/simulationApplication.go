/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package applications

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/x/auth/exported"

	"github.com/cosmos/cosmos-sdk/simapp"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

type SimulationApplication interface {
	Application
	simapp.App

	GetBaseApp() *baseapp.BaseApp
	GetKey(storeKey string) *sdk.KVStoreKey
	GetTKey(storeKey string) *sdk.TransientStoreKey
	GetSubspace(moduleName string) params.Subspace
	GetMaccPerms() map[string][]string
	BlacklistedAccAddrs() map[string]bool

	CheckBalance(*testing.T, sdk.AccAddress, sdk.Coins)

	AddTestAddresses(sdk.Context, int, sdk.Int) []sdk.AccAddress

	Setup(bool) SimulationApplication
	SetupWithGenesisAccounts([]exported.GenesisAccount) SimulationApplication
	NewTestApplication(bool) (SimulationApplication, sdk.Context)
}
