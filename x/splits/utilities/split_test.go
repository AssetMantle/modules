// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"fmt"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	baseTypes "github.com/AssetMantle/schema/go/types/base"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/splits/mappable"
	"github.com/AssetMantle/modules/x/splits/mapper"
)

func createTestInput1(t *testing.T) (sdkTypes.Context, helpers.Mapper) {
	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
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
	testAssetID := baseIDs.NewCoinID(baseIDs.NewStringID("OwnerID"))
	testRate := sdkTypes.OneInt()
	split := baseTypes.NewSplit(testOwnerIdentityID, testAssetID, testRate)
	context, testMapper := createTestInput1(t)
	testSplits := testMapper.NewCollection(sdkTypes.WrapSDKContext(context)).Add(mappable.NewMappable(split))
	type args struct {
		splits  helpers.Collection
		ownerID ids.IdentityID
		assetID ids.AssetID
		value   sdkTypes.Int
	}
	tests := []struct {
		name    string
		args    args
		want    helpers.Collection
		wantErr bool
	}{
		{"+ve", args{testSplits, testOwnerIdentityID, testAssetID, sdkTypes.NewInt(100)}, testSplits.Mutate(mappable.NewMappable(split.Receive(sdkTypes.NewInt(100)))), false},
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
	testAssetID := baseIDs.NewCoinID(baseIDs.NewStringID("OwnerID"))
	testRate := sdkTypes.NewInt(10)
	split := baseTypes.NewSplit(testOwnerIdentityID, testAssetID, testRate)
	context, testMapper := createTestInput1(t)
	testSplits := testMapper.NewCollection(sdkTypes.WrapSDKContext(context)).Add(mappable.NewMappable(split))
	type args struct {
		splits  helpers.Collection
		ownerID ids.IdentityID
		assetID ids.AssetID
		value   sdkTypes.Int
	}
	tests := []struct {
		name    string
		args    args
		want    helpers.Collection
		wantErr bool
	}{
		{"+ve", args{testSplits, testOwnerIdentityID, testAssetID, sdkTypes.NewInt(9)}, testSplits.Mutate(mappable.NewMappable(split)), false},
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
