// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unprovision

import (
	"encoding/json"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/go/data/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
)

func createInputForMessage(t *testing.T) (*baseIDs.IdentityID, string, types.AccAddress, types.Msg, rest.BaseReq) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testIdentityID := baseIDs.NewIdentityID(testClassificationID, immutables)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := types.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	toAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	toAccAddress, err := types.AccAddressFromBech32(toAddress)
	require.Nil(t, err)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: types.NewCoins()}

	testMessage := newMessage(fromAccAddress, toAccAddress, testIdentityID)

	return testIdentityID.(*baseIDs.IdentityID), toAddress, toAccAddress, testMessage, testBaseReq
}

func Test_newTransactionRequest(t *testing.T) {
	testIdentityID, toAddress, _, _, testBaseReq := createInputForMessage(t)
	type args struct {
		baseReq    rest.BaseReq
		to         string
		identityID string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve wit nil", args{}, transactionRequest{}},
		{"+ve", args{testBaseReq, toAddress, testIdentityID.AsString()}, transactionRequest{testBaseReq, toAddress, testIdentityID.AsString()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.baseReq, tt.args.to, tt.args.identityID); !reflect.DeepEqual(got, tt.want) {
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
	var legacyAmino = baseHelpers.CodecPrototype().GetLegacyAmino()

	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.To, constants.IdentityID})

	testIdentityID, toAddress, _, _, testBaseReq := createInputForMessage(t)
	type fields struct {
		BaseReq    rest.BaseReq
		To         string
		IdentityID string
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
		{"+ve", fields{testBaseReq, toAddress, testIdentityID.AsString()}, args{cliCommand, baseHelpers.TestClientContext}, transactionRequest{cliCommand.ReadBaseReq(baseHelpers.TestClientContext), cliCommand.ReadString(constants.To), cliCommand.ReadString(constants.IdentityID)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:    tt.fields.BaseReq,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
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
	testIdentityID, toAddress, toAccAddress, _, testBaseReq := createInputForMessage(t)
	type fields struct {
		BaseReq    rest.BaseReq
		To         string
		IdentityID string
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
		{"+ve", fields{testBaseReq, toAddress, testIdentityID.AsString()}, args{types.MustSortJSON(transaction.RegisterLegacyAminoCodec(messagePrototype).MustMarshalJSON(&Message{types.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c").String(), toAccAddress.String(), testIdentityID}))}, transactionRequest{testBaseReq, toAddress, testIdentityID.AsString()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:    tt.fields.BaseReq,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
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
	testIdentityID, toAddress, _, _, testBaseReq := createInputForMessage(t)
	type fields struct {
		BaseReq    rest.BaseReq
		To         string
		IdentityID string
	}
	tests := []struct {
		name   string
		fields fields
		want   rest.BaseReq
	}{
		{"+ve", fields{testBaseReq, toAddress, testIdentityID.AsString()}, testBaseReq},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:    tt.fields.BaseReq,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			if got := transactionRequest.GetBaseReq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBaseReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	testIdentityID, toAddress, toAccAddress, _, testBaseReq := createInputForMessage(t)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := types.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	type fields struct {
		BaseReq    rest.BaseReq
		To         string
		IdentityID string
	}
	tests := []struct {
		name    string
		fields  fields
		want    types.Msg
		wantErr bool
	}{
		{"+ve", fields{testBaseReq, toAddress, testIdentityID.AsString()}, &Message{fromAccAddress.String(), toAccAddress.String(), testIdentityID}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:    tt.fields.BaseReq,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
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
	testIdentityID, toAddress, _, _, testBaseReq := createInputForMessage(t)
	type fields struct {
		BaseReq    rest.BaseReq
		To         string
		IdentityID string
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve wit nil", fields{}, args{codec.NewLegacyAmino()}},
		{"+ve", fields{testBaseReq, toAddress, testIdentityID.AsString()}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				BaseReq:    tt.fields.BaseReq,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			tr.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	testIdentityID, toAddress, _, _, testBaseReq := createInputForMessage(t)
	type fields struct {
		BaseReq    rest.BaseReq
		To         string
		IdentityID string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve wit nil", fields{}, true},
		{"+ve", fields{testBaseReq, toAddress, testIdentityID.AsString()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:    tt.fields.BaseReq,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
