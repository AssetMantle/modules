// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerate

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/assets/constants"
	recordassets "github.com/AssetMantle/modules/x/assets/record"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	parametersSchema "github.com/AssetMantle/schema/parameters"
	"github.com/AssetMantle/schema/parameters/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	"github.com/cometbft/cometbft/crypto/ed25519"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	authKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/assets/mapper"
	"github.com/AssetMantle/modules/x/assets/parameters"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseMetaProp "github.com/AssetMantle/schema/properties/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	tendermintDB "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
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
	moduleStoreKey = sdkTypes.NewKVStoreKey(constants.ModuleName)

	randomInteger         int64 = 100
	randomNegativeInteger int64 = -1

	SupplyNotRevealedProperty = baseMetaProp.NewMesaProperty(constantProperties.SupplyProperty.GetID().GetKey(), baseData.NewNumberData(sdkTypes.NewInt(randomNegativeInteger)))
	NegativeSupplyProperty    = baseMetaProp.NewMetaProperty(constantProperties.SupplyProperty.GetID().GetKey(), baseData.NewNumberData(sdkTypes.NewInt(randomNegativeInteger)))

	propList = baseQualified.NewMutables(mutableMetaProperties).GetMutablePropertyList().Add(
		baseProperties.NewMesaProperty(constantProperties.SupplyProperty.GetKey(), baseData.NewNumberData(sdkTypes.NewInt(randomInteger))),
	)
	newMutables    = baseQualified.NewMutables(propList)
	testNewAsset   = baseDocuments.NewAsset(baseIDs.NewClassificationID(immutables, newMutables), immutables, newMutables)
	testNewAssetID = baseIDs.NewAssetID(testNewAsset.GetClassificationID(), testNewAsset.GetImmutables()).(*baseIDs.AssetID)

	authenticateAuxiliaryKeeper         = new(MockAuxiliaryKeeper)
	authenticateAuxiliaryFailureAddress = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	_                                   = authenticateAuxiliaryKeeper.On("Help", mock.Anything, authenticate.NewAuxiliaryRequest(authenticateAuxiliaryFailureAddress, baseIDs.PrototypeIdentityID())).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	_                                   = authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	authenticateAuxiliary               = new(MockAuxiliary)
	_                                   = authenticateAuxiliary.On("GetKeeper").Return(authenticateAuxiliaryKeeper)

	authorizeAuxiliaryKeeper                  = new(MockAuxiliaryKeeper)
	authorizeAuxiliaryFailureClassificationID = baseIDs.NewIdentityID(testNewAsset.GetClassificationID(), immutables)
	authorizeAuxiliary                        = new(MockAuxiliary)
	_                                         = authorizeAuxiliary.On("GetKeeper").Return(authorizeAuxiliaryKeeper)

	renumerateAuxiliaryKeeper = new(MockAuxiliaryKeeper)
	renumerateAuxiliary       = new(MockAuxiliary)
	_                         = renumerateAuxiliary.On("GetKeeper").Return(renumerateAuxiliaryKeeper)

	supplementAuxiliaryKeeper = new(MockAuxiliaryKeeper)
	supplementAuxiliary       = new(MockAuxiliary)
	_                         = supplementAuxiliary.On("GetKeeper").Return(supplementAuxiliaryKeeper)

	codec = baseHelpers.TestCodec()

	paramsStoreKey           = sdkTypes.NewKVStoreKey(paramsTypes.StoreKey)
	paramsTransientStoreKeys = sdkTypes.NewTransientStoreKey(paramsTypes.TStoreKey)
	ParamsKeeper             = paramsKeeper.NewKeeper(codec, codec.GetLegacyAmino(), paramsStoreKey, paramsTransientStoreKeys)

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

	parameterManager = parameters.Prototype().Initialize(ParamsKeeper.Subspace(constants.ModuleName).WithKeyTable(parameters.Prototype().GetKeyTable())).
				Set(sdkTypes.WrapSDKContext(Context), []parametersSchema.Parameter{base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(Denom))))}).
				Set(sdkTypes.WrapSDKContext(Context), []parametersSchema.Parameter{base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(true)))}).
				Set(sdkTypes.WrapSDKContext(Context), []parametersSchema.Parameter{base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(true)))}).
				Set(sdkTypes.WrapSDKContext(Context), []parametersSchema.Parameter{base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(true)))}).
				Set(sdkTypes.WrapSDKContext(Context), []parametersSchema.Parameter{base.NewParameter(baseProperties.NewMetaProperty(constantProperties.UnwrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(Denom))))})
	TransactionKeeper = transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey),
		parameterManager,
		authenticateAuxiliary,
		authorizeAuxiliary,
		renumerateAuxiliary,
		supplementAuxiliary,
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
		from    sdkTypes.AccAddress
		fromID  ids.IdentityID
		assetID ids.AssetID
		denom   string
	}

	tests := []struct {
		name    string
		args    args
		setup   func()
		want    *TransactionResponse
		wantErr helpers.Error
	}{
		{
			name: "RenumerateTransactionKeeperSuccess",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: testNewAssetID,
				denom:   Denom,
			},
			setup: func() {
				renumerateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(supplement.NewAuxiliaryResponse(baseLists.NewPropertyList(constantProperties.SupplyProperty)), nil).Once()
			},
			want:    newTransactionResponse(),
			wantErr: nil,
		},
		{
			name: "RenumeratePropertyDisabled",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: testNewAssetID,
				denom:   Denom,
			},
			setup: func() {
				parameterManager.Set(sdkTypes.WrapSDKContext(Context), []parametersSchema.Parameter{base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(false)))})
			},
			want:    nil,
			wantErr: errorConstants.NotAuthorized,
		},
		{
			name: "AuthenticationFailure",
			args: args{
				from:    authenticateAuxiliaryFailureAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: baseIDs.PrototypeAssetID(),
				denom:   Denom,
			},
			setup: func() {
				parameterManager.Set(sdkTypes.WrapSDKContext(Context), []parametersSchema.Parameter{base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(true)))})
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "EntityNotFoundFailure",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: baseIDs.PrototypeAssetID(),
				denom:   Denom,
			},
			setup: func() {
			},
			want:    nil,
			wantErr: errorConstants.EntityNotFound,
		},
		{
			name: "AuthorizeAuxiliaryFailure",
			args: args{
				from:    genesisAddress,
				fromID:  authorizeAuxiliaryFailureClassificationID,
				assetID: testNewAssetID,
				denom:   Denom,
			},
			setup: func() {
				authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "SupplementAuxiliaryFailure",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: testNewAssetID,
			},
			setup: func() {
				authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "SupplyNotRevealedError",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: testNewAssetID,
			},
			setup: func() {
				renumerateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
				authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(supplement.NewAuxiliaryResponse(baseLists.NewPropertyList(SupplyNotRevealedProperty)), nil).Once()
			},
			want:    nil,
			wantErr: errorConstants.MetaDataError,
		},
		{
			name: "NegativeSupplyFailure",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: testNewAssetID,
			},
			setup: func() {
				renumerateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
				authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(supplement.NewAuxiliaryResponse(baseLists.NewPropertyList(NegativeSupplyProperty)), nil).Once()
			},
			want:    nil,
			wantErr: errorConstants.MetaDataError,
		},
		{
			name: "RenumerateAuxiliaryFailure",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: testNewAssetID,
			},
			setup: func() {
				renumerateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
				authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(supplement.NewAuxiliaryResponse(baseLists.NewPropertyList(constantProperties.SupplyProperty)), nil).Once()
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(recordassets.NewRecord(testNewAsset))

			got, err := TransactionKeeper.Transact(sdkTypes.WrapSDKContext(Context),
				NewMessage(tt.args.from,
					tt.args.fromID,
					tt.args.assetID).(helpers.Message))

			if (err != nil) && !tt.wantErr.Is(err) {
				t.Errorf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Error("unexpected response")
			}
		})
	}
}
