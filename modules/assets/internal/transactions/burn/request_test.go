// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/spf13/viper"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/utilities/transaction"
)

var (
	testBaseRequest = rest.BaseReq{From: fromAddress, ChainID: "test", Fees: types.NewCoins()}
)

func Test_newTransactionRequest(t *testing.T) {
	type args struct {
		baseReq rest.BaseReq
		fromID  string
		assetID string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve", args{testBaseRequest, fromID.String(), testAssetID.String()}, TransactionRequest{testBaseRequest, fromID.String(), testAssetID.String()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.baseReq, tt.args.fromID, tt.args.assetID); !reflect.DeepEqual(got, tt.want) {
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
		{"+ve", TransactionRequest{}},
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
	cliCommand := base.NewCLICommand("", "", "", []helpers.CLIFlag{constants.AssetID, constants.FromID})
	viper.Set(constants.AssetID.GetName(), testAssetID.String())
	viper.Set(constants.FromID.GetName(), fromID.String())
	type fields struct {
		BaseReq rest.BaseReq
		FromID  string
		AssetID string
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
		{"+ve", fields{BaseReq: testBaseRequest, FromID: fromID.String(), AssetID: testAssetID.String()}, args{cliCommand, context}, newTransactionRequest(testBaseRequest, fromID.String(), testAssetID.String()), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := TransactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				AssetID: tt.fields.AssetID,
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
		BaseReq rest.BaseReq
		FromID  string
		AssetID string
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
		{"+ve", fields{testBaseRequest, fromID.String(), testAssetID.String()}, args{types.MustSortJSON(transaction.RegisterLegacyAminoCodec(messagePrototype).MustMarshalJSON(message{fromAccAddress, fromID, testAssetID}))}, newTransactionRequest(testBaseRequest, fromID.String(), testAssetID.String()), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := TransactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				AssetID: tt.fields.AssetID,
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
		AssetID string
	}
	tests := []struct {
		name   string
		fields fields
		want   rest.BaseReq
	}{
		{"+ve", fields{testBaseRequest, fromID.String(), testAssetID.String()}, testBaseRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := TransactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				AssetID: tt.fields.AssetID,
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
		AssetID string
	}
	tests := []struct {
		name    string
		fields  fields
		want    types.Msg
		wantErr bool
	}{
		{"+ve", fields{testBaseRequest, fromID.String(), testAssetID.String()}, newMessage(fromAccAddress, fromID, testAssetID), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := TransactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				AssetID: tt.fields.AssetID,
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
		AssetID string
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testBaseRequest, fromID.String(), testAssetID.String()}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TransactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				AssetID: tt.fields.AssetID,
			}
			tr.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	type fields struct {
		BaseReq rest.BaseReq
		FromID  string
		AssetID string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{testBaseRequest, fromID.String(), testAssetID.String()}, false},
		// {"+ve", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := TransactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				AssetID: tt.fields.AssetID,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
