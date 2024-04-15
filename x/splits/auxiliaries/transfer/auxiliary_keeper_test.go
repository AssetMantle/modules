// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transfer

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/utilities/random"
	"github.com/AssetMantle/modules/x/splits/constants"
	"github.com/AssetMantle/modules/x/splits/key"
	"github.com/AssetMantle/modules/x/splits/mappable"
	"github.com/AssetMantle/modules/x/splits/mapper"
	"github.com/AssetMantle/modules/x/splits/parameters"
	"github.com/AssetMantle/modules/x/splits/parameters/transfer_enabled"
	"github.com/AssetMantle/modules/x/splits/record"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents"
	"github.com/AssetMantle/schema/go/documents/base"
	"github.com/AssetMantle/schema/go/errors"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseParameters "github.com/AssetMantle/schema/go/parameters/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/AssetMantle/schema/go/types"
	baseTypes "github.com/AssetMantle/schema/go/types/base"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/mock"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"
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

// mockAuxiliaryRequest
type mockAuxiliaryRequest struct {
}

func (mockAuxiliaryRequest) Validate() error {
	return nil
}

var _ helpers.AuxiliaryRequest = (*mockAuxiliaryRequest)(nil)

const (
	ChainID       = "testChain"
	Denom         = "stake"
	GenesisSupply = 1000000000
)

var (
	testSendAmount = sdkTypes.NewInt(100)

	testFromIdentity   = base.NewNameIdentity(baseIDs.NewStringID(random.GenerateUniqueIdentifier()), baseData.NewListData())
	testFromIdentityID = testFromIdentity.(documents.NameIdentity).GetNameIdentityID()

	testToIdentity   = base.NewNameIdentity(baseIDs.NewStringID(random.GenerateUniqueIdentifier()), baseData.NewListData())
	testToIdentityID = testToIdentity.(documents.NameIdentity).GetNameIdentityID()

	testCoinAsset   = base.NewCoinAsset(Denom)
	testCoinAssetID = testCoinAsset.GetCoinAssetID()

	uninitializedCoinAsset   = base.NewCoinAsset("uninitialized")
	uninitializedCoinAssetID = uninitializedCoinAsset.GetCoinAssetID()

	encodingConfig = simapp.MakeTestEncodingConfig()

	paramsStoreKey           = sdkTypes.NewKVStoreKey(paramsTypes.StoreKey)
	paramsTransientStoreKeys = sdkTypes.NewTransientStoreKey(paramsTypes.TStoreKey)
	ParamsKeeper             = paramsKeeper.NewKeeper(encodingConfig.Marshaler, encodingConfig.Amino, paramsStoreKey, paramsTransientStoreKeys)

	moduleStoreKey  = sdkTypes.NewKVStoreKey(constants.ModuleName)
	AuxiliaryKeeper = auxiliaryKeeper{mapper.Prototype().Initialize(moduleStoreKey), parameterManager}

	setContext = func() sdkTypes.Context {
		memDB := tendermintDB.NewMemDB()
		commitMultiStore := store.NewCommitMultiStore(memDB)
		commitMultiStore.MountStoreWithDB(moduleStoreKey, sdkTypes.StoreTypeIAVL, memDB)
		commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
		commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
		_ = commitMultiStore.LoadLatestVersion()
		return sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{ChainID: ChainID}, false, log.NewNopLogger())

	}

	Context = setContext()

	parameterManager = parameters.Prototype().Initialize(ParamsKeeper.Subspace(constants.ModuleName).WithKeyTable(parameters.Prototype().GetKeyTable())).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(baseParameters.NewParameter(baseProperties.NewMetaProperty(transfer_enabled.ID, baseData.NewBooleanData(true)))))

	_ = AuxiliaryKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).
		Add(record.NewRecord(baseIDs.NewSplitID(testCoinAssetID, testFromIdentityID), baseTypes.NewSplit(sdkTypes.NewInt(GenesisSupply))))
)

func Test_auxiliaryKeeper_Help(t *testing.T) {
	tests := []struct {
		name    string
		setup   func()
		request helpers.AuxiliaryRequest
		want    helpers.AuxiliaryResponse
		wantErr errors.Error
	}{
		{
			"valid request",
			func() {},
			NewAuxiliaryRequest(testFromIdentityID, testToIdentityID, testCoinAssetID, testSendAmount),
			newAuxiliaryResponse(),
			nil,
		},
		{
			"split not present",
			func() {},
			NewAuxiliaryRequest(testFromIdentityID, testToIdentityID, uninitializedCoinAssetID, testSendAmount),
			nil,
			errorConstants.EntityNotFound,
		},
		{
			"insufficient balance",
			func() {},
			NewAuxiliaryRequest(testFromIdentityID, testToIdentityID, testCoinAssetID, sdkTypes.NewInt(GenesisSupply+1)),
			nil,
			errorConstants.InsufficientBalance,
		},
		{
			"send zero",
			func() {},
			NewAuxiliaryRequest(testFromIdentityID, testToIdentityID, testCoinAssetID, sdkTypes.ZeroInt()),
			nil,
			errorConstants.InvalidRequest,
		},
		{
			"send negative",
			func() {},
			NewAuxiliaryRequest(testFromIdentityID, testToIdentityID, testCoinAssetID, sdkTypes.NewInt(-1)),
			nil,
			errorConstants.InvalidRequest,
		},
		{
			"invalid from identity",
			func() {},
			NewAuxiliaryRequest(&baseIDs.IdentityID{HashID: &baseIDs.HashID{IDBytes: []byte("invalid")}}, testToIdentityID, testCoinAssetID, testSendAmount),
			nil,
			errorConstants.InvalidRequest,
		},
		{
			"invalid to identity",
			func() {},
			NewAuxiliaryRequest(testFromIdentityID, &baseIDs.IdentityID{HashID: &baseIDs.HashID{IDBytes: []byte("invalid")}}, testCoinAssetID, testSendAmount),
			nil,
			errorConstants.InvalidRequest,
		},
		{
			"invalid request type",
			func() {},
			mockAuxiliaryRequest{},
			nil,
			errorConstants.InvalidRequest,
		},
		{
			"with many splits",
			func() {
				for i := 0; i < 100000; i++ {
					_ = AuxiliaryKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).
						Add(record.NewRecord(baseIDs.NewSplitID(base.NewCoinAsset(random.GenerateUniqueIdentifier()).GetCoinAssetID(), base.NewNameIdentity(baseIDs.NewStringID(random.GenerateUniqueIdentifier()), baseData.NewListData()).GetNameIdentityID()), baseTypes.NewSplit(sdkTypes.NewInt(int64(rand.Intn(100000000000))))))
				}
			},
			NewAuxiliaryRequest(testFromIdentityID, testToIdentityID, testCoinAssetID, testSendAmount),
			newAuxiliaryResponse(),
			nil,
		},
		{
			"transfer full amount",
			func() {
				_ = AuxiliaryKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).
					Remove(record.NewRecord(baseIDs.NewSplitID(testCoinAssetID, testFromIdentityID), baseTypes.NewSplit(sdkTypes.OneInt()))).
					Add(record.NewRecord(baseIDs.NewSplitID(testCoinAssetID, testFromIdentityID), baseTypes.NewSplit(sdkTypes.NewInt(GenesisSupply))))
			},
			NewAuxiliaryRequest(testFromIdentityID, testToIdentityID, testCoinAssetID, sdkTypes.NewInt(GenesisSupply)),
			newAuxiliaryResponse(),
			nil,
		},
		{
			"transfer not enabled",
			func() {
				parameterManager.Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(baseParameters.NewParameter(baseProperties.NewMetaProperty(transfer_enabled.ID, baseData.NewBooleanData(false)))))
			},
			NewAuxiliaryRequest(testFromIdentityID, testToIdentityID, testCoinAssetID, testSendAmount),
			nil,
			errorConstants.NotAuthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			var fromSplitBefore, toSplitBefore types.Split
			if tt.wantErr == nil {
				fromSplitBefore = mappable.GetSplit(AuxiliaryKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).FetchRecord(key.NewKey(baseIDs.NewSplitID(tt.request.(auxiliaryRequest).AssetID, tt.request.(auxiliaryRequest).FromID))).GetMappable())
				toSplitBefore = mappable.GetSplit(AuxiliaryKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).FetchRecord(key.NewKey(baseIDs.NewSplitID(tt.request.(auxiliaryRequest).AssetID, tt.request.(auxiliaryRequest).ToID))).GetMappable())
			}

			got, err := AuxiliaryKeeper.Help(sdkTypes.WrapSDKContext(Context), tt.request)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Help() got = %v, want %v", got, tt.want)
			}

			if err != nil && tt.wantErr == nil || err == nil && tt.wantErr != nil || err != nil && tt.wantErr != nil && !tt.wantErr.Is(err) {
				t.Errorf("Help() err = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				fromSplitAfter := mappable.GetSplit(AuxiliaryKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).FetchRecord(key.NewKey(baseIDs.NewSplitID(tt.request.(auxiliaryRequest).AssetID, tt.request.(auxiliaryRequest).FromID))).GetMappable())
				toSplitAfter := mappable.GetSplit(AuxiliaryKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).FetchRecord(key.NewKey(baseIDs.NewSplitID(tt.request.(auxiliaryRequest).AssetID, tt.request.(auxiliaryRequest).ToID))).GetMappable())

				if toSplitAfter == nil {
					t.Errorf("to split not created")
				}

				if fromSplitBefore == nil {
					t.Errorf("from split was not present")
				}

				if fromSplitAfter == nil {
					if !fromSplitBefore.GetValue().Equal(tt.request.(auxiliaryRequest).Value) {
						t.Errorf("from split incorrectly deleted")
					}
				} else if !fromSplitBefore.GetValue().Sub(tt.request.(auxiliaryRequest).Value).Equal(fromSplitAfter.GetValue()) {
					t.Errorf("from split not updated")
				}

				if toSplitBefore == nil {
					if !toSplitAfter.GetValue().Equal(tt.request.(auxiliaryRequest).Value) {
						t.Errorf("to split incorrectly created")
					}
				} else if !toSplitBefore.GetValue().Add(tt.request.(auxiliaryRequest).Value).Equal(toSplitAfter.GetValue()) {
					t.Errorf("to split not updated")
				}
			}
		})
	}
}
