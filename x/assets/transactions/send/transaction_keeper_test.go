// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package send

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/utilities/random"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/mapper"
	"github.com/AssetMantle/modules/x/assets/parameters"
	"github.com/AssetMantle/modules/x/assets/record"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
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
	tendermintDB "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/cometbft/cometbft/libs/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/store"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/mock"
	"math/rand"
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
	Denom         = "stake"
	ChainID       = "testChain"
	GenesisSupply = 1000000000000
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

	fromAddress = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())

	asset   = randomAssetGenerator(baseProperties.NewMetaProperty(constantProperties.LockHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(0))), nil)
	assetID = baseIDs.NewAssetID(asset.GetClassificationID(), asset.GetImmutables()).(*baseIDs.AssetID)

	immutableLockAsset   = randomAssetGenerator(baseProperties.NewMetaProperty(constantProperties.LockHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(-1))), nil)
	immutableLockAssetID = baseIDs.NewAssetID(immutableLockAsset.GetClassificationID(), immutableLockAsset.GetImmutables()).(*baseIDs.AssetID)

	mutableLockAsset   = randomAssetGenerator(nil, baseProperties.NewMetaProperty(constantProperties.LockHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(-1))))
	mutableLockAssetID = baseIDs.NewAssetID(mutableLockAsset.GetClassificationID(), mutableLockAsset.GetImmutables()).(*baseIDs.AssetID)

	randomAsset   = randomAssetGenerator(nil, nil)
	randomAssetID = baseIDs.NewAssetID(randomAsset.GetClassificationID(), randomAsset.GetImmutables()).(*baseIDs.AssetID)

	moduleStoreKey = sdkTypes.NewKVStoreKey(constants.ModuleName)

	authenticateAuxiliaryKeeper         = new(MockAuxiliaryKeeper)
	authenticateAuxiliaryFailureAddress = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	_                                   = authenticateAuxiliaryKeeper.On("Help", mock.Anything, authenticate.NewAuxiliaryRequest(authenticateAuxiliaryFailureAddress, baseIDs.PrototypeIdentityID())).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	_                                   = authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	authenticateAuxiliary               = new(MockAuxiliary)
	_                                   = authenticateAuxiliary.On("GetKeeper").Return(authenticateAuxiliaryKeeper)

	supplementAuxiliaryKeeper       = new(MockAuxiliaryKeeper)
	supplementAuxiliaryFailureAsset = randomAssetGenerator(
		baseProperties.NewMesaProperty(constantProperties.LockHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(0))),
		nil,
	)
	supplementAuxiliaryFailureAssetID = baseIDs.NewAssetID(supplementAuxiliaryFailureAsset.GetClassificationID(), supplementAuxiliaryFailureAsset.GetImmutables()).(*baseIDs.AssetID)

	mesaLockAsset   = randomAssetGenerator(baseProperties.NewMesaProperty(constantProperties.LockHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(1))), nil)
	mesaLockAssetID = baseIDs.NewAssetID(mesaLockAsset.GetClassificationID(), mesaLockAsset.GetImmutables()).(*baseIDs.AssetID)
	_               = supplementAuxiliaryKeeper.On("Help", mock.Anything, supplement.NewAuxiliaryRequest(mesaLockAsset.GetProperty(constantProperties.LockHeightProperty.GetID()))).Return(supplement.NewAuxiliaryResponse(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constantProperties.LockHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(1))))), nil)

	unrevealedLockAsset   = randomAssetGenerator(baseProperties.NewMesaProperty(constantProperties.LockHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(2))), nil)
	unrevealedLockAssetID = baseIDs.NewAssetID(unrevealedLockAsset.GetClassificationID(), unrevealedLockAsset.GetImmutables()).(*baseIDs.AssetID)
	_                     = supplementAuxiliaryKeeper.On("Help", mock.Anything, supplement.NewAuxiliaryRequest(unrevealedLockAsset.GetProperty(constantProperties.LockHeightProperty.GetID()))).Return(supplement.NewAuxiliaryResponse(baseLists.NewPropertyList()), nil)

	supplementAuxiliaryAuxiliary = new(MockAuxiliary)
	_                            = supplementAuxiliaryAuxiliary.On("GetKeeper").Return(supplementAuxiliaryKeeper)

	transferAuxiliaryKeeper         = new(MockAuxiliaryKeeper)
	transferAuxiliaryFailureAsset   = randomAssetGenerator(baseProperties.NewMetaProperty(constantProperties.LockHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(0))), nil)
	transferAuxiliaryFailureAssetID = baseIDs.NewAssetID(transferAuxiliaryFailureAsset.GetClassificationID(), transferAuxiliaryFailureAsset.GetImmutables()).(*baseIDs.AssetID)
	_                               = transferAuxiliaryKeeper.On("Help", mock.Anything, transfer.NewAuxiliaryRequest(baseIDs.PrototypeIdentityID(), baseIDs.PrototypeIdentityID(), transferAuxiliaryFailureAssetID, sdkTypes.NewInt(1))).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	_                               = transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	transferAuxiliaryAuxiliary      = new(MockAuxiliary)
	_                               = transferAuxiliaryAuxiliary.On("GetKeeper").Return(transferAuxiliaryKeeper)

	codec = baseHelpers.TestCodec()

	paramsStoreKey           = sdkTypes.NewKVStoreKey(paramsTypes.StoreKey)
	paramsTransientStoreKeys = sdkTypes.NewTransientStoreKey(paramsTypes.TStoreKey)
	ParamsKeeper             = paramsKeeper.NewKeeper(codec, codec.GetLegacyAmino(), paramsStoreKey, paramsTransientStoreKeys)

	setContext = func() sdkTypes.Context {
		memDB := tendermintDB.NewMemDB()
		commitMultiStore := store.NewCommitMultiStore(memDB)
		commitMultiStore.MountStoreWithDB(moduleStoreKey, storeTypes.StoreTypeIAVL, memDB)
		commitMultiStore.MountStoreWithDB(paramsStoreKey, storeTypes.StoreTypeIAVL, memDB)
		commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, storeTypes.StoreTypeTransient, memDB)
		_ = commitMultiStore.LoadLatestVersion()
		return sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{ChainID: ChainID}, false, log.NewNopLogger())
	}
	Context = setContext()

	parameterManager = parameters.Prototype().Initialize(ParamsKeeper.Subspace(constants.ModuleName).WithKeyTable(parameters.Prototype().GetKeyTable())).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(Denom)))))).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(true))))).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(true))))).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(true))))).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.UnwrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(Denom))))))

	TransactionKeeper = transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey), parameterManager, authenticateAuxiliary, supplementAuxiliaryAuxiliary, transferAuxiliaryAuxiliary}

	_ = TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).
		Add(record.NewRecord(asset)).
		Add(record.NewRecord(supplementAuxiliaryFailureAsset)).
		Add(record.NewRecord(transferAuxiliaryFailureAsset)).
		Add(record.NewRecord(immutableLockAsset)).
		Add(record.NewRecord(mutableLockAsset)).
		Add(record.NewRecord(mesaLockAsset)).
		Add(record.NewRecord(unrevealedLockAsset))
)

func TestTransactionKeeperTransact(t *testing.T) {
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
			args{fromAddress, assetID, rand.Intn(GenesisSupply)},
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
			args{fromAddress, randomAssetID, 1},
			func() {},
			nil,
			errorConstants.EntityNotFound,
		},
		{
			"identityAuthenticationFailure",
			args{authenticateAuxiliaryFailureAddress, assetID, 1},
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
			args{fromAddress, immutableLockAssetID, 1},
			func() {
			},
			nil,
			errorConstants.NotAuthorized,
		},
		{
			"sendAssetWithMutableLock",
			args{fromAddress, mutableLockAssetID, 1},
			func() {
			},
			nil,
			errorConstants.NotAuthorized,
		},
		{
			"sendAssetWithMesaLock",
			args{fromAddress, mesaLockAssetID, 1},
			func() {
				Context = Context.WithBlockHeight(2)
			},
			newTransactionResponse(),
			nil,
		},
		{
			"sendAssetWithUnrevealedLock",
			args{fromAddress, unrevealedLockAssetID, 1},
			func() {
			},
			nil,
			errorConstants.MetaDataError,
		},
		{
			"supplementAuxiliaryFailure",
			args{fromAddress, supplementAuxiliaryFailureAssetID, 0},
			func() {
				supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			nil,
			errorConstants.MockError,
		},
		{
			"transferAuxiliaryFailure",
			args{fromAddress, transferAuxiliaryFailureAssetID, 1},
			func() {
			},
			nil,
			errorConstants.MockError,
		},
		{
			"sendInMultiAssetScenario",
			args{fromAddress, assetID, 1},
			func() {
				for i := 0; i < 10000; i++ {
					_ = TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(record.NewRecord(randomAssetGenerator(nil, nil)))
				}
			},
			newTransactionResponse(),
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.setup()

			got, err := TransactionKeeper.Transact(sdkTypes.WrapSDKContext(Context), NewMessage(tt.args.from, baseIDs.PrototypeIdentityID(), baseIDs.PrototypeIdentityID(), tt.args.assetID, sdkTypes.NewInt(int64(tt.args.value))).(helpers.Message))

			if (tt.wantErr != nil && !tt.wantErr.Is(err)) || (tt.wantErr == nil && err != nil) {
				t.Errorf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Error("unexpected response")
			}
		})
	}
}
