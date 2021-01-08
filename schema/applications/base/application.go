/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
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

	keys map[string]*sdkTypes.KVStoreKey

	stakingKeeper      staking.Keeper
	slashingKeeper     slashing.Keeper
	distributionKeeper distribution.Keeper
	crisisKeeper       crisis.Keeper

	moduleManager *sdkTypesModule.Manager
}

var _ applications.Application = (*application)(nil)

func (application *application) LoadHeight(height int64) error {
	return application.LoadVersion(height, application.keys[baseapp.MainStoreKey])
}
func (application *application) ExportApplicationStateAndValidators(forZeroHeight bool, jailWhiteList []string) (json.RawMessage, []tendermintTypes.GenesisValidator, error) {
	context := application.NewContext(true, abciTypes.Header{Height: application.LastBlockHeight()})

	if forZeroHeight {
		applyWhiteList := false

		if len(jailWhiteList) > 0 {
			applyWhiteList = true
		}

		whiteListMap := make(map[string]bool)

		for _, address := range jailWhiteList {
			if _, Error := sdkTypes.ValAddressFromBech32(address); Error != nil {
				panic(Error)
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
	applicationState, Error := codec.MarshalJSONIndent(application.codec, genesisState)
	if Error != nil {
		return nil, nil, Error
	}
	return applicationState, staking.WriteValidators(context, application.stakingKeeper), nil
}

func Prototype(applicationName string, codec *codec.Codec, enabledProposals []wasm.ProposalType, moduleAccountPermissions map[string][]string, tokenReceiveAllowedModules map[string]bool) applications.NewApplication {
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
			assets.Prototype().Name(),
			classifications.Prototype().Name(),
			identities.Prototype().Name(),
			maintainers.Prototype().Name(),
			metas.Prototype().Name(),
			orders.Prototype().Name(),
			splits.Prototype().Name(),
		)

		transientStoreKeys := sdkTypes.NewTransientStoreKeys(params.TStoreKey)

		var application = &application{
			BaseApp: baseApp,
			codec:   codec,
			keys:    keys,
		}
		paramsKeeper := params.NewKeeper(
			codec,
			keys[params.StoreKey],
			transientStoreKeys[params.TStoreKey],
		)

		accountKeeper := auth.NewAccountKeeper(
			codec,
			keys[auth.StoreKey],
			paramsKeeper.Subspace(auth.DefaultParamspace),
			auth.ProtoBaseAccount,
		)

		blacklistedAddresses := make(map[string]bool)
		for account := range moduleAccountPermissions {
			blacklistedAddresses[supply.NewModuleAddress(account).String()] = !tokenReceiveAllowedModules[account]
		}
		bankKeeper := bank.NewBaseKeeper(
			accountKeeper,
			paramsKeeper.Subspace(bank.DefaultParamspace),
			blacklistedAddresses,
		)

		supplyKeeper := supply.NewKeeper(
			codec,
			keys[supply.StoreKey],
			accountKeeper,
			bankKeeper,
			moduleAccountPermissions,
		)

		stakingKeeper := staking.NewKeeper(
			codec,
			keys[staking.StoreKey],
			supplyKeeper,
			paramsKeeper.Subspace(staking.DefaultParamspace),
		)

		mintKeeper := mint.NewKeeper(
			codec,
			keys[mint.StoreKey],
			paramsKeeper.Subspace(mint.DefaultParamspace),
			&stakingKeeper,
			supplyKeeper,
			auth.FeeCollectorName,
		)

		blackListedModuleAddresses := make(map[string]bool)
		for moduleAccount := range moduleAccountPermissions {
			blackListedModuleAddresses[supply.NewModuleAddress(moduleAccount).String()] = true
		}
		application.distributionKeeper = distribution.NewKeeper(
			codec,
			keys[distribution.StoreKey],
			paramsKeeper.Subspace(distribution.DefaultParamspace),
			&stakingKeeper,
			supplyKeeper,
			auth.FeeCollectorName,
			blackListedModuleAddresses,
		)
		application.slashingKeeper = slashing.NewKeeper(
			codec,
			keys[slashing.StoreKey],
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
			keys[upgrade.StoreKey],
			codec,
		)

		evidenceKeeper := evidence.NewKeeper(
			codec,
			keys[evidence.StoreKey],
			paramsKeeper.Subspace(evidence.DefaultParamspace),
			&stakingKeeper,
			application.slashingKeeper,
		)
		evidenceRouter := evidence.NewRouter()
		evidenceKeeper.SetRouter(evidenceRouter)

		govRouter := gov.NewRouter()
		govRouter.AddRoute(
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

		application.stakingKeeper = *stakingKeeper.SetHooks(
			staking.NewMultiStakingHooks(application.distributionKeeper.Hooks(), application.slashingKeeper.Hooks()),
		)

		metasModule := metas.Prototype().Initialize(
			keys[metas.Prototype().Name()],
			paramsKeeper.Subspace(metas.Prototype().Name()),
		)
		maintainersModule := maintainers.Prototype().Initialize(
			keys[metas.Prototype().Name()],
			paramsKeeper.Subspace(maintainers.Prototype().Name()),
		)
		classificationsModule := classifications.Prototype().Initialize(
			keys[classifications.Prototype().Name()],
			paramsKeeper.Subspace(classifications.Prototype().Name()),
			metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
		)
		identitiesModule := identities.Prototype().Initialize(
			keys[identities.Prototype().Name()],
			paramsKeeper.Subspace(identities.Prototype().Name()),
			classificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
			classificationsModule.GetAuxiliary(define.Auxiliary.GetName()),
			maintainersModule.GetAuxiliary(super.Auxiliary.GetName()),
			maintainersModule.GetAuxiliary(maintain.Auxiliary.GetName()),
			metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
		)
		splitsModule := splits.Prototype().Initialize(
			keys[splits.Prototype().Name()],
			paramsKeeper.Subspace(splits.Prototype().Name()),
			supplyKeeper,
			identitiesModule.GetAuxiliary(verify.Auxiliary.GetName()),
		)
		assetsModule := assets.Prototype().Initialize(
			keys[assets.Prototype().Name()],
			paramsKeeper.Subspace(assets.Prototype().Name()),
			classificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
			classificationsModule.GetAuxiliary(define.Auxiliary.GetName()),
			identitiesModule.GetAuxiliary(verify.Auxiliary.GetName()),
			maintainersModule.GetAuxiliary(super.Auxiliary.GetName()),
			maintainersModule.GetAuxiliary(maintain.Auxiliary.GetName()),
			metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
			metasModule.GetAuxiliary(supplement.Auxiliary.GetName()),
			splitsModule.GetAuxiliary(auxiliariesMint.Auxiliary.GetName()),
			splitsModule.GetAuxiliary(burn.Auxiliary.GetName()),
		)
		ordersModule := orders.Prototype().Initialize(
			keys[orders.Prototype().Name()],
			paramsKeeper.Subspace(orders.Prototype().Name()),
			bankKeeper,
			classificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
			classificationsModule.GetAuxiliary(define.Auxiliary.GetName()),
			metasModule.GetAuxiliary(supplement.Auxiliary.GetName()),
			splitsModule.GetAuxiliary(auxiliariesMint.Auxiliary.GetName()),
			maintainersModule.GetAuxiliary(super.Auxiliary.GetName()),
			maintainersModule.GetAuxiliary(maintain.Auxiliary.GetName()),
			metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
			splitsModule.GetAuxiliary(transfer.Auxiliary.GetName()),
			identitiesModule.GetAuxiliary(verify.Auxiliary.GetName()),
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
			keys[wasm.StoreKey],
			paramsKeeper.Subspace(wasm.DefaultParamspace),
			accountKeeper,
			bankKeeper,
			application.stakingKeeper,
			wasmRouter,
			wasmDir,
			wasmConfig,
			"staking",
			&wasm.MessageEncoders{Custom: wasmUtilities.CustomEncoder(assets.Prototype(), classifications.Prototype(), identities.Prototype(), maintainers.Prototype(), metas.Prototype(), orders.Prototype(), splits.Prototype())},
			nil)

		if len(enabledProposals) != 0 {
			govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(wasmKeeper, enabledProposals))
		}

		govKeeper := gov.NewKeeper(
			codec,
			keys[gov.StoreKey],
			paramsKeeper.Subspace(gov.DefaultParamspace).WithKeyTable(gov.ParamKeyTable()),
			supplyKeeper,
			&stakingKeeper,
			govRouter,
		)

		application.moduleManager = sdkTypesModule.NewManager(
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
		application.moduleManager.RegisterRoutes(application.Router(), application.QueryRouter())

		simulationManager := sdkTypesModule.NewSimulationManager(
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
		)

		simulationManager.RegisterStoreDecoders()

		application.MountKVStores(keys)
		application.MountTransientStores(transientStoreKeys)

		application.SetBeginBlocker(application.moduleManager.BeginBlock)
		application.SetEndBlocker(application.moduleManager.EndBlock)
		application.SetInitChainer(func(context sdkTypes.Context, requestInitChain abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
			var genesisState map[string]json.RawMessage
			codec.MustUnmarshalJSON(requestInitChain.AppStateBytes, &genesisState)
			return application.moduleManager.InitGenesis(context, genesisState)
		})
		application.SetAnteHandler(auth.NewAnteHandler(accountKeeper, supplyKeeper, ante.DefaultSigVerificationGasConsumer))

		if loadLatest {
			err := application.LoadLatestVersion(application.keys[baseapp.MainStoreKey])
			if err != nil {
				tendermintOS.Exit(err.Error())
			}
		}

		return application
	}
}
