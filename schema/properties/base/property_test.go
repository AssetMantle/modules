// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/traits"
)

func createTestInputForProperty() (ids.ID, ids.PropertyID, data.Data, properties.Property) {
	testKey := base.NewID("ID")
	testData := NewStringData("Data")
	testPropertyID := base.NewPropertyID(testKey, testData.GetType())
	testProperty := NewProperty(testKey, testData)
	return testKey, testPropertyID, testData, testProperty
}

func TestNewEmptyPropertyFromID(t *testing.T) {
	_, testPropertyID, _, _ := createTestInputForProperty()
	type args struct {
		propertyID ids.PropertyID
	}
	tests := []struct {
		name string
		args args
		want properties.Property
	}{
		// TODO: Add test cases.
		{"+ve", args{testPropertyID}, property{ID: testPropertyID}},
		{"+ve with nil", args{}, property{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmptyPropertyFromID(tt.args.propertyID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmptyPropertyFromID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewProperty(t *testing.T) {
	testKey, testPropertyID, testData, _ := createTestInputForProperty()
	type args struct {
		key  ids.ID
		data data.Data
	}
	tests := []struct {
		name      string
		args      args
		want      properties.Property
		wantPanic bool
	}{
		// TODO: Add test cases.
		{"nil", args{}, property{}, true},
		{"+ve", args{testKey, testData}, property{testPropertyID, testData.GetID()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("NewProperty() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := NewProperty(tt.args.key, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPropertyWithDataID(t *testing.T) {
	_, testPropertyID, testData, _ := createTestInputForProperty()
	type args struct {
		propertyID ids.PropertyID
		dataID     ids.DataID
	}
	tests := []struct {
		name string
		args args
		want properties.Property
	}{
		// TODO: Add test cases.
		{"+ve with nil", args{}, property{}},
		{"+ve", args{testPropertyID, testData.GetID()}, property{testPropertyID, testData.GetID()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPropertyWithDataID(tt.args.propertyID, tt.args.dataID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPropertyWithDataID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyFromInterface(t *testing.T) {
	_, testPropertyID, testData, testProperty := createTestInputForProperty()
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name    string
		args    args
		want    property
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve with nil", args{}, property{}, true},
		{"+ve", args{testProperty}, property{testPropertyID, testData.GetID()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := propertyFromInterface(tt.args.listable)
			if (err != nil) != tt.wantErr {
				t.Errorf("propertyFromInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("propertyFromInterface() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_property_Compare(t *testing.T) {
	_, testPropertyID, testData, testProperty := createTestInputForProperty()
	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
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
		{"+ve compare with property with no Data", fields{testPropertyID, testData.GetID()}, args{property{ID: base.NewPropertyID(base.NewID("ID"), base.NewID("S"))}}, 0},
		{"+ve", fields{testPropertyID, testData.GetID()}, args{property{ID: base.NewPropertyID(base.NewID("ID"), base.NewID("S")), DataID: NewStringData("Data2").GetID()}}, 0},
		{"+ve", fields{testPropertyID, testData.GetID()}, args{testProperty}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			property := property{
				ID:     tt.fields.ID,
				DataID: tt.fields.DataID,
			}
			if got := property.Compare(tt.args.listable); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_property_GetDataID(t *testing.T) {
	_, testPropertyID, testData, _ := createTestInputForProperty()
	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.DataID
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{}, nil},
		{"+ve", fields{testPropertyID, testData.GetID()}, testData.GetID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			property := property{
				ID:     tt.fields.ID,
				DataID: tt.fields.DataID,
			}
			if got := property.GetDataID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDataID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_property_GetHash(t *testing.T) {
	_, testPropertyID, testData, _ := createTestInputForProperty()
	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		// TODO: Add test cases.
		{"+ve", fields{testPropertyID, testData.GetID()}, testData.GetID().GetHash()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			property := property{
				ID:     tt.fields.ID,
				DataID: tt.fields.DataID,
			}
			if got := property.GetHash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_property_GetID(t *testing.T) {
	_, testPropertyID, testData, _ := createTestInputForProperty()
	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.PropertyID
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{}, nil},
		{"+ve", fields{testPropertyID, testData.GetID()}, testPropertyID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			property := property{
				ID:     tt.fields.ID,
				DataID: tt.fields.DataID,
			}
			if got := property.GetID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_property_GetKey(t *testing.T) {
	testKey, testPropertyID, testData, _ := createTestInputForProperty()
	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		// TODO: Add test cases.
		{"+ve", fields{testPropertyID, testData.GetID()}, testKey},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			property := property{
				ID:     tt.fields.ID,
				DataID: tt.fields.DataID,
			}
			if got := property.GetKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_property_GetType(t *testing.T) {
	_, testPropertyID, testData, _ := createTestInputForProperty()
	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		// TODO: Add test cases.
		{"+ve", fields{testPropertyID, testData.GetID()}, base.NewID("S")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			property := property{
				ID:     tt.fields.ID,
				DataID: tt.fields.DataID,
			}
			if got := property.GetType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}
