// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"fmt"
	"github.com/AssetMantle/modules/x/splits/record"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	baseTypes "github.com/AssetMantle/schema/types/base"
	tendermintDB "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/splits/mapper"
)

func createTestInput1(t *testing.T) (sdkTypes.Context, helpers.Mapper) {
	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, storeTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	testMapper := mapper.Prototype().Initialize(storeKey)

	return context, testMapper
}

func TestAddSplits(t *testing.T) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	testOwnerIdentityID := baseIDs.NewIdentityID(classificationID, immutables)
	testAssetID := baseDocuments.NewCoinAsset("OwnerID").GetCoinAssetID()
	testRate := sdkTypes.OneInt()
	split := baseTypes.NewSplit(testRate)
	context, testMapper := createTestInput1(t)
	testSplits := testMapper.NewCollection(sdkTypes.WrapSDKContext(context)).Add(record.NewRecord(baseIDs.NewSplitID(testAssetID, testOwnerIdentityID), split))
	type args struct {
		splits  helpers.Collection
		ownerID ids.IdentityID
		assetID ids.AssetID
		value   math.Int
	}
	tests := []struct {
		name    string
		args    args
		want    helpers.Collection
		wantErr bool
	}{
		{"+ve", args{testSplits, testOwnerIdentityID, testAssetID, sdkTypes.NewInt(100)}, testSplits.Mutate(record.NewRecord(baseIDs.NewSplitID(testAssetID, testOwnerIdentityID), split.Subtract(sdkTypes.NewInt(100)))), false},
		{"+ve Not authorized", args{testSplits, testOwnerIdentityID, testAssetID, sdkTypes.ZeroInt()}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddSplits(tt.args.splits, tt.args.ownerID, tt.args.assetID, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddSplits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("AddSplits() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtractSplits(t *testing.T) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	testOwnerIdentityID := baseIDs.NewIdentityID(classificationID, immutables)
	testAssetID := baseDocuments.NewCoinAsset("OwnerID").GetCoinAssetID()
	testRate := sdkTypes.NewInt(10)
	split := baseTypes.NewSplit(testRate)
	context, testMapper := createTestInput1(t)
	testSplits := testMapper.NewCollection(sdkTypes.WrapSDKContext(context)).Add(record.NewRecord(baseIDs.NewSplitID(testAssetID, testOwnerIdentityID), split))
	type args struct {
		splits  helpers.Collection
		ownerID ids.IdentityID
		assetID ids.AssetID
		value   math.Int
	}
	tests := []struct {
		name    string
		args    args
		want    helpers.Collection
		wantErr bool
	}{
		{"+ve", args{testSplits, testOwnerIdentityID, testAssetID, sdkTypes.NewInt(9)}, testSplits.Mutate(record.NewRecord(baseIDs.NewSplitID(testAssetID, testOwnerIdentityID), split)), false},
		{"+ve Not Authorized", args{testSplits, testOwnerIdentityID, testAssetID, sdkTypes.NewInt(100)}, nil, true},
		{"+ve Not Authorized", args{testSplits, testOwnerIdentityID, testAssetID, sdkTypes.ZeroInt()}, nil, true},
		{"+ve Entity Not found", args{testSplits, baseIDs.PrototypeIdentityID(), testAssetID, testRate}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SubtractSplits(tt.args.splits, tt.args.ownerID, tt.args.assetID, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubtractSplits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("SubtractSplits() got = %v, want %v", got, tt.want)
			}
		})
	}
}
