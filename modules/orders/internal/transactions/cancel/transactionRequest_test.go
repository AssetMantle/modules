// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package cancel

import (
	"encoding/json"
	"fmt"
	"github.com/AssetMantle/modules/utilities/test"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	basehelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

var (
	fromAddress       = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, _ = types.AccAddressFromBech32(fromAddress)
	immutables        = baseQualified.NewImmutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables          = baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))

	testClassificationID = baseIDs.NewClassificationID(immutables, mutables)
	testFromID           = baseIDs.NewIdentityID(testClassificationID, immutables)
	testOrderID          = baseIDs.NewOrderID(testClassificationID, immutables)
	testOrderID1         = baseIDs.PrototypeOrderID()
	testBaseRequest      = rest.BaseReq{From: fromAddress, ChainID: "test", Fees: types.NewCoins()}
)

func Test_newTransactionRequest(t *testing.T) {
	type args struct {
		baseReq rest.BaseReq
		fromID  string
		orderID string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve", args{testBaseRequest, testFromID.AsString(), testOrderID.AsString()}, newTransactionRequest(testBaseRequest, testFromID.AsString(), testOrderID.AsString())},
		{"+ve", args{testBaseRequest, testFromID.AsString(), testOrderID1.AsString()}, newTransactionRequest(testBaseRequest, testFromID.AsString(), testOrderID1.AsString())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.baseReq, tt.args.fromID, tt.args.orderID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTransactionRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_requestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.TransactionRequest
	}{
		{"+ve", transactionRequest{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := requestPrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("requestPrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_FromCLI(t *testing.T) {
	cliCommand := basehelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.OrderID, constants.FromID})

	viper.Set(constants.FromID.GetName(), testFromID.AsString())
	viper.Set(constants.OrderID.GetName(), testOrderID.AsString())
	type fields struct {
		BaseReq rest.BaseReq
		FromID  string
		OrderID string
	}
	type args struct {
		cliCommand helpers.CLICommand
		context    client.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.TransactionRequest
		wantErr bool
	}{
		{"+ve", fields{testBaseRequest, testFromID.AsString(), testOrderID.AsString()}, args{cliCommand, test.TestClientContext}, newTransactionRequest(testBaseRequest, testFromID.AsString(), testOrderID.AsString()), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				OrderID: tt.fields.OrderID,
			}
			got, err := transactionRequest.FromCLI(tt.args.cliCommand, tt.args.context)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromCLI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("FromCLI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_FromJSON(t *testing.T) {
	jsonMessage, err := json.Marshal(newTransactionRequest(testBaseRequest, testFromID.AsString(), testOrderID.AsString()))
	require.NoError(t, err)
	jsonMessage1, err := json.Marshal(newTransactionRequest(testBaseRequest, testFromID.AsString(), testOrderID1.AsString()))
	require.NoError(t, err)

	type fields struct {
		BaseReq rest.BaseReq
		FromID  string
		OrderID string
	}
	type args struct {
		rawMessage json.RawMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.TransactionRequest
		wantErr bool
	}{
		{"+ve", fields{testBaseRequest, testFromID.AsString(), testOrderID.AsString()}, args{jsonMessage}, transactionRequest{testBaseRequest, testFromID.AsString(), testOrderID.AsString()}, false},
		{"+ve with nil", fields{testBaseRequest, testFromID.AsString(), testOrderID1.AsString()}, args{jsonMessage1}, transactionRequest{testBaseRequest, testFromID.AsString(), testOrderID1.AsString()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				OrderID: tt.fields.OrderID,
			}
			got, err := transactionRequest.FromJSON(tt.args.rawMessage)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_GetBaseReq(t *testing.T) {
	type fields struct {
		BaseReq rest.BaseReq
		FromID  string
		OrderID string
	}
	tests := []struct {
		name   string
		fields fields
		want   rest.BaseReq
	}{
		{"+ve", fields{testBaseRequest, testFromID.AsString(), testOrderID.AsString()}, testBaseRequest},
		{"+ve with nil", fields{testBaseRequest, testFromID.AsString(), testOrderID1.AsString()}, testBaseRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				OrderID: tt.fields.OrderID,
			}
			if got := transactionRequest.GetBaseReq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBaseReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	type fields struct {
		BaseReq rest.BaseReq
		FromID  string
		OrderID string
	}
	tests := []struct {
		name    string
		fields  fields
		want    types.Msg
		wantErr bool
	}{
		{"+ve", fields{testBaseRequest, testFromID.AsString(), testOrderID.AsString()}, newMessage(fromAccAddress, testFromID, testOrderID), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				OrderID: tt.fields.OrderID,
			}
			got, err := transactionRequest.MakeMsg()
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeMsg() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_RegisterCodec(t *testing.T) {
	type fields struct {
		BaseReq rest.BaseReq
		FromID  string
		OrderID string
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testBaseRequest, testFromID.AsString(), testOrderID.AsString()}, args{codec.NewLegacyAmino()}},
		{"+ve with nil", fields{testBaseRequest, testFromID.AsString(), testOrderID1.AsString()}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				OrderID: tt.fields.OrderID,
			}
			tr.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	type fields struct {
		BaseReq rest.BaseReq
		FromID  string
		OrderID string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{testBaseRequest, testFromID.AsString(), testOrderID.AsString()}, false},
		{"+ve with nil", fields{testBaseRequest, testFromID.AsString(), testOrderID1.AsString()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				OrderID: tt.fields.OrderID,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
