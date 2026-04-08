// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint

import (
	"cosmossdk.io/math"
	storeTypes "cosmossdk.io/store/types"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	baseProp "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/key"
	"github.com/AssetMantle/modules/x/assets/mapper"
	"github.com/AssetMantle/modules/x/assets/parameters"
	"github.com/AssetMantle/modules/x/assets/record"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/documents"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/parameters/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/AssetMantle/schema/properties"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/mock"
	"math/rand"
	"testing"
	"github.com/stretchr/testify/assert"
)

type testSetup struct {
	Context                            sdkTypes.Context
	genesisAddress                     sdkTypes.AccAddress
	parameterManager                   helpers.ParameterManager
	TransactionKeeper                  transactionKeeper
	authenticateAuxiliaryFailureAddress sdkTypes.AccAddress
	authenticateAuxiliaryKeeper         *testutil.MockAuxiliaryKeeper
	authorizeAuxiliaryKeeper           *testutil.MockAuxiliaryKeeper
	mintAuxiliaryKeeper                *testutil.MockAuxiliaryKeeper
	bondAuxiliaryKeeper                *testutil.MockAuxiliaryKeeper
	conformAuxiliaryKeeper             *testutil.MockAuxiliaryKeeper
	bondAuxiliaryFailureAddress        sdkTypes.AccAddress
	asset                              documents.Asset
	assetID                            *baseIDs.AssetID
	randomMetaProperty                 properties.Property
	randomBondMetaProperty             properties.MetaProperty
}

func setupTest(t *testing.T) *testSetup {
	t.Helper()

	moduleStoreKey := storeTypes.NewKVStoreKey(constants.ModuleName)
	ctx, _, genesisAddress := testutil.NewTestContextWithBankAuth(t, constants.ModuleName, moduleStoreKey)

	randomSeed := int64(99)
	randomInt := int64(1)
	randomMetaProperty := baseProp.GenerateRandomMetaProperty(rand.New(rand.NewSource(randomSeed)))
	randomBondMetaProperty := baseProperties.NewMetaProperty(baseIDs.NewStringID(""), baseData.NewNumberData(math.NewInt(randomInt)))

	newMutables := baseQualified.NewMutables(baseLists.NewPropertyList())
	asset := baseDocuments.NewAsset(baseIDs.NewClassificationID(immutables, newMutables), immutables, newMutables)
	assetID := baseIDs.NewAssetID(asset.GetClassificationID(), asset.GetImmutables()).(*baseIDs.AssetID)

	authenticateAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	authenticateAuxiliaryFailureAddress := testutil.TestAddress()
	authenticateAuxiliaryKeeper.On("Help", mock.Anything, authenticate.NewAuxiliaryRequest(&Message{From: authenticateAuxiliaryFailureAddress.String(), FromID: baseIDs.PrototypeIdentityID().(*baseIDs.IdentityID)})).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	authenticateAuxiliary := new(testutil.MockAuxiliary)
	authenticateAuxiliary.On("GetKeeper").Return(authenticateAuxiliaryKeeper)

	authorizeAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	authorizeAuxiliary := new(testutil.MockAuxiliary)
	authorizeAuxiliary.On("GetKeeper").Return(authorizeAuxiliaryKeeper)

	mintAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	mintAuxiliary := new(testutil.MockAuxiliary)
	mintAuxiliary.On("GetKeeper").Return(mintAuxiliaryKeeper)

	bondAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	bondAuxiliaryFailureAddress := testutil.TestAddress()
	bondAuxiliary := new(testutil.MockAuxiliary)
	bondAuxiliary.On("GetKeeper").Return(bondAuxiliaryKeeper)

	conformAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	conformAuxiliary := new(testutil.MockAuxiliary)
	conformAuxiliary.On("GetKeeper").Return(conformAuxiliaryKeeper)

	parameterManager, _ := parameters.Prototype().Initialize(moduleStoreKey).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(testutil.Denom))))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.UnwrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(testutil.Denom))))).
		Update(ctx)

	TransactionKeeper := transactionKeeper{
		mapper:                mapper.Prototype().Initialize(moduleStoreKey),
		parameterManager:      parameterManager,
		authenticateAuxiliary: authenticateAuxiliary,
		authorizeAuxiliary:    authorizeAuxiliary,
		bondAuxiliary:         bondAuxiliary,
		conformAuxiliary:      conformAuxiliary,
		mintAuxiliary:         mintAuxiliary,
	}

	return &testSetup{
		Context:                            ctx,
		genesisAddress:                     genesisAddress,
		parameterManager:                   parameterManager,
		TransactionKeeper:                  TransactionKeeper,
		authenticateAuxiliaryFailureAddress: authenticateAuxiliaryFailureAddress,
		authenticateAuxiliaryKeeper:         authenticateAuxiliaryKeeper,
		authorizeAuxiliaryKeeper:           authorizeAuxiliaryKeeper,
		mintAuxiliaryKeeper:                mintAuxiliaryKeeper,
		bondAuxiliaryKeeper:                bondAuxiliaryKeeper,
		conformAuxiliaryKeeper:             conformAuxiliaryKeeper,
		bondAuxiliaryFailureAddress:        bondAuxiliaryFailureAddress,
		asset:                              asset,
		assetID:                            assetID,
		randomMetaProperty:                 randomMetaProperty,
		randomBondMetaProperty:             randomBondMetaProperty,
	}
}

func TestTransactionKeeperTransact(t *testing.T) {
	s := setupTest(t)

	type args struct {
		from             sdkTypes.AccAddress
		fromID           ids.IdentityID
		toID             ids.IdentityID
		classificationID ids.ClassificationID
		immutableProps   lists.PropertyList
		mutableProps     lists.PropertyList
	}
	tests := []struct {
		name    string
		args    args
		setup   func(*testing.T)
		want    *TransactionResponse
		wantErr helpers.Error
	}{
		{
			name: "MintTransactionKeeperSuccess",
			args: args{
				from:             s.bondAuxiliaryFailureAddress,
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: s.asset.GetClassificationID(),
				immutableProps:   s.asset.GetImmutables().GetImmutablePropertyList(),
				mutableProps:     baseLists.NewPropertyList(baseProperties.NewMetaProperty(constantProperties.BondAmountProperty.GetID().GetKey(), baseData.NewNumberData(math.NewInt(1)))),
			},
			setup: func(t *testing.T) {
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.mintAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.bondAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			want:    newTransactionResponse(s.assetID),
			wantErr: nil,
		},
		{
			name: "MintPropertyNotEnabled",
			args: args{
				from:             s.genesisAddress,
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: baseIDs.PrototypeClassificationID(),
				immutableProps:   baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(),
			},
			setup: func(t *testing.T) {
				s.parameterManager.Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(false)))).Update(sdkTypes.WrapSDKContext(s.Context))
			},
			wantErr: errorConstants.NotAuthorized,
		},
		{
			name: "AuthorizationFailure",
			args: args{
				from:             s.genesisAddress,
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: baseIDs.PrototypeClassificationID(),
				immutableProps:   baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(),
			},
			setup: func(t *testing.T) {
				s.parameterManager.Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).Update(sdkTypes.WrapSDKContext(s.Context))
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			wantErr: errorConstants.MockError,
		},
		{
			name: "AuthenticationFailure",
			args: args{
				from:             s.authenticateAuxiliaryFailureAddress,
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: baseIDs.PrototypeClassificationID(),
				immutableProps:   baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(),
			},
			setup: func(t *testing.T) {
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			wantErr: errorConstants.MockError,
		},
		{
			name: "EntityAlreadyExistsError",
			args: args{
				from:             s.genesisAddress,
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: baseIDs.PrototypeClassificationID(),
				immutableProps:   baseLists.NewPropertyList(s.randomMetaProperty),
				mutableProps:     baseLists.NewPropertyList(),
			},
			setup: func(t *testing.T) {
				assetID := baseIDs.NewAssetID(baseIDs.PrototypeClassificationID(), baseQualified.NewImmutables(baseLists.NewPropertyList(s.randomMetaProperty)))
				assets := s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Fetch(key.NewKey(assetID))
				assets.Add(record.NewRecord(
					baseDocuments.NewAsset(baseIDs.PrototypeClassificationID(),
						baseQualified.NewImmutables(baseLists.NewPropertyList(s.randomMetaProperty)),
						baseQualified.NewMutables(baseLists.NewPropertyList()))))
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()

			},
			wantErr: errorConstants.EntityAlreadyExists,
		},
		{
			name: "ConformAuxiliaryFailure",
			args: args{
				from:             s.genesisAddress,
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: baseIDs.PrototypeClassificationID(),
				immutableProps:   baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(),
			},
			setup: func(t *testing.T) {
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()

			},
			wantErr: errorConstants.MockError,
		},
		{
			name: "AssetSupplyNegativeIncorrectFormatError",
			args: args{
				from:             s.genesisAddress,
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: baseIDs.PrototypeClassificationID(),
				immutableProps:   baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(baseProperties.NewMetaProperty(constantProperties.SupplyProperty.GetID().GetKey(), baseData.NewNumberData(math.NewInt(-1)))),
			},
			setup: func(t *testing.T) {
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()

			},
			wantErr: errorConstants.IncorrectFormat,
		},
		{
			name: "MintAuxiliaryFailure",
			args: args{
				from:             s.genesisAddress,
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: baseIDs.PrototypeClassificationID(),
				immutableProps:   baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(),
			},
			setup: func(t *testing.T) {
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.mintAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()

			},
			wantErr: errorConstants.MockError,
		},
		{
			name: "BondAuxiliaryFailure",
			args: args{
				from:             s.bondAuxiliaryFailureAddress,
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: baseIDs.PrototypeClassificationID(),
				immutableProps:   baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(baseProperties.NewMetaProperty(constantProperties.BondAmountProperty.GetID().GetKey(), baseData.NewNumberData(math.NewInt(1)))),
			},
			setup: func(t *testing.T) {
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.mintAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.bondAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "BondAmountPropertyMetaDataError",
			args: args{
				from:             s.genesisAddress,
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: baseIDs.PrototypeClassificationID(),
				immutableProps:   baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(s.randomBondMetaProperty),
			},
			setup: func(t *testing.T) {
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.mintAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			wantErr: errorConstants.MetaDataError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s.authorizeAuxiliaryKeeper.ExpectedCalls = nil
			s.authorizeAuxiliaryKeeper.Calls = nil
			s.conformAuxiliaryKeeper.ExpectedCalls = nil
			s.conformAuxiliaryKeeper.Calls = nil
			s.mintAuxiliaryKeeper.ExpectedCalls = nil
			s.mintAuxiliaryKeeper.Calls = nil
			s.bondAuxiliaryKeeper.ExpectedCalls = nil
			s.bondAuxiliaryKeeper.Calls = nil

			tt.setup(t)
			got, err := s.TransactionKeeper.Transact(sdkTypes.WrapSDKContext(s.Context), NewMessage(
				tt.args.from,
				baseIDs.PrototypeIdentityID(),
				tt.args.toID,
				tt.args.classificationID,
				tt.args.immutableProps,
				baseLists.NewPropertyList(),
				tt.args.mutableProps,
				baseLists.NewPropertyList(),
			).(helpers.Message))

			if (err != nil) && !tt.wantErr.Is(err) {
				t.Errorf("unexpected error: %v", err)
			}

			assert.Equal(t, tt.want, got)

		})
	}
}
