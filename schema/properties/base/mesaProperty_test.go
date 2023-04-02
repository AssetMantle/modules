// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-IDentifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/traits"
)

func ValidatedID[V *base.PropertyID | *base.DataID](value any) V {
	if value == nil {
		return nil
	}
	return value.(V)
}

func createTestInputForMesaProperty() (ids.StringID, ids.PropertyID, data.Data, properties.Property) {
	testKey := baseIDs.NewStringID("ID")
	testData := baseData.NewStringData("Data")
	testPropertyID := baseIDs.NewPropertyID(testKey, testData.GetTypeID())
	testProperty := NewMesaProperty(testKey, testData)
	return testKey, testPropertyID, testData, testProperty
}

func TestNewEmptyMesaPropertyFromID(t *testing.T) {
	_, testPropertyID, _, _ := createTestInputForMesaProperty()

	type args struct {
		propertyID ids.PropertyID
	}
	tests := []struct {
		name    string
		args    args
		want    properties.Property
		wantErr bool
	}{
		{"+ve", args{testPropertyID}, &MesaProperty{ID: testPropertyID.(*baseIDs.PropertyID)}, false},
		{"panic with nil", args{}, &MesaProperty{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			if got := NewEmptyMesaPropertyFromID(tt.args.propertyID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmptyMesaPropertyFromID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMesaProperty(t *testing.T) {
	testKey, testPropertyID, testData, _ := createTestInputForMesaProperty()
	type args struct {
		key  ids.StringID
		data data.Data
	}
	tests := []struct {
		name      string
		args      args
		want      properties.Property
		wantPanic bool
	}{
		{"nil", args{}, &MesaProperty{}, true},
		{"+ve", args{testKey, testData}, &MesaProperty{testPropertyID.(*baseIDs.PropertyID), testData.GetID().(*baseIDs.DataID)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("NewMesaProperty() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := NewMesaProperty(tt.args.key, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMesaProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mesaPropertyFromInterface(t *testing.T) {
	_, testMesaPropertyID, testData, testMesaProperty := createTestInputForMesaProperty()
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name    string
		args    args
		want    *MesaProperty
		wantErr bool
	}{
		{"+ve with nil", args{}, &MesaProperty{}, true},
		{"+ve", args{testMesaProperty}, &MesaProperty{testMesaPropertyID.(*baseIDs.PropertyID), testData.GetID().(*baseIDs.DataID)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := propertyFromInterface(tt.args.listable)
			if tt.wantErr {
				assert.ErrorIs(t, err, errorConstants.MetaDataError)
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("mesaPropertyFromInterface() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_mesaProperty_Compare(t *testing.T) {
	_, testMesaPropertyID, testData, testMesaProperty := createTestInputForMesaProperty()
	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"+ve compare with property with no Data", fields{testMesaPropertyID, testData.GetID()}, args{&MesaProperty{ID: baseIDs.NewPropertyID(baseIDs.NewStringID("ID"), baseIDs.NewStringID("S")).(*baseIDs.PropertyID)}}, 0, false},
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, args{&MesaProperty{ID: baseIDs.NewPropertyID(baseIDs.NewStringID("ID"), baseIDs.NewStringID("S")).(*baseIDs.PropertyID), DataID: baseData.NewStringData("Data2").GetID().(*baseIDs.DataID)}}, 0, false},
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, args{testMesaProperty}, 0, false},
		{"+ve nil dataID", fields{testMesaPropertyID, nil}, args{testMesaProperty}, 0, false},
		{"panic nil propertyID", fields{nil, testData.GetID()}, args{testMesaProperty}, 0, true},
		{"panic all nil", fields{}, args{testMesaProperty}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := &MesaProperty{
				ID:     ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				DataID: ValidatedID[*baseIDs.DataID](tt.fields.DataID),
			}
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			if got := mesaProperty.Compare(tt.args.listable); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mesaProperty_GetDataID(t *testing.T) {
	_, testMesaPropertyID, testData, _ := createTestInputForMesaProperty()
	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.DataID
	}{
		{"+ve with nil", fields{}, (*baseIDs.DataID)(nil)},
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, testData.GetID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := &MesaProperty{
				ID:     ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				DataID: ValidatedID[*baseIDs.DataID](tt.fields.DataID),
			}
			if got := mesaProperty.GetDataID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDataID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mesaProperty_GetHash(t *testing.T) {
	_, testMesaPropertyID, testData, _ := createTestInputForMesaProperty()

	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	tests := []struct {
		name    string
		fields  fields
		want    ids.ID
		wantErr bool
	}{
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, testData.GetID().GetHashID(), false},
		{"+ve nil PropertyID", fields{nil, testData.GetID()}, testData.GetID().GetHashID(), false},
		{"panic nil DataID", fields{testMesaPropertyID, nil}, testData.GetID().GetHashID(), true},
		{"panic all nil", fields{nil, nil}, testData.GetID().GetHashID(), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := &MesaProperty{
				ID:     ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				DataID: ValidatedID[*baseIDs.DataID](tt.fields.DataID),
			}
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			if got := mesaProperty.GetHash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mesaProperty_GetID(t *testing.T) {
	_, testMesaPropertyID, testData, _ := createTestInputForMesaProperty()
	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.PropertyID
	}{
		{"+ve with nil", fields{}, (*baseIDs.PropertyID)(nil)},
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, testMesaPropertyID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := &MesaProperty{
				ID:     ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				DataID: ValidatedID[*baseIDs.DataID](tt.fields.DataID),
			}
			if got := mesaProperty.GetID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mesaProperty_GetKey(t *testing.T) {
	testKey, testMesaPropertyID, testData, _ := createTestInputForMesaProperty()

	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	tests := []struct {
		name    string
		fields  fields
		want    ids.StringID
		wantErr bool
	}{
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, testKey, false},
		{"+ve nil dataID", fields{testMesaPropertyID, nil}, testKey, false},
		{"panic nil propertyID", fields{nil, testData.GetID()}, testKey, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := &MesaProperty{
				ID:     ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				DataID: ValidatedID[*baseIDs.DataID](tt.fields.DataID),
			}
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			if got := mesaProperty.GetKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mesaProperty_GetType(t *testing.T) {
	_, testMesaPropertyID, testData, _ := createTestInputForMesaProperty()
	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	tests := []struct {
		name    string
		fields  fields
		want    ids.StringID
		wantErr bool
	}{
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, baseIDs.NewStringID("S"), false},
		{"+ve nil dataID", fields{testMesaPropertyID, nil}, baseIDs.NewStringID("S"), false},
		{"panic nil PropertyID", fields{nil, testData.GetID()}, baseIDs.NewStringID("S"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := &MesaProperty{
				ID:     ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				DataID: ValidatedID[*baseIDs.DataID](tt.fields.DataID),
			}
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			if got := mesaProperty.GetDataTypeID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTypeID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mesaProperty_IsMesa(t *testing.T) {
	_, testMesaPropertyID, testData, _ := createTestInputForMesaProperty()
	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := &MesaProperty{
				ID:     ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				DataID: ValidatedID[*baseIDs.DataID](tt.fields.DataID),
			}
			if got := mesaProperty.IsMesa(); got != tt.want {
				t.Errorf("IsMesa() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mesaProperty_IsMeta(t *testing.T) {
	_, testMesaPropertyID, testData, _ := createTestInputForMesaProperty()
	type fields struct {
		ID     ids.PropertyID
		DataID ids.DataID
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := &MesaProperty{
				ID:     ValidatedID[*baseIDs.PropertyID](tt.fields.ID),
				DataID: ValidatedID[*baseIDs.DataID](tt.fields.DataID),
			}
			if got := mesaProperty.IsMeta(); got != tt.want {
				t.Errorf("IsMeta() = %v, want %v", got, tt.want)
			}
		})
	}
}
