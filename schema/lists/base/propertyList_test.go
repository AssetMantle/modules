// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/errors/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/traits"
)

func TestNewPropertyList(t *testing.T) {
	type args struct {
		properties []properties.Property
	}
	tests := []struct {
		name string
		args args
		want lists.PropertyList
	}{
		{"+ve", args{[]properties.Property{baseProperties.NewMetaProperty(NewStringID("ID1"), NewStringData("ImmutableData")), baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1)))}}, propertyList{List: NewList(propertiesToListables([]properties.Property{baseProperties.NewMetaProperty(NewStringID("ID1"), NewStringData("ImmutableData")), baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1)))}...)...)}},
		{"-ve", args{[]properties.Property{baseProperties.NewMetaProperty(NewStringID("ID1"), NewStringData("ImmutableData")), baseProperties.NewMetaProperty(NewStringID("ID1"), NewStringData("ImmutableData"))}}, propertyList{List: NewList(propertiesToListables([]properties.Property{baseProperties.NewMetaProperty(NewStringID("ID1"), NewStringData("ImmutableData")), baseProperties.NewMetaProperty(NewStringID("ID1"), NewStringData("ImmutableData"))}...)...)}}, // TODO: Should fail for duplicate Property
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPropertyList(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPropertyList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertiesToListables(t *testing.T) {
	type args struct {
		properties []properties.Property
	}
	tests := []struct {
		name string
		args args
		want []traits.Listable
	}{
		{"+ve", args{[]properties.Property{baseProperties.NewMetaProperty(NewStringID("ID1"), NewStringData("ImmutableData")), baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1)))}}, []traits.Listable{baseProperties.NewMetaProperty(NewStringID("ID1"), NewStringData("ImmutableData")), baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1)))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := propertiesToListables(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("propertiesToListables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_Add(t *testing.T) {
	type fields struct {
		List lists.List
	}
	type args struct {
		properties []properties.Property
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   lists.PropertyList
	}{
		{"+ve", fields{NewList(propertiesToListables([]properties.Property{baseProperties.NewMetaProperty(NewStringID("ID1"), NewStringData("ImmutableData"))}...)...)}, args{[]properties.Property{baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1)))}}, NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(NewStringID("ID1"), NewStringData("ImmutableData")), baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1)))}...)},
		{"-ve", fields{NewList(propertiesToListables([]properties.Property{baseProperties.NewMetaProperty(NewStringID("ID1"), NewStringData("ImmutableData"))}...)...)}, args{[]properties.Property{baseProperties.NewMetaProperty(NewStringID("ID1"), NewDecData(sdkTypes.NewDec(1)))}}, NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(NewStringID("ID1"), NewStringData("ImmutableData")), baseProperties.NewMetaProperty(NewStringID("ID1"), NewDecData(sdkTypes.NewDec(1)))}...)}, // TODO: Should fail as add method should not be able to mutate/add property with existing key
		{"+ve with nil", fields{NewList(propertiesToListables([]properties.Property{baseProperties.NewMetaProperty(NewStringID("ID1"), NewStringData("ImmutableData"))}...)...)}, args{nil}, NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(NewStringID("ID1"), NewStringData("ImmutableData"))}...)},                                                                                                                                                                                  // TODO: Should fail as add method should not be able to mutate/add property with existing key
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := propertyList{
				List: tt.fields.List,
			}
			if got := propertyList.Add(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_GetList(t *testing.T) {
	type fields struct {
		List lists.List
	}
	tests := []struct {
		name   string
		fields fields
		want   []properties.Property
	}{
		{"+ve", fields{NewList(propertiesToListables([]properties.Property{baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1)))}...)...)}, []properties.Property{baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1)))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := propertyList{
				List: tt.fields.List,
			}
			if got := propertyList.GetList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_GetProperty(t *testing.T) {
	type fields struct {
		List lists.List
	}
	type args struct {
		propertyID ids.PropertyID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   properties.Property
	}{
		{"+ve Meta", fields{NewList(propertiesToListables([]properties.Property{baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1)))}...)...)}, args{baseIDs.NewPropertyID(NewStringID("supply"), NewStringID("D"))}, baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1)))},
		{"+ve Mesa", fields{NewList(propertiesToListables([]properties.Property{baseProperties.NewMesaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1)))}...)...)}, args{baseIDs.NewPropertyID(NewStringID("supply"), NewStringID("D"))}, baseProperties.NewMesaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1)))},
		{"-ve", fields{NewList(propertiesToListables([]properties.Property{baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1)))}...)...)}, args{propertyID: nil}, baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1)))}, // TODO: panics if propertyID is nil
		{"-ve", fields{NewList(propertiesToListables([]properties.Property{baseProperties.NewEmptyMetaPropertyFromID(baseIDs.NewPropertyID(NewStringID("supply"), NewStringID("D")))}...)...)}, args{baseIDs.NewPropertyID(NewStringID("supply"), NewStringID("D"))}, baseProperties.NewEmptyMetaPropertyFromID(baseIDs.NewPropertyID(NewStringID("supply"), NewStringID("D")))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := propertyList{
				List: tt.fields.List,
			}
			if got := propertyList.GetProperty(tt.args.propertyID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_GetPropertyIDList(t *testing.T) {
	type fields struct {
		List lists.List
	}
	tests := []struct {
		name   string
		fields fields
		want   lists.IDList
	}{
		{"+ve", fields{NewList(propertiesToListables([]properties.Property{baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1)))}...)...)}, NewIDList().Add(baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1))).GetID())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := propertyList{
				List: tt.fields.List,
			}
			if got := propertyList.GetPropertyIDList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPropertyIDList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_Mutate(t *testing.T) {
	type fields struct {
		List lists.List
	}
	type args struct {
		properties []properties.Property
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   lists.PropertyList
	}{
		{"+ve", fields{NewList(propertiesToListables([]properties.Property{baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1)))}...)...)}, args{[]properties.Property{baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(2)))}}, NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(2)))}...)},
		// {"+ve", fields{NewList(propertiesToListables([]properties.Property{baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1)))}...)...)}, args{[]properties.Property{baseProperties.NewMetaProperty(NewStringID("supply"), NewStringData("test"))}}, NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(NewStringID("supply"), NewStringData("test"))}...)}, //TODO: Should handle error if different property data is tried to mutate
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := propertyList{
				List: tt.fields.List,
			}
			if got := propertyList.Mutate(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mutate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_Remove(t *testing.T) {
	type fields struct {
		List lists.List
	}
	type args struct {
		properties []properties.Property
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   lists.PropertyList
	}{
		{"+ve", fields{NewList(propertiesToListables([]properties.Property{baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1))), baseProperties.NewMetaProperty(NewStringID("supply1"), NewDecData(sdkTypes.NewDec(2)))}...)...)}, args{[]properties.Property{baseProperties.NewMetaProperty(NewStringID("supply1"), NewDecData(sdkTypes.NewDec(2)))}}, NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1)))}...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := propertyList{
				List: tt.fields.List,
			}
			if got := propertyList.Remove(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_ScrubData(t *testing.T) {
	type fields struct {
		List lists.List
	}
	tests := []struct {
		name   string
		fields fields
		want   lists.PropertyList
	}{
		{"+ve", fields{NewList(propertiesToListables([]properties.Property{baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1)))}...)...)}, NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1))).ScrubData()}...)},
		{"+ve", fields{NewList(propertiesToListables([]properties.Property{baseProperties.NewMesaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1))), baseProperties.NewMetaProperty(NewStringID("ID"), NewStringData("Test"))}...)...)}, NewPropertyList([]properties.Property{baseProperties.NewMesaProperty(NewStringID("supply"), NewDecData(sdkTypes.NewDec(1))), baseProperties.NewMetaProperty(NewStringID("ID"), NewStringData("Test")).ScrubData()}...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := propertyList{
				List: tt.fields.List,
			}
			if got := propertyList.ScrubData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScrubData() = %v, want %v", got, tt.want)
			}
		})
	}
}

type decData struct {
	Value sdkTypes.Dec `json:"value"`
}

var _ data.DecData = (*decData)(nil)

func (decData decData) GetID() ids.DataID {
	return baseIDs.NewDataID(decData)
}
func (decData decData) Compare(listable traits.Listable) int {
	compareDecData, err := decDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	if decData.Value.GT(compareDecData.Value) {
		return 1
	} else if decData.Value.LT(compareDecData.Value) {
		return -1
	}

	return 0
}
func (decData decData) String() string {
	return decData.Value.String()
}
func (decData decData) Bytes() []byte {
	return decData.Value.Bytes()
}
func (decData decData) GetType() ids.StringID {
	return dataConstants.DecDataID
}
func (decData decData) ZeroValue() data.Data {
	return NewDecData(sdkTypes.ZeroDec())
}
func (decData decData) GenerateHashID() ids.HashID {
	if decData.Compare(decData.ZeroValue()) == 0 {
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(decData.Bytes())
}
func (decData decData) Get() sdkTypes.Dec {
	return decData.Value
}

func decDataFromInterface(listable traits.Listable) (decData, error) {
	switch value := listable.(type) {
	case decData:
		return value, nil
	default:
		return decData{}, constants.MetaDataError
	}
}

func DecDataPrototype() data.DecData {
	return decData{}.ZeroValue().(data.DecData)
}

func NewDecData(value sdkTypes.Dec) data.DecData {
	return decData{
		Value: value,
	}
}
