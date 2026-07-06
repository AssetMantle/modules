// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package send

import (
	"cosmossdk.io/math"
	storeTypes "cosmossdk.io/store/types"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/utilities/random"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/mapper"
	"github.com/AssetMantle/modules/x/assets/parameters"
	"github.com/AssetMantle/modules/x/assets/record"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/compliance"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/documents"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/parameters/base"
	"github.com/AssetMantle/schema/properties"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	baseTypes "github.com/AssetMantle/schema/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

var (
	randomMetaPropertyGenerator = func() properties.MetaProperty {
		return baseProperties.NewMetaProperty(baseIDs.NewStringID(random.GenerateUniqueIdentifier()), baseData.NewStringData(random.GenerateUniqueIdentifier()))
	}
	randomAssetGenerator = func(withImmutable, withMutable properties.Property) documents.Asset {
		immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(withImmutable, randomMetaPropertyGenerator(), randomMetaPropertyGenerator(), randomMetaPropertyGenerator()))
		mutables := baseQualified.NewMutables(baseLists.NewPropertyList(withMutable, randomMetaPropertyGenerator(), randomMetaPropertyGenerator(), randomMetaPropertyGenerator()))
		return baseDocuments.NewAsset(baseIDs.NewClassificationID(immutables, mutables), immutables, mutables)
	}

	fromAddress = testutil.TestAddress()

	asset   = randomAssetGenerator(baseProperties.NewMetaProperty(constantProperties.LockHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(0))), nil)
	assetID = baseIDs.NewAssetID(asset.GetClassificationID(), asset.GetImmutables()).(*baseIDs.AssetID)
)

type testSetup struct {
	Context                            sdkTypes.Context
	TransactionKeeper                  transactionKeeper
	authenticateAuxiliaryFailureAddress sdkTypes.AccAddress
	supplementAuxiliaryKeeper          *testutil.MockAuxiliaryKeeper
	immutableLockAssetID               *baseIDs.AssetID
	mutableLockAssetID                 *baseIDs.AssetID
	randomAssetID                      *baseIDs.AssetID
	supplementAuxiliaryFailureAssetID  *baseIDs.AssetID
	mesaLockAssetID                    *baseIDs.AssetID
	unrevealedLockAssetID              *baseIDs.AssetID
	transferAuxiliaryFailureAssetID    *baseIDs.AssetID
	complianceAuxiliaryKeeper          *testutil.MockAuxiliaryKeeper
	tierBlockedAssetID                 *baseIDs.AssetID
	tierAllowedAssetID                 *baseIDs.AssetID
}

func setupTest(t *testing.T) *testSetup {
	t.Helper()

	moduleStoreKey := storeTypes.NewKVStoreKey(constants.ModuleName)

	immutableLockAsset := randomAssetGenerator(baseProperties.NewMetaProperty(constantProperties.LockHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(-1))), nil)
	immutableLockAssetID := baseIDs.NewAssetID(immutableLockAsset.GetClassificationID(), immutableLockAsset.GetImmutables()).(*baseIDs.AssetID)

	mutableLockAsset := randomAssetGenerator(nil, baseProperties.NewMetaProperty(constantProperties.LockHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(-1))))
	mutableLockAssetID := baseIDs.NewAssetID(mutableLockAsset.GetClassificationID(), mutableLockAsset.GetImmutables()).(*baseIDs.AssetID)

	randomAsset := randomAssetGenerator(nil, nil)
	randomAssetID := baseIDs.NewAssetID(randomAsset.GetClassificationID(), randomAsset.GetImmutables()).(*baseIDs.AssetID)

	authenticateAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	authenticateAuxiliaryFailureAddress := testutil.TestAddress()
	authenticateAuxiliaryKeeper.On("Help", mock.Anything, authenticate.NewAuxiliaryRequest(&Message{From: authenticateAuxiliaryFailureAddress.String(), FromID: baseIDs.PrototypeIdentityID().(*baseIDs.IdentityID)})).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	authenticateAuxiliary := new(testutil.MockAuxiliary)
	authenticateAuxiliary.On("GetKeeper").Return(authenticateAuxiliaryKeeper)

	supplementAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	supplementAuxiliaryFailureAsset := randomAssetGenerator(
		baseProperties.NewMesaProperty(constantProperties.LockHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(0))),
		nil,
	)
	supplementAuxiliaryFailureAssetID := baseIDs.NewAssetID(supplementAuxiliaryFailureAsset.GetClassificationID(), supplementAuxiliaryFailureAsset.GetImmutables()).(*baseIDs.AssetID)

	mesaLockAsset := randomAssetGenerator(baseProperties.NewMesaProperty(constantProperties.LockHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(1))), nil)
	mesaLockAssetID := baseIDs.NewAssetID(mesaLockAsset.GetClassificationID(), mesaLockAsset.GetImmutables()).(*baseIDs.AssetID)
	supplementAuxiliaryKeeper.On("Help", mock.Anything, supplement.NewAuxiliaryRequest(mesaLockAsset.GetProperty(constantProperties.LockHeightProperty.GetID()))).Return(supplement.NewAuxiliaryResponse(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constantProperties.LockHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(1))))), nil)

	unrevealedLockAsset := randomAssetGenerator(baseProperties.NewMesaProperty(constantProperties.LockHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(2))), nil)
	unrevealedLockAssetID := baseIDs.NewAssetID(unrevealedLockAsset.GetClassificationID(), unrevealedLockAsset.GetImmutables()).(*baseIDs.AssetID)
	supplementAuxiliaryKeeper.On("Help", mock.Anything, supplement.NewAuxiliaryRequest(unrevealedLockAsset.GetProperty(constantProperties.LockHeightProperty.GetID()))).Return(supplement.NewAuxiliaryResponse(baseLists.NewPropertyList()), nil)

	supplementAuxiliaryAuxiliary := new(testutil.MockAuxiliary)
	supplementAuxiliaryAuxiliary.On("GetKeeper").Return(supplementAuxiliaryKeeper)

	transferAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	transferAuxiliaryFailureAsset := randomAssetGenerator(baseProperties.NewMetaProperty(constantProperties.LockHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(0))), nil)
	transferAuxiliaryFailureAssetID := baseIDs.NewAssetID(transferAuxiliaryFailureAsset.GetClassificationID(), transferAuxiliaryFailureAsset.GetImmutables()).(*baseIDs.AssetID)
	transferAuxiliaryKeeper.On("Help", mock.Anything, transfer.NewAuxiliaryRequest(baseIDs.PrototypeIdentityID(), baseIDs.PrototypeIdentityID(), transferAuxiliaryFailureAssetID, math.NewInt(1))).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	transferAuxiliaryAuxiliary := new(testutil.MockAuxiliary)
	transferAuxiliaryAuxiliary.On("GetKeeper").Return(transferAuxiliaryKeeper)

	// requiredTier 2 is rejected by the mock, requiredTier 1 passes; assets without requirement properties must never reach the auxiliary
	tierBlockedAsset := randomAssetGenerator(baseProperties.NewMetaProperty(constantProperties.RequiredComplianceTierProperty.GetKey(), baseData.NewNumberData(math.NewInt(2))), nil)
	tierBlockedAssetID := baseIDs.NewAssetID(tierBlockedAsset.GetClassificationID(), tierBlockedAsset.GetImmutables()).(*baseIDs.AssetID)
	tierAllowedAsset := randomAssetGenerator(baseProperties.NewMetaProperty(constantProperties.RequiredComplianceTierProperty.GetKey(), baseData.NewNumberData(math.NewInt(1))), nil)
	tierAllowedAssetID := baseIDs.NewAssetID(tierAllowedAsset.GetClassificationID(), tierAllowedAsset.GetImmutables()).(*baseIDs.AssetID)

	complianceAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	complianceAuxiliaryKeeper.On("Help", mock.Anything, compliance.NewAuxiliaryRequest(baseIDs.PrototypeIdentityID(), math.NewInt(2), "", false)).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	complianceAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	complianceAuxiliaryAuxiliary := new(testutil.MockAuxiliary)
	complianceAuxiliaryAuxiliary.On("GetKeeper").Return(complianceAuxiliaryKeeper)

	paramsStoreKey := storeTypes.NewKVStoreKey("params")

	ctx := testutil.NewTestContext(t, moduleStoreKey, paramsStoreKey)

	parameterManager := parameters.Prototype().Initialize(moduleStoreKey).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(testutil.Denom))))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.UnwrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(testutil.Denom)))))

	TransactionKeeper := transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey), parameterManager, authenticateAuxiliary, complianceAuxiliaryAuxiliary, supplementAuxiliaryAuxiliary, transferAuxiliaryAuxiliary}

	TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(ctx)).
		Add(record.NewRecord(asset)).
		Add(record.NewRecord(supplementAuxiliaryFailureAsset)).
		Add(record.NewRecord(transferAuxiliaryFailureAsset)).
		Add(record.NewRecord(immutableLockAsset)).
		Add(record.NewRecord(mutableLockAsset)).
		Add(record.NewRecord(mesaLockAsset)).
		Add(record.NewRecord(unrevealedLockAsset)).
		Add(record.NewRecord(tierBlockedAsset)).
		Add(record.NewRecord(tierAllowedAsset))

	return &testSetup{
		Context:                            ctx,
		TransactionKeeper:                  TransactionKeeper,
		authenticateAuxiliaryFailureAddress: authenticateAuxiliaryFailureAddress,
		supplementAuxiliaryKeeper:          supplementAuxiliaryKeeper,
		immutableLockAssetID:               immutableLockAssetID,
		mutableLockAssetID:                 mutableLockAssetID,
		randomAssetID:                      randomAssetID,
		supplementAuxiliaryFailureAssetID:  supplementAuxiliaryFailureAssetID,
		mesaLockAssetID:                    mesaLockAssetID,
		unrevealedLockAssetID:              unrevealedLockAssetID,
		transferAuxiliaryFailureAssetID:    transferAuxiliaryFailureAssetID,
		complianceAuxiliaryKeeper:          complianceAuxiliaryKeeper,
		tierBlockedAssetID:                 tierBlockedAssetID,
		tierAllowedAssetID:                 tierAllowedAssetID,
	}
}

func TestTransactionKeeperTransact(t *testing.T) {
	s := setupTest(t)

	type args struct {
		from    sdkTypes.AccAddress
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
		{"sendOne",
			args{fromAddress, assetID, 1},
			func() {},
			newTransactionResponse(),
			nil,
		},
		{"sendRandom",
			args{fromAddress, assetID, rand.Intn(testutil.GenesisSupply)},
			func() {},
			newTransactionResponse(),
			nil,
		},
		{
			"sendNegative",
			args{fromAddress, assetID, -1},
			func() {},
			nil,
			errorConstants.InvalidParameter,
		},
		{
			"sendAssetNotPresent",
			args{fromAddress, s.randomAssetID, 1},
			func() {},
			nil,
			errorConstants.EntityNotFound,
		},
		{
			"identityAuthenticationFailure",
			args{s.authenticateAuxiliaryFailureAddress, assetID, 1},
			func() {},
			nil,
			errorConstants.MockError,
		},
		{
			"sendZero",
			args{fromAddress, assetID, 0},
			func() {},
			newTransactionResponse(),
			nil,
		},
		{
			"sendAssetWithImmutableLock",
			args{fromAddress, s.immutableLockAssetID, 1},
			func() {
			},
			nil,
			errorConstants.NotAuthorized,
		},
		{
			"sendAssetWithMutableLock",
			args{fromAddress, s.mutableLockAssetID, 1},
			func() {
			},
			nil,
			errorConstants.NotAuthorized,
		},
		{
			"sendAssetWithMesaLock",
			args{fromAddress, s.mesaLockAssetID, 1},
			func() {
				s.Context = s.Context.WithBlockHeight(2)
			},
			newTransactionResponse(),
			nil,
		},
		{
			"sendAssetWithUnrevealedLock",
			args{fromAddress, s.unrevealedLockAssetID, 1},
			func() {
			},
			nil,
			errorConstants.MetaDataError,
		},
		{
			"supplementAuxiliaryFailure",
			args{fromAddress, s.supplementAuxiliaryFailureAssetID, 0},
			func() {
				s.supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			nil,
			errorConstants.MockError,
		},
		{
			"transferAuxiliaryFailure",
			args{fromAddress, s.transferAuxiliaryFailureAssetID, 1},
			func() {
			},
			nil,
			errorConstants.MockError,
		},
		{
			"sendComplianceBlocked",
			args{fromAddress, s.tierBlockedAssetID, 1},
			func() {},
			nil,
			errorConstants.MockError,
		},
		{
			"sendComplianceAllowed",
			args{fromAddress, s.tierAllowedAssetID, 1},
			func() {},
			newTransactionResponse(),
			nil,
		},
		{
			"sendInMultiAssetScenario",
			args{fromAddress, assetID, 1},
			func() {
				for i := 0; i < 10000; i++ {
					_ = s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Add(record.NewRecord(randomAssetGenerator(nil, nil)))
				}
			},
			newTransactionResponse(),
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.setup()

			got, err := s.TransactionKeeper.Transact(sdkTypes.WrapSDKContext(s.Context), NewMessage(tt.args.from, baseIDs.PrototypeIdentityID(), baseIDs.PrototypeIdentityID(), tt.args.assetID, math.NewInt(int64(tt.args.value))).(helpers.Message))

			if (tt.wantErr != nil && !tt.wantErr.Is(err)) || (tt.wantErr == nil && err != nil) {
				t.Errorf("unexpected error: %v", err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTransactionKeeperTransactSkipsComplianceWithoutRequirements(t *testing.T) {
	s := setupTest(t)

	got, err := s.TransactionKeeper.Transact(sdkTypes.WrapSDKContext(s.Context), NewMessage(fromAddress, baseIDs.PrototypeIdentityID(), baseIDs.PrototypeIdentityID(), assetID, math.NewInt(1)).(helpers.Message))

	require.NoError(t, err)
	assert.Equal(t, newTransactionResponse(), got)
	s.complianceAuxiliaryKeeper.AssertNotCalled(t, "Help", mock.Anything, mock.Anything)
}

func Test_keeperPrototype(t *testing.T) {
	got := keeperPrototype()
	assert.Equal(t, transactionKeeper{}, got)
}

func Test_transactionKeeper_Initialize(t *testing.T) {
	moduleStoreKey := storeTypes.NewKVStoreKey(constants.ModuleName)
	m := mapper.Prototype().Initialize(moduleStoreKey)
	pm := parameters.Prototype().Initialize(moduleStoreKey)

	authenticateAux, _ := testutil.NewNamedMockAuxiliaryPair(authenticate.Auxiliary.GetName())
	complianceAux, _ := testutil.NewNamedMockAuxiliaryPair(compliance.Auxiliary.GetName())
	supplementAux, _ := testutil.NewNamedMockAuxiliaryPair(supplement.Auxiliary.GetName())
	transferAux, _ := testutil.NewNamedMockAuxiliaryPair(transfer.Auxiliary.GetName())

	keeper := keeperPrototype().Initialize(m, pm, []interface{}{authenticateAux, complianceAux, supplementAux, transferAux})
	require.NotNil(t, keeper)
}
