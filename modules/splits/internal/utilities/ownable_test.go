// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/splits/internal/mappable"
	baseData "github.com/AssetMantle/schema/x/data/base"
	"github.com/AssetMantle/schema/x/helpers"
	"github.com/AssetMantle/schema/x/ids"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	baseLists "github.com/AssetMantle/schema/x/lists/base"
	baseProperties "github.com/AssetMantle/schema/x/properties/base"
	baseQualified "github.com/AssetMantle/schema/x/qualified/base"
	baseTypes "github.com/AssetMantle/schema/x/types/base"
)

func TestGetOwnableTotalSplitsValue(t *testing.T) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	testOwnerIdentityID := baseIDs.NewIdentityID(classificationID, immutables)
	testOwnableID := baseIDs.NewCoinID(baseIDs.NewStringID("OwnerID"))
	testRate := types.NewDec(10)
	split := baseTypes.NewSplit(testOwnerIdentityID, testOwnableID, testRate)
	context, testMapper := createTestInput1(t)
	testSplits := testMapper.NewCollection(types.WrapSDKContext(context)).Add(mappable.NewMappable(split))
	type args struct {
		collection helpers.Collection
		ownableID  ids.ID
	}
	tests := []struct {
		name string
		args args
		want types.Dec
	}{
		{"+ve", args{testSplits, testOwnableID}, testRate},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOwnableTotalSplitsValue(tt.args.collection, tt.args.ownableID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOwnableTotalSplitsValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
