/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	wasmClient "github.com/CosmWasm/wasmd/x/wasm/client"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	paramsClient "github.com/cosmos/cosmos-sdk/x/params/client"
	upgradeClient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	"github.com/persistenceOne/persistenceSDK/modules/assets"
	"github.com/persistenceOne/persistenceSDK/modules/classifications"
	"github.com/persistenceOne/persistenceSDK/modules/identities"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers"
	"github.com/persistenceOne/persistenceSDK/modules/metas"
	"github.com/persistenceOne/persistenceSDK/modules/orders"
	"github.com/persistenceOne/persistenceSDK/modules/splits"

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
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/schema/applications"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	tendermintTypes "github.com/tendermint/tendermint/types"
)

type simulationApplication struct {
	application   *application
	transientKeys map[string]*sdkTypes.TransientStoreKey
	sm            *module.SimulationManager
	subspaces     map[string]params.Subspace

	AccountKeeper  auth.AccountKeeper
	BankKeeper     bank.Keeper
	SupplyKeeper   supply.Keeper
	StakingKeeper  staking.Keeper
	SlashingKeeper slashing.Keeper
	MintKeeper     mint.Keeper
	DistrKeeper    distribution.Keeper
	GovKeeper      gov.Keeper
	CrisisKeeper   crisis.Keeper
	UpgradeKeeper  upgrade.Keeper
	ParamsKeeper   params.Keeper
	EvidenceKeeper evidence.Keeper
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
	panic("implement me")
}

func (simulationApplication simulationApplication) ExportApplicationStateAndValidators(b bool, strings []string) (json.RawMessage, []tendermintTypes.GenesisValidator, error) {
	panic("implement me")
}

func (simulationApplication simulationApplication) Initialize(applicationName string, codec *codec.Codec, enabledProposals []wasm.ProposalType, moduleAccountPermissions map[string][]string, tokenReceiveAllowedModules map[string]bool, logger log.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string, baseAppOptions ...func(*baseapp.BaseApp)) applications.Application {
	panic("implement me")
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
	return simulationApplication.transientKeys[storeKey]
}

func (simulationApplication simulationApplication) GetSubspace(moduleName string) params.Subspace {
	return simulationApplication.subspaces[moduleName]
}

func (simulationApplication simulationApplication) GetMaccPerms() map[string][]string {
	panic("implement me")
}

func (simulationApplication simulationApplication) BlacklistedAccAddrs() map[string]bool {
	panic("implement me")
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
	app := NewApplication().Initialize(app, MakeCodec(), wasm.EnableAllProposals, maccPerms, allowedReceivingModAcc, logger, db, traceStore, loadLatest, invCheckPeriod, skipUpgradeHeights, home, baseAppOptions...)
	return simulationApplication{
		application: app.(*application),
	}
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
