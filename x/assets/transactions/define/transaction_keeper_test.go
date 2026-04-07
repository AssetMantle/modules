// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"cosmossdk.io/math"
	"context"
	"github.com/AssetMantle/modules/helpers/base"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/super"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/lists"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/cometbft/cometbft/crypto/ed25519"
	storeTypes "cosmossdk.io/store/types"
	addressCodec "github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/runtime"
	authKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/mock"
	"math/rand"
	"testing"

	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	cosmosDB "github.com/cosmos/cosmos-db"
	"cosmossdk.io/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"cosmossdk.io/store"
	storeMetrics "cosmossdk.io/store/metrics"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	dataHelper "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/x/assets/mapper"
	permissionHelper "github.com/AssetMantle/modules/x/assets/utilities"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/define"
)

type MockAuxiliary struct {
	mock.Mock
}

var _ helpers.Auxiliary = (*MockAuxiliary)(nil)

func (mockAuxiliary *MockAuxiliary) GetName() string { panic(mockAuxiliary) }
func (mockAuxiliary *MockAuxiliary) GetKeeper() helpers.AuxiliaryKeeper {
	args := mockAuxiliary.Called()
	return args.Get(0).(helpers.AuxiliaryKeeper)
}
func (mockAuxiliary *MockAuxiliary) Initialize(_ helpers.Mapper, _ helpers.ParameterManager, _ ...interface{}) helpers.Auxiliary {
	panic(mockAuxiliary)
}

type MockAuxiliaryKeeper struct {
	mock.Mock
}

var _ helpers.AuxiliaryKeeper = (*MockAuxiliaryKeeper)(nil)

func (mockAuxiliaryKeeper *MockAuxiliaryKeeper) Help(context context.Context, auxiliaryRequest helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	args := mockAuxiliaryKeeper.Called(context, auxiliaryRequest)
	return args.Get(0).(helpers.AuxiliaryResponse), args.Error(1)
}
func (mockAuxiliaryKeeper *MockAuxiliaryKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, i []interface{}) helpers.Keeper {
	args := mockAuxiliaryKeeper.Called(mapper, parameterManager, i)
	return args.Get(0).(helpers.Keeper)
}

const (
	TestMinterModuleName = "testMinter"
	Denom                = "stake"
	ChainID              = "testChain"
	GenesisSupply        = 1000000000000
)

var (
	moduleStoreKey = storeTypes.NewKVStoreKey(constants.ModuleName)

	authenticateAuxiliaryKeeper         = new(MockAuxiliaryKeeper)
	authenticateAuxiliaryFailureAddress = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	_                                   = authenticateAuxiliaryKeeper.On("Help", mock.Anything, authenticate.NewAuxiliaryRequest(&Message{From: authenticateAuxiliaryFailureAddress.String(), FromID: baseIDs.PrototypeIdentityID().(*baseIDs.IdentityID)})).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	_                                   = authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	authenticateAuxiliary               = new(MockAuxiliary)
	_                                   = authenticateAuxiliary.On("GetKeeper").Return(authenticateAuxiliaryKeeper)

	defineAuxiliaryKeeper               = new(MockAuxiliaryKeeper)
	defineAuxiliaryKeeperFailureAddress = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	_                                   = defineAuxiliaryKeeper.On("Help", mock.Anything, define.NewAuxiliaryRequest(defineAuxiliaryKeeperFailureAddress, baseQualified.NewImmutables(baseLists.NewPropertyList()), baseQualified.NewMutables(baseLists.NewPropertyList()))).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	_                                   = defineAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(define.NewAuxiliaryResponse(baseIDs.PrototypeClassificationID()), nil)
	defineAuxiliary                     = new(MockAuxiliary)
	_                                   = defineAuxiliary.On("GetKeeper").Return(defineAuxiliaryKeeper)

	superAuxiliaryKeeper          = new(MockAuxiliaryKeeper)
	superAuxiliaryMutablesFailure = dataHelper.GenerateRandomMetaPropertyListWithoutData(rand.New(rand.NewSource(99)))
	_                             = superAuxiliaryKeeper.On("Help", mock.Anything, super.NewAuxiliaryRequest(baseIDs.PrototypeClassificationID(), baseIDs.PrototypeIdentityID(), baseQualified.NewMutables(superAuxiliaryMutablesFailure), permissionHelper.SetModulePermissions(true, true, true)...)).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	_                             = superAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	superAuxiliary                = new(MockAuxiliary)
	_                             = superAuxiliary.On("GetKeeper").Return(superAuxiliaryKeeper)

	paramsStoreKey           = storeTypes.NewKVStoreKey(paramsTypes.StoreKey)
	paramsTransientStoreKeys = storeTypes.NewTransientStoreKey(paramsTypes.TStoreKey)

	codec = base.TestCodec()

	authStoreKey             = storeTypes.NewKVStoreKey(authTypes.StoreKey)
	moduleAccountPermissions = map[string][]string{TestMinterModuleName: {authTypes.Minter}, constants.ModuleName: nil}
	AuthKeeper               = authKeeper.NewAccountKeeper(codec, runtime.NewKVStoreService(authStoreKey), authTypes.ProtoBaseAccount, moduleAccountPermissions, addressCodec.NewBech32Codec(sdkTypes.GetConfig().GetBech32AccountAddrPrefix()), sdkTypes.GetConfig().GetBech32AccountAddrPrefix(), authTypes.NewModuleAddress(govTypes.ModuleName).String())

	bankStoreKey         = storeTypes.NewKVStoreKey(bankTypes.StoreKey)
	blacklistedAddresses = map[string]bool{authTypes.NewModuleAddress(TestMinterModuleName).String(): false, authTypes.NewModuleAddress(constants.ModuleName).String(): false}
	BankKeeper           = bankKeeper.NewBaseKeeper(codec, runtime.NewKVStoreService(bankStoreKey), AuthKeeper, blacklistedAddresses, authTypes.NewModuleAddress(govTypes.ModuleName).String(), log.NewNopLogger())

	Context = setContext()

	coinSupply = sdkTypes.NewCoins(sdkTypes.NewCoin(Denom, math.NewInt(GenesisSupply)))
	_          = BankKeeper.MintCoins(Context, TestMinterModuleName, coinSupply)

	genesisAddress = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	_              = BankKeeper.SendCoinsFromModuleToAccount(Context, TestMinterModuleName, genesisAddress, coinSupply)

	TransactionKeeper = transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey),
		defineAuxiliary,
		superAuxiliary,
		authenticateAuxiliary,
	}
)

func setContext() sdkTypes.Context {
	memDB := cosmosDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB, log.NewNopLogger(), storeMetrics.NewNoOpMetrics())
	commitMultiStore.MountStoreWithDB(moduleStoreKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(authStoreKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(bankStoreKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, storeTypes.StoreTypeTransient, memDB)
	_ = commitMultiStore.LoadLatestVersion()
	return sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{ChainID: ChainID}, false, log.NewNopLogger())
}

func TestTransactionKeeperTransact(t *testing.T) {
	type args struct {
		from               sdkTypes.AccAddress
		fromID             ids.IdentityID
		immutableMetaProps lists.PropertyList
		immutableProps     lists.PropertyList
		mutableMetaProps   lists.PropertyList
		mutableProps       lists.PropertyList
	}

	tests := []struct {
		name    string
		args    args
		setup   func()
		want    *TransactionResponse
		wantErr helpers.Error
	}{
		{
			name: "DefineTransactionKeeperSuccess",
			args: args{
				from:               genesisAddress,
				fromID:             baseIDs.PrototypeIdentityID(),
				immutableMetaProps: baseLists.NewPropertyList(),
				immutableProps:     baseLists.NewPropertyList(),
				mutableMetaProps:   baseLists.NewPropertyList(),
				mutableProps:       baseLists.NewPropertyList(),
			},
			setup: func() {
			},
			want:    newTransactionResponse(baseIDs.PrototypeClassificationID()),
			wantErr: nil,
		},
		{
			name: "AuthenticationFailure",
			args: args{
				from:               authenticateAuxiliaryFailureAddress,
				fromID:             baseIDs.PrototypeIdentityID(),
				immutableMetaProps: baseLists.NewPropertyList(),
				immutableProps:     baseLists.NewPropertyList(),
				mutableMetaProps:   baseLists.NewPropertyList(),
				mutableProps:       baseLists.NewPropertyList(),
			},
			setup: func() {
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "DefineAuxiliaryFailure",
			args: args{
				from:               defineAuxiliaryKeeperFailureAddress,
				fromID:             baseIDs.PrototypeIdentityID(),
				immutableMetaProps: baseLists.NewPropertyList(),
				immutableProps:     baseLists.NewPropertyList(),
				mutableMetaProps:   baseLists.NewPropertyList(),
				mutableProps:       baseLists.NewPropertyList(),
			},
			setup: func() {

			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "SuperAuxiliaryFailure",
			args: args{
				from:               genesisAddress,
				fromID:             baseIDs.PrototypeIdentityID(),
				immutableMetaProps: baseLists.NewPropertyList(),
				immutableProps:     baseLists.NewPropertyList(),
				mutableMetaProps:   baseLists.NewPropertyList(),
				mutableProps:       superAuxiliaryMutablesFailure,
			},
			setup: func() {

			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			got, err := TransactionKeeper.Transact(sdkTypes.WrapSDKContext(Context),
				NewMessage(tt.args.from,
					tt.args.fromID,
					tt.args.immutableMetaProps.(lists.PropertyList),
					tt.args.immutableProps.(lists.PropertyList),
					tt.args.mutableMetaProps.(lists.PropertyList),
					tt.args.mutableProps.(lists.PropertyList)).(helpers.Message))

			if tt.wantErr != nil {
				require.Error(t, err)
				require.Equal(t, tt.wantErr, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}
