// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testIdentityID  = NewIdentityID(testClassificationID, testImmutables)
	testIdentityID2 = PrototypeIdentityID()
)

func TestNewIdentityID(t *testing.T) {
	type args struct {
		classificationID ids.ClassificationID
		immutables       qualified.Immutables
	}
	tests := []struct {
		name string
		args args
		want ids.IdentityID
	}{
		{"+ve with nil immutables and mutables", args{testClassificationID2, testImmutables2}, identityID{GenerateHashID(testClassificationID2.Bytes(), testImmutables2.GenerateHashID().Bytes())}},
		{"+ve", args{testClassificationID, testImmutables}, identityID{HashID: GenerateHashID(testClassificationID.Bytes(), testImmutables.GenerateHashID().Bytes())}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewIdentityID(tt.args.classificationID, tt.args.immutables), "NewIdentityID(%v, %v)", tt.args.classificationID, tt.args.immutables)
		})
	}
}

func TestPrototypeIdentityID(t *testing.T) {
	tests := []struct {
		name string
		want ids.IdentityID
	}{
		{"+ve", identityID{PrototypeHashID()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PrototypeIdentityID(), "PrototypeIdentityID()")
		})
	}
}

func TestReadIdentityID(t *testing.T) {
	type args struct {
		identityIDString string
	}
	tests := []struct {
		name    string
		args    args
		want    ids.IdentityID
		wantErr assert.ErrorAssertionFunc
	}{
		{"+ve", args{testIdentityID.String()}, testIdentityID, assert.NoError},
		{"-ve with invalid identityIDString", args{"invalid"}, identityID{}, assert.Error},
		{"+ve with empty identityIDString", args{""}, identityID{}, assert.NoError}, //TODO: change the order for emptyString in ReadHashID
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadIdentityID(tt.args.identityIDString)
			if !tt.wantErr(t, err, fmt.Sprintf("ReadIdentityID(%v)", tt.args.identityIDString)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ReadIdentityID(%v)", tt.args.identityIDString)
		})
	}
}

func Test_identityIDFromInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name        string
		args        args
		want        identityID
		shouldPanic bool
	}{
		{"+ve", args{testIdentityID}, NewIdentityID(testClassificationID, testImmutables).(identityID), false},
		{"+ve with nil", args{nil}, identityID{}, true},
		{"-ve with invalid type", args{testClassificationID}, identityID{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				assert.Panicsf(t, func() { identityIDFromInterface(tt.args.i) }, "identityIDFromInterface(%v)", tt.args.i)
			} else {
				assert.Equalf(t, tt.want, identityIDFromInterface(tt.args.i), "identityIDFromInterface(%v)", tt.args.i)
			}

		})
	}
}

func Test_identityID_Bytes(t *testing.T) {
	type fields struct {
		HashID ids.HashID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve", fields{testIdentityID.GetHashID()}, testIdentityID.GetHashID().Bytes()},
		{"+ve with nil", fields{nil}, []byte{}}, //TODO: Sanitize it
		{"+ve with nil mutable and immutable", fields{identityID{HashID: testIdentityID2.GetHashID()}}, testIdentityID2.GetHashID().Bytes()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identityID := identityID{
				HashID: tt.fields.HashID,
			}
			assert.Equalf(t, tt.want, identityID.Bytes(), "Bytes()")
		})
	}
}

func Test_identityID_Compare(t *testing.T) {
	type fields struct {
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
		{"-ve", fields{testIdentityID.GetHashID()}, args{testIdentityID2}, 1},
		{"-ve with nil immutables mutables", fields{testIdentityID2.GetHashID()}, args{testIdentityID}, -1},
		{"+ve", fields{testIdentityID.GetHashID()}, args{testIdentityID}, 0},
		{"+ve with nil immutables mutables", fields{testIdentityID2.GetHashID()}, args{testIdentityID2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identityID := identityID{
				HashID: tt.fields.HashID,
			}
			assert.Equalf(t, tt.want, identityID.Compare(tt.args.listable), "Compare(%v)", tt.args.listable)
		})
	}
}

func Test_identityID_GetHashID(t *testing.T) {
	type fields struct {
		HashID ids.HashID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.HashID
	}{
		{"+ve", fields{GenerateHashID(testClassificationID.Bytes(), testImmutables.GenerateHashID().Bytes())}, GenerateHashID(testClassificationID.Bytes(), testImmutables.GenerateHashID().Bytes())},
		{"+ve with nil mutable immutables", fields{testIdentityID2.GetHashID()}, testIdentityID2.GetHashID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identityID := identityID{
				HashID: tt.fields.HashID,
			}
			assert.Equalf(t, tt.want, identityID.GetHashID(), "GetHashID()")
		})
	}
}

func Test_identityID_String(t *testing.T) {
	type fields struct {
		HashID ids.HashID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{testIdentityID.GetHashID()}, testIdentityID.GetHashID().String()},
		{"+ve with nil mutable immutables", fields{testIdentityID2.GetHashID()}, testIdentityID2.GetHashID().String()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identityID := identityID{
				HashID: tt.fields.HashID,
			}
			assert.Equalf(t, tt.want, identityID.String(), "String()")
		})
	}
}
