/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"
	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypesModule "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
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
	auxiliariesMint "github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/mint"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/transfer"
	"github.com/persistenceOne/persistenceSDK/schema/applications"
	wasmUtilities "github.com/persistenceOne/persistenceSDK/utilities/wasm"
	"github.com/spf13/viper"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintOS "github.com/tendermint/tendermint/libs/os"
	tendermintTypes "github.com/tendermint/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"
	"honnef.co/go/tools/version"
	"io"
	"path/filepath"
)

type application struct {
	*baseapp.BaseApp
	codec *codec.Codec

	invCheckPeriod uint

	keys               map[string]*sdkTypes.KVStoreKey
	transientStoreKeys map[string]*sdkTypes.TransientStoreKey

	subspaces map[string]params.Subspace

	accountKeeper      auth.AccountKeeper
	bankKeeper         bank.Keeper
	supplyKeeper       supply.Keeper
	stakingKeeper      staking.Keeper
	slashingKeeper     slashing.Keeper
	mintKeeper         mint.Keeper
	distributionKeeper distribution.Keeper
	govKeeper          gov.Keeper
	crisisKeeper       crisis.Keeper
	paramsKeeper       params.Keeper
	upgradeKeeper      upgrade.Keeper
	evidenceKeeper     *evidence.Keeper
	wasmKeeper         wasm.Keeper

	basicManager sdkTypesModule.BasicManager

	moduleManager *sdkTypesModule.Manager

	simulationManager *sdkTypesModule.SimulationManager
}

var _ applications.Application = (*application)(nil)

func (application *application) LoadHeight(height int64) error {
	return application.LoadVersion(height, application.keys[baseapp.MainStoreKey])
}
func (application *application) ExportApplicationStateAndValidators(forZeroHeight bool, jailWhiteList []string) (applicationState json.RawMessage, validators []tendermintTypes.GenesisValidator, err error) {
	ctx := application.NewContext(true, abciTypes.Header{Height: application.LastBlockHeight()})

	if forZeroHeight {
		application.prepareForZeroHeightGenesis(ctx, jailWhiteList)
	}

	genesisState := application.moduleManager.ExportGenesis(ctx)
	applicationState, err = codec.MarshalJSONIndent(application.codec, genesisState)
	if err != nil {
		return nil, nil, err
	}
	validators = staking.WriteValidators(ctx, application.stakingKeeper)
	return applicationState, validators, nil
}
func (application *application) prepareForZeroHeightGenesis(ctx sdkTypes.Context, jailWhiteList []string) {
	applyWhiteList := false

	if len(jailWhiteList) > 0 {
		applyWhiteList = true
	}

	whiteListMap := make(map[string]bool)

	for _, address := range jailWhiteList {
		_, err := sdkTypes.ValAddressFromBech32(address)
		if err != nil {
			return
		}
		whiteListMap[address] = true
	}

	application.crisisKeeper.AssertInvariants(ctx)

	application.stakingKeeper.IterateValidators(ctx, func(_ int64, val staking.ValidatorI) (stop bool) {
		_, _ = application.distributionKeeper.WithdrawValidatorCommission(ctx, val.GetOperator())
		return false
	})

	delegations := application.stakingKeeper.GetAllDelegations(ctx)
	for _, delegation := range delegations {
		_, _ = application.distributionKeeper.WithdrawDelegationRewards(ctx, delegation.DelegatorAddress, delegation.ValidatorAddress)
	}

	application.distributionKeeper.DeleteAllValidatorSlashEvents(ctx)

	application.distributionKeeper.DeleteAllValidatorHistoricalRewards(ctx)

	height := ctx.BlockHeight()
	ctx = ctx.WithBlockHeight(0)

	application.stakingKeeper.IterateValidators(ctx, func(_ int64, val staking.ValidatorI) (stop bool) {

		scraps := application.distributionKeeper.GetValidatorOutstandingRewards(ctx, val.GetOperator())
		feePool := application.distributionKeeper.GetFeePool(ctx)
		feePool.CommunityPool = feePool.CommunityPool.Add(scraps...)
		application.distributionKeeper.SetFeePool(ctx, feePool)

		application.distributionKeeper.Hooks().AfterValidatorCreated(ctx, val.GetOperator())
		return false
	})

	for _, delegation := range delegations {
		application.distributionKeeper.Hooks().BeforeDelegationCreated(ctx, delegation.DelegatorAddress, delegation.ValidatorAddress)
		application.distributionKeeper.Hooks().AfterDelegationModified(ctx, delegation.DelegatorAddress, delegation.ValidatorAddress)
	}

	ctx = ctx.WithBlockHeight(height)

	application.stakingKeeper.IterateRedelegations(ctx, func(_ int64, redelegation staking.Redelegation) (stop bool) {
		for i := range redelegation.Entries {
			redelegation.Entries[i].CreationHeight = 0
		}
		application.stakingKeeper.SetRedelegation(ctx, redelegation)
		return false
	})

	application.stakingKeeper.IterateUnbondingDelegations(ctx, func(_ int64, unbondingDelegation staking.UnbondingDelegation) (stop bool) {
		for i := range unbondingDelegation.Entries {
			unbondingDelegation.Entries[i].CreationHeight = 0
		}
		application.stakingKeeper.SetUnbondingDelegation(ctx, unbondingDelegation)
		return false
	})

	store := ctx.KVStore(application.keys[staking.StoreKey])
	kvStoreReversePrefixIterator := sdkTypes.KVStoreReversePrefixIterator(store, staking.ValidatorsKey)
	counter := int16(0)

	for ; kvStoreReversePrefixIterator.Valid(); kvStoreReversePrefixIterator.Next() {
		addr := sdkTypes.ValAddress(kvStoreReversePrefixIterator.Key()[1:])
		validator, found := application.stakingKeeper.GetValidator(ctx, addr)
		if !found {
			panic("Validator not found!")
		}

		validator.UnbondingHeight = 0

		if applyWhiteList && !whiteListMap[addr.String()] {
			validator.Jailed = true
		}

		application.stakingKeeper.SetValidator(ctx, validator)
		counter++
	}

	kvStoreReversePrefixIterator.Close()

	_ = application.stakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)

	application.slashingKeeper.IterateValidatorSigningInfos(
		ctx,
		func(validatorConsAddress sdkTypes.ConsAddress, validatorSigningInfo slashing.ValidatorSigningInfo) (stop bool) {
			validatorSigningInfo.StartHeight = 0
			application.slashingKeeper.SetValidatorSigningInfo(ctx, validatorConsAddress, validatorSigningInfo)
			return false
		},
	)
}

func InitializeApplication(applicationName string, codec *codec.Codec, enabledProposals []wasm.ProposalType, moduleAccountPermissions map[string][]string, tokenReceiveAllowedModules map[string]bool, basicManager sdkTypesModule.BasicManager) applications.NewApplication {
	return func(logger log.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string, baseAppOptions ...func(*baseapp.BaseApp)) applications.Application {

		baseApp := baseapp.NewBaseApp(
			applicationName,
			logger,
			db,
			auth.DefaultTxDecoder(codec),
			baseAppOptions...,
		)
		baseApp.SetCommitMultiStoreTracer(traceStore)
		baseApp.SetAppVersion(version.Version)

		keys := sdkTypes.NewKVStoreKeys(
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
		)
		keys[assets.Module.Name()] = assets.Module.GetKVStoreKey()
		keys[classifications.Module.Name()] = classifications.Module.GetKVStoreKey()
		keys[identities.Module.Name()] = identities.Module.GetKVStoreKey()
		keys[maintainers.Module.Name()] = maintainers.Module.GetKVStoreKey()
		keys[metas.Module.Name()] = metas.Module.GetKVStoreKey()
		keys[orders.Module.Name()] = orders.Module.GetKVStoreKey()
		keys[splits.Module.Name()] = splits.Module.GetKVStoreKey()

		transientStoreKeys := sdkTypes.NewTransientStoreKeys(params.TStoreKey)

		var application = &application{
			BaseApp:            baseApp,
			codec:              codec,
			invCheckPeriod:     invCheckPeriod,
			keys:               keys,
			transientStoreKeys: transientStoreKeys,
			subspaces:          make(map[string]params.Subspace),
			basicManager:       basicManager,
		}
		application.paramsKeeper = params.NewKeeper(
			codec,
			keys[params.StoreKey],
			transientStoreKeys[params.TStoreKey],
		)
		application.subspaces[auth.ModuleName] = application.paramsKeeper.Subspace(auth.DefaultParamspace)
		application.subspaces[bank.ModuleName] = application.paramsKeeper.Subspace(bank.DefaultParamspace)
		application.subspaces[staking.ModuleName] = application.paramsKeeper.Subspace(staking.DefaultParamspace)
		application.subspaces[mint.ModuleName] = application.paramsKeeper.Subspace(mint.DefaultParamspace)
		application.subspaces[distribution.ModuleName] = application.paramsKeeper.Subspace(distribution.DefaultParamspace)
		application.subspaces[slashing.ModuleName] = application.paramsKeeper.Subspace(slashing.DefaultParamspace)
		application.subspaces[gov.ModuleName] = application.paramsKeeper.Subspace(gov.DefaultParamspace).WithKeyTable(gov.ParamKeyTable())
		application.subspaces[crisis.ModuleName] = application.paramsKeeper.Subspace(crisis.DefaultParamspace)
		application.subspaces[evidence.ModuleName] = application.paramsKeeper.Subspace(evidence.DefaultParamspace)
		application.subspaces[wasm.ModuleName] = application.paramsKeeper.Subspace(wasm.DefaultParamspace)

		application.subspaces[assets.Module.Name()] = application.paramsKeeper.Subspace(assets.Module.GetDefaultParamspace())
		application.subspaces[classifications.Module.Name()] = application.paramsKeeper.Subspace(classifications.Module.GetDefaultParamspace())
		application.subspaces[identities.Module.Name()] = application.paramsKeeper.Subspace(identities.Module.GetDefaultParamspace())
		application.subspaces[maintainers.Module.Name()] = application.paramsKeeper.Subspace(maintainers.Module.GetDefaultParamspace())
		application.subspaces[metas.Module.Name()] = application.paramsKeeper.Subspace(metas.Module.GetDefaultParamspace())
		application.subspaces[orders.Module.Name()] = application.paramsKeeper.Subspace(orders.Module.GetDefaultParamspace())
		application.subspaces[splits.Module.Name()] = application.paramsKeeper.Subspace(splits.Module.GetDefaultParamspace())

		application.accountKeeper = auth.NewAccountKeeper(
			codec,
			keys[auth.StoreKey],
			application.subspaces[auth.ModuleName],
			auth.ProtoBaseAccount,
		)

		blacklistedAddresses := make(map[string]bool)
		for account := range moduleAccountPermissions {
			blacklistedAddresses[supply.NewModuleAddress(account).String()] = !tokenReceiveAllowedModules[account]
		}
		application.bankKeeper = bank.NewBaseKeeper(
			application.accountKeeper,
			application.subspaces[bank.ModuleName],
			blacklistedAddresses,
		)

		application.supplyKeeper = supply.NewKeeper(
			application.codec,
			keys[supply.StoreKey],
			application.accountKeeper,
			application.bankKeeper,
			moduleAccountPermissions,
		)

		stakingKeeper := staking.NewKeeper(
			application.codec,
			keys[staking.StoreKey],
			application.supplyKeeper,
			application.subspaces[staking.ModuleName],
		)

		application.mintKeeper = mint.NewKeeper(
			codec,
			keys[mint.StoreKey],
			application.subspaces[mint.ModuleName],
			&stakingKeeper,
			application.supplyKeeper,
			auth.FeeCollectorName,
		)

		blackListedModuleAddresses := make(map[string]bool)
		for moduleAccount := range moduleAccountPermissions {
			blackListedModuleAddresses[supply.NewModuleAddress(moduleAccount).String()] = true
		}
		application.distributionKeeper = distribution.NewKeeper(
			codec,
			keys[distribution.StoreKey],
			application.subspaces[distribution.ModuleName],
			&stakingKeeper,
			application.supplyKeeper,
			auth.FeeCollectorName,
			blackListedModuleAddresses,
		)
		application.slashingKeeper = slashing.NewKeeper(
			codec,
			keys[slashing.StoreKey],
			&stakingKeeper,
			application.subspaces[slashing.ModuleName],
		)
		application.crisisKeeper = crisis.NewKeeper(
			application.subspaces[crisis.ModuleName],
			invCheckPeriod,
			application.supplyKeeper,
			auth.FeeCollectorName,
		)
		application.upgradeKeeper = upgrade.NewKeeper(
			skipUpgradeHeights,
			keys[upgrade.StoreKey],
			codec,
		)

		evidenceKeeper := evidence.NewKeeper(
			codec,
			keys[evidence.StoreKey],
			application.subspaces[evidence.ModuleName],
			&stakingKeeper,
			application.slashingKeeper,
		)
		evidenceRouter := evidence.NewRouter()
		evidenceKeeper.SetRouter(evidenceRouter)
		application.evidenceKeeper = evidenceKeeper

		govRouter := gov.NewRouter()
		govRouter.AddRoute(
			gov.RouterKey,
			gov.ProposalHandler,
		).AddRoute(
			params.RouterKey,
			params.NewParamChangeProposalHandler(application.paramsKeeper),
		).AddRoute(
			distribution.RouterKey,
			distribution.NewCommunityPoolSpendProposalHandler(application.distributionKeeper),
		).AddRoute(
			upgrade.RouterKey,
			upgrade.NewSoftwareUpgradeProposalHandler(application.upgradeKeeper),
		)

		application.stakingKeeper = *stakingKeeper.SetHooks(
			staking.NewMultiStakingHooks(application.distributionKeeper.Hooks(), application.slashingKeeper.Hooks()),
		)

		metasModule := metas.Module.Initialize()
		maintainersModule := maintainers.Module.Initialize()
		classificationsModule := classifications.Module.Initialize(metasModule.GetAuxiliary(scrub.AuxiliaryName))
		identitiesModule := identities.Module.Initialize(
			classificationsModule.GetAuxiliary(conform.AuxiliaryName),
			classificationsModule.GetAuxiliary(define.AuxiliaryName),
			maintainersModule.GetAuxiliary(super.AuxiliaryName),
			maintainersModule.GetAuxiliary(maintain.AuxiliaryName),
			metasModule.GetAuxiliary(scrub.AuxiliaryName),
		)
		splitsModule := splits.Module.Initialize(
			application.supplyKeeper,
			identitiesModule.GetAuxiliary(verify.AuxiliaryName),
		)
		assets.Module.Initialize(
			classificationsModule.GetAuxiliary(conform.AuxiliaryName),
			classificationsModule.GetAuxiliary(define.AuxiliaryName),
			identitiesModule.GetAuxiliary(verify.AuxiliaryName),
			maintainersModule.GetAuxiliary(super.AuxiliaryName),
			maintainersModule.GetAuxiliary(maintain.AuxiliaryName),
			metasModule.GetAuxiliary(scrub.AuxiliaryName),
			metasModule.GetAuxiliary(supplement.AuxiliaryName),
			splitsModule.GetAuxiliary(auxiliariesMint.AuxiliaryName),
			splitsModule.GetAuxiliary(burn.AuxiliaryName),
		)
		orders.Module.Initialize(
			application.bankKeeper,
			classificationsModule.GetAuxiliary(conform.AuxiliaryName),
			classificationsModule.GetAuxiliary(define.AuxiliaryName),
			metasModule.GetAuxiliary(supplement.AuxiliaryName),
			splitsModule.GetAuxiliary(auxiliariesMint.AuxiliaryName),
			maintainersModule.GetAuxiliary(super.AuxiliaryName),
			maintainersModule.GetAuxiliary(maintain.AuxiliaryName),
			metasModule.GetAuxiliary(scrub.AuxiliaryName),
			splitsModule.GetAuxiliary(transfer.AuxiliaryName),
			identitiesModule.GetAuxiliary(verify.AuxiliaryName),
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

		supportedFeatures := "staking"
		application.wasmKeeper = wasm.NewKeeper(
			codec,
			keys[wasm.StoreKey],
			application.subspaces[wasm.ModuleName],
			application.accountKeeper,
			application.bankKeeper,
			application.stakingKeeper,
			wasmRouter,
			wasmDir,
			wasmConfig,
			supportedFeatures,
			&wasm.MessageEncoders{Custom: wasmUtilities.CustomEncoder(assets.Module, classifications.Module, identities.Module, maintainers.Module, metas.Module, orders.Module, splits.Module)},
			nil)

		// The gov proposal types can be individually enabled
		if len(enabledProposals) != 0 {
			govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(application.wasmKeeper, enabledProposals))
		}

		application.govKeeper = gov.NewKeeper(
			codec,
			keys[gov.StoreKey],
			application.subspaces[gov.ModuleName],
			application.supplyKeeper,
			&stakingKeeper,
			govRouter,
		)

		application.moduleManager = sdkTypesModule.NewManager(
			genutil.NewAppModule(application.accountKeeper, application.stakingKeeper, application.BaseApp.DeliverTx),
			auth.NewAppModule(application.accountKeeper),
			bank.NewAppModule(application.bankKeeper, application.accountKeeper),
			crisis.NewAppModule(&application.crisisKeeper),
			supply.NewAppModule(application.supplyKeeper, application.accountKeeper),
			gov.NewAppModule(application.govKeeper, application.accountKeeper, application.supplyKeeper),
			mint.NewAppModule(application.mintKeeper),
			slashing.NewAppModule(application.slashingKeeper, application.accountKeeper, application.stakingKeeper),
			distribution.NewAppModule(application.distributionKeeper, application.accountKeeper, application.supplyKeeper, application.stakingKeeper),
			staking.NewAppModule(application.stakingKeeper, application.accountKeeper, application.supplyKeeper),
			upgrade.NewAppModule(application.upgradeKeeper),
			wasm.NewAppModule(application.wasmKeeper),
			evidence.NewAppModule(*application.evidenceKeeper),

			assets.Module,
			classifications.Module,
			identities.Module,
			metas.Module,
			orders.Module,
			splits.Module,
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
			assets.Module.Name(),
			classifications.Module.Name(),
			identities.Module.Name(),
			metas.Module.Name(),
			orders.Module.Name(),
			splits.Module.Name(),
		)
		application.moduleManager.RegisterInvariants(&application.crisisKeeper)
		application.moduleManager.RegisterRoutes(application.Router(), application.QueryRouter())

		application.simulationManager = sdkTypesModule.NewSimulationManager(
			auth.NewAppModule(application.accountKeeper),
			bank.NewAppModule(application.bankKeeper, application.accountKeeper),
			supply.NewAppModule(application.supplyKeeper, application.accountKeeper),
			gov.NewAppModule(application.govKeeper, application.accountKeeper, application.supplyKeeper),
			mint.NewAppModule(application.mintKeeper),
			staking.NewAppModule(application.stakingKeeper, application.accountKeeper, application.supplyKeeper),
			distribution.NewAppModule(application.distributionKeeper, application.accountKeeper, application.supplyKeeper, application.stakingKeeper),
			slashing.NewAppModule(application.slashingKeeper, application.accountKeeper, application.stakingKeeper),
			params.NewAppModule(),
			assets.Module,
			classifications.Module,
			identities.Module,
			metas.Module,
			orders.Module,
			splits.Module,
		)

		application.simulationManager.RegisterStoreDecoders()

		application.MountKVStores(keys)
		application.MountTransientStores(transientStoreKeys)

		application.SetBeginBlocker(application.moduleManager.BeginBlock)
		application.SetEndBlocker(application.moduleManager.EndBlock)
		application.SetInitChainer(func(context sdkTypes.Context, requestInitChain abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
			var genesisState map[string]json.RawMessage
			application.codec.MustUnmarshalJSON(requestInitChain.AppStateBytes, &genesisState)
			return application.moduleManager.InitGenesis(context, genesisState)
		})
		application.SetAnteHandler(auth.NewAnteHandler(application.accountKeeper, application.supplyKeeper, ante.DefaultSigVerificationGasConsumer))

		if loadLatest {
			err := application.LoadLatestVersion(application.keys[baseapp.MainStoreKey])
			if err != nil {
				tendermintOS.Exit(err.Error())
			}
		}

		return application
	}
}
