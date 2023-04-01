// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/modules/splits/internal/mappable"
	"github.com/AssetMantle/modules/modules/splits/internal/mapper"
	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func createTestInput1(t *testing.T) (sdkTypes.Context, helpers.Mapper) {
	var legacyAmino = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()

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
	testOwnableID := baseIDs.NewCoinID(baseIDs.NewStringID("OwnerID"))
	testRate := sdkTypes.NewDec(1)
	split := baseTypes.NewSplit(testOwnerIdentityID, testOwnableID, testRate)
	context, testMapper := createTestInput1(t)
	testSplits := testMapper.NewCollection(sdkTypes.WrapSDKContext(context)).Add(mappable.NewMappable(split))
	type args struct {
		splits    helpers.Collection
		ownerID   ids.IdentityID
		ownableID ids.OwnableID
		value     sdkTypes.Dec
	}
	tests := []struct {
		name    string
		args    args
		want    helpers.Collection
		wantErr bool
	}{
		{"+ve", args{testSplits, testOwnerIdentityID, testOwnableID, sdkTypes.NewDec(100)}, testSplits.Mutate(mappable.NewMappable(split.Receive(sdkTypes.NewDec(100)))), false},
		{"+ve Not authorized", args{testSplits, testOwnerIdentityID, testOwnableID, sdkTypes.ZeroDec()}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddSplits(tt.args.splits, tt.args.ownerID, tt.args.ownableID, tt.args.value)
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
	testOwnableID := baseIDs.NewCoinID(baseIDs.NewStringID("OwnerID"))
	testRate := sdkTypes.NewDec(10)
	split := baseTypes.NewSplit(testOwnerIdentityID, testOwnableID, testRate)
	context, testMapper := createTestInput1(t)
	testSplits := testMapper.NewCollection(sdkTypes.WrapSDKContext(context)).Add(mappable.NewMappable(split))
	type args struct {
		splits    helpers.Collection
		ownerID   ids.IdentityID
		ownableID ids.OwnableID
		value     sdkTypes.Dec
	}
	tests := []struct {
		name    string
		args    args
		want    helpers.Collection
		wantErr bool
	}{
		{"+ve", args{testSplits, testOwnerIdentityID, testOwnableID, sdkTypes.NewDec(9)}, testSplits.Mutate(mappable.NewMappable(split)), false},
		{"+ve Not Authorized", args{testSplits, testOwnerIdentityID, testOwnableID, sdkTypes.NewDec(100)}, nil, true},
		{"+ve Not Authorized", args{testSplits, testOwnerIdentityID, testOwnableID, sdkTypes.ZeroDec()}, nil, true},
		{"+ve Entity Not found", args{testSplits, baseIDs.PrototypeIdentityID(), testOwnableID, testRate}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SubtractSplits(tt.args.splits, tt.args.ownerID, tt.args.ownableID, tt.args.value)
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
