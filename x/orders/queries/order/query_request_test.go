// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package order

import (
	"github.com/AssetMantle/modules/x/orders/key"
	"testing"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var (
	immutables           = baseQualified.NewImmutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables             = baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID = baseIDs.NewClassificationID(immutables, mutables).(*baseIDs.ClassificationID)
	testOrderID          = baseIDs.NewOrderID(testClassificationID, immutables).(*baseIDs.OrderID)
	testKey              = key.NewKey(testOrderID).(*key.Key)
	testOrderID1         = baseIDs.PrototypeOrderID().(*baseIDs.OrderID)
	testKey1             = key.NewKey(testOrderID1).(*key.Key)
)

func Test_newQueryRequest(t *testing.T) {
	type args struct {
		orderID ids.OrderID
	}
	tests := []struct {
		name string
		args args
		want helpers.QueryRequest
	}{
		{"valid", args{testOrderID}, newQueryRequest(testOrderID)},
		{"nil inputs", args{testOrderID1}, newQueryRequest(testOrderID1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newQueryRequest(tt.args.orderID)
			assert.Equal(t, tt.want, got, "newQueryRequest()")
		})
	}
}

func Test_queryRequest_FromCLI(t *testing.T) {
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.OrderID})

	viper.Set(constants.OrderID.GetName(), testOrderID.AsString())
	type fields struct {
		Key *key.Key
	}
	type args struct {
		cliCommand helpers.CLICommand
		context    client.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.QueryRequest
		wantErr bool
	}{
		{"valid", fields{testKey}, args{cliCommand, client.Context{}.WithCodec(baseHelpers.CodecPrototype())}, newQueryRequest(testOrderID), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qu := &QueryRequest{
				Key: tt.fields.Key,
			}
			got, err := qu.FromCLI(tt.args.cliCommand, tt.args.context)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromCLI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got, "FromCLI() got")
		})
	}
}

func Test_queryRequest_Validate(t *testing.T) {
	type fields struct {
		Key *key.Key
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"valid", fields{testKey}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := &QueryRequest{
				Key: tt.fields.Key,
			}
			if err := queryRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_requestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.QueryRequest
	}{
		{"valid", &QueryRequest{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := requestPrototype()
			assert.Equal(t, tt.want, got, "requestPrototype()")
		})
	}
}
