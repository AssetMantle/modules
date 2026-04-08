// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerate

import (
	"cosmossdk.io/math"
	storeTypes "cosmossdk.io/store/types"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/mapper"
	"github.com/AssetMantle/modules/x/assets/parameters"
	recordassets "github.com/AssetMantle/modules/x/assets/record"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/documents"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/parameters/base"
	"github.com/AssetMantle/schema/properties"
	baseMetaProp "github.com/AssetMantle/schema/properties/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

type testSetup struct {
	Context                                   sdkTypes.Context
	genesisAddress                            sdkTypes.AccAddress
	parameterManager                          helpers.ParameterManager
	TransactionKeeper                         transactionKeeper
	authenticateAuxiliaryFailureAddress        sdkTypes.AccAddress
	authorizeAuxiliaryFailureClassificationID ids.IdentityID
	authorizeAuxiliaryKeeper                  *testutil.MockAuxiliaryKeeper
	renumerateAuxiliaryKeeper                 *testutil.MockAuxiliaryKeeper
	supplementAuxiliaryKeeper                 *testutil.MockAuxiliaryKeeper
	testNewAsset                              documents.Asset
	testNewAssetID                            *baseIDs.AssetID
	SupplyNotRevealedProperty                 properties.Property
	NegativeSupplyProperty                    properties.Property
}

func setupTest(t *testing.T) *testSetup {
	t.Helper()

	moduleStoreKey := storeTypes.NewKVStoreKey(constants.ModuleName)
	ctx, _, genesisAddress := testutil.NewTestContextWithBankAuth(t, constants.ModuleName, moduleStoreKey)

	randomInteger := int64(100)
	randomNegativeInteger := int64(-1)

	SupplyNotRevealedProperty := baseMetaProp.NewMesaProperty(constantProperties.SupplyProperty.GetID().GetKey(), baseData.NewNumberData(math.NewInt(randomNegativeInteger)))
	NegativeSupplyProperty := baseMetaProp.NewMetaProperty(constantProperties.SupplyProperty.GetID().GetKey(), baseData.NewNumberData(math.NewInt(randomNegativeInteger)))

	propList := baseQualified.NewMutables(mutableMetaProperties).GetMutablePropertyList().Add(
		baseProperties.NewMesaProperty(constantProperties.SupplyProperty.GetKey(), baseData.NewNumberData(math.NewInt(randomInteger))),
	)
	newMutables := baseQualified.NewMutables(propList)
	testNewAsset := baseDocuments.NewAsset(baseIDs.NewClassificationID(immutables, newMutables), immutables, newMutables)
	testNewAssetID := baseIDs.NewAssetID(testNewAsset.GetClassificationID(), testNewAsset.GetImmutables()).(*baseIDs.AssetID)

	authenticateAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	authenticateAuxiliaryFailureAddress := testutil.TestAddress()
	authenticateAuxiliaryKeeper.On("Help", mock.Anything, authenticate.NewAuxiliaryRequest(&Message{From: authenticateAuxiliaryFailureAddress.String(), FromID: baseIDs.PrototypeIdentityID().(*baseIDs.IdentityID)})).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	authenticateAuxiliary := new(testutil.MockAuxiliary)
	authenticateAuxiliary.On("GetKeeper").Return(authenticateAuxiliaryKeeper)

	authorizeAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	authorizeAuxiliaryFailureClassificationID := baseIDs.NewIdentityID(testNewAsset.GetClassificationID(), immutables)
	authorizeAuxiliary := new(testutil.MockAuxiliary)
	authorizeAuxiliary.On("GetKeeper").Return(authorizeAuxiliaryKeeper)

	renumerateAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	renumerateAuxiliary := new(testutil.MockAuxiliary)
	renumerateAuxiliary.On("GetKeeper").Return(renumerateAuxiliaryKeeper)

	supplementAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	supplementAuxiliary := new(testutil.MockAuxiliary)
	supplementAuxiliary.On("GetKeeper").Return(supplementAuxiliaryKeeper)

	parameterManager, _ := parameters.Prototype().Initialize(moduleStoreKey).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(testutil.Denom))))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.UnwrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(testutil.Denom))))).
		Update(ctx)

	TransactionKeeper := transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey),
		parameterManager,
		authenticateAuxiliary,
		authorizeAuxiliary,
		renumerateAuxiliary,
		supplementAuxiliary,
	}

	return &testSetup{
		Context:                                   ctx,
		genesisAddress:                            genesisAddress,
		parameterManager:                          parameterManager,
		TransactionKeeper:                         TransactionKeeper,
		authenticateAuxiliaryFailureAddress:        authenticateAuxiliaryFailureAddress,
		authorizeAuxiliaryFailureClassificationID: authorizeAuxiliaryFailureClassificationID,
		authorizeAuxiliaryKeeper:                  authorizeAuxiliaryKeeper,
		renumerateAuxiliaryKeeper:                 renumerateAuxiliaryKeeper,
		supplementAuxiliaryKeeper:                 supplementAuxiliaryKeeper,
		testNewAsset:                              testNewAsset,
		testNewAssetID:                            testNewAssetID,
		SupplyNotRevealedProperty:                 SupplyNotRevealedProperty,
		NegativeSupplyProperty:                    NegativeSupplyProperty,
	}
}

func TestTransactionKeeperTransact(t *testing.T) {
	s := setupTest(t)

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
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: s.testNewAssetID,
				denom:   testutil.Denom,
			},
			setup: func() {
				s.renumerateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(supplement.NewAuxiliaryResponse(baseLists.NewPropertyList(constantProperties.SupplyProperty)), nil).Once()
			},
			want:    newTransactionResponse(),
			wantErr: nil,
		},
		{
			name: "RenumeratePropertyDisabled",
			args: args{
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: s.testNewAssetID,
				denom:   testutil.Denom,
			},
			setup: func() {
				s.parameterManager.Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(false)))).Update(sdkTypes.WrapSDKContext(s.Context))
			},
			want:    nil,
			wantErr: errorConstants.NotAuthorized,
		},
		{
			name: "AuthenticationFailure",
			args: args{
				from:    s.authenticateAuxiliaryFailureAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: baseIDs.PrototypeAssetID(),
				denom:   testutil.Denom,
			},
			setup: func() {
				s.parameterManager.Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).Update(sdkTypes.WrapSDKContext(s.Context))
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "EntityNotFoundFailure",
			args: args{
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: baseIDs.PrototypeAssetID(),
				denom:   testutil.Denom,
			},
			setup: func() {
			},
			want:    nil,
			wantErr: errorConstants.EntityNotFound,
		},
		{
			name: "AuthorizeAuxiliaryFailure",
			args: args{
				from:    s.genesisAddress,
				fromID:  s.authorizeAuxiliaryFailureClassificationID,
				assetID: s.testNewAssetID,
				denom:   testutil.Denom,
			},
			setup: func() {
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "SupplementAuxiliaryFailure",
			args: args{
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: s.testNewAssetID,
			},
			setup: func() {
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "SupplyNotRevealedError",
			args: args{
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: s.testNewAssetID,
			},
			setup: func() {
				s.renumerateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(supplement.NewAuxiliaryResponse(baseLists.NewPropertyList(s.SupplyNotRevealedProperty)), nil).Once()
			},
			want:    nil,
			wantErr: errorConstants.MetaDataError,
		},
		{
			name: "NegativeSupplyFailure",
			args: args{
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: s.testNewAssetID,
			},
			setup: func() {
				s.renumerateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(supplement.NewAuxiliaryResponse(baseLists.NewPropertyList(s.NegativeSupplyProperty)), nil).Once()
			},
			want:    nil,
			wantErr: errorConstants.MetaDataError,
		},
		{
			name: "RenumerateAuxiliaryFailure",
			args: args{
				from:    s.genesisAddress,
				fromID:  baseIDs.PrototypeIdentityID(),
				assetID: s.testNewAssetID,
			},
			setup: func() {
				s.renumerateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(supplement.NewAuxiliaryResponse(baseLists.NewPropertyList(constantProperties.SupplyProperty)), nil).Once()
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s.authorizeAuxiliaryKeeper.ExpectedCalls = nil
			s.authorizeAuxiliaryKeeper.Calls = nil
			s.renumerateAuxiliaryKeeper.ExpectedCalls = nil
			s.renumerateAuxiliaryKeeper.Calls = nil
			s.supplementAuxiliaryKeeper.ExpectedCalls = nil
			s.supplementAuxiliaryKeeper.Calls = nil

			tt.setup()

			s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Add(recordassets.NewRecord(s.testNewAsset))

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
