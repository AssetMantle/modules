// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	serverTypes "github.com/cosmos/cosmos-sdk/server/types"
	authKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzKeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilityKeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	crisisKeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	distributionKeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	evidenceKeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidenceTypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	feeGrantKeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	govKeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	mintKeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	mintTypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingKeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradeKeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradeTypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	icaHostKeeper "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/host/keeper"
	ibcTransferKeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	ibcKeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"
	"io"
	"os"
	"path/filepath"
	"testing"

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
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/modules/assets"
	"github.com/AssetMantle/modules/modules/classifications"
	"github.com/AssetMantle/modules/modules/identities"
	"github.com/AssetMantle/modules/modules/maintainers"
	"github.com/AssetMantle/modules/modules/metas"
	"github.com/AssetMantle/modules/modules/orders"
	"github.com/AssetMantle/modules/modules/splits"
	wasmUtilities "github.com/AssetMantle/modules/utilities/wasm"

	"github.com/AssetMantle/modules/schema/applications"
)

type SimulationApplication struct {
	application

	transientStoreKeys map[string]*sdkTypes.TransientStoreKey
	subspaces          map[string]paramsTypes.Subspace
	simulationManager  *module.SimulationManager

	// keepers
	AccountKeeper    authKeeper.AccountKeeper
	BankKeeper       bankKeeper.Keeper
	CapabilityKeeper *capabilityKeeper.Keeper
	StakingKeeper    stakingKeeper.Keeper
	SlashingKeeper   slashingKeeper.Keeper
	MintKeeper       mintKeeper.Keeper
	DistrKeeper      distributionKeeper.Keeper
	GovKeeper        govKeeper.Keeper
	CrisisKeeper     crisisKeeper.Keeper
	UpgradeKeeper    upgradeKeeper.Keeper
	ParamsKeeper     paramsKeeper.Keeper
	// IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	IBCKeeper      *ibcKeeper.Keeper
	ICAHostKeeper  icaHostKeeper.Keeper
	EvidenceKeeper evidenceKeeper.Keeper
	TransferKeeper ibcTransferKeeper.Keeper
	FeeGrantKeeper feeGrantKeeper.Keeper
	AuthzKeeper    authzKeeper.Keeper

	// make scoped Keepers public for test purposes
	ScopedIBCKeeper      capabilityKeeper.ScopedKeeper
	ScopedTransferKeeper capabilityKeeper.ScopedKeeper
	ScopedICAHostKeeper  capabilityKeeper.ScopedKeeper
}

var _ applications.SimulationApplication = (*SimulationApplication)(nil)

func (simulationApplication SimulationApplication) LegacyAmino() *codec.LegacyAmino {
	return simulationApplication.codec.GetLegacyAmino()
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
func (simulationApplication SimulationApplication) ExportAppStateAndValidators(forZeroHeight bool, jailAllowedAddrs []string) (serverTypes.ExportedApp, error) {
	return simulationApplication.ExportApplicationStateAndValidators(forZeroHeight, jailAllowedAddrs)
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
func (simulationApplication SimulationApplication) GetSubspace(moduleName string) paramsTypes.Subspace {
	return simulationApplication.subspaces[moduleName]
}
func (simulationApplication SimulationApplication) GetModuleAccountPermissions() map[string][]string {
	return simulationApplication.moduleAccountPermissions
}
func (simulationApplication SimulationApplication) GetBlackListedAddresses() map[string]bool {
	blacklistedAddrs := make(map[string]bool)
	for acc := range simulationApplication.moduleAccountPermissions {
		blacklistedAddrs[authTypes.NewModuleAddress(acc).String()] = !simulationApplication.tokenReceiveAllowedModules[acc]
	}

	return blacklistedAddrs
}
func (simulationApplication SimulationApplication) CheckBalance(t *testing.T, address sdkTypes.AccAddress, coins sdkTypes.Coins) {
	ctxCheck := simulationApplication.BaseApp.NewContext(true, abciTypes.Header{})
	res := simulationApplication.BankKeeper.GetAllBalances(ctxCheck, address)

	require.True(t, coins.IsEqual(res))
}
func (simulationApplication SimulationApplication) AddTestAddresses(context sdkTypes.Context, accountNumber int, amount sdkTypes.Int) []sdkTypes.AccAddress {
	testAddresses := make([]sdkTypes.AccAddress, accountNumber)

	for i := 0; i < accountNumber; i++ {
		pk := ed25519.GenPrivKey().PubKey()
		testAddresses[i] = sdkTypes.AccAddress(pk.Address())
	}

	initCoins := sdkTypes.NewCoins(sdkTypes.NewCoin(simulationApplication.StakingKeeper.BondDenom(context), amount))
	totalSupply := sdkTypes.NewCoins(sdkTypes.NewCoin(simulationApplication.StakingKeeper.BondDenom(context), amount.MulRaw(int64(len(testAddresses)))))
	prevSupply := simulationApplication.BankKeeper.GetSupply(context, simulationApplication.StakingKeeper.BondDenom(context))
	simulationApplication.BankKeeper.SetSupply(context, supply.NewSupply(prevSupply.GetTotal().Add(totalSupply...)))

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

		stateBytes, err := codec.MarshalJSONIndent(simulationApplication.GetCodec(), genesisState)
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
	genesisState := simulationApplication.GetModuleBasicManager().DefaultGenesis(simulationApplication.GetCodec())

	authGenesis := authTypes.NewGenesisState(authTypes.DefaultParams(), accounts)
	genesisStateBz := simulationApplication.GetCodec().MustMarshalJSON(authGenesis)
	genesisState[authTypes.ModuleName] = genesisStateBz

	stateBytes, err := codec.MarshalJSONIndent(simulationApplication.GetCodec().GetLegacyAmino(), genesisState)
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
	newSimulationApplication.BeginBlock(abciTypes.RequestBeginBlock{Header: abciTypes.Header{Height: simulationApplication.application.BaseApp.LastBlockHeight() + 1}})

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
	simulationApplication.application = *simulationApplication.Initialize(logger, db, traceStore, loadLatest, invCheckPeriod, skipUpgradeHeights, home, baseAppOptions...).(*application)

	simulationApplication.transientStoreKeys = sdkTypes.NewTransientStoreKeys(paramsTypes.TStoreKey)

	simulationApplication.ParamsKeeper = paramsKeeper.NewKeeper(
		simulationApplication.codec,
		simulationApplication.codec.GetLegacyAmino(),
		simulationApplication.keys[paramsTypes.StoreKey],
		simulationApplication.transientStoreKeys[paramsTypes.TStoreKey],
	)

	simulationApplication.AccountKeeper = authKeeper.NewAccountKeeper(
		simulationApplication.codec,
		simulationApplication.keys[authTypes.StoreKey],
		simulationApplication.ParamsKeeper.Subspace(authTypes.ModuleName),
		authTypes.ProtoBaseAccount,
		simulationApplication.moduleAccountPermissions,
	)

	blacklistedAddresses := make(map[string]bool)
	for account := range simulationApplication.moduleAccountPermissions {
		blacklistedAddresses[authTypes.NewModuleAddress(account).String()] = !simulationApplication.tokenReceiveAllowedModules[account]
	}

	simulationApplication.BankKeeper = bankKeeper.NewBaseKeeper(
		simulationApplication.codec,
		simulationApplication.keys[bankTypes.StoreKey],
		simulationApplication.AccountKeeper,
		simulationApplication.ParamsKeeper.Subspace(bankTypes.ModuleName),
		blacklistedAddresses,
	)

	simulationApplication.StakingKeeper = stakingKeeper.NewKeeper(
		simulationApplication.codec,
		simulationApplication.keys[stakingTypes.StoreKey],
		simulationApplication.AccountKeeper,
		simulationApplication.BankKeeper,
		simulationApplication.ParamsKeeper.Subspace(stakingTypes.ModuleName),
	)

	simulationApplication.MintKeeper = mintKeeper.NewKeeper(
		simulationApplication.codec,
		simulationApplication.keys[mintTypes.StoreKey],
		simulationApplication.ParamsKeeper.Subspace(mintTypes.ModuleName),
		&simulationApplication.StakingKeeper,
		simulationApplication.AccountKeeper,
		simulationApplication.BankKeeper,
		authTypes.FeeCollectorName,
	)

	blackListedModuleAddresses := make(map[string]bool)
	for moduleAccount := range simulationApplication.moduleAccountPermissions {
		blackListedModuleAddresses[authTypes.NewModuleAddress(moduleAccount).String()] = true
	}

	simulationApplication.DistributionKeeper = simulationApplication.distributionKeeper

	simulationApplication.SlashingKeeper = simulationApplication.slashingKeeper

	simulationApplication.CrisisKeeper = simulationApplication.crisisKeeper

	simulationApplication.UpgradeKeeper = upgradeKeeper.NewKeeper(
		skipUpgradeHeights,
		simulationApplication.keys[upgradeTypes.StoreKey],
		simulationApplication.codec,
		simulationApplication.GetDefaultNodeHome(),
	)

	simulationApplication.EvidenceKeeper = *evidenceKeeper.NewKeeper(
		simulationApplication.codec,
		simulationApplication.keys[evidenceTypes.StoreKey],
		&simulationApplication.StakingKeeper,
		simulationApplication.SlashingKeeper,
	)

	evidenceRouter := evidenceTypes.NewRouter()
	simulationApplication.EvidenceKeeper.SetRouter(evidenceRouter)

	simulationApplication.StakingKeeper = simulationApplication.stakingKeeper

	var wasmRouter = simulationApplication.BaseApp.Router()

	wasmDir := filepath.Join(home, wasm.ModuleName)

	wasmWrap := struct {
		Wasm wasm.Config `mapstructure:"wasm"`
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
		simulationApplication.ParamsKeeper.Subspace(wasm.ModuleName),
		simulationApplication.AccountKeeper,
		simulationApplication.BankKeeper,
		simulationApplication.StakingKeeper,
		simulationApplication.DistributionKeeper,
		simulationApplication.
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
