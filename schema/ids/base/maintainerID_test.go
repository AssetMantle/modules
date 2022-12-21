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
	testMaintainerID  = NewMaintainerID(testClassificationID, testImmutables)
	testMaintainerID2 = NewMaintainerID(testClassificationID2, testImmutables2)
)

func TestNewMaintainerID(t *testing.T) {
	type args struct {
		classificationID ids.ClassificationID
		immutables       qualified.Immutables
	}
	tests := []struct {
		name string
		args args
		want ids.MaintainerID
	}{
		{"+ve", args{testClassificationID, testImmutables}, maintainerID{GenerateHashID(testClassificationID.Bytes(), testImmutables.GenerateHashID().Bytes())}},
		{"+ve with PrototypeClassificationID", args{PrototypeClassificationID(), testImmutables}, maintainerID{GenerateHashID(PrototypeClassificationID().Bytes(), testImmutables.GenerateHashID().Bytes())}},
		{"+ve with PrototypeImmutables", args{testClassificationID, testImmutables2}, maintainerID{GenerateHashID(testClassificationID.Bytes(), testImmutables2.GenerateHashID().Bytes())}},
		{"+ve with nil immutables", args{testClassificationID, nil}, maintainerID{GenerateHashID(testClassificationID.Bytes(), nil)}},            //TODO: panics for nil immutables
		{"+ve with nil classificationID", args{nil, testImmutables}, maintainerID{GenerateHashID(nil, testImmutables.GenerateHashID().Bytes())}}, //TODO: panics for nil classificationID
		{"+ve with both nil", args{}, maintainerID{GenerateHashID(nil, nil)}},                                                                    //TODO: panics for nil classificationID and nil immutables
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewMaintainerID(tt.args.classificationID, tt.args.immutables), "NewMaintainerID(%v, %v)", tt.args.classificationID, tt.args.immutables)
		})
	}
}

func TestPrototypeMaintainerID(t *testing.T) {
	tests := []struct {
		name string
		want ids.MaintainerID
	}{
		{"+ve", maintainerID{PrototypeHashID()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PrototypeMaintainerID(), "PrototypeMaintainerID()")
		})
	}
}

func TestReadMaintainerID(t *testing.T) {
	type args struct {
		maintainerIDString string
	}
	tests := []struct {
		name    string
		args    args
		want    ids.MaintainerID
		wantErr assert.ErrorAssertionFunc
	}{
		{"+ve with valid maintainerID", args{testMaintainerID.String()}, testMaintainerID, assert.NoError},
		{"+ve with valid maintainerID with leading and trailing spaces", args{" " + testMaintainerID.String() + " "}, testMaintainerID, assert.NoError}, //TODO: MetaDataError for leading and trailing spaces
		{"+ve with empty maintainerID", args{""}, PrototypeMaintainerID(), assert.NoError},                                                              //TODO: need to refactor the test cases
		{"+ve with PrototypeMaintainerID", args{PrototypeMaintainerID().String()}, PrototypeMaintainerID(), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadMaintainerID(tt.args.maintainerIDString)
			if !tt.wantErr(t, err, fmt.Sprintf("ReadMaintainerID(%v)", tt.args.maintainerIDString)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ReadMaintainerID(%v)", tt.args.maintainerIDString)
		})
	}
}

func Test_maintainerIDFromInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name        string
		args        args
		want        maintainerID
		shouldPanic bool
	}{
		{"+ve with valid maintainerID", args{testMaintainerID}, testMaintainerID.(maintainerID), false},
		{"+ve PrototypeMaintainerID", args{PrototypeMaintainerID()}, PrototypeMaintainerID().(maintainerID), false},
		{"-ve with invalid type", args{testIdentityID}, maintainerID{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				assert.Panicsf(t, func() { maintainerIDFromInterface(tt.args.i) }, "maintainerIDFromInterface(%v)", tt.args.i)
			} else {
				assert.Equalf(t, tt.want, maintainerIDFromInterface(tt.args.i), "maintainerIDFromInterface(%v)", tt.args.i)
			}

		})
	}
}

func Test_maintainerID_Compare(t *testing.T) {
	type fields struct {
		HashID ids.HashID
	}
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        int
		shouldPanic bool
	}{
		{"+ve with same maintainerID", fields{testMaintainerID.(maintainerID).HashID}, args{testMaintainerID}, 0, false},
		{"+ve with different maintainerID", fields{testMaintainerID.(maintainerID).HashID}, args{testMaintainerID2}, 1, false},
		{"+ve with different type", fields{testMaintainerID.(maintainerID).HashID}, args{testIdentityID}, -1, true},
		{"+ce with nil", fields{testMaintainerID.(maintainerID).HashID}, args{nil}, 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maintainerID := maintainerID{
				HashID: tt.fields.HashID,
			}
			if tt.shouldPanic {
				assert.Panicsf(t, func() { maintainerID.Compare(tt.args.listable) }, "maintainerID.Compare(%v)", tt.args.listable)
			} else {
				assert.Equalf(t, tt.want, maintainerID.Compare(tt.args.listable), "Compare(%v)", tt.args.listable)
			}
		})
	}
}
