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

func (application application) GetDefaultHome() string {
	return os.ExpandEnv("$HOME/." + application.name)
}
func (application application) GetModuleBasicManager() module.BasicManager {
	return application.moduleBasicManager
}
func (application application) GetCodec() codec.Marshaler {
	return application.codec
}
func (application application) GetLegacyAminoCodec() *codec.LegacyAmino {
	return application.legacyAminoCodec
}
func (application application) GetInterfaceRegistry() types.InterfaceRegistry {
	return application.interfaceRegistry
}

func (application application) LoadHeight(height int64) error {
	return application.LoadVersion(height) //nolint:typecheck
}

func (application application) RegisterAPIRoutes(apiServer *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiServer.ClientCtx
	rpc.RegisterRoutes(clientCtx, apiServer.Router)
	// Register legacy tx routes.
	authRest.RegisterTxRoutes(clientCtx, apiServer.Router)
	// Register new tx routes from grpc-gateway.
	authTx.RegisterGRPCGatewayRoutes(clientCtx, apiServer.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiServer.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	application.moduleBasicManager.RegisterRESTRoutes(clientCtx, apiServer.Router)
	application.moduleBasicManager.RegisterGRPCGatewayRoutes(clientCtx, apiServer.GRPCGatewayRouter)

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

func (application application) RegisterTxService(clientCtx client.Context) {
	authTx.RegisterTxService(application.BaseApp.GRPCQueryRouter(), clientCtx, application.BaseApp.Simulate, application.interfaceRegistry)
}

func (application application) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(application.BaseApp.GRPCQueryRouter(), clientCtx, application.interfaceRegistry)
}

func (application application) ExportApplicationStateAndValidators(forZeroHeight bool, jailAllowedAddrs []string) (serverTypes.ExportedApp, error) {
	context := application.NewContext(true, tmProto.Header{Height: application.LastBlockHeight()}) //nolint:typecheck

	height := application.LastBlockHeight() + 1 //nolint:typecheck
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

		application.crisisKeeper.AssertInvariants(context)

		application.stakingKeeper.IterateValidators(context, func(_ int64, val sdkStakingTypes.ValidatorI) (stop bool) {
			_, _ = application.distributionKeeper.WithdrawValidatorCommission(context, val.GetOperator())
			return false
		})

		delegations := application.stakingKeeper.GetAllDelegations(context)
		for _, delegation := range delegations {
			delegator, err := sdkTypes.AccAddressFromBech32(delegation.DelegatorAddress)
			if err != nil {
				return serverTypes.ExportedApp{}, err
			}
			validator, err := sdkTypes.ValAddressFromBech32(delegation.ValidatorAddress)
			if err != nil {
				return serverTypes.ExportedApp{}, err
			}
			_, _ = application.distributionKeeper.WithdrawDelegationRewards(context, delegator, validator)
		}

		application.distributionKeeper.DeleteAllValidatorSlashEvents(context)

		application.distributionKeeper.DeleteAllValidatorHistoricalRewards(context)

		contextHeight := context.BlockHeight()
		context = context.WithBlockHeight(0)

		application.stakingKeeper.IterateValidators(context, func(_ int64, val sdkStakingTypes.ValidatorI) (stop bool) {

			scraps := application.distributionKeeper.GetValidatorOutstandingRewards(context, val.GetOperator())
			feePool := application.distributionKeeper.GetFeePool(context)
			feePool.CommunityPool = feePool.CommunityPool.Add(scraps.Rewards...)
			application.distributionKeeper.SetFeePool(context, feePool)

			application.distributionKeeper.Hooks().AfterValidatorCreated(context, val.GetOperator())
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
			application.distributionKeeper.Hooks().BeforeDelegationCreated(context, delegator, validator)
			application.distributionKeeper.Hooks().AfterDelegationModified(context, delegator, validator)
		}

		context = context.WithBlockHeight(contextHeight)

		application.stakingKeeper.IterateRedelegations(context, func(_ int64, redelegation sdkStakingTypes.Redelegation) (stop bool) {
			for i := range redelegation.Entries {
				redelegation.Entries[i].CreationHeight = 0
			}
			application.stakingKeeper.SetRedelegation(context, redelegation)
			return false
		})

		application.stakingKeeper.IterateUnbondingDelegations(context, func(_ int64, unbondingDelegation sdkStakingTypes.UnbondingDelegation) (stop bool) {
			for i := range unbondingDelegation.Entries {
				unbondingDelegation.Entries[i].CreationHeight = 0
			}
			application.stakingKeeper.SetUnbondingDelegation(context, unbondingDelegation)
			return false
		})

		store := context.KVStore(application.keys[sdkStakingTypes.StoreKey])
		kvStoreReversePrefixIterator := sdkTypes.KVStoreReversePrefixIterator(store, sdkStakingTypes.ValidatorsKey)
		counter := int16(0)

		for ; kvStoreReversePrefixIterator.Valid(); kvStoreReversePrefixIterator.Next() {
			addr := sdkTypes.ValAddress(kvStoreReversePrefixIterator.Key()[1:])
			validator, found := application.stakingKeeper.GetValidator(context, addr)

			if !found {
				panic("Validator not found!")
			}

			validator.UnbondingHeight = 0

			if applyAllowedAddrs && !allowedAddrsMap[addr.String()] {
				validator.Jailed = true
			}

			application.stakingKeeper.SetValidator(context, validator)
			counter++
		}

		err := kvStoreReversePrefixIterator.Close()
		if err != nil {
			return serverTypes.ExportedApp{}, err
		}

		_, _ = application.stakingKeeper.ApplyAndReturnValidatorSetUpdates(context)

		application.slashingKeeper.IterateValidatorSigningInfos(
			context,
			func(validatorConsAddress sdkTypes.ConsAddress, validatorSigningInfo sdkSlashingTypes.ValidatorSigningInfo) (stop bool) {
				validatorSigningInfo.StartHeight = 0
				application.slashingKeeper.SetValidatorSigningInfo(context, validatorConsAddress, validatorSigningInfo)
				return false
			},
		)
	}

	genesisState := application.moduleManager.ExportGenesis(context, application.codec)
	appState, err := json.MarshalIndent(genesisState, "", "  ")
	if err != nil {
		return serverTypes.ExportedApp{}, err
	}

	validators, err := staking.WriteValidators(context, application.stakingKeeper)
	return serverTypes.ExportedApp{
		AppState:        appState,
		Validators:      validators,
		Height:          height,
		ConsensusParams: application.GetConsensusParams(context), //nolint:typecheck
	}, err
}

func (application application) Initialize(logger log.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, skipUpgradeHeights map[int64]bool, homePath string, invCheckPeriod uint, encodingConfig encoding.EncodingConfig, appOpts serverTypes.AppOptions, baseAppOptions ...func(*baseapp.BaseApp)) applications.Application {
	application.BaseApp = baseapp.NewBaseApp(
		application.name,
		logger,
		db,
		encodingConfig.TxConfig.TxDecoder(),
		baseAppOptions...,
	)
	application.SetCommitMultiStoreTracer(traceStore)               //nolint:typecheck
	application.SetAppVersion(version.Version)                      //nolint:typecheck
	application.SetInterfaceRegistry(application.interfaceRegistry) //nolint:typecheck

	application.keys = sdkTypes.NewKVStoreKeys(
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

	application.transientStoreKeys = sdkTypes.NewTransientStoreKeys(sdkParamsTypes.TStoreKey)
	application.memoryStoreKeys = sdkTypes.NewMemoryStoreKeys(sdkCapabilityTypes.MemStoreKey)

	paramsKeeper := sdkParamsKeeper.NewKeeper(
		application.codec,
		application.legacyAminoCodec,
		application.keys[sdkParamsTypes.StoreKey],
		application.transientStoreKeys[sdkParamsTypes.TStoreKey],
	)

	application.BaseApp.SetParamStore(paramsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(sdkParamsKeeper.ConsensusParamsKeyTable()))

	capabilityKeeper := sdkCapabilityKeeper.NewKeeper(application.codec, application.keys[sdkCapabilityTypes.StoreKey], application.memoryStoreKeys[sdkCapabilityTypes.MemStoreKey])
	scopedIBCKeeper := capabilityKeeper.ScopeToModule(ibcHost.ModuleName)
	scopedTransferKeeper := capabilityKeeper.ScopeToModule(ibcTransferTypes.ModuleName)
	scopedWasmKeeper := capabilityKeeper.ScopeToModule(wasm.ModuleName)

	accountKeeper := sdkAuthKeeper.NewAccountKeeper(
		application.codec,
		application.keys[sdkAuthTypes.StoreKey],
		paramsKeeper.Subspace(sdkAuthTypes.ModuleName),
		sdkAuthTypes.ProtoBaseAccount,
		application.moduleAccountPermissions,
	)

	moduleAccountAddresses := make(map[string]bool)
	for acc := range application.moduleAccountPermissions {
		moduleAccountAddresses[sdkAuthTypes.NewModuleAddress(acc).String()] = true
	}

	bankKeeper := sdkBankKeeper.NewBaseKeeper(
		application.codec,
		application.keys[sdkBankTypes.StoreKey],
		accountKeeper,
		paramsKeeper.Subspace(sdkBankTypes.ModuleName),
		moduleAccountAddresses,
	)

	stakingKeeper := sdkStakingKeeper.NewKeeper(
		application.codec,
		application.keys[sdkStakingTypes.StoreKey],
		accountKeeper,
		bankKeeper,
		paramsKeeper.Subspace(sdkStakingTypes.ModuleName),
	)

	mintKeeper := sdkMintKeeper.NewKeeper(
		application.codec,
		application.keys[sdkMintTypes.StoreKey],
		paramsKeeper.Subspace(sdkMintTypes.ModuleName),
		&stakingKeeper,
		accountKeeper,
		bankKeeper,
		sdkAuthTypes.FeeCollectorName,
	)

	application.distributionKeeper = sdkDistributionKeeper.NewKeeper(
		application.codec,
		application.keys[sdkDistributionTypes.StoreKey],
		paramsKeeper.Subspace(sdkDistributionTypes.ModuleName),
		accountKeeper,
		bankKeeper,
		&stakingKeeper,
		sdkAuthTypes.FeeCollectorName,
		moduleAccountAddresses,
	)

	application.slashingKeeper = sdkSlashingKeeper.NewKeeper(
		application.codec,
		application.keys[sdkSlashingTypes.StoreKey],
		&stakingKeeper,
		paramsKeeper.Subspace(sdkSlashingTypes.ModuleName),
	)

	application.crisisKeeper = sdkCrisisKeeper.NewKeeper(
		paramsKeeper.Subspace(sdkCrisisTypes.ModuleName),
		invCheckPeriod,
		bankKeeper,
		sdkAuthTypes.FeeCollectorName,
	)

	upgradeKeeper := sdkUpgradeKeeper.NewKeeper(
		skipUpgradeHeights,
		application.keys[sdkUpgradeTypes.StoreKey],
		application.codec,
		application.GetDefaultHome(),
	)

	application.stakingKeeper = *stakingKeeper.SetHooks(
		sdkStakingTypes.NewMultiStakingHooks(application.distributionKeeper.Hooks(), application.slashingKeeper.Hooks()),
	)

	ibcKeeper := sdkIBCKeeper.NewKeeper(
		application.codec, application.keys[ibcHost.StoreKey], paramsKeeper.Subspace(ibcHost.ModuleName), application.stakingKeeper, scopedIBCKeeper,
	)

	transferKeeper := ibcTransferKeeper.NewKeeper(
		application.codec, application.keys[ibcTransferTypes.StoreKey], paramsKeeper.Subspace(ibcTransferTypes.ModuleName),
		ibcKeeper.ChannelKeeper, &ibcKeeper.PortKeeper,
		accountKeeper, bankKeeper, scopedTransferKeeper,
	)
	transferModule := transfer.NewAppModule(transferKeeper)

	ibcRouter := portTypes.NewRouter()
	ibcRouter.AddRoute(ibcTransferTypes.ModuleName, transferModule)
	ibcKeeper.SetRouter(ibcRouter)

	evidenceKeeper := sdkEvidenceKeeper.NewKeeper(
		application.codec,
		application.keys[sdkEvidenceTypes.StoreKey],
		&application.stakingKeeper,
		application.slashingKeeper,
	)
	evidenceKeeper.SetRouter(sdkEvidenceTypes.NewRouter())

	metasModule := metas.Prototype().Initialize(
		application.keys[metas.Prototype().Name()],
		paramsKeeper.Subspace(metas.Prototype().Name()),
	)
	classificationsModule := classifications.Prototype().Initialize(
		application.keys[classifications.Prototype().Name()],
		paramsKeeper.Subspace(classifications.Prototype().Name()),
		metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
	)
	maintainersModule := maintainers.Prototype().Initialize(
		application.keys[metas.Prototype().Name()],
		paramsKeeper.Subspace(maintainers.Prototype().Name()),
		classificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
	)
	identitiesModule := identities.Prototype().Initialize(
		application.keys[identities.Prototype().Name()],
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
		application.keys[splits.Prototype().Name()],
		paramsKeeper.Subspace(splits.Prototype().Name()),
		bankKeeper,
		identitiesModule.GetAuxiliary(verify.Auxiliary.GetName()),
	)
	assetsModule := assets.Prototype().Initialize(
		application.keys[assets.Prototype().Name()],
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
		application.keys[orders.Prototype().Name()],
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

	wasmDir := filepath.Join(application.GetDefaultHome(), wasmTypes.ModuleName)

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
		application.codec,
		application.keys[wasm.StoreKey],
		paramsKeeper.Subspace(wasm.ModuleName),
		accountKeeper,
		bankKeeper,
		application.stakingKeeper,
		application.distributionKeeper,
		ibcKeeper.ChannelKeeper,
		&ibcKeeper.PortKeeper,
		scopedWasmKeeper,
		transferKeeper,
		application.Router(),          //nolint:typecheck
		application.GRPCQueryRouter(), //nolint:typecheck
		wasmDir,
		wasmConfig,
		sdkStakingTypes.ModuleName,
		//TODO &wasm.MessageEncoders{Custom: wasmUtilities.CustomEncoder(assets.Prototype(), classifications.Prototype(), identities.Prototype(), maintainers.Prototype(), metas.Prototype(), orders.Prototype(), splits.Prototype())},
		wasmOpts...)

	govRouter := sdkGovTypes.NewRouter()

	govRouter.AddRoute(sdkGovTypes.RouterKey, sdkGovTypes.ProposalHandler).
		AddRoute(paramsProposal.RouterKey, params.NewParamChangeProposalHandler(paramsKeeper)).
		AddRoute(sdkDistributionTypes.RouterKey, distribution.NewCommunityPoolSpendProposalHandler(application.distributionKeeper)).
		AddRoute(sdkUpgradeTypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(upgradeKeeper)).
		AddRoute(ibcHost.RouterKey, ibcClient.NewClientUpdateProposalHandler(ibcKeeper.ClientKeeper))

	if len(application.enabledWasmProposalTypeList) != 0 {
		govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(wasmKeeper, application.enabledWasmProposalTypeList))
	}

	govKeeper := sdkGovKeeper.NewKeeper(
		application.codec,
		application.keys[sdkGovTypes.StoreKey],
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

	application.moduleManager = module.NewManager(
		genutil.NewAppModule(accountKeeper, application.stakingKeeper, application.BaseApp.DeliverTx, encodingConfig.TxConfig),
		auth.NewAppModule(application.codec, accountKeeper, nil),
		vesting.NewAppModule(accountKeeper, bankKeeper),
		bank.NewAppModule(application.codec, bankKeeper, accountKeeper),
		capability.NewAppModule(application.codec, *capabilityKeeper),
		crisis.NewAppModule(&application.crisisKeeper, skipGenesisInvariants),
		gov.NewAppModule(application.codec, govKeeper, accountKeeper, bankKeeper),
		mint.NewAppModule(application.codec, mintKeeper, accountKeeper),
		slashing.NewAppModule(application.codec, application.slashingKeeper, accountKeeper, bankKeeper, application.stakingKeeper),
		distribution.NewAppModule(application.codec, application.distributionKeeper, accountKeeper, bankKeeper, application.stakingKeeper),
		staking.NewAppModule(application.codec, application.stakingKeeper, accountKeeper, bankKeeper),
		upgrade.NewAppModule(upgradeKeeper),
		evidence.NewAppModule(*evidenceKeeper),
		ibc.NewAppModule(ibcKeeper),
		params.NewAppModule(paramsKeeper),
		transferModule,

		wasm.NewAppModule(application.codec, &wasmKeeper, application.stakingKeeper),

		assetsModule,
		classificationsModule,
		identitiesModule,
		maintainersModule,
		metasModule,
		ordersModule,
		splitsModule,
	)

	application.moduleManager.SetOrderBeginBlockers(
		sdkUpgradeTypes.ModuleName,
		sdkMintTypes.ModuleName,
		sdkDistributionTypes.ModuleName,
		sdkSlashingTypes.ModuleName,
		sdkEvidenceTypes.ModuleName,
		sdkStakingTypes.ModuleName,
		ibcHost.ModuleName,
	)
	application.moduleManager.SetOrderEndBlockers(
		sdkCrisisTypes.ModuleName,
		sdkGovTypes.ModuleName,
		sdkStakingTypes.ModuleName,
		ordersModule.Name(),
	)
	application.moduleManager.SetOrderInitGenesis(
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

	application.moduleManager.RegisterInvariants(&application.crisisKeeper)
	application.moduleManager.RegisterRoutes(application.Router(), application.QueryRouter(), application.legacyAminoCodec)           //nolint:typecheck
	application.moduleManager.RegisterServices(module.NewConfigurator(application.MsgServiceRouter(), application.GRPCQueryRouter())) //nolint:typecheck

	application.moduleSimulationManager = module.NewSimulationManager(
		auth.NewAppModule(application.codec, accountKeeper, authSimulation.RandomGenesisAccounts),
		bank.NewAppModule(application.codec, bankKeeper, accountKeeper),
		capability.NewAppModule(application.codec, *capabilityKeeper),
		gov.NewAppModule(application.codec, govKeeper, accountKeeper, bankKeeper),
		mint.NewAppModule(application.codec, mintKeeper, accountKeeper),
		staking.NewAppModule(application.codec, application.stakingKeeper, accountKeeper, bankKeeper),
		distribution.NewAppModule(application.codec, application.distributionKeeper, accountKeeper, bankKeeper, application.stakingKeeper),
		slashing.NewAppModule(application.codec, application.slashingKeeper, accountKeeper, bankKeeper, application.stakingKeeper),
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
	application.moduleSimulationManager.RegisterStoreDecoders()

	application.MountKVStores(application.keys)                      //nolint:typecheck
	application.MountTransientStores(application.transientStoreKeys) //nolint:typecheck
	application.MountMemoryStores(application.memoryStoreKeys)       //nolint:typecheck

	application.SetInitChainer(func(context sdkTypes.Context, requestInitChain abciTypes.RequestInitChain) abciTypes.ResponseInitChain { //nolint:typecheck
		var genesisState map[string]json.RawMessage
		if err := tmJson.Unmarshal(requestInitChain.AppStateBytes, &genesisState); err != nil {
			panic(err)
		}
		return application.moduleManager.InitGenesis(context, application.codec, genesisState)
	})

	application.SetBeginBlocker(application.moduleManager.BeginBlock) //nolint:typecheck
	application.SetAnteHandler( //nolint:typecheck
		ante.NewAnteHandler(
			accountKeeper, bankKeeper, ante.DefaultSigVerificationGasConsumer,
			encodingConfig.TxConfig.SignModeHandler(),
		))
	application.SetEndBlocker(application.moduleManager.EndBlock) //nolint:typecheck

	if loadLatest {
		err := application.LoadLatestVersion() //nolint:typecheck
		if err != nil {
			tendermintOS.Exit(err.Error())
		}
		ctx := application.NewUncachedContext(true, tmProto.Header{}) //nolint:typecheck
		capabilityKeeper.InitializeAndSeal(ctx)
	}

	return &application
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
