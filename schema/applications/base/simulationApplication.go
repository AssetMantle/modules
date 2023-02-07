// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/json"
	"github.com/AssetMantle/modules/modules/assets"
	"github.com/AssetMantle/modules/modules/classifications"
	"github.com/AssetMantle/modules/modules/identities"
	"github.com/AssetMantle/modules/modules/maintainers"
	"github.com/AssetMantle/modules/modules/metas"
	"github.com/AssetMantle/modules/modules/orders"
	"github.com/AssetMantle/modules/modules/splits"
	wasmUtilities "github.com/AssetMantle/modules/utilities/wasm"
	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
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
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/log"
	tendermintTypes "github.com/tendermint/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/AssetMantle/modules/schema/applications"
)

type SimulationApplication struct {
	Application

	transientStoreKeys map[string]*sdkTypes.TransientStoreKey
	subspaces          map[string]params.Subspace
	simulationManager  *module.SimulationManager

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
	return simulationApplication.codec
}

func (simulationApplication SimulationApplication) BeginBlocker(ctx sdkTypes.Context, req abciTypes.RequestBeginBlock) abciTypes.ResponseBeginBlock {
	return simulationApplication.moduleManager.BeginBlock(ctx, req)
}

func (simulationApplication SimulationApplication) EndBlocker(ctx sdkTypes.Context, req abciTypes.RequestEndBlock) abciTypes.ResponseEndBlock {
	return simulationApplication.moduleManager.EndBlock(ctx, req)
}

func (simulationApplication SimulationApplication) InitChainer(ctx sdkTypes.Context, req abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
	var genesisState simapp.GenesisState

	simulationApplication.codec.MustUnmarshalJSON(req.AppStateBytes, &genesisState)

	return simulationApplication.moduleManager.InitGenesis(ctx, genesisState)
}

func (simulationApplication SimulationApplication) ExportAppStateAndValidators(forZeroHeight bool, jailWhiteList []string) (json.RawMessage, []tendermintTypes.GenesisValidator, error) {
	return simulationApplication.ExportApplicationStateAndValidators(forZeroHeight, jailWhiteList)
}

func (simulationApplication SimulationApplication) ModuleAccountAddrs() map[string]bool {
	return simulationApplication.tokenReceiveAllowedModules
}

func (simulationApplication SimulationApplication) SimulationManager() *module.SimulationManager {
	return simulationApplication.simulationManager
}

func (simulationApplication SimulationApplication) ModuleManager() *module.Manager {
	return simulationApplication.moduleManager
}

func (simulationApplication SimulationApplication) GetBaseApp() *baseapp.BaseApp {
	return &simulationApplication.BaseApp
}

func (simulationApplication SimulationApplication) GetKey(storeKey string) *sdkTypes.KVStoreKey {
	return simulationApplication.keys[storeKey]
}

func (simulationApplication SimulationApplication) GetTKey(storeKey string) *sdkTypes.TransientStoreKey {
	return simulationApplication.transientStoreKeys[storeKey]
}

func (simulationApplication SimulationApplication) GetSubspace(moduleName string) params.Subspace {
	return simulationApplication.subspaces[moduleName]
}

func (simulationApplication SimulationApplication) GetModuleAccountPermissions() map[string][]string {
	return simulationApplication.moduleAccountPermissions
}

func (simulationApplication SimulationApplication) GetBlackListedAddresses() map[string]bool {
	blacklistedAddrs := make(map[string]bool)
	for acc := range simulationApplication.moduleAccountPermissions {
		blacklistedAddrs[supply.NewModuleAddress(acc).String()] = !simulationApplication.tokenReceiveAllowedModules[acc]
	}

	return blacklistedAddrs
}

func (simulationApplication SimulationApplication) CheckBalance(t *testing.T, address sdkTypes.AccAddress, coins sdkTypes.Coins) {
	ctxCheck := simulationApplication.BaseApp.NewContext(true, abciTypes.Header{})
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
	newSimulationApplication.BeginBlock(abciTypes.RequestBeginBlock{Header: abciTypes.Header{Height: simulationApplication.Application.BaseApp.LastBlockHeight() + 1}})

	return newSimulationApplication
}

func (simulationApplication SimulationApplication) NewTestApplication(isCheckTx bool) (applications.SimulationApplication, sdkTypes.Context) {
	app := simulationApplication.Setup(isCheckTx)
	ctx := simulationApplication.GetBaseApp().NewContext(isCheckTx, abciTypes.Header{})

	return app, ctx
}

func (simulationApplication SimulationApplication) InitializeSimulationApplication(logger log.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string, baseAppOptions ...func(*baseapp.BaseApp)) applications.SimulationApplication {
	cache := store.NewCommitKVStoreCacheManager()
	baseAppOptions = append(baseAppOptions, baseapp.SetInterBlockCache(cache), baseapp.SetMinGasPrices(viper.GetString("minimum-gas-prices")))
	simulationApplication.Application = *simulationApplication.Initialize(logger, db, traceStore, loadLatest, invCheckPeriod, skipUpgradeHeights, home, baseAppOptions...).(*Application)

	simulationApplication.transientStoreKeys = sdkTypes.NewTransientStoreKeys(params.TStoreKey)

	simulationApplication.ParamsKeeper = params.NewKeeper(
		simulationApplication.codec,
		simulationApplication.keys[params.StoreKey],
		simulationApplication.transientStoreKeys[params.TStoreKey],
	)

	simulationApplication.AccountKeeper = auth.NewAccountKeeper(
		simulationApplication.codec,
		simulationApplication.keys[auth.StoreKey],
		simulationApplication.ParamsKeeper.Subspace(auth.DefaultParamspace),
		auth.ProtoBaseAccount,
	)

	blacklistedAddresses := make(map[string]bool)
	for account := range simulationApplication.moduleAccountPermissions {
		blacklistedAddresses[supply.NewModuleAddress(account).String()] = !simulationApplication.tokenReceiveAllowedModules[account]
	}

	simulationApplication.BankKeeper = bank.NewBaseKeeper(
		simulationApplication.AccountKeeper,
		simulationApplication.ParamsKeeper.Subspace(bank.DefaultParamspace),
		blacklistedAddresses,
	)

	simulationApplication.SupplyKeeper = supply.NewKeeper(
		simulationApplication.codec,
		simulationApplication.keys[supply.StoreKey],
		simulationApplication.AccountKeeper,
		simulationApplication.BankKeeper,
		simulationApplication.moduleAccountPermissions,
	)

	simulationApplication.StakingKeeper = simulationApplication.stakingKeeper

	simulationApplication.MintKeeper = mint.NewKeeper(
		simulationApplication.codec,
		simulationApplication.keys[mint.StoreKey],
		simulationApplication.ParamsKeeper.Subspace(mint.DefaultParamspace),
		&simulationApplication.StakingKeeper,
		simulationApplication.SupplyKeeper,
		auth.FeeCollectorName,
	)

	blackListedModuleAddresses := make(map[string]bool)
	for moduleAccount := range simulationApplication.moduleAccountPermissions {
		blackListedModuleAddresses[supply.NewModuleAddress(moduleAccount).String()] = true
	}

	simulationApplication.DistributionKeeper = simulationApplication.distributionKeeper

	simulationApplication.SlashingKeeper = simulationApplication.slashingKeeper

	simulationApplication.CrisisKeeper = simulationApplication.crisisKeeper

	simulationApplication.UpgradeKeeper = upgrade.NewKeeper(
		skipUpgradeHeights,
		simulationApplication.keys[upgrade.StoreKey],
		simulationApplication.codec,
	)

	simulationApplication.EvidenceKeeper = *evidence.NewKeeper(
		simulationApplication.codec,
		simulationApplication.keys[evidence.StoreKey],
		simulationApplication.ParamsKeeper.Subspace(evidence.DefaultParamspace),
		&simulationApplication.StakingKeeper,
		simulationApplication.SlashingKeeper,
	)

	evidenceRouter := evidence.NewRouter()
	simulationApplication.EvidenceKeeper.SetRouter(evidenceRouter)

	simulationApplication.StakingKeeper = simulationApplication.stakingKeeper

	var wasmRouter = simulationApplication.BaseApp.Router()

	wasmDir := filepath.Join(home, wasm.ModuleName)

	wasmWrap := struct {
		Wasm wasm.WasmConfig `mapstructure:"wasm"`
	}{
		Wasm: wasm.DefaultWasmConfig(),
	}

	err := viper.Unmarshal(&wasmWrap)
	if err != nil {
		panic("error while reading wasm config: " + err.Error())
	}

	wasmConfig := wasmWrap.Wasm

	wasmKeeper := wasm.NewKeeper(
		simulationApplication.codec,
		simulationApplication.keys[wasm.StoreKey],
		simulationApplication.ParamsKeeper.Subspace(wasm.DefaultParamspace),
		simulationApplication.AccountKeeper,
		simulationApplication.BankKeeper,
		simulationApplication.StakingKeeper,
		wasmRouter,
		wasmDir,
		wasmConfig,
		staking.ModuleName,
		&wasm.MessageEncoders{Custom: wasmUtilities.CustomEncoder(assets.Prototype(), classifications.Prototype(), identities.Prototype(), maintainers.Prototype(), metas.Prototype(), orders.Prototype(), splits.Prototype())},
		nil)

	govRouter := gov.NewRouter().AddRoute(
		gov.RouterKey,
		gov.ProposalHandler,
	).AddRoute(
		params.RouterKey,
		params.NewParamChangeProposalHandler(simulationApplication.ParamsKeeper),
	).AddRoute(
		distribution.RouterKey,
		distribution.NewCommunityPoolSpendProposalHandler(simulationApplication.DistributionKeeper),
	).AddRoute(
		upgrade.RouterKey,
		upgrade.NewSoftwareUpgradeProposalHandler(simulationApplication.UpgradeKeeper),
	)

	if len(simulationApplication.enabledWasmProposalTypeList) != 0 {
		govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(wasmKeeper, simulationApplication.enabledWasmProposalTypeList))
	}

	simulationApplication.GovKeeper = gov.NewKeeper(
		simulationApplication.codec,
		simulationApplication.keys[gov.StoreKey],
		simulationApplication.ParamsKeeper.Subspace(gov.DefaultParamspace).WithKeyTable(gov.ParamKeyTable()),
		simulationApplication.SupplyKeeper,
		&simulationApplication.StakingKeeper,
		govRouter,
	)

	simulationApplication.simulationManager = module.NewSimulationManager(
		auth.NewAppModule(simulationApplication.AccountKeeper),
		bank.NewAppModule(simulationApplication.BankKeeper, simulationApplication.AccountKeeper),
		supply.NewAppModule(simulationApplication.SupplyKeeper, simulationApplication.AccountKeeper),
		gov.NewAppModule(simulationApplication.GovKeeper, simulationApplication.AccountKeeper, simulationApplication.SupplyKeeper),
		mint.NewAppModule(simulationApplication.MintKeeper),
		staking.NewAppModule(simulationApplication.StakingKeeper, simulationApplication.AccountKeeper, simulationApplication.SupplyKeeper),
		distribution.NewAppModule(simulationApplication.DistributionKeeper, simulationApplication.AccountKeeper, simulationApplication.SupplyKeeper, simulationApplication.StakingKeeper),
		slashing.NewAppModule(simulationApplication.SlashingKeeper, simulationApplication.AccountKeeper, simulationApplication.StakingKeeper),
		params.NewAppModule(),
		assets.Prototype(),
		classifications.Prototype(),
		identities.Prototype(),
		maintainers.Prototype(),
		metas.Prototype(),
		orders.Prototype(),
		splits.Prototype(),
	)

	return &simulationApplication
}

func NewSimulationApplication(name string, moduleBasicManager module.BasicManager, enabledWasmProposalTypeList []wasm.ProposalType, moduleAccountPermissions map[string][]string, tokenReceiveAllowedModules map[string]bool) applications.SimulationApplication {
	return &SimulationApplication{
		Application: Application{
			name:                        name,
			moduleBasicManager:          moduleBasicManager,
			codec:                       makeCodec(moduleBasicManager),
			enabledWasmProposalTypeList: enabledWasmProposalTypeList,
			moduleAccountPermissions:    moduleAccountPermissions,
			tokenReceiveAllowedModules:  tokenReceiveAllowedModules,
		},
	}
}
