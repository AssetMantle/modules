// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mutate

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/mapper"
	recordassets "github.com/AssetMantle/modules/x/assets/record"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/maintain"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	tendermintDB "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/cometbft/cometbft/libs/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/store"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	authKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
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

	maintainFailureClassificationID = testAsset.GetClassificationID()
	maintainAuxiliaryKeeper         = new(MockAuxiliaryKeeper)
	maintainFailureID               = baseIDs.NewIdentityID(classificationID, immutables)
	maintainAuxiliary               = new(MockAuxiliary)
	_                               = maintainAuxiliary.On("GetKeeper").Return(maintainAuxiliaryKeeper)

	conformAuxiliaryKeeper = new(MockAuxiliaryKeeper)
	conformAuxiliary       = new(MockAuxiliary)
	_                      = conformAuxiliary.On("GetKeeper").Return(conformAuxiliaryKeeper)

	testAsset      = baseDocuments.NewAsset(baseIDs.NewClassificationID(immutables, mutables), baseQualified.NewImmutables(baseLists.NewPropertyList()), baseQualified.NewMutables(baseLists.NewPropertyList()))
	testNewAssetID = baseIDs.NewAssetID(testAsset.GetClassificationID(), testAsset.GetImmutables()).(*baseIDs.AssetID)

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
		authenticateAuxiliary,
		maintainAuxiliary,
		conformAuxiliary,
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
		from             sdkTypes.AccAddress
		fromID           ids.IdentityID
		assetID          ids.AssetID
		mutableMetaProps lists.PropertyList
		mutableProps     lists.PropertyList
	}

	tests := []struct {
		name    string
		args    args
		setup   func()
		want    *TransactionResponse
		wantErr helpers.Error
	}{
		{
			"MutateTransactionKeeeperSuccess",
			args{
				from:             genesisAddress,
				fromID:           baseIDs.PrototypeIdentityID(),
				assetID:          testNewAssetID,
				mutableMetaProps: baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(),
			},
			func() {
				conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				maintainAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			newTransactionResponse(),
			nil,
		},
		{
			"MutateValidAsset",
			args{
				from:             genesisAddress,
				fromID:           baseIDs.PrototypeIdentityID(),
				assetID:          testNewAssetID,
				mutableMetaProps: baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(),
			},
			func() {
				conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				maintainAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			newTransactionResponse(),
			nil,
		},
		{
			"EntityNotFoundError",
			args{
				from:             genesisAddress,
				fromID:           baseIDs.PrototypeIdentityID(),
				assetID:          baseIDs.PrototypeAssetID(),
				mutableMetaProps: baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(),
			},
			func() {},
			nil,
			errorConstants.EntityNotFound,
		},
		{
			"AuthenticateAuxiliaryFailure",
			args{
				from:             authenticateAuxiliaryFailureAddress,
				fromID:           baseIDs.PrototypeIdentityID(),
				assetID:          baseIDs.PrototypeAssetID(),
				mutableMetaProps: baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(),
			},
			func() {},
			nil,
			errorConstants.MockError,
		},
		{
			"MaintainAuxiliaryFailure",
			args{
				from:             genesisAddress,
				fromID:           maintainFailureID,
				assetID:          testNewAssetID,
				mutableMetaProps: baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(),
			},
			func() {
				maintainAuxiliaryKeeper.On("Help", mock.Anything, maintain.NewAuxiliaryRequest(maintainFailureClassificationID, maintainFailureID, baseQualified.NewMutables(baseLists.NewPropertyList()))).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
			},
			nil,
			errorConstants.MockError,
		},
		{
			"ConformAuxiliaryFailure",
			args{
				from:             genesisAddress,
				fromID:           baseIDs.PrototypeIdentityID(),
				assetID:          testNewAssetID,
				mutableMetaProps: baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(),
			},
			func() {
				maintainAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				conformAuxiliaryKeeper.On("Help", mock.Anything, conform.NewAuxiliaryRequest(testAsset.GetClassificationID(), testAsset.GetImmutables(), baseQualified.NewMutables(baseLists.NewPropertyList()))).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
			},
			nil,
			errorConstants.MockError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(recordassets.NewRecord(testAsset))

			got, err := TransactionKeeper.Transact(sdkTypes.WrapSDKContext(Context),
				NewMessage(tt.args.from,
					tt.args.fromID,
					tt.args.assetID,
					tt.args.mutableMetaProps.(lists.PropertyList),
					tt.args.mutableProps.(lists.PropertyList)).(helpers.Message),
			)

			if (tt.wantErr != nil && !tt.wantErr.Is(err)) || (tt.wantErr == nil && err != nil) {
				t.Errorf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Error("unexpected response")
			}
		})
	}
}
