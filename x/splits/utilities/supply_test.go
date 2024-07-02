// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"github.com/AssetMantle/modules/x/splits/record"
	"github.com/AssetMantle/schema/documents/base"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	baseTypes "github.com/AssetMantle/schema/types/base"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

func TestGetTotalSupply(t *testing.T) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	testOwnerIdentityID := baseIDs.NewIdentityID(classificationID, immutables)
	testAssetID := base.NewCoinAsset("OwnerID").GetCoinAssetID()
	testRate := types.NewInt(10)
	split := baseTypes.NewSplit(testRate)
	context, testMapper := createTestInput1(t)
	testSplits := testMapper.NewCollection(types.WrapSDKContext(context)).Add(record.NewRecord(baseIDs.NewSplitID(testAssetID, testOwnerIdentityID), split))
	type args struct {
		collection helpers.Collection
		assetID    ids.AssetID
	}
	tests := []struct {
		name string
		args args
		want types.Int
	}{
		{"+ve", args{testSplits, testAssetID}, testRate},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTotalSupply(tt.args.collection, tt.args.assetID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTotalSupply() = %v, want %v", got, tt.want)
			}
		})
	}
}
