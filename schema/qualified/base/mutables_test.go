// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
)

func Test_Mutables(t *testing.T) {

	testProperty := baseProperties.NewMesaProperty(baseIDs.NewStringID("ID"), baseData.NewStringData("Data"))
	testProperties := baseLists.NewPropertyList(testProperty)
	testMutables := mutables{testProperties}
	require.Equal(t, mutables{PropertyList: testProperties}, testMutables)
	require.Equal(t, testProperties, testMutables.GetMutablePropertyList())
	mutatedTestProperty := baseProperties.NewMesaProperty(baseIDs.NewStringID("ID"), baseData.NewStringData("Data2"))
	require.Equal(t, mutables{PropertyList: baseLists.NewPropertyList(mutatedTestProperty)}, testMutables.Mutate(mutatedTestProperty))

}

func TestMutables_GetMutablePropertyList(t *testing.T) {

	testProperty := baseProperties.NewProperty(baseIDs.NewID("ID"), baseData.NewStringData("Data"))
	testProperties := baseLists.NewPropertyList(testProperty)

	type fields struct {
		Properties lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   lists.PropertyList
	}{
		{"Test for GetMutablePropertyList", fields{testProperties}, testProperties},
		{"Test for nil case", fields{nil}, baseLists.NewPropertyList()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mutables := Mutables{
				Properties: tt.fields.Properties,
			}
			if got := mutables.GetMutablePropertyList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMutablePropertyList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMutables_Mutate(t *testing.T) {

	testProperty := baseProperties.NewProperty(baseIDs.NewID("ID"), baseData.NewStringData("Data"))
	testProperties := baseLists.NewPropertyList(testProperty)
	mutatedTestProperty := baseProperties.NewProperty(baseIDs.NewID("ID"), baseData.NewStringData("Data2"))
	mutatedTestProperties := baseLists.NewPropertyList(mutatedTestProperty)

	type fields struct {
		Properties lists.PropertyList
	}
	type args struct {
		propertyList []properties.Property
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   qualified.Mutables
	}{
		{"Test for Mutate", fields{testProperties}, args{[]properties.Property{mutatedTestProperty}}, Mutables{mutatedTestProperties}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mutables := Mutables{
				Properties: tt.fields.Properties,
			}
			if got := mutables.Mutate(tt.args.propertyList...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mutate() = %v, want %v", got, tt.want)
			}
		})
	}
}
