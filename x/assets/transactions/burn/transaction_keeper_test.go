// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/assets/constants"
	recordassets "github.com/AssetMantle/modules/x/assets/record"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/parameters/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	base2 "github.com/AssetMantle/schema/types/base"
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

	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	tendermintDB "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/assets/mapper"
	"github.com/AssetMantle/modules/x/assets/parameters"
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

	//Mesa Property List variables
	immutablesMesaMock = baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("mockMesaProp"), baseData.NewListData())).(*baseLists.PropertyList))

	//Burn Height Asset - Mesa
	burnHeightMesaPropList = baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("mockMesaProp"), baseData.NewListData())).(*baseLists.PropertyList)
	burnHeightMesaMutables = baseQualified.NewMutables(burnHeightMesaPropList).GetMutablePropertyList().Add(
		baseProperties.NewMesaProperty(constantProperties.BurnHeightProperty.GetKey(), baseData.NewHeightData(base2.NewHeight(13))),
		baseProperties.NewMesaProperty(constantProperties.BondAmountProperty.GetKey(), baseData.NewNumberData(sdkTypes.NewInt(34))),
	)
	burnHeightAssetMesaMutables = baseQualified.NewMutables(burnHeightMesaMutables)
	burnHeightMesaAsset         = baseDocuments.NewAsset(baseIDs.NewClassificationID(immutablesMesaMock, burnHeightAssetMesaMutables), immutablesMesaMock, burnHeightAssetMesaMutables)
	burnHeightMesaAssetID       = baseIDs.NewAssetID(burnHeightMesaAsset.GetClassificationID(), burnHeightMesaAsset.GetImmutables()).(*baseIDs.AssetID)

	//Supply Asset - Mesa
	supplyMesaPropList = baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("mockMesaProp"), baseData.NewListData())).(*baseLists.PropertyList)
	supplyMesaMutables = baseQualified.NewMutables(supplyMesaPropList).GetMutablePropertyList().Add(
		baseProperties.NewMesaProperty(constantProperties.SupplyProperty.GetKey(), baseData.NewNumberData(sdkTypes.NewInt(34))),
	)
	supplyAssetMesaMutables = baseQualified.NewMutables(supplyMesaMutables)
	supplyMesaAsset         = baseDocuments.NewAsset(baseIDs.NewClassificationID(immutablesMesaMock, supplyAssetMesaMutables), immutablesMesaMock, supplyAssetMesaMutables)
	supplyMesaAssetID       = baseIDs.NewAssetID(supplyMesaAsset.GetClassificationID(), supplyMesaAsset.GetImmutables()).(*baseIDs.AssetID)

	//Meta Property List variables
	burnHeightMetaPropList = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData())).(*baseLists.PropertyList)
	burnHeightMetaMutables = baseQualified.NewMutables(burnHeightMetaPropList).GetMutablePropertyList().Add(
		baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewNumberData(sdkTypes.NewInt(100))),
		baseProperties.NewMetaProperty(constantProperties.BondAmountProperty.GetKey(), baseData.NewNumberData(sdkTypes.NewInt(100))),
		baseProperties.NewMetaProperty(constantProperties.BurnHeightProperty.GetKey(), baseData.NewHeightData(base2.NewHeight(20))),
	)

	supplyMetaPropList = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData())).(*baseLists.PropertyList)
	supplyMetaMutables = baseQualified.NewMutables(supplyMetaPropList).GetMutablePropertyList().Add(
		baseProperties.NewMetaProperty(constantProperties.SupplyProperty.GetKey(), baseData.NewNumberData(sdkTypes.NewInt(100))),
	)

	//Burn Enabled Asset - Meta
	mutableMetaMock        = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData())).(*baseLists.PropertyList)
	BurnEnableMetaMutables = baseQualified.NewMutables(mutableMetaMock).GetMutablePropertyList().Add(
		baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewNumberData(sdkTypes.NewInt(100))),
		baseProperties.NewMetaProperty(constantProperties.BondAmountProperty.GetKey(), baseData.NewNumberData(sdkTypes.NewInt(100))),
	)
	BurnEnabledAssetMetaMutable = baseQualified.NewMutables(BurnEnableMetaMutables)
	asset                       = baseDocuments.NewAsset(baseIDs.NewClassificationID(immutables, BurnEnabledAssetMetaMutable), immutables, BurnEnabledAssetMetaMutable)
	assetID                     = baseIDs.NewAssetID(asset.GetClassificationID(), asset.GetImmutables()).(*baseIDs.AssetID)

	authenticateAuxiliaryKeeper = new(MockAuxiliaryKeeper)
	authenticateAuxiliary       = new(MockAuxiliary)
	_                           = authenticateAuxiliary.On("GetKeeper").Return(authenticateAuxiliaryKeeper)

	authorizeAuxiliaryKeeper = new(MockAuxiliaryKeeper)
	authorizeAuxiliary       = new(MockAuxiliary)
	_                        = authorizeAuxiliary.On("GetKeeper").Return(authorizeAuxiliaryKeeper)

	purgeAuxiliaryKeeper = new(MockAuxiliaryKeeper)
	purgeAuxiliary       = new(MockAuxiliary)
	_                    = purgeAuxiliary.On("GetKeeper").Return(purgeAuxiliaryKeeper)

	supplementAuxiliaryKeeper = new(MockAuxiliaryKeeper)
	supplementAuxiliary       = new(MockAuxiliary)
	_                         = supplementAuxiliary.On("GetKeeper").Return(supplementAuxiliaryKeeper)

	unbondAuxiliaryKeeper = new(MockAuxiliaryKeeper)
	unbondAuxiliary       = new(MockAuxiliary)
	_                     = unbondAuxiliary.On("GetKeeper").Return(unbondAuxiliaryKeeper)

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

	coinSupply = sdkTypes.NewCoins(sdkTypes.NewCoin(Denom, sdkTypes.NewInt(GenesisSupply)))
	_          = BankKeeper.MintCoins(Context, TestMinterModuleName, coinSupply)

	genesisAddress = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	_              = BankKeeper.SendCoinsFromModuleToAccount(Context, TestMinterModuleName, genesisAddress, coinSupply)

	Context = setContext()

	parameterManager = parameters.Prototype().Initialize(ParamsKeeper.Subspace(constants.ModuleName).WithKeyTable(parameters.Prototype().GetKeyTable())).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(Denom)))))).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(true))))).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(true))))).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(true))))).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.UnwrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(Denom))))))

	TransactionKeeper = transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey), parameterManager, authenticateAuxiliary, authorizeAuxiliary, purgeAuxiliary, supplementAuxiliary, unbondAuxiliary}
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
		value   int
	}
	tests := []struct {
		name    string
		args    args
		setup   func()
		want    *TransactionResponse
		wantErr helpers.Error
	}{
		{
			name: "BurnPropertyDisabled",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: assetID,
			},
			setup: func() {
				parameterManager.Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(false)))))
				authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			want:    nil,
			wantErr: errorConstants.NotAuthorized,
		},
		{
			name: "BurnTransactionKeeperSuccess",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: assetID,
			},
			setup: func() {
				parameterManager.Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(true)))))
				TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(recordassets.NewRecord(asset))
				authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				purgeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				unbondAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			want:    newTransactionResponse(),
			wantErr: nil,
		},
		{
			name: "AuthenticationFailure",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: assetID,
			},
			setup: func() {
				authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "AuthorizationFailure",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: assetID,
			},
			setup: func() {
				TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(recordassets.NewRecord(asset))
				authenticateAuxiliaryKeeper.On("Help", mock.Anything, authenticate.NewAuxiliaryRequest(genesisAddress, baseIDs.PrototypeIdentityID())).Return(new(helpers.AuxiliaryResponse), nil).Once()
				authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.NotAuthorized).Once()
			},
			want:    nil,
			wantErr: errorConstants.NotAuthorized,
		},
		{
			name: "BurnAssetEntityNotFound",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: baseIDs.PrototypeAssetID(),
			},
			setup: func() {
				TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(recordassets.NewRecord(asset))
				authenticateAuxiliaryKeeper.On("Help", mock.Anything, authenticate.NewAuxiliaryRequest(genesisAddress, baseIDs.PrototypeIdentityID())).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			want:    nil,
			wantErr: errorConstants.EntityNotFound,
		},
		{
			name: "BurnHeightSupplementAuxiliaryFailure",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: burnHeightMesaAssetID,
			},
			setup: func() {
				TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(recordassets.NewRecord(burnHeightMesaAsset))
				authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()

			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "BurnHeightMetaDataError",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: burnHeightMesaAssetID,
			},
			setup: func() {
				TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(recordassets.NewRecord(burnHeightMesaAsset))
				authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(supplement.NewAuxiliaryResponse(burnHeightMetaMutables), nil).Once()

			},
			want:    nil,
			wantErr: errorConstants.NotAuthorized,
		},
		{
			name: "BurnHeightPropertyNotRevealed",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: burnHeightMesaAssetID,
			},
			setup: func() {
				TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(recordassets.NewRecord(burnHeightMesaAsset))
				authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(supplement.NewAuxiliaryResponse(mutableMetaMock), nil).Once()

			},
			want:    nil,
			wantErr: errorConstants.MetaDataError,
		},
		{
			name: "SupplyPropertySupplementAuxiliaryFailure",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: supplyMesaAssetID,
			},
			setup: func() {
				TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(recordassets.NewRecord(supplyMesaAsset))
				authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()

			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "SupplyMetaPropertyAuxiliaryResponseSuccess",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: supplyMesaAssetID,
			},
			setup: func() {
				TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(recordassets.NewRecord(supplyMesaAsset))
				authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(supplement.NewAuxiliaryResponse(supplyMetaMutables), nil).Once()
				purgeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()

			},
			want:    nil,
			wantErr: errorConstants.MetaDataError,
		},
		{
			name: "AssetsWithoutSupplyCannotBeBurnedMetaError",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: supplyMesaAssetID,
			},
			setup: func() {
				TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(recordassets.NewRecord(supplyMesaAsset))
				authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(supplement.NewAuxiliaryResponse(mutableMetaMock), nil).Once()

			},
			want:    nil,
			wantErr: errorConstants.MetaDataError,
		},
		{
			name: "PurgeAuxiliaryFailure",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: assetID,
			},
			setup: func() {
				TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(recordassets.NewRecord(asset))
				authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				purgeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		//{
		//	name: "noRevealedBondAmountFailure",
		//	args: args{
		//		from:    genesisAddress,
		//		fromID:  baseIDs.PrototypeIdentityID(),
		//		assetID: burnHeightMesaAssetID,
		//	},
		//	setup: func() {
		//		TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(recordassets.NewRecord(burnHeightMesaAsset))
		//		authenticateAuxiliaryKeeper.On("Help", mock.Anything, authenticate.NewAuxiliaryRequest(genesisAddress, baseIDs.PrototypeIdentityID())).Return(new(helpers.AuxiliaryResponse), nil).Once()
		//		purgeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
		//		supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
		//		authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
		//	},
		//	want:    nil,
		//	wantErr: errorConstants.MetaDataError,
		//},
		{
			name: "unbondAuxiliaryFailure",
			args: args{
				from:    genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: assetID,
			},
			setup: func() {
				TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(recordassets.NewRecord(asset))
				authenticateAuxiliaryKeeper.On("Help", mock.Anything, authenticate.NewAuxiliaryRequest(genesisAddress, baseIDs.PrototypeIdentityID())).Return(new(helpers.AuxiliaryResponse), nil).Once()
				purgeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				unbondAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
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

//sdkTypes.NewInt(10)
