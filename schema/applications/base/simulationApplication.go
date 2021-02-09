/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/cosmos/cosmos-sdk/x/upgrade"

	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/schema/applications"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	tendermintTypes "github.com/tendermint/tendermint/types"
)

type simulationApplication struct {
	application   application
	transientKeys map[string]*sdkTypes.TransientStoreKey
	sm            *module.SimulationManager
	subspaces     map[string]params.Subspace

	AccountKeeper  auth.AccountKeeper
	BankKeeper     bank.Keeper
	SupplyKeeper   supply.Keeper
	StakingKeeper  staking.Keeper
	SlashingKeeper slashing.Keeper
	MintKeeper     mint.Keeper
	DistrKeeper    distr.Keeper
	GovKeeper      gov.Keeper
	CrisisKeeper   crisis.Keeper
	UpgradeKeeper  upgrade.Keeper
	ParamsKeeper   params.Keeper
	EvidenceKeeper evidence.Keeper
}

func (simulationApplication simulationApplication) Info(info abciTypes.RequestInfo) abciTypes.ResponseInfo {
	panic("implement me")
}

func (simulationApplication simulationApplication) SetOption(option abciTypes.RequestSetOption) abciTypes.ResponseSetOption {
	panic("implement me")
}

func (simulationApplication simulationApplication) Query(query abciTypes.RequestQuery) abciTypes.ResponseQuery {
	panic("implement me")
}

func (simulationApplication simulationApplication) CheckTx(tx abciTypes.RequestCheckTx) abciTypes.ResponseCheckTx {
	panic("implement me")
}

func (simulationApplication simulationApplication) InitChain(chain abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
	panic("implement me")
}

func (simulationApplication simulationApplication) BeginBlock(block abciTypes.RequestBeginBlock) abciTypes.ResponseBeginBlock {
	panic("implement me")
}

func (simulationApplication simulationApplication) DeliverTx(tx abciTypes.RequestDeliverTx) abciTypes.ResponseDeliverTx {
	panic("implement me")
}

func (simulationApplication simulationApplication) EndBlock(block abciTypes.RequestEndBlock) abciTypes.ResponseEndBlock {
	panic("implement me")
}

func (simulationApplication simulationApplication) Commit() abciTypes.ResponseCommit {
	panic("implement me")
}

func (simulationApplication simulationApplication) LoadHeight(i int64) error {
	panic("implement me")
}

func (simulationApplication simulationApplication) ExportApplicationStateAndValidators(b bool, strings []string) (json.RawMessage, []tendermintTypes.GenesisValidator, error) {
	panic("implement me")
}

func (simulationApplication simulationApplication) Initialize(applicationName string, codec *codec.Codec, enabledProposals []wasm.ProposalType, moduleAccountPermissions map[string][]string, tokenReceiveAllowedModules map[string]bool, logger log.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string, baseAppOptions ...func(*baseapp.BaseApp)) applications.Application {
	panic("implement me")
}

func (simulationApplication simulationApplication) Name() string {
	panic("implement me")
}

func (simulationApplication simulationApplication) Codec() *codec.Codec {
	panic("implement me")
}

func (simulationApplication simulationApplication) BeginBlocker(ctx sdkTypes.Context, req abciTypes.RequestBeginBlock) abciTypes.ResponseBeginBlock {
	panic("implement me")
}

func (simulationApplication simulationApplication) EndBlocker(ctx sdkTypes.Context, req abciTypes.RequestEndBlock) abciTypes.ResponseEndBlock {
	panic("implement me")
}

func (simulationApplication simulationApplication) InitChainer(ctx sdkTypes.Context, req abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
	panic("implement me")
}

func (simulationApplication simulationApplication) ExportAppStateAndValidators(forZeroHeight bool, jailWhiteList []string) (json.RawMessage, []tendermintTypes.GenesisValidator, error) {
	panic("implement me")
}

func (simulationApplication simulationApplication) ModuleAccountAddrs() map[string]bool {
	panic("implement me")
}

func (simulationApplication simulationApplication) SimulationManager() *module.SimulationManager {
	panic("implement me")
}

func (simulationApplication simulationApplication) GetBaseApp() *baseapp.BaseApp {
	panic("implement me")
}

func (simulationApplication simulationApplication) GetKey(storeKey string) *sdkTypes.KVStoreKey {
	panic("implement me")
}

func (simulationApplication simulationApplication) GetTKey(storeKey string) *sdkTypes.TransientStoreKey {
	panic("implement me")
}

func (simulationApplication simulationApplication) GetSubspace(moduleName string) params.Subspace {
	panic("implement me")
}

func (simulationApplication simulationApplication) GetMaccPerms() map[string][]string {
	panic("implement me")
}

func (simulationApplication simulationApplication) BlacklistedAccAddrs() map[string]bool {
	panic("implement me")
}

func (simulationApplication simulationApplication) CheckBalance(t *testing.T, address sdkTypes.AccAddress, coins sdkTypes.Coins) {
	panic("implement me")
}

func (simulationApplication simulationApplication) AddTestAddresses(context sdkTypes.Context, i int, s sdkTypes.Int) []sdkTypes.AccAddress {
	panic("implement me")
}

func (simulationApplication simulationApplication) Setup(b bool) applications.SimulationApplication {
	panic("implement me")
}

func (simulationApplication simulationApplication) SetupWithGenesisAccounts(accounts []exported.GenesisAccount) applications.SimulationApplication {
	panic("implement me")
}

func (simulationApplication simulationApplication) NewTestApplication(b bool) (applications.SimulationApplication, sdkTypes.Context) {
	panic("implement me")
}

var _ applications.SimulationApplication = (*simulationApplication)(nil)
