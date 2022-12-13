// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"strings"
	"testing"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

func createTestInputForPropertyID() (ids.StringID, ids.StringID, ids.PropertyID) {
	testKey := NewStringID("ID")
	testType := NewStringID("ID2")
	testPropertyID := NewPropertyID(testKey, testType)
	return testKey, testType, testPropertyID
}

func TestNewPropertyID(t *testing.T) {
	type args struct {
		key  ids.StringID
		Type ids.StringID
	}
	tests := []struct {
		name string
		args args
		want ids.PropertyID
	}{

		{"+ve", args{NewStringID("ID"), NewStringID("ID2")}, propertyID{NewStringID("ID"), NewStringID("ID2")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPropertyID(tt.args.key, tt.args.Type); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPropertyID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyIDFromInterface(t *testing.T) {
	testKey, testType, testPropertyID := createTestInputForPropertyID()
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name    string
		args    args
		want    propertyID
		wantErr bool
	}{

		{"+ve", args{testPropertyID}, propertyID{testKey, testType}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := propertyIDFromInterface(tt.args.listable)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("propertyIDFromInterface() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyID_Bytes(t *testing.T) {
	testKey, testType, _ := createTestInputForPropertyID()
	type fields struct {
		Key  ids.StringID
		Type ids.StringID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{

		{"+ve", fields{testKey, testType}, append([]byte(testKey.String()), []byte(testType.String())...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyID := propertyID{
				Key:  tt.fields.Key,
				Type: tt.fields.Type,
			}
			if got := propertyID.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyID_Compare(t *testing.T) {
	testKey, testType, testPropertyID := createTestInputForPropertyID()
	type fields struct {
		Key  ids.StringID
		Type ids.StringID
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

		{"+ve", fields{testKey, testType}, args{testPropertyID}, 0},
		{"-ve", fields{testType, testType}, args{testPropertyID}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyID := propertyID{
				Key:  tt.fields.Key,
				Type: tt.fields.Type,
			}
			if got := propertyID.Compare(tt.args.listable); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyID_GetKey(t *testing.T) {
	testKey, testType, _ := createTestInputForPropertyID()
	type fields struct {
		Key  ids.StringID
		Type ids.StringID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.StringID
	}{

		{"+ve", fields{testKey, testType}, testKey},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyID := propertyID{
				Key:  tt.fields.Key,
				Type: tt.fields.Type,
			}
			if got := propertyID.GetKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyID_GetType(t *testing.T) {
	testKey, testType, _ := createTestInputForPropertyID()
	type fields struct {
		Key  ids.StringID
		Type ids.StringID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.StringID
	}{

		{"+ve", fields{testKey, testType}, testType},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyID := propertyID{
				Key:  tt.fields.Key,
				Type: tt.fields.Type,
			}
			if got := propertyID.GetType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyID_String(t *testing.T) {
	testKey, testType, _ := createTestInputForPropertyID()
	type fields struct {
		Key  ids.StringID
		Type ids.StringID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{

		{"+ve", fields{testKey, testType}, strings.Join([]string{testKey.String(), testType.String()}, ".")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyID := propertyID{
				Key:  tt.fields.Key,
				Type: tt.fields.Type,
			}
			if got := propertyID.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
