// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

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
		{"+ve", args{nil}, dataID{}}, //TODO: report for panic
		{"+ve", args{NewBooleanData(true)}, dataID{NewID("B"), NewID(strconv.FormatBool(true))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDataID(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDataID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataIDFromInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    dataID
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", args{dataID{NewID("B"), NewID(strconv.FormatBool(true))}}, dataID{NewID("B"), NewID(strconv.FormatBool(true))}, false},
		{"-ve", args{id{}}, dataID{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dataIDFromInterface(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("dataIDFromInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dataIDFromInterface() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataID_Bytes(t *testing.T) {
	type fields struct {
		Type ids.ID
		Hash ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{}, []byte{}}, //TODO: report for panic
		{"+ve", fields{NewID("B"), NewID(strconv.FormatBool(true))}, append(append([]byte{}, NewID("B").Bytes()...), NewID(strconv.FormatBool(true)).Bytes()...)},
		{"+ve", fields{NewID("B"), NewID(strconv.FormatBool(false))}, append(append([]byte{}, NewID("B").Bytes()...), NewID(strconv.FormatBool(false)).Bytes()...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataID := dataID{
				Type: tt.fields.Type,
				Hash: tt.fields.Hash,
			}
			if got := dataID.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataID_Compare(t *testing.T) {
	type fields struct {
		Type ids.ID
		Hash ids.ID
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
		{"+ve with nil", fields{}, args{dataID{}}, 0}, //TODO: Panics for nil
		{"+ve", fields{NewID("B"), NewID(strconv.FormatBool(true))}, args{dataID{NewID("B"), NewID(strconv.FormatBool(true))}}, 0},
		{"+ve", fields{NewID("B"), NewID(strconv.FormatBool(false))}, args{dataID{NewID("B"), NewID(strconv.FormatBool(true))}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataID := dataID{
				Type: tt.fields.Type,
				Hash: tt.fields.Hash,
			}
			if got := dataID.Compare(tt.args.listable); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataID_GetHash(t *testing.T) {
	type fields struct {
		Type ids.ID
		Hash ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		// TODO: Add test cases.
		{"+ve", fields{}, dataID{}.Hash},
		{"+ve", fields{NewID("B"), NewID(strconv.FormatBool(true))}, NewID(strconv.FormatBool(true))},
		{"+ve", fields{NewID("B"), NewID(strconv.FormatBool(false))}, NewID(strconv.FormatBool(false))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataID := dataID{
				Type: tt.fields.Type,
				Hash: tt.fields.Hash,
			}
			if got := dataID.GetHash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataID_String(t *testing.T) {
	type fields struct {
		Type ids.ID
		Hash ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{NewID("B"), NewID(strconv.FormatBool(true))}, strings.Join(append(append([]string{}, NewID("B").String()), NewID(strconv.FormatBool(true)).String()), constants.FirstOrderCompositeIDSeparator)},
		{"+ve", fields{NewID("B"), NewID(strconv.FormatBool(false))}, strings.Join(append(append([]string{}, NewID("B").String()), NewID(strconv.FormatBool(false)).String()), constants.FirstOrderCompositeIDSeparator)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataID := dataID{
				Type: tt.fields.Type,
				Hash: tt.fields.Hash,
			}
			if got := dataID.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
