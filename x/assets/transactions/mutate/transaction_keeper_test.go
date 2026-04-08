// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mutate

import (
	storeTypes "cosmossdk.io/store/types"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/mapper"
	recordassets "github.com/AssetMantle/modules/x/assets/record"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/maintain"
	"github.com/AssetMantle/schema/documents"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

type testSetup struct {
	Context                            sdkTypes.Context
	genesisAddress                     sdkTypes.AccAddress
	TransactionKeeper                  transactionKeeper
	authenticateAuxiliaryFailureAddress sdkTypes.AccAddress
	testAsset                          documents.Asset
	testNewAssetID                     *baseIDs.AssetID
	maintainFailureID                  ids.IdentityID
	maintainFailureClassificationID    ids.ClassificationID
	maintainAuxiliaryKeeper            *testutil.MockAuxiliaryKeeper
	conformAuxiliaryKeeper             *testutil.MockAuxiliaryKeeper
}

func setupTest(t *testing.T) *testSetup {
	t.Helper()

	moduleStoreKey := storeTypes.NewKVStoreKey(constants.ModuleName)
	ctx, _, genesisAddress := testutil.NewTestContextWithBankAuth(t, constants.ModuleName, moduleStoreKey)

	authenticateAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	authenticateAuxiliaryFailureAddress := testutil.TestAddress()
	authenticateAuxiliaryKeeper.On("Help", mock.Anything, authenticate.NewAuxiliaryRequest(&Message{From: authenticateAuxiliaryFailureAddress.String(), FromID: baseIDs.PrototypeIdentityID().(*baseIDs.IdentityID)})).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	authenticateAuxiliary := new(testutil.MockAuxiliary)
	authenticateAuxiliary.On("GetKeeper").Return(authenticateAuxiliaryKeeper)

	testAsset := baseDocuments.NewAsset(baseIDs.NewClassificationID(immutables, mutables), baseQualified.NewImmutables(baseLists.NewPropertyList()), baseQualified.NewMutables(baseLists.NewPropertyList()))
	testNewAssetID := baseIDs.NewAssetID(testAsset.GetClassificationID(), testAsset.GetImmutables()).(*baseIDs.AssetID)

	maintainFailureClassificationID := testAsset.GetClassificationID()
	maintainAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	maintainFailureID := baseIDs.NewIdentityID(classificationID, immutables)
	maintainAuxiliary := new(testutil.MockAuxiliary)
	maintainAuxiliary.On("GetKeeper").Return(maintainAuxiliaryKeeper)

	conformAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	conformAuxiliary := new(testutil.MockAuxiliary)
	conformAuxiliary.On("GetKeeper").Return(conformAuxiliaryKeeper)

	TransactionKeeper := transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey),
		authenticateAuxiliary,
		maintainAuxiliary,
		conformAuxiliary,
	}

	return &testSetup{
		Context:                            ctx,
		genesisAddress:                     genesisAddress,
		TransactionKeeper:                  TransactionKeeper,
		authenticateAuxiliaryFailureAddress: authenticateAuxiliaryFailureAddress,
		testAsset:                          testAsset,
		testNewAssetID:                     testNewAssetID,
		maintainFailureID:                  maintainFailureID,
		maintainFailureClassificationID:    maintainFailureClassificationID,
		maintainAuxiliaryKeeper:            maintainAuxiliaryKeeper,
		conformAuxiliaryKeeper:             conformAuxiliaryKeeper,
	}
}

func TestTransactionKeeperTransact(t *testing.T) {
	s := setupTest(t)

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
				from:             s.genesisAddress,
				fromID:           baseIDs.PrototypeIdentityID(),
				assetID:          s.testNewAssetID,
				mutableMetaProps: baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(),
			},
			func() {
				s.conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.maintainAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			newTransactionResponse(),
			nil,
		},
		{
			"MutateValidAsset",
			args{
				from:             s.genesisAddress,
				fromID:           baseIDs.PrototypeIdentityID(),
				assetID:          s.testNewAssetID,
				mutableMetaProps: baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(),
			},
			func() {
				s.conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.maintainAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			newTransactionResponse(),
			nil,
		},
		{
			"EntityNotFoundError",
			args{
				from:             s.genesisAddress,
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
				from:             s.authenticateAuxiliaryFailureAddress,
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
				from:             s.genesisAddress,
				fromID:           s.maintainFailureID,
				assetID:          s.testNewAssetID,
				mutableMetaProps: baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(),
			},
			func() {
				s.maintainAuxiliaryKeeper.On("Help", mock.Anything, maintain.NewAuxiliaryRequest(s.maintainFailureClassificationID, s.maintainFailureID, baseQualified.NewMutables(baseLists.NewPropertyList()))).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
			},
			nil,
			errorConstants.MockError,
		},
		{
			"ConformAuxiliaryFailure",
			args{
				from:             s.genesisAddress,
				fromID:           baseIDs.PrototypeIdentityID(),
				assetID:          s.testNewAssetID,
				mutableMetaProps: baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(),
			},
			func() {
				s.maintainAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.conformAuxiliaryKeeper.On("Help", mock.Anything, conform.NewAuxiliaryRequest(s.testAsset.GetClassificationID(), s.testAsset.GetImmutables(), baseQualified.NewMutables(baseLists.NewPropertyList()))).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
			},
			nil,
			errorConstants.MockError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Add(recordassets.NewRecord(s.testAsset))

			got, err := s.TransactionKeeper.Transact(sdkTypes.WrapSDKContext(s.Context),
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
