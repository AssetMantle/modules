// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package send

import (
	"encoding/json"
	"fmt"
	"github.com/AssetMantle/modules/utilities/rest"
	"reflect"
	"testing"

	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/client"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
)

var (
	fromAccAddress, _ = types.AccAddressFromBech32(fromAddress.String())
	testBaseRequest   = rest.BaseReq{From: fromAddress.String(), ChainID: "test", Fees: types.NewCoins()}
	fromID            = baseIDs.PrototypeIdentityID().(*baseIDs.IdentityID)
	testRate          = types.OneInt()
)

func Test_newTransactionRequest(t *testing.T) {
	type args struct {
		baseReq rest.BaseReq
		fromID  string
		toID    string
		assetID string
		value   string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve", args{testBaseRequest, fromID.AsString(), fromID.AsString(), assetID.AsString(), testRate.String()}, transactionRequest{testBaseRequest, fromID.AsString(), fromID.AsString(), assetID.AsString(), testRate.String()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.baseReq, tt.args.fromID, tt.args.toID, tt.args.assetID, tt.args.value); !reflect.DeepEqual(got, tt.want) {
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
	cliCommand := base.NewCLICommand("", "", "", []helpers.CLIFlag{constants.ToIdentityID, constants.FromIdentityID, constants.AssetID, constants.Value})

	viper.Set(constants.FromIdentityID.GetName(), fromID.AsString())
	viper.Set(constants.ToIdentityID.GetName(), fromID.AsString())
	viper.Set(constants.AssetID.GetName(), assetID.AsString())
	viper.Set(constants.Value.GetName(), testRate.String())
	type fields struct {
		BaseReq rest.BaseReq
		FromID  string
		ToID    string
		AssetID string
		Value   string
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
		{"+ve", fields{testBaseRequest, fromID.AsString(), fromID.AsString(), assetID.AsString(), testRate.String()}, args{cliCommand, client.Context{}.WithCodec(base.CodecPrototype())}, transactionRequest{testBaseRequest, fromID.AsString(), fromID.AsString(), assetID.AsString(), testRate.String()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				ToID:    tt.fields.ToID,
				AssetID: tt.fields.AssetID,
				Value:   tt.fields.Value,
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
	jsonMessage, err := json.Marshal(newTransactionRequest(testBaseRequest, fromID.AsString(), fromID.AsString(), assetID.AsString(), testRate.String()))
	require.NoError(t, err)
	type fields struct {
		BaseReq rest.BaseReq
		FromID  string
		ToID    string
		AssetID string
		Value   string
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
		{"+ve", fields{testBaseRequest, fromID.AsString(), fromID.AsString(), assetID.AsString(), testRate.String()}, args{jsonMessage}, transactionRequest{testBaseRequest, fromID.AsString(), fromID.AsString(), assetID.AsString(), testRate.String()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				ToID:    tt.fields.ToID,
				AssetID: tt.fields.AssetID,
				Value:   tt.fields.Value,
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
		ToID    string
		AssetID string
		Value   string
	}
	tests := []struct {
		name   string
		fields fields
		want   rest.BaseReq
	}{
		{"+ve", fields{testBaseRequest, fromID.AsString(), fromID.AsString(), assetID.AsString(), testRate.String()}, testBaseRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				ToID:    tt.fields.ToID,
				AssetID: tt.fields.AssetID,
				Value:   tt.fields.Value,
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
		ToID    string
		AssetID string
		Value   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    types.Msg
		wantErr bool
	}{
		{"+ve", fields{testBaseRequest, fromID.AsString(), fromID.AsString(), assetID.AsString(), testRate.String()}, NewMessage(fromAccAddress, fromID, fromID, assetID, testRate), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				ToID:    tt.fields.ToID,
				AssetID: tt.fields.AssetID,
				Value:   tt.fields.Value,
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
		ToID    string
		AssetID string
		Value   string
	}
	type args struct {
		legacyAmino *sdkCodec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testBaseRequest, fromID.AsString(), fromID.AsString(), assetID.AsString(), testRate.String()}, args{codec.GetLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				ToID:    tt.fields.ToID,
				AssetID: tt.fields.AssetID,
				Value:   tt.fields.Value,
			}
			tr.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	type fields struct {
		BaseReq rest.BaseReq
		FromID  string
		ToID    string
		AssetID string
		Value   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{testBaseRequest, fromID.AsString(), fromID.AsString(), assetID.AsString(), testRate.String()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				ToID:    tt.fields.ToID,
				AssetID: tt.fields.AssetID,
				Value:   tt.fields.Value,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
