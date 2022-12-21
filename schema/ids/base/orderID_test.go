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
	testOrderID  = NewOrderID(testClassificationID, testImmutables)
	testOrderID2 = NewOrderID(testClassificationID2, testImmutables2)
)

func TestNewOrderID(t *testing.T) {
	type args struct {
		classificationID ids.ClassificationID
		immutables       qualified.Immutables
	}
	tests := []struct {
		name string
		args args
		want ids.OrderID
	}{
		{"+ve with Prototype ClasificationID", args{testClassificationID2, testImmutables}, orderID{GenerateHashID(testClassificationID2.Bytes(), testImmutables.GenerateHashID().Bytes())}},
		{"+ve", args{testClassificationID, testImmutables}, orderID{GenerateHashID(testClassificationID.Bytes(), testImmutables.GenerateHashID().Bytes())}},
		{"+ve with nil ClassificationID", args{nil, testImmutables}, orderID{GenerateHashID(nil, testImmutables.GenerateHashID().Bytes())}}, //TODO: implement Sanitize()
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewOrderID(tt.args.classificationID, tt.args.immutables), "NewOrderID(%v, %v)", tt.args.classificationID, tt.args.immutables)
		})
	}
}

func TestPrototypeOrderID(t *testing.T) {
	tests := []struct {
		name string
		want ids.OrderID
	}{
		{"+ve", orderID{GenerateHashID(testClassificationID2.Bytes(), testImmutables2.GenerateHashID().Bytes())}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PrototypeOrderID(), "PrototypeOrderID()")
		})
	}
}

func TestReadOrderID(t *testing.T) {
	type args struct {
		orderIDString string
	}
	tests := []struct {
		name    string
		args    args
		want    ids.OrderID
		wantErr assert.ErrorAssertionFunc
	}{
		{"+ve", args{testOrderID.String()}, testOrderID, assert.NoError},
		{"+ve with Prototype", args{PrototypeOrderID().String()}, PrototypeOrderID(), assert.NoError}, //TODO: Restructure the logic
		{"+ve with empty string", args{""}, nil, assert.Error},                                        //TODO: Restructure the logic
		{"+ve with invalid string", args{"invalid"}, nil, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadOrderID(tt.args.orderIDString)
			if !tt.wantErr(t, err, fmt.Sprintf("ReadOrderID(%v)", tt.args.orderIDString)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ReadOrderID(%v)", tt.args.orderIDString)
		})
	}
}

func Test_orderIDFromInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name        string
		args        args
		want        orderID
		shouldPanic bool
	}{
		{"+ve", args{testOrderID}, testOrderID.(orderID), false},
		{"+ve with nil", args{nil}, orderID{}, true},
		{"+ve with Prototype", args{PrototypeOrderID()}, PrototypeOrderID().(orderID), false},
		{"+ve with invalid type", args{testClassificationID}, orderID{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				assert.Panicsf(t, func() { orderIDFromInterface(tt.args.i) }, "orderIDFromInterface(%v)", tt.args.i)
			} else {
				assert.Equalf(t, tt.want, orderIDFromInterface(tt.args.i), "orderIDFromInterface(%v)", tt.args.i)
			}

		})
	}
}

func Test_orderID_Compare(t *testing.T) {
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
		{"+ve", fields{testOrderID.(orderID).HashID}, args{testOrderID}, 0},
		{"+ve with Prototype", fields{PrototypeOrderID().(orderID).HashID}, args{PrototypeOrderID()}, 0},
		{"-ve", fields{testOrderID.(orderID).HashID}, args{testOrderID2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderID := orderID{
				HashID: tt.fields.HashID,
			}
			assert.Equalf(t, tt.want, orderID.Compare(tt.args.listable), "Compare(%v)", tt.args.listable)
		})
	}
}
