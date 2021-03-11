/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmClient "github.com/CosmWasm/wasmd/x/wasm/client"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptoCodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	"github.com/cosmos/cosmos-sdk/server/types"
	serverTypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilityTypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisisKeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisisTypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	distributionClient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distributionKeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distributionTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidenceKeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidenceTypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutilTypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govKeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ibctransferTypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	ibchost "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintKeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	mintTypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsClient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramsProposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingKeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingTypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeClient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradeKeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradeTypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/gogo/protobuf/grpc"
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
	"github.com/persistenceOne/persistenceSDK/schema/applications"
	"github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	wasmUtilities "github.com/persistenceOne/persistenceSDK/utilities/wasm"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/log"
	tendermintLog "github.com/tendermint/tendermint/libs/log"
	tendermintOS "github.com/tendermint/tendermint/libs/os"
	tendermintProto "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"
	"io"
	"os"
	"path/filepath"
	"testing"
)

type SimulationApplication struct {
	application        *application
	transientStoreKeys map[string]*sdkTypes.TransientStoreKey
	memoryKeys         map[string]*sdkTypes.MemoryStoreKey
	sm                 *module.SimulationManager
	subspaces          map[string]paramsTypes.Subspace

	moduleAddressPermissions   map[string][]string
	tokenReceiveAllowedModules map[string]bool

	AccountKeeper      authKeeper.AccountKeeper
	BankKeeper         bankKeeper.Keeper
	StakingKeeper      stakingKeeper.Keeper
	SlashingKeeper     slashingKeeper.Keeper
	MintKeeper         mintKeeper.Keeper
	DistributionKeeper distributionKeeper.Keeper
	GovKeeper          govKeeper.Keeper
	CrisisKeeper       crisisKeeper.Keeper
	UpgradeKeeper      upgradeKeeper.Keeper
	ParamsKeeper       paramsKeeper.Keeper
	EvidenceKeeper     evidenceKeeper.Keeper
}

func (simulationApplication SimulationApplication) RegisterAPIRoutes(server *api.Server, config config.APIConfig) {
	panic("implement me")
}

func (simulationApplication SimulationApplication) RegisterGRPCServer(context client.Context, server grpc.Server) {
	panic("implement me")
}

func (simulationApplication SimulationApplication) RegisterTxService(clientCtx client.Context) {
	panic("implement me")
}

func (simulationApplication SimulationApplication) RegisterTendermintService(clientCtx client.Context) {
	panic("implement me")
}

func (simulationApplication SimulationApplication) ListSnapshots(snapshots abciTypes.RequestListSnapshots) abciTypes.ResponseListSnapshots {
	panic("implement me")
}

func (simulationApplication SimulationApplication) OfferSnapshot(snapshot abciTypes.RequestOfferSnapshot) abciTypes.ResponseOfferSnapshot {
	panic("implement me")
}

func (simulationApplication SimulationApplication) LoadSnapshotChunk(chunk abciTypes.RequestLoadSnapshotChunk) abciTypes.ResponseLoadSnapshotChunk {
	panic("implement me")
}

func (simulationApplication SimulationApplication) ApplySnapshotChunk(chunk abciTypes.RequestApplySnapshotChunk) abciTypes.ResponseApplySnapshotChunk {
	panic("implement me")
}

func (simulationApplication SimulationApplication) LegacyAmino() *codec.LegacyAmino {
	panic("implement me")
}

var _ applications.SimulationApplication = (*SimulationApplication)(nil)

func (simulationApplication SimulationApplication) Info(requestInfo abciTypes.RequestInfo) abciTypes.ResponseInfo {
	return simulationApplication.application.baseApp.Info(requestInfo)
}

func (simulationApplication SimulationApplication) SetOption(requestSetOption abciTypes.RequestSetOption) abciTypes.ResponseSetOption {
	return simulationApplication.application.baseApp.SetOption(requestSetOption)
}

func (simulationApplication SimulationApplication) Query(requestQuery abciTypes.RequestQuery) abciTypes.ResponseQuery {
	return simulationApplication.application.baseApp.Query(requestQuery)
}

func (simulationApplication SimulationApplication) CheckTx(requestCheckTx abciTypes.RequestCheckTx) abciTypes.ResponseCheckTx {
	return simulationApplication.application.baseApp.CheckTx(requestCheckTx)
}

func (simulationApplication SimulationApplication) InitChain(requestInitChain abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
	return simulationApplication.application.baseApp.InitChain(requestInitChain)
}

func (simulationApplication SimulationApplication) BeginBlock(requestBeginBlock abciTypes.RequestBeginBlock) abciTypes.ResponseBeginBlock {
	return simulationApplication.application.baseApp.BeginBlock(requestBeginBlock)
}

func (simulationApplication SimulationApplication) DeliverTx(requestDeliverTx abciTypes.RequestDeliverTx) abciTypes.ResponseDeliverTx {
	return simulationApplication.application.baseApp.DeliverTx(requestDeliverTx)
}

func (simulationApplication SimulationApplication) EndBlock(requestEndBlock abciTypes.RequestEndBlock) abciTypes.ResponseEndBlock {
	return simulationApplication.application.baseApp.EndBlock(requestEndBlock)
}

func (simulationApplication SimulationApplication) Commit() abciTypes.ResponseCommit {
	return simulationApplication.application.baseApp.Commit()
}

func (simulationApplication SimulationApplication) LoadHeight(i int64) error {
	return simulationApplication.application.LoadHeight(i)
}

func (simulationApplication SimulationApplication) ExportApplicationStateAndValidators(forZeroHeight bool, jailWhiteList []string) (types.ExportedApp, error) {
	return simulationApplication.application.ExportApplicationStateAndValidators(forZeroHeight, jailWhiteList)
}

func (simulationApplication SimulationApplication) Initialize(applicationName string, encodingConfig applications.EncodingConfig, enabledProposals []wasm.ProposalType, moduleAccountPermissions map[string][]string, tokenReceiveAllowedModules map[string]bool, logger tendermintLog.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string, applicationOptions serverTypes.AppOptions, baseAppOptions ...func(*baseapp.BaseApp)) applications.Application {
	simulationApplication.moduleAddressPermissions = moduleAccountPermissions
	simulationApplication.tokenReceiveAllowedModules = tokenReceiveAllowedModules

	applicationCodec := encodingConfig.Marshaler
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	baseApp := baseapp.NewBaseApp(
		applicationName,
		logger,
		db,
		encodingConfig.TxConfig.TxDecoder(),
		baseAppOptions...,
	)
	baseApp.SetCommitMultiStoreTracer(traceStore)
	baseApp.SetAppVersion(version.Version)
	baseApp.SetInterfaceRegistry(interfaceRegistry)

	simulationApplication.application = &application{}
	simulationApplication.application.keys = sdkTypes.NewKVStoreKeys(
		authTypes.StoreKey, bankTypes.StoreKey, stakingTypes.StoreKey,
		mintTypes.StoreKey, distributionTypes.StoreKey, slashingTypes.StoreKey,
		govTypes.StoreKey, paramsTypes.StoreKey, ibchost.StoreKey, upgradeTypes.StoreKey,
		evidenceTypes.StoreKey, ibctransferTypes.StoreKey, capabilityTypes.StoreKey,
		wasm.StoreKey,
		assets.Prototype().Name(),
		classifications.Prototype().Name(),
		identities.Prototype().Name(),
		maintainers.Prototype().Name(),
		metas.Prototype().Name(),
		orders.Prototype().Name(),
		splits.Prototype().Name(),
	)

	simulationApplication.transientStoreKeys = sdkTypes.NewTransientStoreKeys(paramsTypes.TStoreKey)
	simulationApplication.memoryKeys = sdkTypes.NewMemoryStoreKeys(capabilityTypes.MemStoreKey)

	simulationApplication.application.baseApp = baseApp
	simulationApplication.application.legacyAmino = legacyAmino
	simulationApplication.application.applicationCodec = applicationCodec

	simulationApplication.ParamsKeeper = paramsKeeper.NewKeeper(
		applicationCodec,
		legacyAmino,
		simulationApplication.application.keys[paramsTypes.StoreKey],
		simulationApplication.transientStoreKeys[paramsTypes.TStoreKey],
	)

	simulationApplication.AccountKeeper = authKeeper.NewAccountKeeper(
		applicationCodec,
		simulationApplication.application.keys[authTypes.StoreKey],
		simulationApplication.ParamsKeeper.Subspace(authTypes.ModuleName),
		authTypes.ProtoBaseAccount,
		moduleAccountPermissions,
	)

	blacklistedAddresses := make(map[string]bool)
	for account := range moduleAccountPermissions {
		blacklistedAddresses[authTypes.NewModuleAddress(account).String()] = !tokenReceiveAllowedModules[account]
	}

	simulationApplication.BankKeeper = bankKeeper.NewBaseKeeper(
		applicationCodec,
		simulationApplication.application.keys[bankTypes.StoreKey],
		simulationApplication.AccountKeeper,
		simulationApplication.ParamsKeeper.Subspace(bankTypes.ModuleName),
		blacklistedAddresses,
	)

	simulationApplication.application.stakingKeeper = stakingKeeper.NewKeeper(
		applicationCodec,
		simulationApplication.application.keys[stakingTypes.StoreKey],
		simulationApplication.AccountKeeper,
		simulationApplication.BankKeeper,
		simulationApplication.ParamsKeeper.Subspace(stakingTypes.ModuleName),
	)
	simulationApplication.StakingKeeper = simulationApplication.application.stakingKeeper

	simulationApplication.MintKeeper = mintKeeper.NewKeeper(
		applicationCodec,
		simulationApplication.application.keys[mintTypes.StoreKey],
		simulationApplication.ParamsKeeper.Subspace(mintTypes.ModuleName),
		&simulationApplication.application.stakingKeeper,
		simulationApplication.AccountKeeper,
		simulationApplication.BankKeeper,
		authTypes.FeeCollectorName,
	)

	blackListedModuleAddresses := make(map[string]bool)
	for moduleAccount := range moduleAccountPermissions {
		blackListedModuleAddresses[authTypes.NewModuleAddress(moduleAccount).String()] = true
	}

	simulationApplication.application.distributionKeeper = distributionKeeper.NewKeeper(
		applicationCodec,
		simulationApplication.application.keys[distributionTypes.StoreKey],
		simulationApplication.ParamsKeeper.Subspace(distributionTypes.ModuleName),
		simulationApplication.AccountKeeper,
		simulationApplication.BankKeeper,
		&simulationApplication.application.stakingKeeper,
		authTypes.FeeCollectorName,
		blackListedModuleAddresses,
	)
	simulationApplication.DistributionKeeper = simulationApplication.application.distributionKeeper

	simulationApplication.application.slashingKeeper = slashingKeeper.NewKeeper(
		applicationCodec,
		simulationApplication.application.keys[slashingTypes.StoreKey],
		&simulationApplication.application.stakingKeeper,
		simulationApplication.ParamsKeeper.Subspace(slashingTypes.ModuleName),
	)
	simulationApplication.SlashingKeeper = simulationApplication.application.slashingKeeper

	simulationApplication.application.crisisKeeper = crisisKeeper.NewKeeper(
		simulationApplication.ParamsKeeper.Subspace(crisisTypes.ModuleName),
		invCheckPeriod,
		simulationApplication.BankKeeper,
		authTypes.FeeCollectorName,
	)
	simulationApplication.CrisisKeeper = simulationApplication.application.crisisKeeper

	simulationApplication.UpgradeKeeper = upgradeKeeper.NewKeeper(
		skipUpgradeHeights,
		simulationApplication.application.keys[upgradeTypes.StoreKey],
		applicationCodec,
		home,
	)

	evidenceKeeper := evidenceKeeper.NewKeeper(
		applicationCodec,
		simulationApplication.application.keys[evidenceTypes.StoreKey],
		&simulationApplication.application.stakingKeeper,
		simulationApplication.application.slashingKeeper,
	)
	simulationApplication.EvidenceKeeper = *evidenceKeeper
	govRouter := govTypes.NewRouter()
	govRouter.AddRoute(
		govTypes.RouterKey,
		govTypes.ProposalHandler,
	).AddRoute(
		paramsProposal.RouterKey,
		params.NewParamChangeProposalHandler(simulationApplication.ParamsKeeper),
	).AddRoute(
		distributionTypes.RouterKey,
		distribution.NewCommunityPoolSpendProposalHandler(simulationApplication.application.distributionKeeper),
	).AddRoute(
		upgradeTypes.RouterKey,
		upgrade.NewSoftwareUpgradeProposalHandler(simulationApplication.UpgradeKeeper),
	)

	simulationApplication.application.stakingKeeper = *simulationApplication.StakingKeeper.SetHooks(
		stakingTypes.NewMultiStakingHooks(simulationApplication.application.distributionKeeper.Hooks(), simulationApplication.application.slashingKeeper.Hooks()),
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
		simulationApplication.BankKeeper,
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

	wasmConfig, err := wasm.ReadWasmConfig(applicationOptions)
	if err != nil {
		panic("error while reading wasm config: " + err.Error())
	}

	wasmKeeper := wasm.NewKeeper(
		applicationCodec,
		simulationApplication.application.keys[wasm.StoreKey],
		simulationApplication.ParamsKeeper.Subspace(wasm.DefaultParamspace),
		simulationApplication.AccountKeeper,
		simulationApplication.BankKeeper,
		simulationApplication.application.stakingKeeper,
		simulationApplication.application.distributionKeeper,
		wasmRouter,
		wasmDir,
		wasmConfig,
		stakingTypes.ModuleName,
		&wasm.MessageEncoders{Custom: wasmUtilities.CustomEncoder(assets.Prototype(), classifications.Prototype(), identities.Prototype(), maintainers.Prototype(), metas.Prototype(), orders.Prototype(), splits.Prototype())},
		nil)

	if len(enabledProposals) != 0 {
		govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(wasmKeeper, enabledProposals))
	}

	simulationApplication.GovKeeper = govKeeper.NewKeeper(
		applicationCodec,
		simulationApplication.application.keys[govTypes.StoreKey],
		simulationApplication.ParamsKeeper.Subspace(govTypes.ModuleName).WithKeyTable(govTypes.ParamKeyTable()),
		simulationApplication.AccountKeeper,
		simulationApplication.BankKeeper,
		&simulationApplication.application.stakingKeeper,
		govRouter,
	)
	/****  Module Options ****/
	var skipGenesisInvariants = false
	opt := applicationOptions.Get(crisis.FlagSkipGenesisInvariants)
	if opt, ok := opt.(bool); ok {
		skipGenesisInvariants = opt
	}

	simulationApplication.application.moduleManager = module.NewManager(
		genutil.NewAppModule(simulationApplication.AccountKeeper, simulationApplication.application.stakingKeeper, simulationApplication.application.baseApp.DeliverTx, encodingConfig.TxConfig),
		auth.NewAppModule(applicationCodec, simulationApplication.AccountKeeper, nil),
		bank.NewAppModule(applicationCodec, simulationApplication.BankKeeper, simulationApplication.AccountKeeper),
		crisis.NewAppModule(&simulationApplication.application.crisisKeeper, skipGenesisInvariants),
		gov.NewAppModule(applicationCodec, simulationApplication.GovKeeper, simulationApplication.AccountKeeper, simulationApplication.BankKeeper),
		mint.NewAppModule(applicationCodec, simulationApplication.MintKeeper, simulationApplication.AccountKeeper),
		slashing.NewAppModule(applicationCodec, simulationApplication.application.slashingKeeper, simulationApplication.AccountKeeper, simulationApplication.BankKeeper, simulationApplication.application.stakingKeeper),
		distribution.NewAppModule(applicationCodec, simulationApplication.application.distributionKeeper, simulationApplication.AccountKeeper, simulationApplication.BankKeeper, simulationApplication.application.stakingKeeper),
		staking.NewAppModule(applicationCodec, simulationApplication.application.stakingKeeper, simulationApplication.AccountKeeper, simulationApplication.BankKeeper),
		upgrade.NewAppModule(simulationApplication.UpgradeKeeper),
		wasm.NewAppModule(&wasmKeeper, simulationApplication.application.stakingKeeper),
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
		upgradeTypes.ModuleName,
		mintTypes.ModuleName,
		distributionTypes.ModuleName,
		slashingTypes.ModuleName,
	)
	simulationApplication.application.moduleManager.SetOrderEndBlockers(
		crisisTypes.ModuleName,
		govTypes.ModuleName,
		stakingTypes.ModuleName,
		ordersModule.Name(),
	)
	simulationApplication.application.moduleManager.SetOrderInitGenesis(
		authTypes.ModuleName,
		distributionTypes.ModuleName,
		stakingTypes.ModuleName,
		bankTypes.ModuleName,
		slashingTypes.ModuleName,
		govTypes.ModuleName,
		mintTypes.ModuleName,
		crisisTypes.ModuleName,
		genutilTypes.ModuleName,
		evidenceTypes.ModuleName,
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
	simulationApplication.application.moduleManager.RegisterRoutes(simulationApplication.application.baseApp.Router(), simulationApplication.application.baseApp.QueryRouter(), legacyAmino)

	simulationApplication.sm = module.NewSimulationManager(
		auth.NewAppModule(applicationCodec, simulationApplication.AccountKeeper, nil),
		bank.NewAppModule(applicationCodec, simulationApplication.BankKeeper, simulationApplication.AccountKeeper),
		gov.NewAppModule(applicationCodec, simulationApplication.GovKeeper, simulationApplication.AccountKeeper, simulationApplication.BankKeeper),
		mint.NewAppModule(applicationCodec, simulationApplication.MintKeeper, simulationApplication.AccountKeeper),
		staking.NewAppModule(applicationCodec, simulationApplication.application.stakingKeeper, simulationApplication.AccountKeeper, simulationApplication.BankKeeper),
		distribution.NewAppModule(applicationCodec, simulationApplication.application.distributionKeeper, simulationApplication.AccountKeeper, simulationApplication.BankKeeper, simulationApplication.application.stakingKeeper),
		slashing.NewAppModule(applicationCodec, simulationApplication.application.slashingKeeper, simulationApplication.AccountKeeper, simulationApplication.BankKeeper, simulationApplication.application.stakingKeeper),
		params.NewAppModule(simulationApplication.ParamsKeeper),
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
		legacyAmino.MustUnmarshalJSON(requestInitChain.AppStateBytes, &genesisState)
		return simulationApplication.application.moduleManager.InitGenesis(context, applicationCodec, genesisState)
	})
	simulationApplication.application.baseApp.SetAnteHandler(ante.NewAnteHandler(simulationApplication.AccountKeeper, simulationApplication.BankKeeper, ante.DefaultSigVerificationGasConsumer, encodingConfig.TxConfig.SignModeHandler()))

	if loadLatest {
		err := simulationApplication.application.baseApp.LoadLatestVersion()
		if err != nil {
			tendermintOS.Exit(err.Error())
		}
	}

	return simulationApplication
}

func (simulationApplication SimulationApplication) Name() string {
	return simulationApplication.application.baseApp.Name()
}

func (simulationApplication SimulationApplication) Codec() *codec.LegacyAmino {
	return simulationApplication.application.legacyAmino
}

func (simulationApplication SimulationApplication) BeginBlocker(ctx sdkTypes.Context, req abciTypes.RequestBeginBlock) abciTypes.ResponseBeginBlock {
	return simulationApplication.application.moduleManager.BeginBlock(ctx, req)
}

func (simulationApplication SimulationApplication) EndBlocker(ctx sdkTypes.Context, req abciTypes.RequestEndBlock) abciTypes.ResponseEndBlock {
	return simulationApplication.application.moduleManager.EndBlock(ctx, req)
}

func (simulationApplication SimulationApplication) InitChainer(ctx sdkTypes.Context, req abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
	var genesisState simapp.GenesisState

	simulationApplication.application.legacyAmino.MustUnmarshalJSON(req.AppStateBytes, &genesisState)

	return simulationApplication.application.moduleManager.InitGenesis(ctx, simulationApplication.application.applicationCodec, genesisState)
}

func (simulationApplication SimulationApplication) ExportAppStateAndValidators(forZeroHeight bool, jailWhiteList []string) (serverTypes.ExportedApp, error) {
	return simulationApplication.application.ExportApplicationStateAndValidators(forZeroHeight, jailWhiteList)
}

func (simulationApplication SimulationApplication) ModuleAccountAddrs() map[string]bool {
	return simulationApplication.tokenReceiveAllowedModules
}

func (simulationApplication SimulationApplication) SimulationManager() *module.SimulationManager {
	return simulationApplication.sm
}

func (simulationApplication SimulationApplication) ModuleManager() *module.Manager {
	return simulationApplication.application.moduleManager
}

func (simulationApplication SimulationApplication) GetBaseApp() *baseapp.BaseApp {
	return simulationApplication.application.baseApp
}

func (simulationApplication SimulationApplication) GetKey(storeKey string) *sdkTypes.KVStoreKey {
	return simulationApplication.application.keys[storeKey]
}

func (simulationApplication SimulationApplication) GetTKey(storeKey string) *sdkTypes.TransientStoreKey {
	return simulationApplication.transientStoreKeys[storeKey]
}

func (simulationApplication SimulationApplication) GetSubspace(moduleName string) paramsTypes.Subspace {
	return simulationApplication.subspaces[moduleName]
}

func (simulationApplication SimulationApplication) GetModuleAccountPermissions() map[string][]string {
	return simulationApplication.moduleAddressPermissions
}

func (simulationApplication SimulationApplication) GetBlackListedAddresses() map[string]bool {
	blacklistedAddrs := make(map[string]bool)
	for acc := range ModuleAccountPermissions {
		blacklistedAddrs[authTypes.NewModuleAddress(acc).String()] = !AllowedReceivingModuleAccounts[acc]
	}

	return blacklistedAddrs
}

func (simulationApplication SimulationApplication) CheckBalance(t *testing.T, address sdkTypes.AccAddress, coins sdkTypes.Coins) {
	ctxCheck := simulationApplication.application.baseApp.NewContext(true, tendermintProto.Header{})
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
	prevSupply := simulationApplication.BankKeeper.GetSupply(context)
	simulationApplication.BankKeeper.SetSupply(context, bankTypes.NewSupply(prevSupply.GetTotal().Add(totalSupply...)))

	// fill all the addresses with some coins, set the loose pool tokens simultaneously
	for _, addr := range testAddresses {
		err := simulationApplication.BankKeeper.AddCoins(context, addr, initCoins)
		if err != nil {
			panic(err)
		}
	}

	return testAddresses
}

func (simulationApplication SimulationApplication) Setup(isCheckTx bool) applications.SimulationApplication {
	db := tendermintDB.NewMemDB()
	app := NewSimApp().Initialize(ApplicationName, base.MakeEncodingConfig(), wasm.EnableAllProposals, ModuleAccountPermissions, AllowedReceivingModuleAccounts, log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0, map[int64]bool{}, DefaultNodeHome, simapp.EmptyAppOptions{})

	if !isCheckTx {
		// init chain must be called to stop deliverState from being nil
		genesisState := ModuleBasics.DefaultGenesis(simulationApplication.application.applicationCodec)

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

	return app.(SimulationApplication)
}

func (simulationApplication SimulationApplication) SetupWithGenesisAccounts(accounts []authTypes.GenesisAccount) applications.SimulationApplication {
	db := tendermintDB.NewMemDB()
	app := NewSimApp().Initialize(ApplicationName, base.MakeEncodingConfig(), wasm.EnableAllProposals, ModuleAccountPermissions, AllowedReceivingModuleAccounts, log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0, map[int64]bool{}, DefaultNodeHome, simapp.EmptyAppOptions{})

	// initialize the chain with the passed in genesis accounts
	genesisState := ModuleBasics.DefaultGenesis(simulationApplication.application.applicationCodec)

	authGenesis := authTypes.NewGenesisState(authTypes.DefaultParams(), accounts)
	genesisStateBz := simulationApplication.Codec().MustMarshalJSON(authGenesis)
	genesisState[authTypes.ModuleName] = genesisStateBz

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
	simulationApplication.BeginBlock(abciTypes.RequestBeginBlock{Header: tendermintProto.Header{Height: simulationApplication.application.baseApp.LastBlockHeight() + 1}})

	return app.(SimulationApplication)
}

func (simulationApplication SimulationApplication) NewTestApplication(isCheckTx bool) (applications.SimulationApplication, sdkTypes.Context) {
	app := simulationApplication.Setup(isCheckTx)
	ctx := simulationApplication.GetBaseApp().NewContext(isCheckTx, tendermintProto.Header{})

	return app, ctx
}

func NewSimApp() SimulationApplication {
	return SimulationApplication{}
}

var (
	ApplicationName = "SimulationApplication"
	DefaultNodeHome = os.ExpandEnv("$HOME/.simapp")

	ModuleBasics = module.NewBasicManager(
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distribution.AppModuleBasic{},
		gov.NewAppModuleBasic(append(wasmClient.ProposalHandlers, paramsClient.ProposalHandler, distributionClient.ProposalHandler, upgradeClient.ProposalHandler)...),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		wasm.AppModuleBasic{},
		slashing.AppModuleBasic{},
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

	// ModuleAccountPermissions module account permissions
	ModuleAccountPermissions = map[string][]string{
		authTypes.FeeCollectorName:     nil,
		distributionTypes.ModuleName:   nil,
		mintTypes.ModuleName:           {authTypes.Minter},
		stakingTypes.BondedPoolName:    {authTypes.Burner, authTypes.Staking},
		stakingTypes.NotBondedPoolName: {authTypes.Burner, authTypes.Staking},
		govTypes.ModuleName:            {authTypes.Burner},
	}

	// AllowedReceivingModuleAccounts module accounts that are allowed to receive tokens
	AllowedReceivingModuleAccounts = map[string]bool{
		distributionTypes.ModuleName: true,
	}
)

func MakeCodec() *codec.LegacyAmino {
	Codec := codec.NewLegacyAmino()

	ModuleBasics.RegisterLegacyAminoCodec(Codec)
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterLegacyAminoCodec(Codec)
	cryptoCodec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vestingTypes.RegisterLegacyAminoCodec(Codec)

	return Codec
}
