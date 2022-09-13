// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	base2 "github.com/AssetMantle/modules/schema/properties/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	metaUtilities "github.com/AssetMantle/modules/utilities/string"
)

//func Test_Immutables(t *testing.T) {
//	testProperty := base2.NewProperty(baseIDs.NewID("ID"), baseData.NewHeightData(baseTypes.NewHeight(123)))
//	testImmutables := Immutables{base.NewPropertyList(testProperty)}
//
//	require.Equal(t, Immutables{PropertyList: base.NewPropertyList(testProperty)}, testImmutables)
//	require.Equal(t, base.NewPropertyList(testProperty), testImmutables.GetImmutablePropertyList())
//	require.Equal(t, baseIDs.NewID(metaUtilities.Hash([]string{testProperty.GetHash().String()}...)), testImmutables.GenerateHashID())
//	require.Equal(t, baseIDs.NewID(""), Immutables{base.NewPropertyList()}.GenerateHashID())
//}

func TestImmutables_GenerateHashID(t *testing.T) {

	testProperty := base2.NewProperty(baseIDs.NewID("ID"), baseData.NewHeightData(baseTypes.NewHeight(123)))

	type fields struct {
		PropertyList lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		{"Test for Generate Hash", fields{base.NewPropertyList(testProperty)}, baseIDs.NewID(metaUtilities.Hash([]string{testProperty.GetHash().String()}...))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			immutables := Immutables{
				PropertyList: tt.fields.PropertyList,
			}
			if got := immutables.GenerateHashID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateHashID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImmutables_GetImmutablePropertyList(t *testing.T) {

	testProperty := base2.NewProperty(baseIDs.NewID("ID"), baseData.NewHeightData(baseTypes.NewHeight(123)))

	type fields struct {
		PropertyList lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   lists.PropertyList
	}{
		{"Test for GetImmutablePropertyList", fields{base.NewPropertyList(testProperty)}, base.NewPropertyList(testProperty)},
		{"Test for nil case", fields{base.NewPropertyList()}, base.NewPropertyList()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			immutables := Immutables{
				PropertyList: tt.fields.PropertyList,
			}
			if got := immutables.GetImmutablePropertyList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetImmutablePropertyList() = %v, want %v", got, tt.want)
			}
		})
	}
}
