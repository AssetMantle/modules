// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-IDentifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/traits"
)

func ValidatedData(value data.Data) *baseData.AnyData {
	if value == nil {
		return nil
	}
	return value.ToAnyData().(*baseData.AnyData)
}

func createTestInput() (ids.StringID, ids.PropertyID, data.Data, properties.MetaProperty) {
	testKey := baseIDs.NewStringID("ID")
	testData := baseData.NewStringData("Data")
	testPropertyID := baseIDs.NewPropertyID(testKey, testData.GetTypeID())
	testMetaProperty := NewMetaProperty(testKey, testData)
	return testKey, testPropertyID, testData, testMetaProperty
}

func TestNewEmptyMetaPropertyFromID(t *testing.T) {
	_, testPropertyID, _, _ := createTestInput()
	type args struct {
		propertyID ids.PropertyID
	}
	tests := []struct {
		name    string
		args    args
		want    properties.MetaProperty
		wantErr bool
	}{
		{"+ve", args{testPropertyID}, &MetaProperty{ID: testPropertyID.(*baseIDs.PropertyID)}, false},
		{"panic with nil", args{}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
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
		name        string
		args        args
		want        properties.MetaProperty
		shouldPanic bool
	}{
		{"+ve", args{testKey, testData}, &MetaProperty{testPropertyID.(*baseIDs.PropertyID), testData.ToAnyData().(*baseData.AnyData)}, false},
		{"+ve", args{nil, nil}, &MetaProperty{nil, nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				assert.Panics(t, func() { NewMetaProperty(tt.args.key, tt.args.data) }, "The code did not panic, but it should panic")
			} else if got := NewMetaProperty(tt.args.key, tt.args.data); !reflect.DeepEqual(got, tt.want) {
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
		{"+ve", fields{testPropertyID, testData}, args{testMetaProperty}, 0},
		{"+ve compare with metaProperty with no Data", fields{testPropertyID, testData}, args{&MetaProperty{ID: baseIDs.NewPropertyID(baseIDs.NewStringID("ID"), baseIDs.NewStringID("S")).(*baseIDs.PropertyID)}}, 0},
		{"+ve", fields{testPropertyID, testData}, args{&MetaProperty{baseIDs.NewPropertyID(baseIDs.NewStringID("ID"), baseIDs.NewStringID("S")).(*baseIDs.PropertyID), baseData.NewStringData("Data2").ToAnyData().(*baseData.AnyData)}}, 0}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metaProperty := &MetaProperty{
				ID:   ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				Data: ValidatedData(tt.fields.Data),
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
		{"+ve", fields{testPropertyID, testData}, testData.ToAnyData()},
		{"+ve with nil", fields{}, (*baseData.AnyData)(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metaProperty := &MetaProperty{
				ID:   ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				Data: ValidatedData(tt.fields.Data),
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
		{"+ve", fields{testPropertyID, testData}, testMetaProperty.GetData().GetID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metaProperty := &MetaProperty{
				ID:   ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				Data: ValidatedData(tt.fields.Data),
			}
			if got := metaProperty.GetDataID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDataID() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		{"+ve", fields{testPropertyID, testData}, testPropertyID},
		{"+ve", fields{}, (&MetaProperty{}).ID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metaProperty := &MetaProperty{
				ID:   ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				Data: ValidatedData(tt.fields.Data),
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
		{"+ve", fields{testPropertyID, testData}, testKey},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metaProperty := &MetaProperty{
				ID:   ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				Data: ValidatedData(tt.fields.Data),
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
		{"+ve", fields{testPropertyID, testData}, testData.GetTypeID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metaProperty := &MetaProperty{
				ID:   ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				Data: ValidatedData(tt.fields.Data),
			}
			if got := metaProperty.GetDataTypeID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTypeID() = %v, want %v", got, tt.want)
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
		{"+ve", fields{testPropertyID, testData}, &MesaProperty{ID: testPropertyID.(*baseIDs.PropertyID), DataID: testData.GetID().(*baseIDs.DataID)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metaProperty := &MetaProperty{
				ID:   ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				Data: ValidatedData(tt.fields.Data),
			}
			if got := metaProperty.ScrubData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScrubData() = %v, want %v", got, tt.want)
			}
		})
	}
}
