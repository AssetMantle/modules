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

func getListBytes(immutables qualified.Immutables, mutables qualified.Mutables) ([][]byte, [][]byte) {
	immutableIDByteList := make([][]byte, len(immutables.GetImmutablePropertyList().GetList()))
	for i, property := range immutables.GetImmutablePropertyList().GetList() {
		immutableIDByteList[i] = property.GetID().Bytes()
	}

	mutableIDByteList := make([][]byte, len(mutables.GetMutablePropertyList().GetList()))
	for i, property := range mutables.GetMutablePropertyList().GetList() {
		mutableIDByteList[i] = property.GetID().Bytes()
	}
	return immutableIDByteList, mutableIDByteList
}

func TestNewClassificationID(t *testing.T) {
	immutableIDByteList, mutableIDByteList := getListBytes(testImmutables, testMutables)
	immutableIDByteList2, mutableIDByteList2 := getListBytes(testImmutables2, testMutables2)
	type args struct {
		immutables qualified.Immutables
		mutables   qualified.Mutables
	}
	tests := []struct {
		name string
		args args
		want ids.ClassificationID
	}{
		{"+ve with nil immutables and mutables", args{testImmutables2, testMutables2}, classificationID{GenerateHashID(GenerateHashID(immutableIDByteList2...).Bytes(), GenerateHashID(mutableIDByteList2...).Bytes(), testImmutables2.GenerateHashID().Bytes())}},
		{"+ve", args{testImmutables, testMutables}, classificationID{GenerateHashID(GenerateHashID(immutableIDByteList...).Bytes(), GenerateHashID(mutableIDByteList...).Bytes(), testImmutables.GenerateHashID().Bytes())}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewClassificationID(tt.args.immutables, tt.args.mutables), "NewClassificationID(%v, %v)", tt.args.immutables, tt.args.mutables)
		})
	}
}

func TestPrototypeClassificationID(t *testing.T) {
	tests := []struct {
		name string
		want ids.ClassificationID
	}{
		{"+ve", classificationID{PrototypeHashID()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PrototypeClassificationID(), "PrototypeClassificationID()")
		})
	}
}

func TestReadClassificationID(t *testing.T) {
	type args struct {
		classificationIDString string
	}
	tests := []struct {
		name    string
		args    args
		want    ids.ClassificationID
		wantErr assert.ErrorAssertionFunc
	}{
		{"+ve", args{testClassificationID.String()}, testClassificationID, assert.NoError},
		{"+ve with nil immutables and mutables", args{testClassificationID2.String()}, testClassificationID2, assert.NoError},
		{"-ve", args{"invalid"}, classificationID{}, assert.Error},
		{"+ve with empty string", args{""}, PrototypeClassificationID(), assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadClassificationID(tt.args.classificationIDString)
			if !tt.wantErr(t, err, fmt.Sprintf("ReadClassificationID(%v)", tt.args.classificationIDString)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ReadClassificationID(%v)", tt.args.classificationIDString)
		})
	}
}

func Test_classificationIDFromInterface(t *testing.T) {
	immutableIDByteList, mutableIDByteList := getListBytes(testImmutables, testMutables)
	immutableIDByteList2, mutableIDByteList2 := getListBytes(testImmutables2, testMutables2)
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want classificationID
	}{
		{"+ve with nil immutables and mutables", args{testClassificationID2}, classificationID{GenerateHashID(GenerateHashID(immutableIDByteList2...).Bytes(), GenerateHashID(mutableIDByteList2...).Bytes(), testImmutables2.GenerateHashID().Bytes())}},
		{"+ve", args{testClassificationID}, classificationID{GenerateHashID(GenerateHashID(immutableIDByteList...).Bytes(), GenerateHashID(mutableIDByteList...).Bytes(), testImmutables.GenerateHashID().Bytes())}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, classificationIDFromInterface(tt.args.i), "classificationIDFromInterface(%v)", tt.args.i)
		})
	}
}

func Test_classificationID_Compare(t *testing.T) {
	immutableIDByteList, mutableIDByteList := getListBytes(testImmutables, testMutables)
	immutableIDByteList2, mutableIDByteList2 := getListBytes(testImmutables2, testMutables2)
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
		{"+ve with nil immutables and mutables", fields{GenerateHashID(GenerateHashID(immutableIDByteList2...).Bytes(), GenerateHashID(mutableIDByteList2...).Bytes(), testImmutables2.GenerateHashID().Bytes())}, args{testClassificationID2}, 0},
		{"-ve", fields{GenerateHashID(GenerateHashID(immutableIDByteList...).Bytes(), GenerateHashID(mutableIDByteList...).Bytes(), testImmutables.GenerateHashID().Bytes())}, args{testClassificationID2}, 1},
		{"+ve", fields{GenerateHashID(GenerateHashID(immutableIDByteList...).Bytes(), GenerateHashID(mutableIDByteList...).Bytes(), testImmutables.GenerateHashID().Bytes())}, args{testClassificationID}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			classificationID := classificationID{
				HashID: tt.fields.HashID,
			}
			assert.Equalf(t, tt.want, classificationID.Compare(tt.args.listable), "Compare(%v)", tt.args.listable)
		})
	}
}
