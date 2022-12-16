// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var (
	testIdentityIDPrototype = PrototypeIdentityID()
	testOwnableIDPrototype  = PrototypeOwnableID()
	testIdentityID, _       = ReadIdentityID("CBepOLnJFnKO9NEyZlSv7r80nKNZFFXRqHfnsObZ_KU=")
	testOwnableID           = NewOwnableID(NewStringID("ownableID"))
)

func TestNewSplitID(t *testing.T) {
	type args struct {
		ownerID   ids.IdentityID
		ownableID ids.OwnableID
	}
	tests := []struct {
		name string
		args args
		want ids.SplitID
	}{
		{"+ve", args{testIdentityID, testOwnableID}, splitID{testIdentityID, testOwnableID}},
		{"+ve test", args{testIdentityIDPrototype, testOwnableIDPrototype}, splitID{testIdentityIDPrototype, testOwnableIDPrototype}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSplitID(tt.args.ownerID, tt.args.ownableID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSplitID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrototypeSplitID(t *testing.T) {
	tests := []struct {
		name string
		want ids.SplitID
	}{
		{"+ve", splitID{testIdentityIDPrototype, testOwnableIDPrototype}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrototypeSplitID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrototypeSplitID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadSplitID(t *testing.T) {
	type args struct {
		splitIDString string
	}
	tests := []struct {
		name    string
		args    args
		want    ids.SplitID
		wantErr bool
	}{
		{"+ve", args{NewSplitID(testIdentityID, testOwnableID).String()}, splitID{testIdentityID, testOwnableID}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadSplitID(tt.args.splitIDString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadSplitID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadSplitID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitIDFromInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name        string
		args        args
		want        splitID
		shouldPanic bool
	}{
		{"+ve for PrototypeSplitID", args{PrototypeSplitID()}, splitID{testIdentityIDPrototype, testOwnableIDPrototype}, false},
		{"+ve", args{splitID{testIdentityID, testOwnableID}}, splitID{testIdentityID, testOwnableID}, false},
		{"-ve for panic case", args{testOwnableID}, splitID{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				assert.Panics(t, func() { splitIDFromInterface(tt.args.i) })
			} else if got := splitIDFromInterface(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitIDFromInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitID_Bytes(t *testing.T) {
	type fields struct {
		OwnerID   ids.IdentityID
		OwnableID ids.OwnableID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve with Prototype", fields{testIdentityIDPrototype, testOwnableID}, append(testIdentityIDPrototype.Bytes(), testOwnableID.Bytes()...)},
		{"+ve", fields{testIdentityID, testOwnableID}, append(testIdentityID.Bytes(), testOwnableID.Bytes()...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			splitID := splitID{
				OwnerID:   tt.fields.OwnerID,
				OwnableID: tt.fields.OwnableID,
			}
			if got := splitID.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitID_Compare(t *testing.T) {
	type fields struct {
		OwnerID   ids.IdentityID
		OwnableID ids.OwnableID
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
		{"+ve", fields{testIdentityID, testOwnableID}, args{splitID{testIdentityID, testOwnableID}}, 0},
		{"-ve", fields{testIdentityIDPrototype, testOwnableID}, args{splitID{testIdentityIDPrototype, testOwnableIDPrototype}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			splitID := splitID{
				OwnerID:   tt.fields.OwnerID,
				OwnableID: tt.fields.OwnableID,
			}
			if got := splitID.Compare(tt.args.listable); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitID_GetOwnableID(t *testing.T) {
	type fields struct {
		OwnerID   ids.IdentityID
		OwnableID ids.OwnableID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		{"+ve", fields{testIdentityID, testOwnableID}, testOwnableID},
		{"+ve", fields{testIdentityIDPrototype, testOwnableID}, testOwnableID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			splitID := splitID{
				OwnerID:   tt.fields.OwnerID,
				OwnableID: tt.fields.OwnableID,
			}
			if got := splitID.GetOwnableID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOwnableID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitID_IsSplitID(t *testing.T) {
	type fields struct {
		OwnerID   ids.IdentityID
		OwnableID ids.OwnableID
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"+ve", fields{testIdentityID, testOwnableID}},
		{"+ve", fields{testIdentityIDPrototype, testOwnableIDPrototype}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			splitID := splitID{
				OwnerID:   tt.fields.OwnerID,
				OwnableID: tt.fields.OwnableID,
			}
			splitID.IsSplitID()
		})
	}
}

func Test_splitID_String(t *testing.T) {
	type fields struct {
		OwnerID   ids.IdentityID
		OwnableID ids.OwnableID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve with prototype", fields{testIdentityIDPrototype, testOwnableIDPrototype}, stringUtilities.JoinIDStrings(testIdentityIDPrototype.String(), testOwnableIDPrototype.String())},
		{"+ve", fields{testIdentityID, testOwnableID}, stringUtilities.JoinIDStrings(testIdentityID.String(), testOwnableID.String())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			splitID := splitID{
				OwnerID:   tt.fields.OwnerID,
				OwnableID: tt.fields.OwnableID,
			}
			if got := splitID.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
