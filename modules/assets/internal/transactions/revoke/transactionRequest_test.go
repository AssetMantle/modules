// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/spf13/viper"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/utilities/transaction"
)

var (
	testBaseRequest = rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
)

func Test_newTransactionRequest(t *testing.T) {
	type args struct {
		baseReq          rest.BaseReq
		fromID           string
		toID             string
		classificationID string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve", args{testBaseRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString()}, transactionRequest{testBaseRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.baseReq, tt.args.fromID, tt.args.toID, tt.args.classificationID); !reflect.DeepEqual(got, tt.want) {
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
	cliCommand := base.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromID, constants.ToID, constants.ClassificationID})

	viper.Set(constants.FromID.GetName(), fromID.AsString())
	viper.Set(constants.ToID.GetName(), fromID.AsString())
	viper.Set(constants.ClassificationID.GetName(), classificationID.AsString())
	type fields struct {
		BaseReq          rest.BaseReq
		FromID           string
		ToID             string
		ClassificationID string
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
		{"+ve", fields{testBaseRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString()}, args{cliCommand, base.TestClientContext}, newTransactionRequest(testBaseRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString()), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:          tt.fields.BaseReq,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
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
	type fields struct {
		BaseReq          rest.BaseReq
		FromID           string
		ToID             string
		ClassificationID string
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
		{"+ve", fields{testBaseRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString()}, args{sdkTypes.MustSortJSON(transaction.RegisterLegacyAminoCodec(messagePrototype).MustMarshalJSON(&Message{fromAccAddress.String(), fromID, fromID, classificationID}))}, newTransactionRequest(testBaseRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString()), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:          tt.fields.BaseReq,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
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
		BaseReq          rest.BaseReq
		FromID           string
		ToID             string
		ClassificationID string
	}
	tests := []struct {
		name   string
		fields fields
		want   rest.BaseReq
	}{
		{"+ve", fields{testBaseRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString()}, testBaseRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:          tt.fields.BaseReq,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
			}
			if got := transactionRequest.GetBaseReq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBaseReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	type fields struct {
		BaseReq          rest.BaseReq
		FromID           string
		ToID             string
		ClassificationID string
	}
	tests := []struct {
		name    string
		fields  fields
		want    sdkTypes.Msg
		wantErr bool
	}{
		{"+ve", fields{testBaseRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString()}, newMessage(fromAccAddress, fromID, fromID, classificationID), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:          tt.fields.BaseReq,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
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
		BaseReq          rest.BaseReq
		FromID           string
		ToID             string
		ClassificationID string
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testBaseRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString()}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				BaseReq:          tt.fields.BaseReq,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
			}
			tr.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	type fields struct {
		BaseReq          rest.BaseReq
		FromID           string
		ToID             string
		ClassificationID string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{testBaseRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:          tt.fields.BaseReq,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
