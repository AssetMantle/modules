// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

func TestNewStringID(t *testing.T) {
	type args struct {
		stringIDString string
	}
	tests := []struct {
		name string
		args args
		want ids.StringID
	}{

		{"+ve", args{"ID"}, NewStringID("ID")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStringID(tt.args.stringIDString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStringID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringIDFromInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    stringID
		wantErr bool
	}{

		{"+ve", args{NewStringID("ID")}, stringID{IDString: "ID"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringIDFromInterface(tt.args.i)
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("stringIDFromInterface() error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringIDFromInterface() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringID_Bytes(t *testing.T) {
	type fields struct {
		IDString string
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{

		{"+ve", fields{"ID"}, []byte("ID")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringID := stringID{
				IDString: tt.fields.IDString,
			}
			if got := stringID.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringID_Compare(t *testing.T) {
	type fields struct {
		IDString string
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

		{"+ve", fields{"ID"}, args{NewStringID("ID")}, 0},
		// TODO: It Should fail
		{"-ve", fields{"ID"}, args{NewStringID("ID2")}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringID := stringID{
				IDString: tt.fields.IDString,
			}
			if got := stringID.Compare(tt.args.listable); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringID_String(t *testing.T) {
	type fields struct {
		IDString string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{

		{"+ve", fields{"ID"}, "ID"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringID := stringID{
				IDString: tt.fields.IDString,
			}
			if got := stringID.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
