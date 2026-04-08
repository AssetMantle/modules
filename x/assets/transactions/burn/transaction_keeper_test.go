// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn

import (
	"cosmossdk.io/math"
	storeTypes "cosmossdk.io/store/types"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	baseProp "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/mapper"
	"github.com/AssetMantle/modules/x/assets/parameters"
	recordassets "github.com/AssetMantle/modules/x/assets/record"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/documents"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	"github.com/AssetMantle/schema/parameters/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	baseType "github.com/AssetMantle/schema/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/mock"
	"math/rand"
	"reflect"
	"testing"
)

type testSetup struct {
	Context                     sdkTypes.Context
	genesisAddress              sdkTypes.AccAddress
	parameterManager            helpers.ParameterManager
	TransactionKeeper           transactionKeeper
	authenticateAuxiliaryKeeper *testutil.MockAuxiliaryKeeper
	authorizeAuxiliaryKeeper    *testutil.MockAuxiliaryKeeper
	purgeAuxiliaryKeeper        *testutil.MockAuxiliaryKeeper
	supplementAuxiliaryKeeper   *testutil.MockAuxiliaryKeeper
	unbondAuxiliaryKeeper       *testutil.MockAuxiliaryKeeper
	asset                       documents.Asset
	assetID                     *baseIDs.AssetID
	burnHeightMesaAsset         documents.Asset
	burnHeightMesaAssetID       *baseIDs.AssetID
	supplyMesaAsset             documents.Asset
	supplyMesaAssetID           *baseIDs.AssetID
	burnHeightMetaMutables      lists.PropertyList
	supplyMetaMutables          lists.PropertyList
	mutableMetaMock             lists.PropertyList
}

func setupTest(t *testing.T) *testSetup {
	t.Helper()

	moduleStoreKey := storeTypes.NewKVStoreKey(constants.ModuleName)
	ctx, _, genesisAddress := testutil.NewTestContextWithBankAuth(t, constants.ModuleName, moduleStoreKey)

	randomInteger := int64(100)
	randomSeed := int64(99)

	immutablesMesaMock := baseQualified.NewImmutables(baseProp.GenerateRandomPropertyList(rand.New(rand.NewSource(randomSeed))))

	burnHeightMesaPropList := baseProp.GenerateRandomPropertyList(rand.New(rand.NewSource(randomSeed)))
	burnHeightMesaMutables := baseQualified.NewMutables(burnHeightMesaPropList).GetMutablePropertyList().Add(
		baseProperties.NewMesaProperty(constantProperties.BurnHeightProperty.GetKey(), baseData.NewHeightData(baseType.NewHeight(randomInteger))),
		baseProperties.NewMesaProperty(constantProperties.BondAmountProperty.GetKey(), baseData.NewNumberData(math.NewInt(randomInteger))),
	)
	burnHeightAssetMesaMutables := baseQualified.NewMutables(burnHeightMesaMutables)
	burnHeightMesaAsset := baseDocuments.NewAsset(baseIDs.NewClassificationID(immutablesMesaMock, burnHeightAssetMesaMutables), immutablesMesaMock, burnHeightAssetMesaMutables)
	burnHeightMesaAssetID := baseIDs.NewAssetID(burnHeightMesaAsset.GetClassificationID(), burnHeightMesaAsset.GetImmutables()).(*baseIDs.AssetID)

	supplyMesaPropList := baseProp.GenerateRandomPropertyList(rand.New(rand.NewSource(randomSeed)))
	supplyMesaMutables := baseQualified.NewMutables(supplyMesaPropList).GetMutablePropertyList().Add(
		baseProperties.NewMesaProperty(constantProperties.SupplyProperty.GetKey(), baseData.NewNumberData(math.NewInt(randomInteger))),
	)
	supplyAssetMesaMutables := baseQualified.NewMutables(supplyMesaMutables)
	supplyMesaAsset := baseDocuments.NewAsset(baseIDs.NewClassificationID(immutablesMesaMock, supplyAssetMesaMutables), immutablesMesaMock, supplyAssetMesaMutables)
	supplyMesaAssetID := baseIDs.NewAssetID(supplyMesaAsset.GetClassificationID(), supplyMesaAsset.GetImmutables()).(*baseIDs.AssetID)

	burnHeightMetaPropList := baseProp.GenerateRandomPropertyList(rand.New(rand.NewSource(randomSeed)))
	burnHeightMetaMutables := baseQualified.NewMutables(burnHeightMetaPropList).GetMutablePropertyList().Add(
		baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewNumberData(math.NewInt(randomInteger))),
		baseProperties.NewMetaProperty(constantProperties.BondAmountProperty.GetKey(), baseData.NewNumberData(math.NewInt(randomInteger))),
		baseProperties.NewMetaProperty(constantProperties.BurnHeightProperty.GetKey(), baseData.NewHeightData(baseType.NewHeight(randomInteger))),
	)

	supplyMetaPropList := baseProp.GenerateRandomPropertyList(rand.New(rand.NewSource(randomSeed)))
	supplyMetaMutables := baseQualified.NewMutables(supplyMetaPropList).GetMutablePropertyList().Add(
		baseProperties.NewMetaProperty(constantProperties.SupplyProperty.GetKey(), baseData.NewNumberData(math.NewInt(randomInteger))),
	)

	mutableMetaMock := baseProp.GenerateRandomPropertyList(rand.New(rand.NewSource(randomSeed)))
	BurnEnableMetaMutables := baseQualified.NewMutables(mutableMetaMock).GetMutablePropertyList().Add(
		baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewNumberData(math.NewInt(randomInteger))),
		baseProperties.NewMetaProperty(constantProperties.BondAmountProperty.GetKey(), baseData.NewNumberData(math.NewInt(randomInteger))),
	)
	BurnEnabledAssetMetaMutable := baseQualified.NewMutables(BurnEnableMetaMutables)
	asset := baseDocuments.NewAsset(baseIDs.NewClassificationID(immutables, BurnEnabledAssetMetaMutable), immutables, BurnEnabledAssetMetaMutable)
	assetID := baseIDs.NewAssetID(asset.GetClassificationID(), asset.GetImmutables()).(*baseIDs.AssetID)

	authenticateAuxiliary, authenticateAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()
	authorizeAuxiliary, authorizeAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()
	purgeAuxiliary, purgeAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()
	supplementAuxiliary, supplementAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()
	unbondAuxiliary, unbondAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()

	parameterManager, _ := parameters.Prototype().Initialize(moduleStoreKey).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(testutil.Denom))))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.UnwrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(testutil.Denom))))).
		Update(ctx)

	TransactionKeeper := transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey), parameterManager, authenticateAuxiliary, authorizeAuxiliary, purgeAuxiliary, supplementAuxiliary, unbondAuxiliary}

	return &testSetup{
		Context:                     ctx,
		genesisAddress:              genesisAddress,
		parameterManager:            parameterManager,
		TransactionKeeper:           TransactionKeeper,
		authenticateAuxiliaryKeeper: authenticateAuxiliaryKeeper,
		authorizeAuxiliaryKeeper:    authorizeAuxiliaryKeeper,
		purgeAuxiliaryKeeper:        purgeAuxiliaryKeeper,
		supplementAuxiliaryKeeper:   supplementAuxiliaryKeeper,
		unbondAuxiliaryKeeper:       unbondAuxiliaryKeeper,
		asset:                       asset,
		assetID:                     assetID,
		burnHeightMesaAsset:         burnHeightMesaAsset,
		burnHeightMesaAssetID:       burnHeightMesaAssetID,
		supplyMesaAsset:             supplyMesaAsset,
		supplyMesaAssetID:           supplyMesaAssetID,
		burnHeightMetaMutables:      burnHeightMetaMutables,
		supplyMetaMutables:          supplyMetaMutables,
		mutableMetaMock:             mutableMetaMock,
	}
}

func TestTransactionKeeperTransact(t *testing.T) {
	s := setupTest(t)

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
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: s.assetID,
			},
			setup: func() {
				s.parameterManager.Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(false)))).Update(sdkTypes.WrapSDKContext(s.Context))
			},
			want:    nil,
			wantErr: errorConstants.NotAuthorized,
		},
		{
			name: "BurnTransactionKeeperSuccess",
			args: args{
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: s.assetID,
			},
			setup: func() {
				s.parameterManager.Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).Update(sdkTypes.WrapSDKContext(s.Context))
				s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Add(recordassets.NewRecord(s.asset))
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.purgeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.unbondAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			want:    newTransactionResponse(),
			wantErr: nil,
		},
		{
			name: "AuthenticationFailure",
			args: args{
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: s.assetID,
			},
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "AuthorizationFailure",
			args: args{
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: s.assetID,
			},
			setup: func() {
				s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Add(recordassets.NewRecord(s.asset))
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.NotAuthorized).Once()
			},
			want:    nil,
			wantErr: errorConstants.NotAuthorized,
		},
		{
			name: "BurnAssetEntityNotFound",
			args: args{
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: baseIDs.PrototypeAssetID(),
			},
			setup: func() {
				s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Add(recordassets.NewRecord(s.asset))
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			want:    nil,
			wantErr: errorConstants.EntityNotFound,
		},
		{
			name: "BurnHeightSupplementAuxiliaryFailure",
			args: args{
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: s.burnHeightMesaAssetID,
			},
			setup: func() {
				s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Add(recordassets.NewRecord(s.burnHeightMesaAsset))
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "BurnHeightMetaDataError",
			args: args{
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: s.burnHeightMesaAssetID,
			},
			setup: func() {
				s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Add(recordassets.NewRecord(s.burnHeightMesaAsset))
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(supplement.NewAuxiliaryResponse(s.burnHeightMetaMutables), nil).Once()
			},
			want:    nil,
			wantErr: errorConstants.NotAuthorized,
		},
		{
			name: "BurnHeightPropertyNotRevealed",
			args: args{
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: s.burnHeightMesaAssetID,
			},
			setup: func() {
				s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Add(recordassets.NewRecord(s.burnHeightMesaAsset))
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(supplement.NewAuxiliaryResponse(s.mutableMetaMock), nil).Once()
			},
			want:    nil,
			wantErr: errorConstants.MetaDataError,
		},
		{
			name: "SupplyPropertySupplementAuxiliaryFailure",
			args: args{
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: s.supplyMesaAssetID,
			},
			setup: func() {
				s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Add(recordassets.NewRecord(s.supplyMesaAsset))
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "SupplyMetaPropertyAuxiliaryResponseSuccess",
			args: args{
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: s.supplyMesaAssetID,
			},
			setup: func() {
				s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Add(recordassets.NewRecord(s.supplyMesaAsset))
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(supplement.NewAuxiliaryResponse(s.supplyMetaMutables), nil).Once()
				s.purgeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			want:    nil,
			wantErr: errorConstants.MetaDataError,
		},
		{
			name: "AssetsWithoutSupplyCannotBeBurnedMetaError",
			args: args{
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: s.supplyMesaAssetID,
			},
			setup: func() {
				s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Add(recordassets.NewRecord(s.supplyMesaAsset))
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(supplement.NewAuxiliaryResponse(s.mutableMetaMock), nil).Once()
			},
			want:    nil,
			wantErr: errorConstants.MetaDataError,
		},
		{
			name: "PurgeAuxiliaryFailure",
			args: args{
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: s.assetID,
			},
			setup: func() {
				s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Add(recordassets.NewRecord(s.asset))
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.purgeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "unbondAuxiliaryFailure",
			args: args{
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: s.assetID,
			},
			setup: func() {
				s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Add(recordassets.NewRecord(s.asset))
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.purgeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.unbondAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s.authenticateAuxiliaryKeeper.ExpectedCalls = nil
			s.authenticateAuxiliaryKeeper.Calls = nil
			s.authorizeAuxiliaryKeeper.ExpectedCalls = nil
			s.authorizeAuxiliaryKeeper.Calls = nil
			s.purgeAuxiliaryKeeper.ExpectedCalls = nil
			s.purgeAuxiliaryKeeper.Calls = nil
			s.supplementAuxiliaryKeeper.ExpectedCalls = nil
			s.supplementAuxiliaryKeeper.Calls = nil
			s.unbondAuxiliaryKeeper.ExpectedCalls = nil
			s.unbondAuxiliaryKeeper.Calls = nil

			tt.setup()

			got, err := s.TransactionKeeper.Transact(sdkTypes.WrapSDKContext(s.Context),
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
