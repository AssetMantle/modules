// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package order

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/orders/internal/common"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	base2 "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

var (
	immutables = baseQualified.NewImmutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables   = baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))

	testClassificationID = baseIDs.NewClassificationID(immutables, mutables)
	testFromID           = baseIDs.NewIdentityID(testClassificationID, immutables)
	testOrderID          = baseIDs.NewOrderID(testClassificationID, immutables)
	testOrderID1         = baseIDs.PrototypeOrderID()
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
		{"+ve", args{testOrderID}, newQueryRequest(testOrderID)},
		{"+ve with nil", args{testOrderID1}, newQueryRequest(testOrderID1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newQueryRequest(tt.args.orderID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newQueryRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryRequestFromInterface(t *testing.T) {
	type args struct {
		request helpers.QueryRequest
	}
	tests := []struct {
		name string
		args args
		want queryRequest
	}{
		{"+ve", args{newQueryRequest(testOrderID)}, queryRequest{testOrderID}},
		{"+ve with nil", args{newQueryRequest(testOrderID1)}, queryRequest{testOrderID1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := queryRequestFromInterface(tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryRequestFromInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryRequest_Decode(t *testing.T) {
	encodedReq, err := common.LegacyAmino.MarshalJSON(newQueryRequest(testOrderID))
	require.NoError(t, err)
	encodedReq1, err1 := common.LegacyAmino.MarshalJSON(newQueryRequest(testOrderID1))
	require.NoError(t, err1)
	type fields struct {
		OrderID ids.OrderID
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.QueryRequest
		wantErr bool
	}{
		{"+ve", fields{testOrderID}, args{encodedReq}, newQueryRequest(testOrderID), false},
		{"+ve", fields{testOrderID1}, args{encodedReq1}, newQueryRequest(testOrderID1), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := queryRequest{
				OrderID: tt.fields.OrderID,
			}
			got, err := queryRequest.Decode(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryRequest_Encode(t *testing.T) {
	encodedReq, err := common.LegacyAmino.MarshalJSON(newQueryRequest(testOrderID))
	require.NoError(t, err)
	encodedReq1, err1 := common.LegacyAmino.MarshalJSON(newQueryRequest(testOrderID1))
	require.NoError(t, err1)
	type fields struct {
		OrderID ids.OrderID
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{"+ve", fields{testOrderID}, encodedReq, false},
		{"+ve", fields{testOrderID1}, encodedReq1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := queryRequest{
				OrderID: tt.fields.OrderID,
			}
			got, err := queryRequest.Encode()
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryRequest_FromCLI(t *testing.T) {
	cliCommand := base2.NewCLICommand("", "", "", []helpers.CLIFlag{constants.OrderID})
	viper.Set(constants.OrderID.GetName(), testOrderID.String())
	type fields struct {
		OrderID ids.OrderID
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
		{"+ve", fields{testOrderID}, args{cliCommand, context}, newQueryRequest(testOrderID), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qu := queryRequest{
				OrderID: tt.fields.OrderID,
			}
			got, err := qu.FromCLI(tt.args.cliCommand, tt.args.context)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromCLI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromCLI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryRequest_FromMap(t *testing.T) {
	vars := make(map[string]string)
	vars[Query.GetName()] = testOrderID.String()
	type fields struct {
		OrderID ids.OrderID
	}
	type args struct {
		vars map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.QueryRequest
		wantErr bool
	}{
		{"+ve", fields{testOrderID}, args{vars}, newQueryRequest(testOrderID), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qu := queryRequest{
				OrderID: tt.fields.OrderID,
			}
			got, err := qu.FromMap(tt.args.vars)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryRequest_Validate(t *testing.T) {
	type fields struct {
		OrderID ids.OrderID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{testOrderID}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := queryRequest{
				OrderID: tt.fields.OrderID,
			}
			if err := queryRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_requestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.QueryRequest
	}{
		{"+ve", queryRequest{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := requestPrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("requestPrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
