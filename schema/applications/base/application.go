// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
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
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	"github.com/spf13/viper"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintOS "github.com/tendermint/tendermint/libs/os"
	tendermintTypes "github.com/tendermint/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"
	"honnef.co/go/tools/version"

	"github.com/AssetMantle/modules/modules/assets"
	"github.com/AssetMantle/modules/modules/classifications"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/define"
	"github.com/AssetMantle/modules/modules/identities"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/maintainers"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/deputize"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/maintain"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/revoke"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/super"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/metas"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/orders"
	"github.com/AssetMantle/modules/modules/splits"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/burn"
	splitsMint "github.com/AssetMantle/modules/modules/splits/auxiliaries/mint"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/renumerate"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/transfer"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/applications"
	wasmUtilities "github.com/AssetMantle/modules/utilities/wasm"
)

type application struct {
	name string

	moduleBasicManager module.BasicManager

	codec *codec.Codec

	enabledWasmProposalTypeList []wasm.ProposalType
	moduleAccountPermissions    map[string][]string
	tokenReceiveAllowedModules  map[string]bool

	keys map[string]*sdkTypes.KVStoreKey

	stakingKeeper      staking.Keeper
	slashingKeeper     slashing.Keeper
	distributionKeeper distribution.Keeper
	crisisKeeper       crisis.Keeper

	moduleManager *module.Manager

	baseapp.BaseApp
}

var _ applications.Application = (*application)(nil)

func (application application) GetDefaultNodeHome() string {
	return os.ExpandEnv("$HOME/." + application.name + "/Node")
}
func (application application) GetDefaultClientHome() string {
	return os.ExpandEnv("$HOME/." + application.name + "/Client")
}
func (application application) GetModuleBasicManager() module.BasicManager {
	return application.moduleBasicManager
}
func (application application) GetCodec() *codec.Codec {
	return application.codec
}
func (application application) LoadHeight(height int64) error {
	return application.BaseApp.LoadVersion(height, application.keys[baseapp.MainStoreKey])
}
func (application application) ExportApplicationStateAndValidators(forZeroHeight bool, jailWhiteList []string) (json.RawMessage, []tendermintTypes.GenesisValidator, error) {
	context := application.BaseApp.NewContext(true, abciTypes.Header{Height: application.BaseApp.LastBlockHeight()})

	if forZeroHeight {
		applyWhiteList := false

		if len(jailWhiteList) > 0 {
			applyWhiteList = true
		}

		whiteListMap := make(map[string]bool)

		for _, address := range jailWhiteList {
			if _, err := sdkTypes.ValAddressFromBech32(address); err != nil {
				panic(err)
			}

			whiteListMap[address] = true
		}

		application.crisisKeeper.AssertInvariants(context)

		application.stakingKeeper.IterateValidators(context, func(_ int64, val staking.ValidatorI) (stop bool) {
			_, _ = application.distributionKeeper.WithdrawValidatorCommission(context, val.GetOperator())
			return false
		})

		delegations := application.stakingKeeper.GetAllDelegations(context)
		for _, delegation := range delegations {
			_, _ = application.distributionKeeper.WithdrawDelegationRewards(context, delegation.DelegatorAddress, delegation.ValidatorAddress)
		}

		application.distributionKeeper.DeleteAllValidatorSlashEvents(context)

		application.distributionKeeper.DeleteAllValidatorHistoricalRewards(context)

		height := context.BlockHeight()
		context = context.WithBlockHeight(0)

		application.stakingKeeper.IterateValidators(context, func(_ int64, val staking.ValidatorI) (stop bool) {
			scraps := application.distributionKeeper.GetValidatorOutstandingRewards(context, val.GetOperator())
			feePool := application.distributionKeeper.GetFeePool(context)
			feePool.CommunityPool = feePool.CommunityPool.Add(scraps...)
			application.distributionKeeper.SetFeePool(context, feePool)

			application.distributionKeeper.Hooks().AfterValidatorCreated(context, val.GetOperator())
			return false
		})

		for _, delegation := range delegations {
			application.distributionKeeper.Hooks().BeforeDelegationCreated(context, delegation.DelegatorAddress, delegation.ValidatorAddress)
			application.distributionKeeper.Hooks().AfterDelegationModified(context, delegation.DelegatorAddress, delegation.ValidatorAddress)
		}

		context = context.WithBlockHeight(height)

		application.stakingKeeper.IterateRedelegations(context, func(_ int64, redelegation staking.Redelegation) (stop bool) {
			for i := range redelegation.Entries {
				redelegation.Entries[i].CreationHeight = 0
			}
			application.stakingKeeper.SetRedelegation(context, redelegation)
			return false
		})

		application.stakingKeeper.IterateUnbondingDelegations(context, func(_ int64, unbondingDelegation staking.UnbondingDelegation) (stop bool) {
			for i := range unbondingDelegation.Entries {
				unbondingDelegation.Entries[i].CreationHeight = 0
			}
			application.stakingKeeper.SetUnbondingDelegation(context, unbondingDelegation)
			return false
		})

		store := context.KVStore(application.keys[staking.StoreKey])
		kvStoreReversePrefixIterator := sdkTypes.KVStoreReversePrefixIterator(store, staking.ValidatorsKey)
		counter := int16(0)

		for ; kvStoreReversePrefixIterator.Valid(); kvStoreReversePrefixIterator.Next() {
			addr := sdkTypes.ValAddress(kvStoreReversePrefixIterator.Key()[1:])
			validator, found := application.stakingKeeper.GetValidator(context, addr)

			if !found {
				panic("Validator not found!")
			}

			validator.UnbondingHeight = 0

			if applyWhiteList && !whiteListMap[addr.String()] {
				validator.Jailed = true
			}

			application.stakingKeeper.SetValidator(context, validator)
			counter++
		}

		kvStoreReversePrefixIterator.Close()

		_ = application.stakingKeeper.ApplyAndReturnValidatorSetUpdates(context)

		application.slashingKeeper.IterateValidatorSigningInfos(
			context,
			func(validatorConsAddress sdkTypes.ConsAddress, validatorSigningInfo slashing.ValidatorSigningInfo) (stop bool) {
				validatorSigningInfo.StartHeight = 0
				application.slashingKeeper.SetValidatorSigningInfo(context, validatorConsAddress, validatorSigningInfo)
				return false
			},
		)
	}

	genesisState := application.moduleManager.ExportGenesis(context)
	applicationState, err := codec.MarshalJSONIndent(application.codec, genesisState)

	if err != nil {
		return nil, nil, err
	}

	return applicationState, staking.WriteValidators(context, application.stakingKeeper), nil
}

func (application application) Initialize(logger log.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string, baseAppOptions ...func(*baseapp.BaseApp)) applications.Application {
	application.BaseApp = *baseapp.NewBaseApp(
		application.name,
		logger,
		db,
		auth.DefaultTxDecoder(application.codec),
		baseAppOptions...,
	)
	application.BaseApp.SetCommitMultiStoreTracer(traceStore)
	application.BaseApp.SetAppVersion(version.Version)

	application.keys = sdkTypes.NewKVStoreKeys(
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

	transientStoreKeys := sdkTypes.NewTransientStoreKeys(params.TStoreKey)

	paramsKeeper := params.NewKeeper(
		application.codec,
		application.keys[params.StoreKey],
		transientStoreKeys[params.TStoreKey],
	)

	accountKeeper := auth.NewAccountKeeper(
		application.codec,
		application.keys[auth.StoreKey],
		paramsKeeper.Subspace(auth.DefaultParamspace),
		auth.ProtoBaseAccount,
	)

	blacklistedAddresses := make(map[string]bool)
	for account := range application.moduleAccountPermissions {
		blacklistedAddresses[supply.NewModuleAddress(account).String()] = !application.tokenReceiveAllowedModules[account]
	}

	bankKeeper := bank.NewBaseKeeper(
		accountKeeper,
		paramsKeeper.Subspace(bank.DefaultParamspace),
		blacklistedAddresses,
	)

	supplyKeeper := supply.NewKeeper(
		application.codec,
		application.keys[supply.StoreKey],
		accountKeeper,
		bankKeeper,
		application.moduleAccountPermissions,
	)

	stakingKeeper := staking.NewKeeper(
		application.codec,
		application.keys[staking.StoreKey],
		supplyKeeper,
		paramsKeeper.Subspace(staking.DefaultParamspace),
	)

	mintKeeper := mint.NewKeeper(
		application.codec,
		application.keys[mint.StoreKey],
		paramsKeeper.Subspace(mint.DefaultParamspace),
		&stakingKeeper,
		supplyKeeper,
		auth.FeeCollectorName,
	)

	blackListedModuleAddresses := make(map[string]bool)
	for moduleAccount := range application.moduleAccountPermissions {
		blackListedModuleAddresses[supply.NewModuleAddress(moduleAccount).String()] = true
	}

	application.distributionKeeper = distribution.NewKeeper(
		application.codec,
		application.keys[distribution.StoreKey],
		paramsKeeper.Subspace(distribution.DefaultParamspace),
		&stakingKeeper,
		supplyKeeper,
		auth.FeeCollectorName,
		blackListedModuleAddresses,
	)
	application.slashingKeeper = slashing.NewKeeper(
		application.codec,
		application.keys[slashing.StoreKey],
		&stakingKeeper,
		paramsKeeper.Subspace(slashing.DefaultParamspace),
	)
	application.crisisKeeper = crisis.NewKeeper(
		paramsKeeper.Subspace(crisis.DefaultParamspace),
		invCheckPeriod,
		supplyKeeper,
		auth.FeeCollectorName,
	)
	upgradeKeeper := upgrade.NewKeeper(
		skipUpgradeHeights,
		application.keys[upgrade.StoreKey],
		application.codec,
	)

	evidenceKeeper := evidence.NewKeeper(
		application.codec,
		application.keys[evidence.StoreKey],
		paramsKeeper.Subspace(evidence.DefaultParamspace),
		&stakingKeeper,
		application.slashingKeeper,
	)
	evidenceRouter := evidence.NewRouter()
	evidenceKeeper.SetRouter(evidenceRouter)

	application.stakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(application.distributionKeeper.Hooks(), application.slashingKeeper.Hooks()),
	)

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
		metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(super.Auxiliary.GetName()),
		metasModule.GetAuxiliary(supplement.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(revoke.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(verify.Auxiliary.GetName()),
	)
	splitsModule := splits.Prototype().Initialize(
		application.keys[splits.Prototype().Name()],
		paramsKeeper.Subspace(splits.Prototype().Name()),
		supplyKeeper,
		identitiesModule.GetAuxiliary(authenticate.Auxiliary.GetName()),
	)
	assetsModule := assets.Prototype().Initialize(
		application.keys[assets.Prototype().Name()],
		paramsKeeper.Subspace(assets.Prototype().Name()),
		identitiesModule.GetAuxiliary(authenticate.Auxiliary.GetName()),
		splitsModule.GetAuxiliary(burn.Auxiliary.GetName()),
		classificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
		classificationsModule.GetAuxiliary(define.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(deputize.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(maintain.Auxiliary.GetName()),
		splitsModule.GetAuxiliary(renumerate.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(revoke.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(super.Auxiliary.GetName()),
		metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
		metasModule.GetAuxiliary(supplement.Auxiliary.GetName()),
		splitsModule.GetAuxiliary(splitsMint.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(verify.Auxiliary.GetName()),
	)
	ordersModule := orders.Prototype().Initialize(
		application.keys[orders.Prototype().Name()],
		paramsKeeper.Subspace(orders.Prototype().Name()),
		identitiesModule.GetAuxiliary(authenticate.Auxiliary.GetName()),
		classificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
		classificationsModule.GetAuxiliary(define.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(deputize.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(maintain.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(revoke.Auxiliary.GetName()),
		metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(super.Auxiliary.GetName()),
		metasModule.GetAuxiliary(supplement.Auxiliary.GetName()),
		splitsModule.GetAuxiliary(transfer.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(verify.Auxiliary.GetName()),
	)

	var wasmRouter = application.BaseApp.Router()

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
		application.codec,
		application.keys[wasm.StoreKey],
		paramsKeeper.Subspace(wasm.DefaultParamspace),
		accountKeeper,
		bankKeeper,
		application.stakingKeeper,
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
		params.NewParamChangeProposalHandler(paramsKeeper),
	).AddRoute(
		distribution.RouterKey,
		distribution.NewCommunityPoolSpendProposalHandler(application.distributionKeeper),
	).AddRoute(
		upgrade.RouterKey,
		upgrade.NewSoftwareUpgradeProposalHandler(upgradeKeeper),
	)

	if len(application.enabledWasmProposalTypeList) != 0 {
		govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(wasmKeeper, application.enabledWasmProposalTypeList))
	}

	govKeeper := gov.NewKeeper(
		application.codec,
		application.keys[gov.StoreKey],
		paramsKeeper.Subspace(gov.DefaultParamspace).WithKeyTable(gov.ParamKeyTable()),
		supplyKeeper,
		&stakingKeeper,
		govRouter,
	)

	application.moduleManager = module.NewManager(
		genutil.NewAppModule(accountKeeper, application.stakingKeeper, application.BaseApp.DeliverTx),
		auth.NewAppModule(accountKeeper),
		bank.NewAppModule(bankKeeper, accountKeeper),
		crisis.NewAppModule(&application.crisisKeeper),
		supply.NewAppModule(supplyKeeper, accountKeeper),
		gov.NewAppModule(govKeeper, accountKeeper, supplyKeeper),
		mint.NewAppModule(mintKeeper),
		slashing.NewAppModule(application.slashingKeeper, accountKeeper, application.stakingKeeper),
		distribution.NewAppModule(application.distributionKeeper, accountKeeper, supplyKeeper, application.stakingKeeper),
		staking.NewAppModule(application.stakingKeeper, accountKeeper, supplyKeeper),
		upgrade.NewAppModule(upgradeKeeper),
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

	application.moduleManager.SetOrderBeginBlockers(
		upgrade.ModuleName,
		mint.ModuleName,
		distribution.ModuleName,
		slashing.ModuleName,
	)
	application.moduleManager.SetOrderEndBlockers(
		crisis.ModuleName,
		gov.ModuleName,
		staking.ModuleName,
		ordersModule.Name(),
	)
	application.moduleManager.SetOrderInitGenesis(
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
	application.moduleManager.RegisterInvariants(&application.crisisKeeper)
	application.moduleManager.RegisterRoutes(application.BaseApp.Router(), application.BaseApp.QueryRouter())

	module.NewSimulationManager(
		auth.NewAppModule(accountKeeper),
		bank.NewAppModule(bankKeeper, accountKeeper),
		supply.NewAppModule(supplyKeeper, accountKeeper),
		gov.NewAppModule(govKeeper, accountKeeper, supplyKeeper),
		mint.NewAppModule(mintKeeper),
		staking.NewAppModule(application.stakingKeeper, accountKeeper, supplyKeeper),
		distribution.NewAppModule(application.distributionKeeper, accountKeeper, supplyKeeper, application.stakingKeeper),
		slashing.NewAppModule(application.slashingKeeper, accountKeeper, application.stakingKeeper),
		params.NewAppModule(),
		assets.Prototype(),
		classifications.Prototype(),
		identities.Prototype(),
		maintainers.Prototype(),
		metas.Prototype(),
		orders.Prototype(),
		splits.Prototype(),
	).RegisterStoreDecoders()

	application.BaseApp.MountKVStores(application.keys)
	application.BaseApp.MountTransientStores(transientStoreKeys)

	application.BaseApp.SetBeginBlocker(application.moduleManager.BeginBlock)
	application.BaseApp.SetEndBlocker(application.moduleManager.EndBlock)
	application.BaseApp.SetInitChainer(func(context sdkTypes.Context, requestInitChain abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
		var genesisState map[string]json.RawMessage
		application.codec.MustUnmarshalJSON(requestInitChain.AppStateBytes, &genesisState)
		return application.moduleManager.InitGenesis(context, genesisState)
	})
	application.BaseApp.SetAnteHandler(auth.NewAnteHandler(accountKeeper, supplyKeeper, ante.DefaultSigVerificationGasConsumer))

	if loadLatest {
		err := application.BaseApp.LoadLatestVersion(application.keys[baseapp.MainStoreKey])
		if err != nil {
			tendermintOS.Exit(err.Error())
		}
	}

	return &application
}
func makeCodec(moduleBasicManager module.BasicManager) *codec.Codec {
	Codec := codec.New()
	moduleBasicManager.RegisterCodec(Codec)
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	return Codec
}

func NewApplication(name string, moduleBasicManager module.BasicManager, enabledWasmProposalTypeList []wasm.ProposalType, moduleAccountPermissions map[string][]string, tokenReceiveAllowedModules map[string]bool) applications.Application {
	return &application{
		name:                        name,
		moduleBasicManager:          moduleBasicManager,
		codec:                       makeCodec(moduleBasicManager),
		enabledWasmProposalTypeList: enabledWasmProposalTypeList,
		moduleAccountPermissions:    moduleAccountPermissions,
		tokenReceiveAllowedModules:  tokenReceiveAllowedModules,
	}
}
