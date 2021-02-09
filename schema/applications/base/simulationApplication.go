/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/spf13/viper"
	tendermintOS "github.com/tendermint/tendermint/libs/os"

	wasmClient "github.com/CosmWasm/wasmd/x/wasm/client"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	paramsClient "github.com/cosmos/cosmos-sdk/x/params/client"
	upgradeClient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	"github.com/persistenceOne/persistenceSDK/modules/assets"
	"github.com/persistenceOne/persistenceSDK/modules/classifications"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/conform"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/define"
	"github.com/persistenceOne/persistenceSDK/modules/identities"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/maintain"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/super"
	"github.com/persistenceOne/persistenceSDK/modules/metas"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/orders"
	"github.com/persistenceOne/persistenceSDK/modules/splits"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/burn"
	splitsMint "github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/mint"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/renumerate"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/transfer"
	wasmUtilities "github.com/persistenceOne/persistenceSDK/utilities/wasm"

	authVesting "github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/persistenceOne/persistenceSDK/schema"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/cosmos/cosmos-sdk/simapp"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"
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
	sdkTypesModule "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/schema/applications"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	tendermintTypes "github.com/tendermint/tendermint/types"
)

type simulationApplication struct {
	application        *application
	transientStoreKeys map[string]*sdkTypes.TransientStoreKey
	sm                 *module.SimulationManager
	subspaces          map[string]params.Subspace

	moduleAddressPermissions   map[string][]string
	tokenReceiveAllowedModules map[string]bool

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

var _ applications.SimulationApplication = (*simulationApplication)(nil)

func (simulationApplication simulationApplication) Info(info abciTypes.RequestInfo) abciTypes.ResponseInfo {
	return simulationApplication.application.baseApp.Info(info)
}

func (simulationApplication simulationApplication) SetOption(option abciTypes.RequestSetOption) abciTypes.ResponseSetOption {
	return simulationApplication.application.baseApp.SetOption(option)
}

func (simulationApplication simulationApplication) Query(query abciTypes.RequestQuery) abciTypes.ResponseQuery {
	return simulationApplication.application.baseApp.Query(query)
}

func (simulationApplication simulationApplication) CheckTx(tx abciTypes.RequestCheckTx) abciTypes.ResponseCheckTx {
	return simulationApplication.application.baseApp.CheckTx(tx)
}

func (simulationApplication simulationApplication) InitChain(chain abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
	return simulationApplication.application.baseApp.InitChain(chain)
}

func (simulationApplication simulationApplication) BeginBlock(block abciTypes.RequestBeginBlock) abciTypes.ResponseBeginBlock {
	return simulationApplication.application.baseApp.BeginBlock(block)
}

func (simulationApplication simulationApplication) DeliverTx(tx abciTypes.RequestDeliverTx) abciTypes.ResponseDeliverTx {
	return simulationApplication.application.baseApp.DeliverTx(tx)
}

func (simulationApplication simulationApplication) EndBlock(block abciTypes.RequestEndBlock) abciTypes.ResponseEndBlock {
	return simulationApplication.application.baseApp.EndBlock(block)
}

func (simulationApplication simulationApplication) Commit() abciTypes.ResponseCommit {
	return simulationApplication.application.baseApp.Commit()
}

func (simulationApplication simulationApplication) LoadHeight(i int64) error {
	return simulationApplication.application.LoadHeight(i)
}

func (simulationApplication simulationApplication) ExportApplicationStateAndValidators(b bool, strings []string) (json.RawMessage, []tendermintTypes.GenesisValidator, error) {
	return simulationApplication.application.ExportApplicationStateAndValidators(b, strings)
}

func (simulationApplication simulationApplication) Initialize(applicationName string, codec *codec.Codec, enabledProposals []wasm.ProposalType, moduleAccountPermissions map[string][]string, tokenReceiveAllowedModules map[string]bool, logger log.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string, baseAppOptions ...func(*baseapp.BaseApp)) applications.Application {

	//simulationApplication.moduleAddressPermissions = moduleAccountPermissions
	//simulationApplication.tokenReceiveAllowedModules = tokenReceiveAllowedModules
	baseApp := baseapp.NewBaseApp(
		applicationName,
		logger,
		db,
		auth.DefaultTxDecoder(codec),
		baseAppOptions...,
	)
	baseApp.SetCommitMultiStoreTracer(traceStore)
	baseApp.SetAppVersion(version.Version)

	simulationApplication.application.keys = sdkTypes.NewKVStoreKeys(
		baseapp.MainStoreKey,
		auth.StoreKey,
		supply.StoreKey,
		staking.StoreKey,
		mint.StoreKey,
		distribution.StoreKey,
		slashing.StoreKey,
		gov.StoreKey,
		params.StoreKey,
		upgrade.StoreKey,
		evidence.StoreKey,
		wasm.StoreKey,
		assets.Prototype().Name(),
		classifications.Prototype().Name(),
		identities.Prototype().Name(),
		maintainers.Prototype().Name(),
		metas.Prototype().Name(),
		orders.Prototype().Name(),
		splits.Prototype().Name(),
	)

	simulationApplication.transientStoreKeys = sdkTypes.NewTransientStoreKeys(params.TStoreKey)

	simulationApplication.application.baseApp = baseApp
	simulationApplication.application.codec = codec

	simulationApplication.ParamsKeeper = params.NewKeeper(
		codec,
		simulationApplication.application.keys[params.StoreKey],
		simulationApplication.transientStoreKeys[params.TStoreKey],
	)

	simulationApplication.AccountKeeper = auth.NewAccountKeeper(
		codec,
		simulationApplication.application.keys[auth.StoreKey],
		simulationApplication.ParamsKeeper.Subspace(auth.DefaultParamspace),
		auth.ProtoBaseAccount,
	)

	blacklistedAddresses := make(map[string]bool)
	for account := range moduleAccountPermissions {
		blacklistedAddresses[supply.NewModuleAddress(account).String()] = !tokenReceiveAllowedModules[account]
	}

	simulationApplication.BankKeeper = bank.NewBaseKeeper(
		simulationApplication.AccountKeeper,
		simulationApplication.ParamsKeeper.Subspace(bank.DefaultParamspace),
		blacklistedAddresses,
	)

	simulationApplication.SupplyKeeper = supply.NewKeeper(
		codec,
		simulationApplication.application.keys[supply.StoreKey],
		simulationApplication.AccountKeeper,
		simulationApplication.BankKeeper,
		moduleAccountPermissions,
	)

	simulationApplication.application.stakingKeeper = staking.NewKeeper(
		codec,
		simulationApplication.application.keys[staking.StoreKey],
		simulationApplication.SupplyKeeper,
		simulationApplication.ParamsKeeper.Subspace(staking.DefaultParamspace),
	)
	simulationApplication.StakingKeeper = simulationApplication.application.stakingKeeper

	simulationApplication.MintKeeper = mint.NewKeeper(
		codec,
		simulationApplication.application.keys[mint.StoreKey],
		simulationApplication.ParamsKeeper.Subspace(mint.DefaultParamspace),
		&simulationApplication.StakingKeeper,
		simulationApplication.SupplyKeeper,
		auth.FeeCollectorName,
	)

	blackListedModuleAddresses := make(map[string]bool)
	for moduleAccount := range moduleAccountPermissions {
		blackListedModuleAddresses[supply.NewModuleAddress(moduleAccount).String()] = true
	}

	simulationApplication.application.distributionKeeper = distribution.NewKeeper(
		codec,
		simulationApplication.application.keys[distribution.StoreKey],
		simulationApplication.ParamsKeeper.Subspace(distribution.DefaultParamspace),
		&simulationApplication.application.stakingKeeper,
		simulationApplication.SupplyKeeper,
		auth.FeeCollectorName,
		blackListedModuleAddresses,
	)
	simulationApplication.DistributionKeeper = simulationApplication.application.distributionKeeper

	simulationApplication.application.slashingKeeper = slashing.NewKeeper(
		codec,
		simulationApplication.application.keys[slashing.StoreKey],
		&simulationApplication.application.stakingKeeper,
		simulationApplication.ParamsKeeper.Subspace(slashing.DefaultParamspace),
	)
	simulationApplication.SlashingKeeper = simulationApplication.application.slashingKeeper

	simulationApplication.application.crisisKeeper = crisis.NewKeeper(
		simulationApplication.ParamsKeeper.Subspace(crisis.DefaultParamspace),
		invCheckPeriod,
		simulationApplication.SupplyKeeper,
		auth.FeeCollectorName,
	)
	simulationApplication.CrisisKeeper = simulationApplication.application.crisisKeeper

	simulationApplication.UpgradeKeeper = upgrade.NewKeeper(
		skipUpgradeHeights,
		simulationApplication.application.keys[upgrade.StoreKey],
		codec,
	)

	evidenceKeeper := evidence.NewKeeper(
		codec,
		simulationApplication.application.keys[evidence.StoreKey],
		simulationApplication.ParamsKeeper.Subspace(evidence.DefaultParamspace),
		&simulationApplication.application.stakingKeeper,
		simulationApplication.application.slashingKeeper,
	)
	evidenceRouter := evidence.NewRouter()
	evidenceKeeper.SetRouter(evidenceRouter)
	simulationApplication.EvidenceKeeper = *evidenceKeeper
	govRouter := gov.NewRouter()
	govRouter.AddRoute(
		gov.RouterKey,
		gov.ProposalHandler,
	).AddRoute(
		params.RouterKey,
		params.NewParamChangeProposalHandler(simulationApplication.ParamsKeeper),
	).AddRoute(
		distribution.RouterKey,
		distribution.NewCommunityPoolSpendProposalHandler(simulationApplication.application.distributionKeeper),
	).AddRoute(
		upgrade.RouterKey,
		upgrade.NewSoftwareUpgradeProposalHandler(simulationApplication.UpgradeKeeper),
	)

	simulationApplication.application.stakingKeeper = *simulationApplication.application.stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(simulationApplication.application.distributionKeeper.Hooks(), simulationApplication.application.slashingKeeper.Hooks()),
	)

	metasModule := metas.Prototype().Initialize(
		simulationApplication.application.keys[metas.Prototype().Name()],
		simulationApplication.ParamsKeeper.Subspace(metas.Prototype().Name()),
	)
	maintainersModule := maintainers.Prototype().Initialize(
		simulationApplication.application.keys[metas.Prototype().Name()],
		simulationApplication.ParamsKeeper.Subspace(maintainers.Prototype().Name()),
	)
	classificationsModule := classifications.Prototype().Initialize(
		simulationApplication.application.keys[classifications.Prototype().Name()],
		simulationApplication.ParamsKeeper.Subspace(classifications.Prototype().Name()),
		metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
	)
	identitiesModule := identities.Prototype().Initialize(
		simulationApplication.application.keys[identities.Prototype().Name()],
		simulationApplication.ParamsKeeper.Subspace(identities.Prototype().Name()),
		classificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
		classificationsModule.GetAuxiliary(define.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(super.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(maintain.Auxiliary.GetName()),
		metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
	)
	splitsModule := splits.Prototype().Initialize(
		simulationApplication.application.keys[splits.Prototype().Name()],
		simulationApplication.ParamsKeeper.Subspace(splits.Prototype().Name()),
		simulationApplication.SupplyKeeper,
		identitiesModule.GetAuxiliary(verify.Auxiliary.GetName()),
	)
	assetsModule := assets.Prototype().Initialize(
		simulationApplication.application.keys[assets.Prototype().Name()],
		simulationApplication.ParamsKeeper.Subspace(assets.Prototype().Name()),
		classificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
		classificationsModule.GetAuxiliary(define.Auxiliary.GetName()),
		identitiesModule.GetAuxiliary(verify.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(super.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(maintain.Auxiliary.GetName()),
		metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
		metasModule.GetAuxiliary(supplement.Auxiliary.GetName()),
		splitsModule.GetAuxiliary(splitsMint.Auxiliary.GetName()),
		splitsModule.GetAuxiliary(burn.Auxiliary.GetName()),
		splitsModule.GetAuxiliary(renumerate.Auxiliary.GetName()),
	)
	ordersModule := orders.Prototype().Initialize(
		simulationApplication.application.keys[orders.Prototype().Name()],
		simulationApplication.ParamsKeeper.Subspace(orders.Prototype().Name()),
		classificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
		classificationsModule.GetAuxiliary(define.Auxiliary.GetName()),
		identitiesModule.GetAuxiliary(verify.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(super.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(maintain.Auxiliary.GetName()),
		metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
		metasModule.GetAuxiliary(supplement.Auxiliary.GetName()),
		splitsModule.GetAuxiliary(splitsMint.Auxiliary.GetName()),
		splitsModule.GetAuxiliary(transfer.Auxiliary.GetName()),
	)

	var wasmRouter = baseApp.Router()

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
		codec,
		simulationApplication.application.keys[wasm.StoreKey],
		simulationApplication.ParamsKeeper.Subspace(wasm.DefaultParamspace),
		simulationApplication.AccountKeeper,
		simulationApplication.BankKeeper,
		simulationApplication.application.stakingKeeper,
		wasmRouter,
		wasmDir,
		wasmConfig,
		staking.ModuleName,
		&wasm.MessageEncoders{Custom: wasmUtilities.CustomEncoder(assets.Prototype(), classifications.Prototype(), identities.Prototype(), maintainers.Prototype(), metas.Prototype(), orders.Prototype(), splits.Prototype())},
		nil)

	if len(enabledProposals) != 0 {
		govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(wasmKeeper, enabledProposals))
	}

	simulationApplication.GovKeeper = gov.NewKeeper(
		codec,
		simulationApplication.application.keys[gov.StoreKey],
		simulationApplication.ParamsKeeper.Subspace(gov.DefaultParamspace).WithKeyTable(gov.ParamKeyTable()),
		simulationApplication.SupplyKeeper,
		&simulationApplication.application.stakingKeeper,
		govRouter,
	)

	simulationApplication.application.moduleManager = sdkTypesModule.NewManager(
		genutil.NewAppModule(simulationApplication.AccountKeeper, simulationApplication.application.stakingKeeper, simulationApplication.application.baseApp.DeliverTx),
		auth.NewAppModule(simulationApplication.AccountKeeper),
		bank.NewAppModule(simulationApplication.BankKeeper, simulationApplication.AccountKeeper),
		crisis.NewAppModule(&simulationApplication.application.crisisKeeper),
		supply.NewAppModule(simulationApplication.SupplyKeeper, simulationApplication.AccountKeeper),
		gov.NewAppModule(simulationApplication.GovKeeper, simulationApplication.AccountKeeper, simulationApplication.SupplyKeeper),
		mint.NewAppModule(simulationApplication.MintKeeper),
		slashing.NewAppModule(simulationApplication.application.slashingKeeper, simulationApplication.AccountKeeper, simulationApplication.application.stakingKeeper),
		distribution.NewAppModule(simulationApplication.application.distributionKeeper, simulationApplication.AccountKeeper, simulationApplication.SupplyKeeper, simulationApplication.application.stakingKeeper),
		staking.NewAppModule(simulationApplication.application.stakingKeeper, simulationApplication.AccountKeeper, simulationApplication.SupplyKeeper),
		upgrade.NewAppModule(simulationApplication.UpgradeKeeper),
		wasm.NewAppModule(wasmKeeper),
		evidence.NewAppModule(*evidenceKeeper),

		assetsModule,
		classificationsModule,
		identitiesModule,
		maintainersModule,
		metasModule,
		ordersModule,
		splitsModule,
	)

	simulationApplication.application.moduleManager.SetOrderBeginBlockers(
		upgrade.ModuleName,
		mint.ModuleName,
		distribution.ModuleName,
		slashing.ModuleName,
		ordersModule.Name(),
	)
	simulationApplication.application.moduleManager.SetOrderEndBlockers(
		crisis.ModuleName,
		gov.ModuleName,
		staking.ModuleName,
	)
	simulationApplication.application.moduleManager.SetOrderInitGenesis(
		auth.ModuleName,
		distribution.ModuleName,
		staking.ModuleName,
		bank.ModuleName,
		slashing.ModuleName,
		gov.ModuleName,
		mint.ModuleName,
		supply.ModuleName,
		crisis.ModuleName,
		genutil.ModuleName,
		evidence.ModuleName,
		wasm.ModuleName,
		assets.Prototype().Name(),
		classifications.Prototype().Name(),
		identities.Prototype().Name(),
		maintainers.Prototype().Name(),
		metas.Prototype().Name(),
		orders.Prototype().Name(),
		splits.Prototype().Name(),
	)
	simulationApplication.application.moduleManager.RegisterInvariants(&simulationApplication.application.crisisKeeper)
	simulationApplication.application.moduleManager.RegisterRoutes(simulationApplication.application.baseApp.Router(), simulationApplication.application.baseApp.QueryRouter())

	simulationApplication.sm = sdkTypesModule.NewSimulationManager(
		auth.NewAppModule(simulationApplication.AccountKeeper),
		bank.NewAppModule(simulationApplication.BankKeeper, simulationApplication.AccountKeeper),
		supply.NewAppModule(simulationApplication.SupplyKeeper, simulationApplication.AccountKeeper),
		gov.NewAppModule(simulationApplication.GovKeeper, simulationApplication.AccountKeeper, simulationApplication.SupplyKeeper),
		mint.NewAppModule(simulationApplication.MintKeeper),
		staking.NewAppModule(simulationApplication.application.stakingKeeper, simulationApplication.AccountKeeper, simulationApplication.SupplyKeeper),
		distribution.NewAppModule(simulationApplication.application.distributionKeeper, simulationApplication.AccountKeeper, simulationApplication.SupplyKeeper, simulationApplication.application.stakingKeeper),
		slashing.NewAppModule(simulationApplication.application.slashingKeeper, simulationApplication.AccountKeeper, simulationApplication.application.stakingKeeper),
		params.NewAppModule(),
		assets.Prototype(),
		classifications.Prototype(),
		identities.Prototype(),
		maintainers.Prototype(),
		metas.Prototype(),
		orders.Prototype(),
		splits.Prototype(),
	)

	simulationApplication.sm.RegisterStoreDecoders()

	simulationApplication.application.baseApp.MountKVStores(simulationApplication.application.keys)
	simulationApplication.application.baseApp.MountTransientStores(simulationApplication.transientStoreKeys)

	simulationApplication.application.baseApp.SetBeginBlocker(simulationApplication.application.moduleManager.BeginBlock)
	simulationApplication.application.baseApp.SetEndBlocker(simulationApplication.application.moduleManager.EndBlock)
	simulationApplication.application.baseApp.SetInitChainer(func(context sdkTypes.Context, requestInitChain abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
		var genesisState map[string]json.RawMessage
		codec.MustUnmarshalJSON(requestInitChain.AppStateBytes, &genesisState)
		return simulationApplication.application.moduleManager.InitGenesis(context, genesisState)
	})
	simulationApplication.application.baseApp.SetAnteHandler(auth.NewAnteHandler(simulationApplication.AccountKeeper, simulationApplication.SupplyKeeper, ante.DefaultSigVerificationGasConsumer))

	if loadLatest {
		err := simulationApplication.application.baseApp.LoadLatestVersion(simulationApplication.application.keys[baseapp.MainStoreKey])
		if err != nil {
			tendermintOS.Exit(err.Error())
		}
	}

	return simulationApplication
}

func (simulationApplication simulationApplication) Name() string {
	return simulationApplication.application.baseApp.Name()
}

func (simulationApplication simulationApplication) Codec() *codec.Codec {
	return simulationApplication.application.codec
}

func (simulationApplication simulationApplication) BeginBlocker(ctx sdkTypes.Context, req abciTypes.RequestBeginBlock) abciTypes.ResponseBeginBlock {
	return simulationApplication.application.moduleManager.BeginBlock(ctx, req)
}

func (simulationApplication simulationApplication) EndBlocker(ctx sdkTypes.Context, req abciTypes.RequestEndBlock) abciTypes.ResponseEndBlock {
	return simulationApplication.application.moduleManager.EndBlock(ctx, req)
}

func (simulationApplication simulationApplication) InitChainer(ctx sdkTypes.Context, req abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
	var genesisState simapp.GenesisState
	simulationApplication.application.codec.MustUnmarshalJSON(req.AppStateBytes, &genesisState)
	return simulationApplication.application.moduleManager.InitGenesis(ctx, genesisState)
}

func (simulationApplication simulationApplication) ExportAppStateAndValidators(forZeroHeight bool, jailWhiteList []string) (json.RawMessage, []tendermintTypes.GenesisValidator, error) {
	panic("implement me")
}

func (simulationApplication simulationApplication) ModuleAccountAddrs() map[string]bool {
	panic("implement me")
}

func (simulationApplication simulationApplication) SimulationManager() *module.SimulationManager {
	return simulationApplication.sm
}

func (simulationApplication simulationApplication) GetBaseApp() *baseapp.BaseApp {
	return simulationApplication.application.baseApp
}

func (simulationApplication simulationApplication) GetKey(storeKey string) *sdkTypes.KVStoreKey {
	return simulationApplication.application.keys[storeKey]
}

func (simulationApplication simulationApplication) GetTKey(storeKey string) *sdkTypes.TransientStoreKey {
	return simulationApplication.transientStoreKeys[storeKey]
}

func (simulationApplication simulationApplication) GetSubspace(moduleName string) params.Subspace {
	return simulationApplication.subspaces[moduleName]
}

func (simulationApplication simulationApplication) GetMaccPerms() map[string][]string {
	return simulationApplication.moduleAddressPermissions
}

func (simulationApplication simulationApplication) BlacklistedAccAddrs() map[string]bool {
	blacklistedAddrs := make(map[string]bool)
	for acc := range maccPerms {
		blacklistedAddrs[supply.NewModuleAddress(acc).String()] = !allowedReceivingModAcc[acc]
	}

	return blacklistedAddrs
}

func (simulationApplication simulationApplication) CheckBalance(t *testing.T, address sdkTypes.AccAddress, coins sdkTypes.Coins) {
	ctxCheck := simulationApplication.application.baseApp.NewContext(true, abciTypes.Header{})
	res := simulationApplication.AccountKeeper.GetAccount(ctxCheck, address)

	require.True(t, coins.IsEqual(res.GetCoins()))
}

func (simulationApplication simulationApplication) AddTestAddresses(context sdkTypes.Context, i int, s sdkTypes.Int) []sdkTypes.AccAddress {
	testAddrs := make([]sdkTypes.AccAddress, i)
	for i := 0; i < i; i++ {
		pk := ed25519.GenPrivKey().PubKey()
		testAddrs[i] = sdkTypes.AccAddress(pk.Address())
	}

	initCoins := sdkTypes.NewCoins(sdkTypes.NewCoin(simulationApplication.StakingKeeper.BondDenom(context), s))
	totalSupply := sdkTypes.NewCoins(sdkTypes.NewCoin(simulationApplication.StakingKeeper.BondDenom(context), s.MulRaw(int64(len(testAddrs)))))
	prevSupply := simulationApplication.SupplyKeeper.GetSupply(context)
	simulationApplication.SupplyKeeper.SetSupply(context, supply.NewSupply(prevSupply.GetTotal().Add(totalSupply...)))

	// fill all the addresses with some coins, set the loose pool tokens simultaneously
	for _, addr := range testAddrs {
		_, err := simulationApplication.BankKeeper.AddCoins(context, addr, initCoins)
		if err != nil {
			panic(err)
		}
	}
	return testAddrs
}

func (simulationApplication simulationApplication) Setup(isCheckTx bool) applications.SimulationApplication {
	db := tendermintDB.NewMemDB()
	app := NewSimApp(log.NewNopLogger(), db, nil, true, 0, map[int64]bool{}, DefaultNodeHome)
	if !isCheckTx {
		// init chain must be called to stop deliverState from being nil
		genesisState := ModuleBasics.DefaultGenesis()
		stateBytes, err := codec.MarshalJSONIndent(simulationApplication.Codec(), genesisState)
		if err != nil {
			panic(err)
		}

		// Initialize the chain
		simulationApplication.InitChain(
			abciTypes.RequestInitChain{
				Validators:    []abciTypes.ValidatorUpdate{},
				AppStateBytes: stateBytes,
			},
		)
	}

	return app
}

func (simulationApplication simulationApplication) SetupWithGenesisAccounts(accounts []exported.GenesisAccount) applications.SimulationApplication {
	db := tendermintDB.NewMemDB()
	app := NewSimApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0, map[int64]bool{}, DefaultNodeHome)

	// initialize the chain with the passed in genesis accounts
	genesisState := ModuleBasics.DefaultGenesis()

	authGenesis := auth.NewGenesisState(auth.DefaultParams(), accounts)
	genesisStateBz := simulationApplication.Codec().MustMarshalJSON(authGenesis)
	genesisState[auth.ModuleName] = genesisStateBz

	stateBytes, err := codec.MarshalJSONIndent(simulationApplication.Codec(), genesisState)
	if err != nil {
		panic(err)
	}

	// Initialize the chain
	simulationApplication.InitChain(
		abciTypes.RequestInitChain{
			Validators:    []abciTypes.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)

	simulationApplication.Commit()
	simulationApplication.BeginBlock(abciTypes.RequestBeginBlock{Header: abciTypes.Header{Height: simulationApplication.application.baseApp.LastBlockHeight() + 1}})

	return app
}

func (simulationApplication simulationApplication) NewTestApplication(isCheckTx bool) (applications.SimulationApplication, sdkTypes.Context) {
	app := simulationApplication.Setup(isCheckTx)
	ctx := simulationApplication.GetBaseApp().NewContext(isCheckTx, abciTypes.Header{})
	return app, ctx
}

func NewSimApp(logger log.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string, baseAppOptions ...func(*baseapp.BaseApp)) simulationApplication {
	return simulationApplication{}
}

var (
	app             = "simapp"
	DefaultCLIHome  = os.ExpandEnv("$HOME/.simapp")
	DefaultNodeHome = os.ExpandEnv("$HOME/.simapp")

	ModuleBasics = module.NewBasicManager(
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distribution.AppModuleBasic{},
		gov.NewAppModuleBasic(append(wasmClient.ProposalHandlers, paramsClient.ProposalHandler, distribution.ProposalHandler, upgradeClient.ProposalHandler)...),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		wasm.AppModuleBasic{},
		slashing.AppModuleBasic{},
		supply.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},

		assets.Prototype(),
		classifications.Prototype(),
		identities.Prototype(),
		maintainers.Prototype(),
		metas.Prototype(),
		orders.Prototype(),
		splits.Prototype(),
	)

	// module account permissions
	maccPerms = map[string][]string{
		auth.FeeCollectorName:     nil,
		distribution.ModuleName:   nil,
		mint.ModuleName:           {supply.Minter},
		staking.BondedPoolName:    {supply.Burner, supply.Staking},
		staking.NotBondedPoolName: {supply.Burner, supply.Staking},
		gov.ModuleName:            {supply.Burner},
	}

	// module accounts that are allowed to receive tokens
	allowedReceivingModAcc = map[string]bool{
		distribution.ModuleName: true,
	}
)

func MakeCodec() *codec.Codec {
	Codec := codec.New()
	ModuleBasics.RegisterCodec(Codec)
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	authVesting.RegisterCodec(Codec)
	return Codec
}
