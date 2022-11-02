// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package nub

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/utilities/transaction"
)

func CreateTestInputForRequest(t *testing.T) (*codec.Codec, helpers.CLICommand, context.CLIContext, string, sdkTypes.AccAddress, rest.BaseReq) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.NubID})
	cliContext := context.NewCLIContext().WithCodec(Codec)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}

	return Codec, cliCommand, cliContext, fromAddress, fromAccAddress, testBaseReq
}

func Test_newTransactionRequest(t *testing.T) {
	_, _, _, _, _, testBaseReq := CreateTestInputForRequest(t)
	type args struct {
		baseReq rest.BaseReq
		nubID   string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		// TODO: Add test cases.
		{"+ve", args{testBaseReq, "nubID"}, transactionRequest{testBaseReq, "nubID"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.baseReq, tt.args.nubID); !reflect.DeepEqual(got, tt.want) {
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
		// TODO: Add test cases.
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
	_, cliCommand, cliContext, _, _, testBaseReq := CreateTestInputForRequest(t)

	type fields struct {
		BaseReq rest.BaseReq
		NubID   string
	}
	type args struct {
		cliCommand helpers.CLICommand
		cliContext context.CLIContext
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.TransactionRequest
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{testBaseReq, "nubID"}, args{cliCommand, cliContext}, transactionRequest{cliCommand.ReadBaseReq(cliContext), cliCommand.ReadString(constants.NubID)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				NubID:   tt.fields.NubID,
			}
			got, err := transactionRequest.FromCLI(tt.args.cliCommand, tt.args.cliContext)
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

func Test_transactionRequest_FromJSON(t *testing.T) {
	_, _, _, _, fromAccAddress, testBaseReq := CreateTestInputForRequest(t)

	type fields struct {
		BaseReq rest.BaseReq
		NubID   string
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
		// TODO: Add test cases.
		{"+ve", fields{testBaseReq, "nubID"}, args{sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(message{fromAccAddress, base.NewStringID("nubID")}))}, transactionRequest{testBaseReq, "nubID"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				NubID:   tt.fields.NubID,
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
	_, _, _, _, _, testBaseReq := CreateTestInputForRequest(t)

	type fields struct {
		BaseReq rest.BaseReq
		NubID   string
	}
	tests := []struct {
		name   string
		fields fields
		want   rest.BaseReq
	}{
		// TODO: Add test cases.
		{"+ve", fields{testBaseReq, "nubID"}, testBaseReq},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				NubID:   tt.fields.NubID,
			}
			if got := transactionRequest.GetBaseReq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBaseReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	_, _, _, _, fromAccAddress, testBaseReq := CreateTestInputForRequest(t)

	type fields struct {
		BaseReq rest.BaseReq
		NubID   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    sdkTypes.Msg
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{testBaseReq, "nubID"}, message{fromAccAddress, base.NewStringID("nubID")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				NubID:   tt.fields.NubID,
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
	_, _, _, _, _, testBaseReq := CreateTestInputForRequest(t)

	type fields struct {
		BaseReq rest.BaseReq
		NubID   string
	}
	type args struct {
		codec *codec.Codec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"+ve", fields{testBaseReq, "nubID"}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				NubID:   tt.fields.NubID,
			}
			tr.RegisterCodec(tt.args.codec)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	_, _, _, _, _, testBaseReq := CreateTestInputForRequest(t)

	type fields struct {
		BaseReq rest.BaseReq
		NubID   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{testBaseReq, "nubID"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				NubID:   tt.fields.NubID,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
