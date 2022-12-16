// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	testStringID          = NewStringID("ownableID")
	testOwnableID         = NewOwnableID(testStringID)
	testStringIDPrototype = PrototypeStringID()
)

func TestNewOwnableID(t *testing.T) {
	type args struct {
		stringID ids.StringID
	}
	tests := []struct {
		name string
		args args
		want ids.OwnableID
	}{
		{"+ve", args{testStringID}, testOwnableID},
		{"+ve with prototype", args{testStringIDPrototype}, testOwnableIDPrototype},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewOwnableID(tt.args.stringID), "NewOwnableID(%v)", tt.args.stringID)
		})
	}
}

func TestPrototypeOwnableID(t *testing.T) {
	tests := []struct {
		name string
		want ids.OwnableID
	}{
		{"+ve", ownableID{testStringIDPrototype}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PrototypeOwnableID(), "PrototypeOwnableID()")
		})
	}
}

func TestReadOwnableID(t *testing.T) {
	assetID, err := ReadAssetID("")
	require.NoError(t, err)
	type args struct {
		ownableIDString string
	}
	tests := []struct {
		name    string
		args    args
		want    ids.OwnableID
		wantErr assert.ErrorAssertionFunc
	}{
		{"+ve", args{"ownableID"}, testOwnableID, assert.NoError},
		{"+ve with prototype", args{""}, assetID, assert.NoError},
		{"-ve with empty string", args{"ownableID"}, testOwnableID, assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadOwnableID(tt.args.ownableIDString)
			if !tt.wantErr(t, err, fmt.Sprintf("ReadOwnableID(%v)", tt.args.ownableIDString)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ReadOwnableID(%v)", tt.args.ownableIDString)
		})
	}
}

func Test_ownableIDFromInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name        string
		args        args
		want        ids.OwnableID
		shouldPanic bool
	}{
		{"+ve", args{testOwnableID}, testOwnableID, false},
		{"+ve with prototype", args{testOwnableIDPrototype}, testOwnableIDPrototype, false},
		{"+ve Panic case", args{testOwnableIDPrototype}, testOwnableIDPrototype, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ownableIDFromInterface(tt.args.i), "ownableIDFromInterface(%v)", tt.args.i)
		})
	}
}

func Test_ownableID_Compare(t *testing.T) {
	type fields struct {
		StringID ids.StringID
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
		{"+ve", fields{testStringID}, args{testOwnableID}, 0},
		{"+ve with prototype", fields{testStringIDPrototype}, args{testOwnableIDPrototype}, 0},
		{"-ve", fields{testStringID}, args{testOwnableIDPrototype}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ownableID := ownableID{
				StringID: tt.fields.StringID,
			}
			assert.Equalf(t, tt.want, ownableID.Compare(tt.args.listable), "Compare(%v)", tt.args.listable)
		})
	}
}

func Test_ownableID_IsOwnableID(t *testing.T) {
	type fields struct {
		StringID ids.StringID
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"+ve", fields{testStringID}},
		{"+ve with prototype", fields{testStringIDPrototype}},
		{"-ve", fields{testStringIDPrototype}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ownableID := ownableID{
				StringID: tt.fields.StringID,
			}
			ownableID.IsOwnableID()
		})
	}
}
