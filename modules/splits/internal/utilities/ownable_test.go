// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"github.com/AssetMantle/modules/modules/splits/internal/mappable"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIds "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	"github.com/cosmos/cosmos-sdk/types"
	"reflect"
	"testing"
)

func TestGetOwnableTotalSplitsValue(t *testing.T) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIds.NewClassificationID(immutables, mutables)
	testOwnerIdentityID := baseIds.NewIdentityID(classificationID, immutables)
	testOwnableID := baseIds.NewOwnableID(baseIds.NewStringID("OwnerID"))
	testRate := types.NewDec(10)
	split := baseTypes.NewSplit(testOwnerIdentityID, testOwnableID, testRate)
	context, testMapper := createTestInput1(t)
	testSplits := testMapper.NewCollection(context).Add(mappable.NewMappable(split))
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
