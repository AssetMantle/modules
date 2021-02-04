package simapp

import (
	"io"
	"os"
	"path/filepath"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmClient "github.com/CosmWasm/wasmd/x/wasm/client"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsClient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
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
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	wasmUtilities "github.com/persistenceOne/persistenceSDK/utilities/wasm"
	"github.com/spf13/viper"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintOS "github.com/tendermint/tendermint/libs/os"
	tendermintDB "github.com/tendermint/tm-db"
)

const appName = "SimulationApplication"

var (
	// DefaultCLIHome default home directories for the application CLI
	DefaultCLIHome = os.ExpandEnv("$HOME/.simulationApplication")

	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome = os.ExpandEnv("$HOME/.simulationApplication")

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
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

		wasm.AppModuleBasic{},
		assets.Prototype(),
		classifications.Prototype(),
		identities.Prototype(),
		maintainers.Prototype(),
		metas.Prototype(),
		orders.Prototype(),
		splits.Prototype(),
	)

	// module account permissions
	moduleAccountPermissions = map[string][]string{
		auth.FeeCollectorName:     nil,
		distribution.ModuleName:   nil,
		mint.ModuleName:           {supply.Minter},
		staking.BondedPoolName:    {supply.Burner, supply.Staking},
		staking.NotBondedPoolName: {supply.Burner, supply.Staking},
		gov.ModuleName:            {supply.Burner},
		splits.Prototype().Name(): nil,
	}

	// module accounts that are allowed to receive tokens
	allowedReceivingModuleAccounts = map[string]bool{
		distribution.ModuleName: true,
	}
)

// MakeCodec - custom tx codec
func MakeCodec() *codec.Codec {
	var cdc = codec.New()

	ModuleBasics.RegisterCodec(cdc)
	vesting.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	schema.RegisterCodec(cdc)

	return cdc
}

// Verify app interface at compile time
var _ simapp.App = (*SimulationApplication)(nil)

// SimulationApplication extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type SimulationApplication struct {
	*baseapp.BaseApp
	cdc *codec.Codec

	invariantsCheckPeriod uint

	// keys to access the sub stores
	keys          map[string]*sdk.KVStoreKey
	transientKeys map[string]*sdk.TransientStoreKey

	// subspaces
	subspaces map[string]params.Subspace

	// keepers
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

	WasmKeeper wasm.Keeper

	AssetsModule          helpers.Module
	SplitsModule          helpers.Module
	IdentitiesModule      helpers.Module
	OrdersModule          helpers.Module
	ClassificationsModule helpers.Module
	MaintainersModule     helpers.Module
	MetasModule           helpers.Module

	// the module manager
	moduleManager *module.Manager

	// simulation manager
	simulationManager *module.SimulationManager
}

// NewSimApp returns a reference to an initialized SimulationApplication.
func NewSimApp(
	logger log.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, skipUpgradeHeights map[int64]bool,
	invariantsCheckPeriod uint, baseAppOptions ...func(*baseapp.BaseApp),
) *SimulationApplication {
	cdc := MakeCodec()

	baseApplication := baseapp.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...)
	baseApplication.SetCommitMultiStoreTracer(traceStore)
	baseApplication.SetAppVersion(version.Version)

	keys := sdk.NewKVStoreKeys(
		baseapp.MainStoreKey, auth.StoreKey, staking.StoreKey,
		supply.StoreKey, mint.StoreKey, distribution.StoreKey, slashing.StoreKey,
		gov.StoreKey, params.StoreKey, upgrade.StoreKey, evidence.StoreKey,
		wasm.StoreKey,
		assets.Prototype().Name(),
		classifications.Prototype().Name(),
		identities.Prototype().Name(),
		maintainers.Prototype().Name(),
		metas.Prototype().Name(),
		orders.Prototype().Name(),
		splits.Prototype().Name(),
	)
	transientStoreKeys := sdk.NewTransientStoreKeys(params.TStoreKey)

	application := &SimulationApplication{
		BaseApp:               baseApplication,
		cdc:                   cdc,
		invariantsCheckPeriod: invariantsCheckPeriod,
		keys:                  keys,
		transientKeys:         transientStoreKeys,
		subspaces:             make(map[string]params.Subspace),
	}

	// init params keeper and subspaces
	application.ParamsKeeper = params.NewKeeper(application.cdc, keys[params.StoreKey], transientStoreKeys[params.TStoreKey])
	application.subspaces[auth.ModuleName] = application.ParamsKeeper.Subspace(auth.DefaultParamspace)
	application.subspaces[bank.ModuleName] = application.ParamsKeeper.Subspace(bank.DefaultParamspace)
	application.subspaces[staking.ModuleName] = application.ParamsKeeper.Subspace(staking.DefaultParamspace)
	application.subspaces[mint.ModuleName] = application.ParamsKeeper.Subspace(mint.DefaultParamspace)
	application.subspaces[distribution.ModuleName] = application.ParamsKeeper.Subspace(distribution.DefaultParamspace)
	application.subspaces[slashing.ModuleName] = application.ParamsKeeper.Subspace(slashing.DefaultParamspace)
	application.subspaces[gov.ModuleName] = application.ParamsKeeper.Subspace(gov.DefaultParamspace).WithKeyTable(gov.ParamKeyTable())
	application.subspaces[crisis.ModuleName] = application.ParamsKeeper.Subspace(crisis.DefaultParamspace)
	application.subspaces[evidence.ModuleName] = application.ParamsKeeper.Subspace(evidence.DefaultParamspace)

	// add keepers
	application.AccountKeeper = auth.NewAccountKeeper(
		application.cdc, keys[auth.StoreKey], application.subspaces[auth.ModuleName], auth.ProtoBaseAccount,
	)
	application.BankKeeper = bank.NewBaseKeeper(
		application.AccountKeeper, application.subspaces[bank.ModuleName], application.BlacklistedAccAddrs(),
	)
	application.SupplyKeeper = supply.NewKeeper(
		application.cdc, keys[supply.StoreKey], application.AccountKeeper, application.BankKeeper, moduleAccountPermissions,
	)
	stakingKeeper := staking.NewKeeper(
		application.cdc, keys[staking.StoreKey], application.SupplyKeeper, application.subspaces[staking.ModuleName],
	)
	application.MintKeeper = mint.NewKeeper(
		application.cdc, keys[mint.StoreKey], application.subspaces[mint.ModuleName], &stakingKeeper,
		application.SupplyKeeper, auth.FeeCollectorName,
	)
	application.DistributionKeeper = distribution.NewKeeper(
		application.cdc, keys[distribution.StoreKey], application.subspaces[distribution.ModuleName], &stakingKeeper,
		application.SupplyKeeper, auth.FeeCollectorName, application.ModuleAccountAddrs(),
	)
	application.SlashingKeeper = slashing.NewKeeper(
		application.cdc, keys[slashing.StoreKey], &stakingKeeper, application.subspaces[slashing.ModuleName],
	)
	application.CrisisKeeper = crisis.NewKeeper(
		application.subspaces[crisis.ModuleName], invariantsCheckPeriod, application.SupplyKeeper, auth.FeeCollectorName,
	)
	application.UpgradeKeeper = upgrade.NewKeeper(skipUpgradeHeights, keys[upgrade.StoreKey], application.cdc)

	// create evidence keeper with router
	evidenceKeeper := evidence.NewKeeper(
		application.cdc, keys[evidence.StoreKey], application.subspaces[evidence.ModuleName], &application.StakingKeeper, application.SlashingKeeper,
	)
	evidenceRouter := evidence.NewRouter()
	// TODO: Register evidence routes. - from cosmos
	evidenceKeeper.SetRouter(evidenceRouter)
	application.EvidenceKeeper = *evidenceKeeper

	// register the proposal types
	govRouter := gov.NewRouter()
	govRouter.AddRoute(gov.RouterKey, gov.ProposalHandler).
		AddRoute(params.RouterKey, params.NewParamChangeProposalHandler(application.ParamsKeeper)).
		AddRoute(distribution.RouterKey, distribution.NewCommunityPoolSpendProposalHandler(application.DistributionKeeper)).
		AddRoute(upgrade.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(application.UpgradeKeeper))

	application.MetasModule = metas.Prototype().Initialize(
		keys[metas.Prototype().Name()],
		application.ParamsKeeper.Subspace(metas.Prototype().Name()),
	)
	application.MaintainersModule = maintainers.Prototype().Initialize(
		keys[metas.Prototype().Name()],
		application.ParamsKeeper.Subspace(maintainers.Prototype().Name()),
	)
	application.ClassificationsModule = classifications.Prototype().Initialize(
		keys[classifications.Prototype().Name()],
		application.ParamsKeeper.Subspace(classifications.Prototype().Name()),
		application.MetasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
	)
	application.IdentitiesModule = identities.Prototype().Initialize(
		keys[identities.Prototype().Name()],
		application.ParamsKeeper.Subspace(identities.Prototype().Name()),
		application.ClassificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
		application.ClassificationsModule.GetAuxiliary(define.Auxiliary.GetName()),
		application.MaintainersModule.GetAuxiliary(super.Auxiliary.GetName()),
		application.MaintainersModule.GetAuxiliary(maintain.Auxiliary.GetName()),
		application.MetasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
	)
	application.SplitsModule = splits.Prototype().Initialize(
		keys[splits.Prototype().Name()],
		application.ParamsKeeper.Subspace(splits.Prototype().Name()),
		application.SupplyKeeper,
		application.IdentitiesModule.GetAuxiliary(verify.Auxiliary.GetName()),
	)
	application.AssetsModule = assets.Prototype().Initialize(
		keys[assets.Prototype().Name()],
		application.ParamsKeeper.Subspace(assets.Prototype().Name()),
		application.ClassificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
		application.ClassificationsModule.GetAuxiliary(define.Auxiliary.GetName()),
		application.IdentitiesModule.GetAuxiliary(verify.Auxiliary.GetName()),
		application.MaintainersModule.GetAuxiliary(super.Auxiliary.GetName()),
		application.MaintainersModule.GetAuxiliary(maintain.Auxiliary.GetName()),
		application.MetasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
		application.MetasModule.GetAuxiliary(supplement.Auxiliary.GetName()),
		application.SplitsModule.GetAuxiliary(splitsMint.Auxiliary.GetName()),
		application.SplitsModule.GetAuxiliary(burn.Auxiliary.GetName()),
		application.SplitsModule.GetAuxiliary(renumerate.Auxiliary.GetName()),
	)
	application.OrdersModule = orders.Prototype().Initialize(
		keys[orders.Prototype().Name()],
		application.ParamsKeeper.Subspace(orders.Prototype().Name()),
		application.ClassificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
		application.ClassificationsModule.GetAuxiliary(define.Auxiliary.GetName()),
		application.IdentitiesModule.GetAuxiliary(verify.Auxiliary.GetName()),
		application.MaintainersModule.GetAuxiliary(super.Auxiliary.GetName()),
		application.MaintainersModule.GetAuxiliary(maintain.Auxiliary.GetName()),
		application.MetasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
		application.MetasModule.GetAuxiliary(supplement.Auxiliary.GetName()),
		application.SplitsModule.GetAuxiliary(splitsMint.Auxiliary.GetName()),
		application.SplitsModule.GetAuxiliary(transfer.Auxiliary.GetName()),
	)

	var wasmRouter = application.Router()

	wasmDir := filepath.Join(DefaultNodeHome, wasm.ModuleName)

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

	application.WasmKeeper = wasm.NewKeeper(
		cdc,
		keys[wasm.StoreKey],
		application.ParamsKeeper.Subspace(wasm.DefaultParamspace),
		application.AccountKeeper,
		application.BankKeeper,
		application.StakingKeeper,
		wasmRouter,
		wasmDir,
		wasmConfig,
		staking.ModuleName,
		&wasm.MessageEncoders{Custom: wasmUtilities.CustomEncoder(assets.Prototype(), classifications.Prototype(), identities.Prototype(), maintainers.Prototype(), metas.Prototype(), orders.Prototype(), splits.Prototype())},
		nil)

	if len(wasm.EnableAllProposals) != 0 {
		govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(application.WasmKeeper, wasm.EnableAllProposals))
	}

	application.GovKeeper = gov.NewKeeper(
		application.cdc, keys[gov.StoreKey], application.subspaces[gov.ModuleName], application.SupplyKeeper,
		&stakingKeeper, govRouter,
	)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	application.StakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(application.DistributionKeeper.Hooks(), application.SlashingKeeper.Hooks()),
	)

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	application.moduleManager = module.NewManager(
		genutil.NewAppModule(application.AccountKeeper, application.StakingKeeper, application.BaseApp.DeliverTx),
		auth.NewAppModule(application.AccountKeeper),
		bank.NewAppModule(application.BankKeeper, application.AccountKeeper),
		crisis.NewAppModule(&application.CrisisKeeper),
		supply.NewAppModule(application.SupplyKeeper, application.AccountKeeper),
		gov.NewAppModule(application.GovKeeper, application.AccountKeeper, application.SupplyKeeper),
		mint.NewAppModule(application.MintKeeper),
		slashing.NewAppModule(application.SlashingKeeper, application.AccountKeeper, application.StakingKeeper),
		distribution.NewAppModule(application.DistributionKeeper, application.AccountKeeper, application.SupplyKeeper, application.StakingKeeper),
		staking.NewAppModule(application.StakingKeeper, application.AccountKeeper, application.SupplyKeeper),
		upgrade.NewAppModule(application.UpgradeKeeper),
		evidence.NewAppModule(application.EvidenceKeeper),
		wasm.NewAppModule(application.WasmKeeper),
		application.AssetsModule,
		application.ClassificationsModule,
		application.IdentitiesModule,
		application.MaintainersModule,
		application.MetasModule,
		application.OrdersModule,
		application.SplitsModule,
	)

	// During begin block slashing happens after distribution.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	application.moduleManager.SetOrderBeginBlockers(upgrade.ModuleName, mint.ModuleName, distribution.ModuleName, slashing.ModuleName, evidence.ModuleName, application.OrdersModule.Name())
	application.moduleManager.SetOrderEndBlockers(crisis.ModuleName, gov.ModuleName, staking.ModuleName)

	// NOTE: The genesisUtilities module must occur after staking so that pools arel
	// properly initialized with tokens from genesis accounts.
	application.moduleManager.SetOrderInitGenesis(
		auth.ModuleName, distribution.ModuleName, staking.ModuleName, bank.ModuleName,
		slashing.ModuleName, gov.ModuleName, mint.ModuleName, supply.ModuleName,
		crisis.ModuleName, genutil.ModuleName, evidence.ModuleName,
		assets.Prototype().Name(),
		classifications.Prototype().Name(),
		identities.Prototype().Name(),
		maintainers.Prototype().Name(),
		metas.Prototype().Name(),
		orders.Prototype().Name(),
		splits.Prototype().Name(),
	)

	application.moduleManager.RegisterInvariants(&application.CrisisKeeper)
	application.moduleManager.RegisterRoutes(application.Router(), application.QueryRouter())

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	application.simulationManager = module.NewSimulationManager(
		auth.NewAppModule(application.AccountKeeper),
		bank.NewAppModule(application.BankKeeper, application.AccountKeeper),
		supply.NewAppModule(application.SupplyKeeper, application.AccountKeeper),
		gov.NewAppModule(application.GovKeeper, application.AccountKeeper, application.SupplyKeeper),
		mint.NewAppModule(application.MintKeeper),
		staking.NewAppModule(application.StakingKeeper, application.AccountKeeper, application.SupplyKeeper),
		distribution.NewAppModule(application.DistributionKeeper, application.AccountKeeper, application.SupplyKeeper, application.StakingKeeper),
		slashing.NewAppModule(application.SlashingKeeper, application.AccountKeeper, application.StakingKeeper),
		params.NewAppModule(), // NOTE: only used for simulation to generate randomized param change proposals
		assets.Prototype(),
		classifications.Prototype(),
		identities.Prototype(),
		maintainers.Prototype(),
		metas.Prototype(),
		orders.Prototype(),
		splits.Prototype(),
	)

	application.simulationManager.RegisterStoreDecoders()

	// initialize stores
	application.MountKVStores(keys)
	application.MountTransientStores(transientStoreKeys)

	// initialize BaseApp
	application.SetInitChainer(application.InitChainer)
	application.SetBeginBlocker(application.BeginBlocker)
	application.SetAnteHandler(ante.NewAnteHandler(application.AccountKeeper, application.SupplyKeeper, auth.DefaultSigVerificationGasConsumer))
	application.SetEndBlocker(application.EndBlocker)

	if loadLatest {
		err := application.LoadLatestVersion(application.keys[baseapp.MainStoreKey])
		if err != nil {
			tendermintOS.Exit(err.Error())
		}
	}

	return application
}

// Name returns the name of the App
func (app *SimulationApplication) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *SimulationApplication) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.moduleManager.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *SimulationApplication) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.moduleManager.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *SimulationApplication) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState simapp.GenesisState

	app.cdc.MustUnmarshalJSON(req.AppStateBytes, &genesisState)

	return app.moduleManager.InitGenesis(ctx, genesisState)
}

// LoadHeight loads a particular height
func (app *SimulationApplication) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keys[baseapp.MainStoreKey])
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *SimulationApplication) ModuleAccountAddrs() map[string]bool {
	moduleAccountAddresses := make(map[string]bool)
	for acc := range moduleAccountPermissions {
		moduleAccountAddresses[supply.NewModuleAddress(acc).String()] = true
	}

	return moduleAccountAddresses
}

// BlacklistedAccAddrs returns all the app's module account addresses black listed for receiving tokens.
func (app *SimulationApplication) BlacklistedAccAddrs() map[string]bool {
	blacklistedAddresses := make(map[string]bool)
	for acc := range moduleAccountPermissions {
		blacklistedAddresses[supply.NewModuleAddress(acc).String()] = !allowedReceivingModuleAccounts[acc]
	}

	return blacklistedAddresses
}

// Codec returns SimulationApplication's codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *SimulationApplication) Codec() *codec.Codec {
	return app.cdc
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimulationApplication) GetKey(storeKey string) *sdk.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimulationApplication) GetTKey(storeKey string) *sdk.TransientStoreKey {
	return app.transientKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimulationApplication) GetSubspace(moduleName string) params.Subspace {
	return app.subspaces[moduleName]
}

// SimulationManager implements the SimulationApp interface
func (app *SimulationApplication) SimulationManager() *module.SimulationManager {
	return app.simulationManager
}

// GetModuleAccountPermissions returns a copy of the module account permissions
func GetModuleAccountPermissions() map[string][]string {
	duplicateModuleAccountPermissions := make(map[string][]string)
	for k, v := range moduleAccountPermissions {
		duplicateModuleAccountPermissions[k] = v
	}

	return duplicateModuleAccountPermissions
}
