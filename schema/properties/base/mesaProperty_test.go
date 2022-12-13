// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"reflect"
	"strings"
	"testing"

	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/traits"
)

func createTestInputForMesaProperty() (ids.StringID, ids.PropertyID, data.Data, properties.Property) {
	testKey := base.NewStringID("ID")
	testData := NewStringData("Data")
	testPropertyID := base.NewPropertyID(testKey, testData.GetType())
	testProperty := NewMesaProperty(testKey, testData)
	return testKey, testPropertyID, testData, testProperty
}

func TestNewEmptyMesaPropertyFromID(t *testing.T) {
	_, testPropertyID, _, _ := createTestInputForMesaProperty()

	type args struct {
		propertyID ids.PropertyID
	}
	tests := []struct {
		name string
		args args
		want properties.Property
	}{
		{"+ve", args{testPropertyID}, mesaProperty{ID: testPropertyID}},
		{"+ve with nil", args{}, mesaProperty{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
		{"nil", args{}, mesaProperty{}, true},
		{"+ve", args{testKey, testData}, mesaProperty{testPropertyID, testData.GetID()}, false},
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
		want    mesaProperty
		wantErr bool
	}{
		{"+ve with nil", args{}, mesaProperty{}, true},
		{"+ve", args{testMesaProperty}, mesaProperty{testMesaPropertyID, testData.GetID()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := propertyFromInterface(tt.args.listable)
			if (err != nil) != tt.wantErr {
				t.Errorf("mesaPropertyFromInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mesaPropertyFromInterface() got = %v, want %v", got, tt.want)
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
		name   string
		fields fields
		args   args
		want   int
	}{
		{"+ve compare with property with no Data", fields{testMesaPropertyID, testData.GetID()}, args{mesaProperty{ID: base.NewPropertyID(base.NewStringID("ID"), base.NewStringID("S"))}}, 0},
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, args{mesaProperty{ID: base.NewPropertyID(base.NewStringID("ID"), base.NewStringID("S")), DataID: NewStringData("Data2").GetID()}}, 0},
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, args{testMesaProperty}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := mesaProperty{
				ID:     tt.fields.ID,
				DataID: tt.fields.DataID,
			}
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
		{"+ve with nil", fields{}, nil},
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, testData.GetID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := mesaProperty{
				ID:     tt.fields.ID,
				DataID: tt.fields.DataID,
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
		name   string
		fields fields
		want   ids.ID
	}{
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, testData.GetID().GetHashID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := mesaProperty{
				ID:     tt.fields.ID,
				DataID: tt.fields.DataID,
			}
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
		{"+ve with nil", fields{}, nil},
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, testMesaPropertyID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := mesaProperty{
				ID:     tt.fields.ID,
				DataID: tt.fields.DataID,
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
		name   string
		fields fields
		want   ids.StringID
	}{
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, testKey},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := mesaProperty{
				ID:     tt.fields.ID,
				DataID: tt.fields.DataID,
			}
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
		name   string
		fields fields
		want   ids.StringID
	}{
		{"+ve", fields{testMesaPropertyID, testData.GetID()}, base.NewStringID("S")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mesaProperty := mesaProperty{
				ID:     tt.fields.ID,
				DataID: tt.fields.DataID,
			}
			if got := mesaProperty.GetType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
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
			mesaProperty := mesaProperty{
				ID:     tt.fields.ID,
				DataID: tt.fields.DataID,
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
			mesaProperty := mesaProperty{
				ID:     tt.fields.ID,
				DataID: tt.fields.DataID,
			}
			if got := mesaProperty.IsMeta(); got != tt.want {
				t.Errorf("IsMeta() = %v, want %v", got, tt.want)
			}
		})
	}
}

// MOCKS
type stringData struct {
	Value string `json:"value"`
}

var _ data.StringData = (*stringData)(nil)

func (stringData stringData) GetID() ids.DataID {
	return base.NewDataID(stringData)
}
func (stringData stringData) Compare(listable traits.Listable) int {
	compareStringData, err := stringDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return strings.Compare(stringData.Value, compareStringData.Value)
}
func (stringData stringData) String() string {
	return stringData.Value
}
func (stringData stringData) Bytes() []byte {
	return []byte(stringData.Value)
}
func (stringData stringData) GetType() ids.StringID {
	return dataConstants.StringDataID
}
func (stringData stringData) ZeroValue() data.Data {
	return NewStringData("")
}
func (stringData stringData) GenerateHashID() ids.HashID {
	return base.GenerateHashID(stringData.Bytes())
}
func (stringData stringData) Get() string {
	return stringData.Value
}

func stringDataFromInterface(listable traits.Listable) (stringData, error) {
	switch value := listable.(type) {
	case stringData:
		return value, nil
	default:
		return stringData{}, constants.MetaDataError
	}
}

func NewStringData(value string) data.StringData {
	return stringData{
		Value: value,
	}
}
