/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmTypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	serverTypes "github.com/cosmos/cosmos-sdk/server/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authRest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	sdkAuthKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authSimulation "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authTx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	sdkAuthTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/bank"
	sdkBankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	sdkBankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	sdkCapabilityKeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	sdkCapabilityTypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	sdkCrisisKeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	sdkCrisisTypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	sdkDistributionKeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	sdkDistributionTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	sdkEvidenceKeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	sdkEvidenceTypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	sdkGenUtilTypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	sdkGovKeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	sdkGovTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer"
	ibcTransferKeeper "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/keeper"
	ibcTransferTypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	ibc "github.com/cosmos/cosmos-sdk/x/ibc/core"
	ibcClient "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client"
	portTypes "github.com/cosmos/cosmos-sdk/x/ibc/core/05-port/types"
	ibcHost "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	sdkIBCKeeper "github.com/cosmos/cosmos-sdk/x/ibc/core/keeper"
	"github.com/cosmos/cosmos-sdk/x/mint"
	sdkMintKeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	sdkMintTypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	sdkParamsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	sdkParamsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramsProposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	sdkSlashingKeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	sdkSlashingTypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	sdkStakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	sdkStakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	sdkUpgradeKeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	sdkUpgradeTypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/persistenceOne/persistenceSDK/modules/assets"
	"github.com/persistenceOne/persistenceSDK/modules/classifications"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/conform"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/define"
	"github.com/persistenceOne/persistenceSDK/modules/identities"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/deputize"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/maintain"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/revoke"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/super"
	"github.com/persistenceOne/persistenceSDK/modules/metas"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/orders"
	"github.com/persistenceOne/persistenceSDK/modules/splits"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/burn"
	splitsMint "github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/mint"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/renumerate"
	splitsTransfer "github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/transfer"
	"github.com/persistenceOne/persistenceSDK/schema/applications"
	"github.com/persistenceOne/persistenceSDK/schema/applications/base/encoding"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	tmJson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tendermintOS "github.com/tendermint/tendermint/libs/os"
	tmProto "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"
	"honnef.co/go/tools/version"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type application struct {
	name string

	*baseapp.BaseApp

	moduleBasicManager module.BasicManager

	codec             codec.Marshaler
	legacyAminoCodec  *codec.LegacyAmino
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	enabledWasmProposalTypeList []wasm.ProposalType
	moduleAccountPermissions    map[string][]string

	keys               map[string]*sdkTypes.KVStoreKey
	transientStoreKeys map[string]*sdkTypes.TransientStoreKey
	memoryStoreKeys    map[string]*sdkTypes.MemoryStoreKey

	stakingKeeper      sdkStakingKeeper.Keeper
	slashingKeeper     sdkSlashingKeeper.Keeper
	distributionKeeper sdkDistributionKeeper.Keeper
	crisisKeeper       sdkCrisisKeeper.Keeper

	moduleManager           *module.Manager
	moduleSimulationManager *module.SimulationManager
}

var _ applications.Application = (*application)(nil)

func (app application) GetDefaultHome() string {
	return os.ExpandEnv("$HOME/." + app.name)
}
func (app application) GetModuleBasicManager() module.BasicManager {
	return app.moduleBasicManager
}
func (app application) GetCodec() codec.Marshaler {
	return app.codec
}
func (app application) GetLegacyAminoCodec() *codec.LegacyAmino {
	return app.legacyAminoCodec
}
func (app application) GetInterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

func (app application) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

func (app application) RegisterAPIRoutes(apiServer *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiServer.ClientCtx
	rpc.RegisterRoutes(clientCtx, apiServer.Router)
	// Register legacy tx routes.
	authRest.RegisterTxRoutes(clientCtx, apiServer.Router)
	// Register new tx routes from grpc-gateway.
	authTx.RegisterGRPCGatewayRoutes(clientCtx, apiServer.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiServer.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	app.moduleBasicManager.RegisterRESTRoutes(clientCtx, apiServer.Router)
	app.moduleBasicManager.RegisterGRPCGatewayRoutes(clientCtx, apiServer.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if apiConfig.Swagger {
		statikFS, err := fs.New()
		if err != nil {
			panic(err)
		}

		staticServer := http.FileServer(statikFS)
		apiServer.Router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))
	}
}

func (app application) RegisterTxService(clientCtx client.Context) {
	authTx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

func (app application) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.interfaceRegistry)
}

func (app application) ExportApplicationStateAndValidators(forZeroHeight bool, jailAllowedAddrs []string) (serverTypes.ExportedApp, error) {
	context := app.NewContext(true, tmProto.Header{Height: app.LastBlockHeight()})

	height := app.LastBlockHeight() + 1
	if forZeroHeight {
		height = 0
		applyAllowedAddrs := false

		if len(jailAllowedAddrs) > 0 {
			applyAllowedAddrs = true
		}

		allowedAddrsMap := make(map[string]bool)

		for _, addr := range jailAllowedAddrs {
			_, err := sdkTypes.ValAddressFromBech32(addr)
			if err != nil {
				panic(err)
			}
			allowedAddrsMap[addr] = true
		}

		app.crisisKeeper.AssertInvariants(context)

		app.stakingKeeper.IterateValidators(context, func(_ int64, val sdkStakingTypes.ValidatorI) (stop bool) {
			_, _ = app.distributionKeeper.WithdrawValidatorCommission(context, val.GetOperator())
			return false
		})

		delegations := app.stakingKeeper.GetAllDelegations(context)
		for _, delegation := range delegations {
			delegator, err := sdkTypes.AccAddressFromBech32(delegation.DelegatorAddress)
			if err != nil {
				return serverTypes.ExportedApp{}, err
			}
			validator, err := sdkTypes.ValAddressFromBech32(delegation.ValidatorAddress)
			if err != nil {
				return serverTypes.ExportedApp{}, err
			}
			_, _ = app.distributionKeeper.WithdrawDelegationRewards(context, delegator, validator)
		}

		app.distributionKeeper.DeleteAllValidatorSlashEvents(context)

		app.distributionKeeper.DeleteAllValidatorHistoricalRewards(context)

		contextHeight := context.BlockHeight()
		context = context.WithBlockHeight(0)

		app.stakingKeeper.IterateValidators(context, func(_ int64, val sdkStakingTypes.ValidatorI) (stop bool) {

			scraps := app.distributionKeeper.GetValidatorOutstandingRewards(context, val.GetOperator())
			feePool := app.distributionKeeper.GetFeePool(context)
			feePool.CommunityPool = feePool.CommunityPool.Add(scraps.Rewards...)
			app.distributionKeeper.SetFeePool(context, feePool)

			app.distributionKeeper.Hooks().AfterValidatorCreated(context, val.GetOperator())
			return false
		})

		for _, delegation := range delegations {
			delegator, err := sdkTypes.AccAddressFromBech32(delegation.DelegatorAddress)
			if err != nil {
				return serverTypes.ExportedApp{}, err
			}
			validator, err := sdkTypes.ValAddressFromBech32(delegation.ValidatorAddress)
			if err != nil {
				return serverTypes.ExportedApp{}, err
			}
			app.distributionKeeper.Hooks().BeforeDelegationCreated(context, delegator, validator)
			app.distributionKeeper.Hooks().AfterDelegationModified(context, delegator, validator)
		}

		context = context.WithBlockHeight(contextHeight)

		app.stakingKeeper.IterateRedelegations(context, func(_ int64, redelegation sdkStakingTypes.Redelegation) (stop bool) {
			for i := range redelegation.Entries {
				redelegation.Entries[i].CreationHeight = 0
			}
			app.stakingKeeper.SetRedelegation(context, redelegation)
			return false
		})

		app.stakingKeeper.IterateUnbondingDelegations(context, func(_ int64, unbondingDelegation sdkStakingTypes.UnbondingDelegation) (stop bool) {
			for i := range unbondingDelegation.Entries {
				unbondingDelegation.Entries[i].CreationHeight = 0
			}
			app.stakingKeeper.SetUnbondingDelegation(context, unbondingDelegation)
			return false
		})

		store := context.KVStore(app.keys[sdkStakingTypes.StoreKey])
		kvStoreReversePrefixIterator := sdkTypes.KVStoreReversePrefixIterator(store, sdkStakingTypes.ValidatorsKey)
		counter := int16(0)

		for ; kvStoreReversePrefixIterator.Valid(); kvStoreReversePrefixIterator.Next() {
			addr := sdkTypes.ValAddress(kvStoreReversePrefixIterator.Key()[1:])
			validator, found := app.stakingKeeper.GetValidator(context, addr)

			if !found {
				panic("Validator not found!")
			}

			validator.UnbondingHeight = 0

			if applyAllowedAddrs && !allowedAddrsMap[addr.String()] {
				validator.Jailed = true
			}

			app.stakingKeeper.SetValidator(context, validator)
			counter++
		}

		err := kvStoreReversePrefixIterator.Close()
		if err != nil {
			return serverTypes.ExportedApp{}, err
		}

		_, _ = app.stakingKeeper.ApplyAndReturnValidatorSetUpdates(context)

		app.slashingKeeper.IterateValidatorSigningInfos(
			context,
			func(validatorConsAddress sdkTypes.ConsAddress, validatorSigningInfo sdkSlashingTypes.ValidatorSigningInfo) (stop bool) {
				validatorSigningInfo.StartHeight = 0
				app.slashingKeeper.SetValidatorSigningInfo(context, validatorConsAddress, validatorSigningInfo)
				return false
			},
		)
	}

	genesisState := app.moduleManager.ExportGenesis(context, app.codec)
	appState, err := json.MarshalIndent(genesisState, "", "  ")
	if err != nil {
		return serverTypes.ExportedApp{}, err
	}

	validators, err := staking.WriteValidators(context, app.stakingKeeper)
	return serverTypes.ExportedApp{
		AppState:        appState,
		Validators:      validators,
		Height:          height,
		ConsensusParams: app.GetConsensusParams(context),
	}, err
}

func (app application) Initialize(logger log.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, skipUpgradeHeights map[int64]bool, homePath string, invCheckPeriod uint, encodingConfig encoding.EncodingConfig, appOpts serverTypes.AppOptions, baseAppOptions ...func(*baseapp.BaseApp)) applications.Application {
	app.BaseApp = baseapp.NewBaseApp(
		app.name,
		logger,
		db,
		encodingConfig.TxConfig.TxDecoder(),
		baseAppOptions...,
	)
	app.SetCommitMultiStoreTracer(traceStore)
	app.SetAppVersion(version.Version)
	app.SetInterfaceRegistry(app.interfaceRegistry)

	app.keys = sdkTypes.NewKVStoreKeys(
		sdkAuthTypes.StoreKey,
		sdkBankTypes.StoreKey,
		sdkStakingTypes.StoreKey,
		sdkMintTypes.StoreKey,
		sdkDistributionTypes.StoreKey,
		sdkSlashingTypes.StoreKey,
		sdkGovTypes.StoreKey,
		sdkParamsTypes.StoreKey,
		ibcHost.StoreKey,
		sdkUpgradeTypes.StoreKey,
		sdkEvidenceTypes.StoreKey,
		ibcTransferTypes.StoreKey,
		sdkCapabilityTypes.StoreKey,

		wasm.StoreKey,

		assets.Prototype().Name(),
		classifications.Prototype().Name(),
		identities.Prototype().Name(),
		maintainers.Prototype().Name(),
		metas.Prototype().Name(),
		orders.Prototype().Name(),
		splits.Prototype().Name(),
	)

	app.transientStoreKeys = sdkTypes.NewTransientStoreKeys(sdkParamsTypes.TStoreKey)
	app.memoryStoreKeys = sdkTypes.NewMemoryStoreKeys(sdkCapabilityTypes.MemStoreKey)

	paramsKeeper := sdkParamsKeeper.NewKeeper(
		app.codec,
		app.legacyAminoCodec,
		app.keys[sdkParamsTypes.StoreKey],
		app.transientStoreKeys[sdkParamsTypes.TStoreKey],
	)

	app.BaseApp.SetParamStore(paramsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(sdkParamsKeeper.ConsensusParamsKeyTable()))

	capabilityKeeper := sdkCapabilityKeeper.NewKeeper(app.codec, app.keys[sdkCapabilityTypes.StoreKey], app.memoryStoreKeys[sdkCapabilityTypes.MemStoreKey])
	scopedIBCKeeper := capabilityKeeper.ScopeToModule(ibcHost.ModuleName)
	scopedTransferKeeper := capabilityKeeper.ScopeToModule(ibcTransferTypes.ModuleName)
	scopedWasmKeeper := capabilityKeeper.ScopeToModule(wasm.ModuleName)

	accountKeeper := sdkAuthKeeper.NewAccountKeeper(
		app.codec,
		app.keys[sdkAuthTypes.StoreKey],
		paramsKeeper.Subspace(sdkAuthTypes.ModuleName),
		sdkAuthTypes.ProtoBaseAccount,
		app.moduleAccountPermissions,
	)

	moduleAccountAddresses := make(map[string]bool)
	for acc := range app.moduleAccountPermissions {
		moduleAccountAddresses[sdkAuthTypes.NewModuleAddress(acc).String()] = true
	}

	bankKeeper := sdkBankKeeper.NewBaseKeeper(
		app.codec,
		app.keys[sdkBankTypes.StoreKey],
		accountKeeper,
		paramsKeeper.Subspace(sdkBankTypes.ModuleName),
		moduleAccountAddresses,
	)

	stakingKeeper := sdkStakingKeeper.NewKeeper(
		app.codec,
		app.keys[sdkStakingTypes.StoreKey],
		accountKeeper,
		bankKeeper,
		paramsKeeper.Subspace(sdkStakingTypes.ModuleName),
	)

	mintKeeper := sdkMintKeeper.NewKeeper(
		app.codec,
		app.keys[sdkMintTypes.StoreKey],
		paramsKeeper.Subspace(sdkMintTypes.ModuleName),
		&stakingKeeper,
		accountKeeper,
		bankKeeper,
		sdkAuthTypes.FeeCollectorName,
	)

	app.distributionKeeper = sdkDistributionKeeper.NewKeeper(
		app.codec,
		app.keys[sdkDistributionTypes.StoreKey],
		paramsKeeper.Subspace(sdkDistributionTypes.ModuleName),
		accountKeeper,
		bankKeeper,
		&stakingKeeper,
		sdkAuthTypes.FeeCollectorName,
		moduleAccountAddresses,
	)

	app.slashingKeeper = sdkSlashingKeeper.NewKeeper(
		app.codec,
		app.keys[sdkSlashingTypes.StoreKey],
		&stakingKeeper,
		paramsKeeper.Subspace(sdkSlashingTypes.ModuleName),
	)

	app.crisisKeeper = sdkCrisisKeeper.NewKeeper(
		paramsKeeper.Subspace(sdkCrisisTypes.ModuleName),
		invCheckPeriod,
		bankKeeper,
		sdkAuthTypes.FeeCollectorName,
	)

	upgradeKeeper := sdkUpgradeKeeper.NewKeeper(
		skipUpgradeHeights,
		app.keys[sdkUpgradeTypes.StoreKey],
		app.codec,
		app.GetDefaultHome(),
	)

	app.stakingKeeper = *stakingKeeper.SetHooks(
		sdkStakingTypes.NewMultiStakingHooks(app.distributionKeeper.Hooks(), app.slashingKeeper.Hooks()),
	)

	ibcKeeper := sdkIBCKeeper.NewKeeper(
		app.codec, app.keys[ibcHost.StoreKey], paramsKeeper.Subspace(ibcHost.ModuleName), app.stakingKeeper, scopedIBCKeeper,
	)

	transferKeeper := ibcTransferKeeper.NewKeeper(
		app.codec, app.keys[ibcTransferTypes.StoreKey], paramsKeeper.Subspace(ibcTransferTypes.ModuleName),
		ibcKeeper.ChannelKeeper, &ibcKeeper.PortKeeper,
		accountKeeper, bankKeeper, scopedTransferKeeper,
	)
	transferModule := transfer.NewAppModule(transferKeeper)

	ibcRouter := portTypes.NewRouter()
	ibcRouter.AddRoute(ibcTransferTypes.ModuleName, transferModule)
	ibcKeeper.SetRouter(ibcRouter)

	evidenceKeeper := sdkEvidenceKeeper.NewKeeper(
		app.codec,
		app.keys[sdkEvidenceTypes.StoreKey],
		&app.stakingKeeper,
		app.slashingKeeper,
	)
	evidenceKeeper.SetRouter(sdkEvidenceTypes.NewRouter())

	metasModule := metas.Prototype().Initialize(
		app.keys[metas.Prototype().Name()],
		paramsKeeper.Subspace(metas.Prototype().Name()),
	)
	classificationsModule := classifications.Prototype().Initialize(
		app.keys[classifications.Prototype().Name()],
		paramsKeeper.Subspace(classifications.Prototype().Name()),
		metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
	)
	maintainersModule := maintainers.Prototype().Initialize(
		app.keys[metas.Prototype().Name()],
		paramsKeeper.Subspace(maintainers.Prototype().Name()),
		classificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
	)
	identitiesModule := identities.Prototype().Initialize(
		app.keys[identities.Prototype().Name()],
		paramsKeeper.Subspace(identities.Prototype().Name()),
		classificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
		classificationsModule.GetAuxiliary(define.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(deputize.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(maintain.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(revoke.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(super.Auxiliary.GetName()),
		metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
		metasModule.GetAuxiliary(supplement.Auxiliary.GetName()),
	)
	splitsModule := splits.Prototype().Initialize(
		app.keys[splits.Prototype().Name()],
		paramsKeeper.Subspace(splits.Prototype().Name()),
		bankKeeper,
		identitiesModule.GetAuxiliary(verify.Auxiliary.GetName()),
	)
	assetsModule := assets.Prototype().Initialize(
		app.keys[assets.Prototype().Name()],
		paramsKeeper.Subspace(assets.Prototype().Name()),
		classificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
		classificationsModule.GetAuxiliary(define.Auxiliary.GetName()),
		identitiesModule.GetAuxiliary(verify.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(deputize.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(maintain.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(revoke.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(super.Auxiliary.GetName()),
		metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
		metasModule.GetAuxiliary(supplement.Auxiliary.GetName()),
		splitsModule.GetAuxiliary(splitsMint.Auxiliary.GetName()),
		splitsModule.GetAuxiliary(burn.Auxiliary.GetName()),
		splitsModule.GetAuxiliary(renumerate.Auxiliary.GetName()),
	)
	ordersModule := orders.Prototype().Initialize(
		app.keys[orders.Prototype().Name()],
		paramsKeeper.Subspace(orders.Prototype().Name()),
		classificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
		classificationsModule.GetAuxiliary(define.Auxiliary.GetName()),
		identitiesModule.GetAuxiliary(verify.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(super.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(maintain.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(deputize.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(revoke.Auxiliary.GetName()),
		metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
		metasModule.GetAuxiliary(supplement.Auxiliary.GetName()),
		splitsModule.GetAuxiliary(splitsTransfer.Auxiliary.GetName()),
	)

	wasmDir := filepath.Join(app.GetDefaultHome(), wasmTypes.ModuleName)

	wasmWrap := struct {
		Wasm wasmTypes.WasmConfig `mapstructure:"wasm"`
	}{
		Wasm: wasmTypes.DefaultWasmConfig(),
	}

	err := viper.Unmarshal(&wasmWrap)
	if err != nil {
		panic("error while reading wasm config: " + err.Error())
	}

	wasmConfig := wasmWrap.Wasm

	var wasmOpts []wasm.Option
	if cast.ToBool(appOpts.Get("telemetry.enabled")) {
		wasmOpts = append(wasmOpts, wasmkeeper.WithVMCacheMetrics(prometheus.DefaultRegisterer))
	}

	wasmKeeper := wasm.NewKeeper(
		app.codec,
		app.keys[wasm.StoreKey],
		paramsKeeper.Subspace(wasm.ModuleName),
		accountKeeper,
		bankKeeper,
		app.stakingKeeper,
		app.distributionKeeper,
		ibcKeeper.ChannelKeeper,
		&ibcKeeper.PortKeeper,
		scopedWasmKeeper,
		transferKeeper,
		app.Router(),
		app.GRPCQueryRouter(),
		wasmDir,
		wasmConfig,
		sdkStakingTypes.ModuleName,
		//TODO &wasm.MessageEncoders{Custom: wasmUtilities.CustomEncoder(assets.Prototype(), classifications.Prototype(), identities.Prototype(), maintainers.Prototype(), metas.Prototype(), orders.Prototype(), splits.Prototype())},
		wasmOpts...)

	govRouter := sdkGovTypes.NewRouter()

	govRouter.AddRoute(sdkGovTypes.RouterKey, sdkGovTypes.ProposalHandler).
		AddRoute(paramsProposal.RouterKey, params.NewParamChangeProposalHandler(paramsKeeper)).
		AddRoute(sdkDistributionTypes.RouterKey, distribution.NewCommunityPoolSpendProposalHandler(app.distributionKeeper)).
		AddRoute(sdkUpgradeTypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(upgradeKeeper)).
		AddRoute(ibcHost.RouterKey, ibcClient.NewClientUpdateProposalHandler(ibcKeeper.ClientKeeper))

	if len(app.enabledWasmProposalTypeList) != 0 {
		govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(wasmKeeper, app.enabledWasmProposalTypeList))
	}

	govKeeper := sdkGovKeeper.NewKeeper(
		app.codec,
		app.keys[sdkGovTypes.StoreKey],
		paramsKeeper.Subspace(sdkGovTypes.ModuleName).WithKeyTable(sdkGovTypes.ParamKeyTable()),
		accountKeeper,
		bankKeeper,
		&stakingKeeper,
		govRouter,
	)

	var skipGenesisInvariants = false
	opt := appOpts.Get(crisis.FlagSkipGenesisInvariants)
	if opt, ok := opt.(bool); ok {
		skipGenesisInvariants = opt
	}

	app.moduleManager = module.NewManager(
		genutil.NewAppModule(accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx, encodingConfig.TxConfig),
		auth.NewAppModule(app.codec, accountKeeper, nil),
		vesting.NewAppModule(accountKeeper, bankKeeper),
		bank.NewAppModule(app.codec, bankKeeper, accountKeeper),
		capability.NewAppModule(app.codec, *capabilityKeeper),
		crisis.NewAppModule(&app.crisisKeeper, skipGenesisInvariants),
		gov.NewAppModule(app.codec, govKeeper, accountKeeper, bankKeeper),
		mint.NewAppModule(app.codec, mintKeeper, accountKeeper),
		slashing.NewAppModule(app.codec, app.slashingKeeper, accountKeeper, bankKeeper, app.stakingKeeper),
		distribution.NewAppModule(app.codec, app.distributionKeeper, accountKeeper, bankKeeper, app.stakingKeeper),
		staking.NewAppModule(app.codec, app.stakingKeeper, accountKeeper, bankKeeper),
		upgrade.NewAppModule(upgradeKeeper),
		evidence.NewAppModule(*evidenceKeeper),
		ibc.NewAppModule(ibcKeeper),
		params.NewAppModule(paramsKeeper),
		transferModule,

		wasm.NewAppModule(app.codec, &wasmKeeper, app.stakingKeeper),

		assetsModule,
		classificationsModule,
		identitiesModule,
		maintainersModule,
		metasModule,
		ordersModule,
		splitsModule,
	)

	app.moduleManager.SetOrderBeginBlockers(
		sdkUpgradeTypes.ModuleName,
		sdkMintTypes.ModuleName,
		sdkDistributionTypes.ModuleName,
		sdkSlashingTypes.ModuleName,
		sdkEvidenceTypes.ModuleName,
		sdkStakingTypes.ModuleName,
		ibcHost.ModuleName,
	)
	app.moduleManager.SetOrderEndBlockers(
		sdkCrisisTypes.ModuleName,
		sdkGovTypes.ModuleName,
		sdkStakingTypes.ModuleName,
		ordersModule.Name(),
	)
	app.moduleManager.SetOrderInitGenesis(
		sdkCapabilityTypes.ModuleName,
		sdkAuthTypes.ModuleName,
		sdkBankTypes.ModuleName,
		sdkDistributionTypes.ModuleName,
		sdkStakingTypes.ModuleName,
		sdkSlashingTypes.ModuleName,
		sdkGovTypes.ModuleName,
		sdkMintTypes.ModuleName,
		sdkCrisisTypes.ModuleName,
		ibcHost.ModuleName,
		sdkGenUtilTypes.ModuleName,
		sdkEvidenceTypes.ModuleName,
		ibcTransferTypes.ModuleName,

		wasm.ModuleName,

		assets.Prototype().Name(),
		classifications.Prototype().Name(),
		identities.Prototype().Name(),
		maintainers.Prototype().Name(),
		metas.Prototype().Name(),
		orders.Prototype().Name(),
		splits.Prototype().Name(),
	)

	app.moduleManager.RegisterInvariants(&app.crisisKeeper)
	app.moduleManager.RegisterRoutes(app.Router(), app.QueryRouter(), app.legacyAminoCodec)
	app.moduleManager.RegisterServices(module.NewConfigurator(app.MsgServiceRouter(), app.GRPCQueryRouter()))

	app.moduleSimulationManager = module.NewSimulationManager(
		auth.NewAppModule(app.codec, accountKeeper, authSimulation.RandomGenesisAccounts),
		bank.NewAppModule(app.codec, bankKeeper, accountKeeper),
		capability.NewAppModule(app.codec, *capabilityKeeper),
		gov.NewAppModule(app.codec, govKeeper, accountKeeper, bankKeeper),
		mint.NewAppModule(app.codec, mintKeeper, accountKeeper),
		staking.NewAppModule(app.codec, app.stakingKeeper, accountKeeper, bankKeeper),
		distribution.NewAppModule(app.codec, app.distributionKeeper, accountKeeper, bankKeeper, app.stakingKeeper),
		slashing.NewAppModule(app.codec, app.slashingKeeper, accountKeeper, bankKeeper, app.stakingKeeper),
		params.NewAppModule(paramsKeeper),
		evidence.NewAppModule(*evidenceKeeper),
		ibc.NewAppModule(ibcKeeper),
		transferModule,

		assets.Prototype(),
		classifications.Prototype(),
		identities.Prototype(),
		maintainers.Prototype(),
		metas.Prototype(),
		orders.Prototype(),
		splits.Prototype(),
	)
	app.moduleSimulationManager.RegisterStoreDecoders()

	app.MountKVStores(app.keys)
	app.MountTransientStores(app.transientStoreKeys)
	app.MountMemoryStores(app.memoryStoreKeys)

	app.SetInitChainer(func(context sdkTypes.Context, requestInitChain abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
		var genesisState map[string]json.RawMessage
		if err := tmJson.Unmarshal(requestInitChain.AppStateBytes, &genesisState); err != nil {
			panic(err)
		}
		return app.moduleManager.InitGenesis(context, app.codec, genesisState)
	})

	app.SetBeginBlocker(app.moduleManager.BeginBlock)
	app.SetAnteHandler(
		ante.NewAnteHandler(
			accountKeeper, bankKeeper, ante.DefaultSigVerificationGasConsumer,
			encodingConfig.TxConfig.SignModeHandler(),
		))
	app.SetEndBlocker(app.moduleManager.EndBlock)

	if loadLatest {
		err := app.LoadLatestVersion()
		if err != nil {
			tendermintOS.Exit(err.Error())
		}
		ctx := app.NewUncachedContext(true, tmProto.Header{})
		capabilityKeeper.InitializeAndSeal(ctx)
	}

	return &app
}

func NewApplication(name string, moduleBasicManager module.BasicManager, encodingConfig encoding.EncodingConfig, enabledWasmProposalTypeList []wasm.ProposalType, moduleAccountPermissions map[string][]string) applications.Application {
	return &application{
		name:                        name,
		moduleBasicManager:          moduleBasicManager,
		legacyAminoCodec:            encodingConfig.LegacyAmino,
		codec:                       encodingConfig.Marshaler,
		interfaceRegistry:           encodingConfig.InterfaceRegistry,
		enabledWasmProposalTypeList: enabledWasmProposalTypeList,
		moduleAccountPermissions:    moduleAccountPermissions,
	}
}
