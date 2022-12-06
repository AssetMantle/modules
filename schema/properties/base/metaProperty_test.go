// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/traits"
)

func createTestInput() (ids.StringID, ids.PropertyID, data.Data, properties.MetaProperty) {
	testKey := base.NewStringID("PropertyID")
	testData := baseData.NewStringData("Data")
	testPropertyID := base.GeneratePropertyID(testKey, testData.GetType())
	testMetaProperty := NewMetaProperty(testKey, testData)
	return testKey, testPropertyID, testData, testMetaProperty
}

func TestNewEmptyMetaPropertyFromID(t *testing.T) {
	_, testPropertyID, _, _ := createTestInput()
	type args struct {
		propertyID ids.PropertyID
	}
	tests := []struct {
		name string
		args args
		want properties.MetaProperty
	}{
		// TODO: Add test cases.
		{"+ve", args{testPropertyID}, metaProperty{PropertyID: testPropertyID}},
		{"+ve with nil", args{}, metaProperty{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmptyMetaPropertyFromID(tt.args.propertyID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmptyMetaPropertyFromID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMetaProperty(t *testing.T) {
	testKey, testPropertyID, testData, _ := createTestInput()
	type args struct {
		key  ids.StringID
		data data.Data
	}
	tests := []struct {
		name string
		args args
		want properties.MetaProperty
	}{
		// TODO: Add test cases.
		{"+ve", args{testKey, testData}, metaProperty{testPropertyID, testData}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMetaProperty(tt.args.key, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMetaProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metaProperty_Compare(t *testing.T) {
	_, testPropertyID, testData, testMetaProperty := createTestInput()
	type fields struct {
		ID   ids.PropertyID
		Data data.Data
	}
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{"+ve", fields{testPropertyID, testData}, args{testMetaProperty}, 0},
		{"+ve compare with metaProperty with no Data", fields{testPropertyID, testData}, args{metaProperty{PropertyID: base.GeneratePropertyID(base.NewStringID("PropertyID"), base.NewStringID("S"))}}, 0},
		{"+ve", fields{testPropertyID, testData}, args{metaProperty{PropertyID: base.GeneratePropertyID(base.NewStringID("PropertyID"), base.NewStringID("S")), Data: baseData.NewStringData("Data2")}}, 0}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metaProperty := metaProperty{
				PropertyID: tt.fields.ID,
				Data:       tt.fields.Data,
			}
			if got := metaProperty.Compare(tt.args.listable); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metaProperty_GetData(t *testing.T) {
	_, testPropertyID, testData, _ := createTestInput()
	type fields struct {
		ID   ids.PropertyID
		Data data.Data
	}
	tests := []struct {
		name   string
		fields fields
		want   data.Data
	}{
		// TODO: Add test cases.
		{"+ve", fields{testPropertyID, testData}, testData},
		{"+ve with nil", fields{}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metaProperty := metaProperty{
				PropertyID: tt.fields.ID,
				Data:       tt.fields.Data,
			}
			if got := metaProperty.GetData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metaProperty_GetDataID(t *testing.T) {
	_, testPropertyID, testData, testMetaProperty := createTestInput()
	type fields struct {
		ID   ids.PropertyID
		Data data.Data
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.DataID
	}{
		// TODO: Add test cases.
		{"+ve", fields{testPropertyID, testData}, testMetaProperty.GetData().GetID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metaProperty := metaProperty{
				PropertyID: tt.fields.ID,
				Data:       tt.fields.Data,
			}
			if got := metaProperty.GetDataID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDataID() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_metaProperty_GetHash(t *testing.T) {
//	_, testPropertyID, testData, _ := createTestInput()
//	type fields struct {
//		PropertyID   ids.PropertyID
//		Data data.Data
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   ids.PropertyID
//	}{
//		// TODO: Add test cases.
//		{"+ve", fields{testPropertyID, testData}, testData.GenerateHash()},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			metaProperty := metaProperty{
//				PropertyID:   tt.fields.PropertyID,
//				Data: tt.fields.Data,
//			}
//			if got := metaProperty.GetHash(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetHash() = %v, want %v", got, tt.want)
//			}
//		})
//	}
// }

func Test_metaProperty_GetID(t *testing.T) {
	_, testPropertyID, testData, _ := createTestInput()
	type fields struct {
		ID   ids.PropertyID
		Data data.Data
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.PropertyID
	}{
		// TODO: Add test cases.
		{"+ve", fields{testPropertyID, testData}, testPropertyID},
		{"+ve", fields{}, metaProperty{}.PropertyID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metaProperty := metaProperty{
				PropertyID: tt.fields.ID,
				Data:       tt.fields.Data,
			}
			if got := metaProperty.GetID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metaProperty_GetKey(t *testing.T) {
	testKey, testPropertyID, testData, _ := createTestInput()
	type fields struct {
		ID   ids.PropertyID
		Data data.Data
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		// TODO: Add test cases.
		{"+ve", fields{testPropertyID, testData}, testKey},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metaProperty := metaProperty{
				PropertyID: tt.fields.ID,
				Data:       tt.fields.Data,
			}
			if got := metaProperty.GetKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metaProperty_GetType(t *testing.T) {
	_, testPropertyID, testData, _ := createTestInput()
	type fields struct {
		ID   ids.PropertyID
		Data data.Data
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		// TODO: Add test cases.
		{"+ve", fields{testPropertyID, testData}, testData.GetType()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metaProperty := metaProperty{
				PropertyID: tt.fields.ID,
				Data:       tt.fields.Data,
			}
			if got := metaProperty.GetType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metaProperty_RemoveData(t *testing.T) {
	_, testPropertyID, testData, _ := createTestInput()
	type fields struct {
		ID   ids.PropertyID
		Data data.Data
	}
	tests := []struct {
		name   string
		fields fields
		want   properties.Property
	}{
		// TODO: Add test cases.
		{"+ve", fields{testPropertyID, testData}, mesaProperty{PropertyID: testPropertyID, DataID: testData.GetID()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metaProperty := metaProperty{
				PropertyID: tt.fields.ID,
				Data:       tt.fields.Data,
			}
			if got := metaProperty.ScrubData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScrubData() = %v, want %v", got, tt.want)
			}
		})
	}
}
