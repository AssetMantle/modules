// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/data/base"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

// func TestNewDataID(t *testing.T) {
//	type args struct {
//		data data.Data
//	}
//	tests := []struct {
//		name string
//		args args
//		want ids.DataID
//	}{
//		// TODO: Add test cases.
//		{"+ve", args{base.NewBooleanData(true)}, dataID{NewStringID("B"), base.NewBooleanData(true).GenerateHashID()}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewDataID(tt.args.data); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewDataID() = %v, want %v", got, tt.want)
//			}
//		})
//	}
// }

func Test_dataIDFromInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want dataID
	}{
		// TODO: Add test cases.
		// {"+ve", args{dataID{NewStringID("B"), base.NewBooleanData(true).GenerateHashID()}}, dataID{NewStringID("B"), base.NewBooleanData(true).GenerateHashID()}},
		{"-ve", args{dataID{}}, dataID{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
		// TODO: Add test cases.
		// {"+ve", fields{NewStringID("B"), base.NewBooleanData(true).GenerateHashID()}, append(append([]byte{}, NewStringID("B").Bytes()...), base.NewBooleanData(true).GenerateHashID().Bytes()...)},
		// {"+ve", fields{NewStringID("B"), base.NewBooleanData(false).GenerateHashID()}, append(append([]byte{}, NewStringID("B").Bytes()...), base.NewBooleanData(false).GenerateHashID().Bytes()...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataID := dataID{
				Type:   tt.fields.Type,
				HashID: tt.fields.HashID,
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
		// TODO: Add test cases.
		// {"+ve", fields{NewStringID("B"), base.NewBooleanData(true).GenerateHashID()}, args{dataID{NewStringID("B"), base.NewBooleanData(true).GenerateHashID()}}, 0},
		// {"+ve", fields{NewStringID("B"), base.NewBooleanData(false).GenerateHashID()}, args{dataID{NewStringID("B"), base.NewBooleanData(true).GenerateHashID()}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataID := dataID{
				Type:   tt.fields.Type,
				HashID: tt.fields.HashID,
			}
			if got := dataID.Compare(tt.args.listable); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataID_GetHashID(t *testing.T) {
	type fields struct {
		Type   ids.StringID
		HashID ids.HashID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.HashID
	}{
		// TODO: Add test cases.
		{"+ve", fields{}, dataID{}.HashID},
		// {"+ve", fields{NewStringID("B"), base.NewBooleanData(true).GenerateHashID()}, base.NewBooleanData(true).GenerateHashID()},
		// {"+ve", fields{NewStringID("B"), base.NewBooleanData(false).GenerateHashID()}, base.NewBooleanData(false).GenerateHashID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataID := dataID{
				Type:   tt.fields.Type,
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
		Type   ids.StringID
		HashID ids.HashID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		// {"+ve", fields{NewStringID("B"), base.NewBooleanData(true).GenerateHashID()}, stringUtilities.JoinIDStrings(NewStringID("B").String(), base.NewBooleanData(true).GenerateHashID().String())},
		// {"+ve", fields{NewStringID("B"), base.NewBooleanData(false).GenerateHashID()}, stringUtilities.JoinIDStrings(NewStringID("B").String(), base.NewBooleanData(false).GenerateHashID().String())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataID := dataID{
				Type:   tt.fields.Type,
				HashID: tt.fields.HashID,
			}
			if got := dataID.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDataID(t *testing.T) {
	type args struct {
		data data.Data
	}
	tests := []struct {
		name string
		args args
		want ids.DataID
	}{
		// TODO: Add test cases.
		{"+ve", args{}, dataID{}},
		{"+ve", args{base.NewStringData("Data")}, dataID{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDataID(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDataID() = %v, want %v", got, tt.want)
			}
		})
	}
}
