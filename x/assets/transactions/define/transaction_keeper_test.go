// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/super"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/lists"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/cometbft/cometbft/crypto/ed25519"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
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
	tendermintDB "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
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

func (mockAuxiliaryKeeper *MockAuxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	args := mockAuxiliaryKeeper.Called(context, request)
	return args.Get(0).(helpers.AuxiliaryResponse), args.Error(1)
}
func (mockAuxiliaryKeeper *MockAuxiliaryKeeper) Initialize(m2 helpers.Mapper, manager helpers.ParameterManager, i []interface{}) helpers.Keeper {
	args := mockAuxiliaryKeeper.Called(m2, manager, i)
	return args.Get(0).(helpers.Keeper)
}

const (
	TestMinterModuleName = "testMinter"
	Denom                = "stake"
	ChainID              = "testChain"
	GenesisSupply        = 1000000000000
)

var (
	moduleStoreKey = sdkTypes.NewKVStoreKey(constants.ModuleName)

	authenticateAuxiliaryKeeper         = new(MockAuxiliaryKeeper)
	authenticateAuxiliaryFailureAddress = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	_                                   = authenticateAuxiliaryKeeper.On("Help", mock.Anything, authenticate.NewAuxiliaryRequest(authenticateAuxiliaryFailureAddress, baseIDs.PrototypeIdentityID())).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
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

	codec = baseHelpers.TestCodec()

	paramsStoreKey           = sdkTypes.NewKVStoreKey(paramsTypes.StoreKey)
	paramsTransientStoreKeys = sdkTypes.NewTransientStoreKey(paramsTypes.TStoreKey)

	authStoreKey             = sdkTypes.NewKVStoreKey(authTypes.StoreKey)
	moduleAccountPermissions = map[string][]string{TestMinterModuleName: {authTypes.Minter}, constants.ModuleName: nil}
	AuthKeeper               = authKeeper.NewAccountKeeper(codec, authStoreKey, authTypes.ProtoBaseAccount, moduleAccountPermissions, sdkTypes.GetConfig().GetBech32AccountAddrPrefix(), authTypes.NewModuleAddress(govTypes.ModuleName).String())

	bankStoreKey         = sdkTypes.NewKVStoreKey(bankTypes.StoreKey)
	blacklistedAddresses = map[string]bool{authTypes.NewModuleAddress(TestMinterModuleName).String(): false, authTypes.NewModuleAddress(constants.ModuleName).String(): false}
	BankKeeper           = bankKeeper.NewBaseKeeper(codec, bankStoreKey, AuthKeeper, blacklistedAddresses, authTypes.NewModuleAddress(govTypes.ModuleName).String())

	Context = setContext()

	coinSupply = sdkTypes.NewCoins(sdkTypes.NewCoin(Denom, sdkTypes.NewInt(GenesisSupply)))
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
	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
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
