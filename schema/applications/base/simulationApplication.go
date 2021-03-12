/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	"github.com/persistenceOne/persistenceSDK/schema/applications"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/log"
	tendermintTypes "github.com/tendermint/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"
)

type SimulationApplication struct {
	application

	transientStoreKeys map[string]*sdkTypes.TransientStoreKey
	simulationManager  *module.SimulationManager
	subspaces          map[string]params.Subspace

	AccountKeeper      auth.AccountKeeper
	BankKeeper         bank.Keeper
	SupplyKeeper       supply.Keeper
	StakingKeeper      staking.Keeper
	SlashingKeeper     slashing.Keeper
	MintKeeper         mint.Keeper
	DistributionKeeper distribution.Keeper
	GovKeeper          gov.Keeper
	CrisisKeeper       crisis.Keeper
	UpgradeKeeper      upgrade.Keeper
	ParamsKeeper       params.Keeper
	EvidenceKeeper     evidence.Keeper
}

var _ applications.SimulationApplication = (*SimulationApplication)(nil)

func (simulationApplication SimulationApplication) Codec() *codec.Codec {
	return simulationApplication.application.codec
}

func (simulationApplication SimulationApplication) BeginBlocker(ctx sdkTypes.Context, req abciTypes.RequestBeginBlock) abciTypes.ResponseBeginBlock {
	return simulationApplication.application.moduleManager.BeginBlock(ctx, req)
}

func (simulationApplication SimulationApplication) EndBlocker(ctx sdkTypes.Context, req abciTypes.RequestEndBlock) abciTypes.ResponseEndBlock {
	return simulationApplication.application.moduleManager.EndBlock(ctx, req)
}

func (simulationApplication SimulationApplication) InitChainer(ctx sdkTypes.Context, req abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
	var genesisState simapp.GenesisState

	simulationApplication.application.codec.MustUnmarshalJSON(req.AppStateBytes, &genesisState)

	return simulationApplication.application.moduleManager.InitGenesis(ctx, genesisState)
}

func (simulationApplication SimulationApplication) ExportAppStateAndValidators(forZeroHeight bool, jailWhiteList []string) (json.RawMessage, []tendermintTypes.GenesisValidator, error) {
	return simulationApplication.application.ExportApplicationStateAndValidators(forZeroHeight, jailWhiteList)
}

func (simulationApplication SimulationApplication) ModuleAccountAddrs() map[string]bool {
	return simulationApplication.tokenReceiveAllowedModules
}

func (simulationApplication SimulationApplication) SimulationManager() *module.SimulationManager {
	return simulationApplication.simulationManager
}

func (simulationApplication SimulationApplication) ModuleManager() *module.Manager {
	return simulationApplication.application.moduleManager
}

func (simulationApplication SimulationApplication) GetBaseApp() *baseapp.BaseApp {
	return &simulationApplication.application.BaseApp
}

func (simulationApplication SimulationApplication) GetKey(storeKey string) *sdkTypes.KVStoreKey {
	return simulationApplication.application.keys[storeKey]
}

func (simulationApplication SimulationApplication) GetTKey(storeKey string) *sdkTypes.TransientStoreKey {
	return simulationApplication.transientStoreKeys[storeKey]
}

func (simulationApplication SimulationApplication) GetSubspace(moduleName string) params.Subspace {
	return simulationApplication.subspaces[moduleName]
}

func (simulationApplication SimulationApplication) GetModuleAccountPermissions() map[string][]string {
	return simulationApplication.application.moduleAccountPermissions
}

func (simulationApplication SimulationApplication) GetBlackListedAddresses() map[string]bool {
	blacklistedAddrs := make(map[string]bool)
	for acc := range simulationApplication.moduleAccountPermissions {
		blacklistedAddrs[supply.NewModuleAddress(acc).String()] = !simulationApplication.tokenReceiveAllowedModules[acc]
	}

	return blacklistedAddrs
}

func (simulationApplication SimulationApplication) CheckBalance(t *testing.T, address sdkTypes.AccAddress, coins sdkTypes.Coins) {
	ctxCheck := simulationApplication.application.NewContext(true, abciTypes.Header{})
	res := simulationApplication.AccountKeeper.GetAccount(ctxCheck, address)

	require.True(t, coins.IsEqual(res.GetCoins()))
}

func (simulationApplication SimulationApplication) AddTestAddresses(context sdkTypes.Context, accountNumber int, amount sdkTypes.Int) []sdkTypes.AccAddress {
	testAddresses := make([]sdkTypes.AccAddress, accountNumber)

	for i := 0; i < accountNumber; i++ {
		pk := ed25519.GenPrivKey().PubKey()
		testAddresses[i] = sdkTypes.AccAddress(pk.Address())
	}

	initCoins := sdkTypes.NewCoins(sdkTypes.NewCoin(simulationApplication.StakingKeeper.BondDenom(context), amount))
	totalSupply := sdkTypes.NewCoins(sdkTypes.NewCoin(simulationApplication.StakingKeeper.BondDenom(context), amount.MulRaw(int64(len(testAddresses)))))
	prevSupply := simulationApplication.SupplyKeeper.GetSupply(context)
	simulationApplication.SupplyKeeper.SetSupply(context, supply.NewSupply(prevSupply.GetTotal().Add(totalSupply...)))

	// fill all the addresses with some coins, set the loose pool tokens simultaneously
	for _, addr := range testAddresses {
		_, err := simulationApplication.BankKeeper.AddCoins(context, addr, initCoins)
		if err != nil {
			panic(err)
		}
	}

	return testAddresses
}

func (simulationApplication SimulationApplication) Setup(isCheckTx bool) applications.SimulationApplication {
	db := tendermintDB.NewMemDB()
	newSimulationApplication := simulationApplication.Initialize(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0, map[int64]bool{}, simulationApplication.GetDefaultNodeHome()).(*SimulationApplication)

	if !isCheckTx {
		// init chain must be called to stop deliverState from being nil
		genesisState := simulationApplication.GetModuleBasicManager().DefaultGenesis()

		stateBytes, err := codec.MarshalJSONIndent(simulationApplication.Codec(), genesisState)
		if err != nil {
			panic(err)
		}

		// Initialize the chain
		newSimulationApplication.InitChain(
			abciTypes.RequestInitChain{
				Validators:    []abciTypes.ValidatorUpdate{},
				AppStateBytes: stateBytes,
			},
		)
	}

	return newSimulationApplication
}

func (simulationApplication SimulationApplication) SetupWithGenesisAccounts(accounts []exported.GenesisAccount) applications.SimulationApplication {
	db := tendermintDB.NewMemDB()
	newSimulationApplication := simulationApplication.Initialize(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0, map[int64]bool{}, simulationApplication.GetDefaultNodeHome()).(*SimulationApplication)

	// initialize the chain with the passed in genesis accounts
	genesisState := simulationApplication.GetModuleBasicManager().DefaultGenesis()

	authGenesis := auth.NewGenesisState(auth.DefaultParams(), accounts)
	genesisStateBz := simulationApplication.Codec().MustMarshalJSON(authGenesis)
	genesisState[auth.ModuleName] = genesisStateBz

	stateBytes, err := codec.MarshalJSONIndent(simulationApplication.Codec(), genesisState)
	if err != nil {
		panic(err)
	}

	// Initialize the chain
	newSimulationApplication.InitChain(
		abciTypes.RequestInitChain{
			Validators:    []abciTypes.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)

	newSimulationApplication.Commit()
	newSimulationApplication.BeginBlock(abciTypes.RequestBeginBlock{Header: abciTypes.Header{Height: simulationApplication.application.LastBlockHeight() + 1}})

	return newSimulationApplication
}

func (simulationApplication SimulationApplication) NewTestApplication(isCheckTx bool) (applications.SimulationApplication, sdkTypes.Context) {
	app := simulationApplication.Setup(isCheckTx)
	ctx := simulationApplication.GetBaseApp().NewContext(isCheckTx, abciTypes.Header{})

	return app, ctx
}

func NewSimulationApplication(name string, moduleBasicManager module.BasicManager, enabledWasmProposalTypeList []wasm.ProposalType, moduleAccountPermissions map[string][]string, tokenReceiveAllowedModules map[string]bool) applications.SimulationApplication {
	return &SimulationApplication{
		application: application{
			name:                        name,
			moduleBasicManager:          moduleBasicManager,
			codec:                       makeCodec(moduleBasicManager),
			enabledWasmProposalTypeList: enabledWasmProposalTypeList,
			moduleAccountPermissions:    moduleAccountPermissions,
			tokenReceiveAllowedModules:  tokenReceiveAllowedModules,
		},
	}
}
