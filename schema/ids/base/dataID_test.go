// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/data/utilities"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/traits"
)

// TODO: Test GetID for all Data types; If every data tests GetID() then GenerateID() is automatically tested
// func TestNewDataID(t *testing.T) {
//	type args struct {
//		data data.Data
//	}
//	tests := []struct {
//		name      string
//		args      args
//		want      ids.DataID
//		wantError bool
//	}{
//		{"-ve with nil", args{}, &DataID{}, true},
//		{"+ve", args{NewBooleanData(true)}, &DataID{NewStringID("B").(*StringID), NewBooleanData(true).GenerateHashID().(*HashID)}, false},
//		{"-ve with invalid data", args{nil}, &DataID{}, true},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			defer func() {
//				r := recover()
//
//				if (r != nil) != tt.wantError {
//					t.Errorf("GenerateDataID() error = %v wantError = %v", r, tt.wantError)
//				}
//			}()
//			if got := GenerateDataID(tt.args.data); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GenerateDataID() = %v, want %v", got, tt.want)
//			}
//		})
//	}
// }
func Test_dataIDFromInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name      string
		args      args
		want      *DataID
		wantError bool
	}{
		{"+ve", args{&DataID{NewStringID("B").(*StringID), NewBooleanData(true).GenerateHashID().(*HashID)}}, &DataID{NewStringID("B").(*StringID), NewBooleanData(true).GenerateHashID().(*HashID)}, false},
		{"-ve", args{&DataID{}}, &DataID{}, false},
		{"-ve", args{nil}, &DataID{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantError {
					t.Errorf("MetadataFromInterface() error = %v, wantError %v", r, tt.wantError)
				}
			}()
			got := dataIDFromInterface(tt.args.i)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dataIDFromInterface() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataID_Bytes(t *testing.T) {
	type fields struct {
		Type   ids.StringID
		HashID ids.HashID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve", fields{NewStringID("B"), NewBooleanData(true).GenerateHashID()}, append(append([]byte{}, NewStringID("B").Bytes()...), NewBooleanData(true).GenerateHashID().Bytes()...)},
		{"+ve", fields{NewStringID("B"), NewBooleanData(false).GenerateHashID()}, append(append([]byte{}, NewStringID("B").Bytes()...), NewBooleanData(false).GenerateHashID().Bytes()...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataID := &DataID{
				TypeID: tt.fields.Type.(*StringID),
				HashID: tt.fields.HashID.(*HashID),
			}
			if got := dataID.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataID_Compare(t *testing.T) {
	type fields struct {
		Type   ids.StringID
		HashID ids.HashID
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
		{"+ve", fields{NewStringID("B"), NewBooleanData(true).GenerateHashID()}, args{&DataID{NewStringID("B").(*StringID), NewBooleanData(true).GenerateHashID().(*HashID)}}, 0},
		{"+ve", fields{NewStringID("B"), NewBooleanData(false).GenerateHashID()}, args{&DataID{NewStringID("B").(*StringID), NewBooleanData(true).GenerateHashID().(*HashID)}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataID := &DataID{
				TypeID: tt.fields.Type.(*StringID),
				HashID: tt.fields.HashID.(*HashID),
			}
			if got := dataID.Compare(tt.args.listable); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataID_GetHashID(t *testing.T) {
	type fields struct {
		Type   *StringID
		HashID *HashID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.HashID
	}{
		{"+ve", fields{}, (&DataID{}).HashID},
		{"+ve", fields{NewStringID("B").(*StringID), NewBooleanData(true).GenerateHashID().(*HashID)}, NewBooleanData(true).GenerateHashID()},
		{"+ve", fields{NewStringID("B").(*StringID), NewBooleanData(false).GenerateHashID().(*HashID)}, NewBooleanData(false).GenerateHashID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataID := &DataID{
				TypeID: tt.fields.Type,
				HashID: tt.fields.HashID,
			}
			if got := dataID.GetHashID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHashID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataID_String(t *testing.T) {
	type fields struct {
		Type   *StringID
		HashID *HashID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{NewStringID("B").(*StringID), NewBooleanData(true).GenerateHashID().(*HashID)}, stringUtilities.JoinIDStrings(NewStringID("B").AsString(), NewBooleanData(true).GenerateHashID().AsString())},
		{"+ve", fields{NewStringID("B").(*StringID), NewBooleanData(false).GenerateHashID().(*HashID)}, stringUtilities.JoinIDStrings(NewStringID("B").AsString(), NewBooleanData(false).GenerateHashID().AsString())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataID := &DataID{
				TypeID: tt.fields.Type,
				HashID: tt.fields.HashID,
			}
			if got := dataID.AsString(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

// mocks for decData
type booleanData struct {
	Value bool `json:"value"`
}

func (booleanData booleanData) GetWidth() int {
	// TODO implement me
	panic("implement me")
}

func (booleanData booleanData) Unmarshal(bytes []byte) error {
	// TODO implement me
	panic("implement me")
}

func (booleanData booleanData) MarshalTo(bytes []byte) (int, error) {
	// TODO implement me
	panic("implement me")
}

func (booleanData booleanData) ToAnyData() data.AnyData {
	// TODO implement me
	panic("implement me")
}

var _ data.BooleanData = (*booleanData)(nil)

func (booleanData booleanData) GetID() ids.DataID {
	return GenerateDataID(booleanData)
}
func (booleanData booleanData) GetBondWeight() int64 {
	return int64(1)
}
func (booleanData booleanData) Compare(listable traits.Listable) int {
	compareBooleanData, err := booleanDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	if booleanData.Value == compareBooleanData.Value {
		return 0
	} else if booleanData.Value == true {
		return 1
	}

	return -1
}
func (booleanData booleanData) AsString() string {
	return utilities.JoinDataTypeAndValueStrings(booleanData.GetType().AsString(), strconv.FormatBool(booleanData.Value))
}
func (booleanData booleanData) FromString(dataString string) (data.Data, error) {
	dataTypeString, dataString := splitDataTypeAndValueStrings(dataTypeAndValueString)

	if dataTypeString != stringData.GetType().AsString() {
		return PrototypeStringData(), errorConstants.IncorrectFormat
	}

	if dataString == "" {
		return BooleanDataPrototype(), nil
	}

	Bool, err := strconv.ParseBool(dataString)
	if err != nil {
		return BooleanDataPrototype(), err
	}

	return NewBooleanData(Bool), nil
}
func (booleanData booleanData) Bytes() []byte {
	if booleanData.Get() {
		return []byte{0x1}
	}
	return []byte{0x0}
}
func (booleanData booleanData) GetType() ids.StringID {
	return NewStringID("B")
}
func (booleanData booleanData) ZeroValue() data.Data {
	return NewBooleanData(false)
}
func (booleanData booleanData) GenerateHashID() ids.HashID {
	if booleanData.Compare(booleanData.ZeroValue()) == 0 {
		return GenerateHashID()
	}

	return GenerateHashID(booleanData.Bytes())
}
func (booleanData booleanData) Get() bool {
	return booleanData.Value
}

func booleanDataFromInterface(listable traits.Listable) (booleanData, error) {
	switch value := listable.(type) {
	case booleanData:
		return value, nil
	default:
		return booleanData{}, errorConstants.MetaDataError
	}
}

func BooleanDataPrototype() data.BooleanData {
	return booleanData{}.ZeroValue().(data.BooleanData)
}

func NewBooleanData(value bool) data.BooleanData {
	return booleanData{
		Value: value,
	}
}

func TestReadDataID(t *testing.T) {
	type args struct {
		dataIDString string
	}
	tests := []struct {
		name    string
		args    args
		want    ids.DataID
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", args{stringUtilities.JoinIDStrings(NewStringID("B").AsString(), NewBooleanData(true).GenerateHashID().AsString())}, GenerateDataID(NewBooleanData(true)), false},
		{"+ve with empty string", args{""}, PrototypeDataID(), false},
		{"+ve with nil", args{}, PrototypeDataID(), false},
		{"-ve", args{stringUtilities.JoinIDStrings(NewStringID("j").AsString(), "0")}, &DataID{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadDataID(tt.args.dataIDString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDataID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadDataID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
