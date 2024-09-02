// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package name

import (
	"encoding/json"
	"reflect"
	"testing"

	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
)

func CreateTestInputForRequest(t *testing.T) (*codec.LegacyAmino, helpers.CLICommand, client.Context, string, sdkTypes.AccAddress, helpers.CommonTransactionRequest) {
	var legacyAmino = baseHelpers.CodecPrototype().GetLegacyAmino()

	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.Name})

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	commonTransactionRequest := helpers.PrototypeCommonTransactionRequest()

	return legacyAmino, cliCommand, client.Context{}.WithCodec(baseHelpers.CodecPrototype()), fromAddress, fromAccAddress, commonTransactionRequest
}
func Test_newTransactionRequest(t *testing.T) {
	_, _, _, _, _, commonTransactionRequest := CreateTestInputForRequest(t)
	type args struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		nubID                    string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve", args{commonTransactionRequest, "nubID"}, transactionRequest{commonTransactionRequest, "nubID"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.commonTransactionRequest, tt.args.nubID); !reflect.DeepEqual(got, tt.want) {
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
	_, cliCommand, context, _, _, commonTransactionRequest := CreateTestInputForRequest(t)

	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		NubID                    string
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
		{"+ve", fields{commonTransactionRequest, "nubID"}, args{cliCommand, client.Context{}.WithCodec(baseHelpers.CodecPrototype())}, transactionRequest{cliCommand.ReadCommonTransactionRequest(context), cliCommand.ReadString(constants.Name)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				Name:                     tt.fields.NubID,
			}
			got, err := transactionRequest.FromCLI(tt.args.cliCommand, tt.args.context)
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
	_, _, _, _, fromAccAddress, commonTransactionRequest := CreateTestInputForRequest(t)

	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		NubID                    string
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
		{"+ve", fields{commonTransactionRequest, "nubID"}, args{sdkTypes.MustSortJSON(baseHelpers.CodecPrototype().MustMarshalJSON(&Message{fromAccAddress.String(), baseIDs.NewStringID("nubID").(*baseIDs.StringID)}))}, transactionRequest{commonTransactionRequest, "nubID"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				Name:                     tt.fields.NubID,
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
	_, _, _, _, _, commonTransactionRequest := CreateTestInputForRequest(t)

	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		NubID                    string
	}
	tests := []struct {
		name   string
		fields fields
		want   helpers.CommonTransactionRequest
	}{
		{"+ve", fields{commonTransactionRequest, "nubID"}, commonTransactionRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				Name:                     tt.fields.NubID,
			}
			if got := transactionRequest.GetCommonTransactionRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommonTransactionRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	_, _, _, _, fromAccAddress, commonTransactionRequest := CreateTestInputForRequest(t)

	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		NubID                    string
	}
	tests := []struct {
		name    string
		fields  fields
		want    sdkTypes.Msg
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, "nubID"}, &Message{fromAccAddress.String(), baseIDs.NewStringID("nubID").(*baseIDs.StringID)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				Name:                     tt.fields.NubID,
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
	_, _, _, _, _, commonTransactionRequest := CreateTestInputForRequest(t)

	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		NubID                    string
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{commonTransactionRequest, "nubID"}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				Name:                     tt.fields.NubID,
			}
			tr.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	_, _, _, _, _, commonTransactionRequest := CreateTestInputForRequest(t)

	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		NubID                    string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, "nubID"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				Name:                     tt.fields.NubID,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
