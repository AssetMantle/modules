// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package provision

import (
	"encoding/json"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
)

func createTestInput(t *testing.T) (helpers.CommonTransactionRequest, string, *baseIDs.IdentityID) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	commonTransactionRequest := helpers.PrototypeCommonTransactionRequest().SetFrom(fromAddress)
	testToAddress := "cosmos1vx8knpllrj7n963p9ttd80w47kpacrhuts497x"
	return commonTransactionRequest, testToAddress, testFromID.(*baseIDs.IdentityID)
}

func Test_newTransactionRequest(t *testing.T) {
	commonTransactionRequest, testToAddress, testFromID := createTestInput(t)
	type args struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		to                       string
		identityID               string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve with nil", args{}, transactionRequest{}},
		{"+ve", args{commonTransactionRequest, testToAddress, testFromID.AsString()}, transactionRequest{commonTransactionRequest, testToAddress, testFromID.AsString()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.commonTransactionRequest, tt.args.to, tt.args.identityID); !reflect.DeepEqual(got, tt.want) {
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
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.To, constants.IdentityID})

	commonTransactionRequest, testToAddress, testFromID := createTestInput(t)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		To                       string
		IdentityID               string
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
		{"+ve", fields{commonTransactionRequest, testToAddress, testFromID.AsString()}, args{cliCommand: cliCommand, context: client.Context{}.WithCodec(baseHelpers.CodecPrototype())}, transactionRequest{cliCommand.ReadCommonTransactionRequest(client.Context{}.WithCodec(baseHelpers.CodecPrototype())), cliCommand.ReadString(constants.To), cliCommand.ReadString(constants.IdentityID)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				To:                       tt.fields.To,
				IdentityID:               tt.fields.IdentityID,
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
	commonTransactionRequest, testToAddress, testFromID := createTestInput(t)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		To                       string
		IdentityID               string
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
		{"+ve", fields{commonTransactionRequest, testToAddress, testFromID.AsString()}, args{sdkTypes.MustSortJSON(baseHelpers.CodecPrototype().MustMarshalJSON(&Message{sdkTypes.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c").String(), sdkTypes.AccAddress(testToAddress).String(), testFromID}))}, transactionRequest{commonTransactionRequest, testToAddress, testFromID.AsString()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				To:                       tt.fields.To,
				IdentityID:               tt.fields.IdentityID,
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
	commonTransactionRequest, testToAddress, testFromID := createTestInput(t)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		To                       string
		IdentityID               string
	}
	tests := []struct {
		name   string
		fields fields
		want   helpers.CommonTransactionRequest
	}{
		{"+ve", fields{commonTransactionRequest, testToAddress, testFromID.AsString()}, commonTransactionRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				To:                       tt.fields.To,
				IdentityID:               tt.fields.IdentityID,
			}
			if got := transactionRequest.GetCommonTransactionRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommonTransactionRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	commonTransactionRequest, testToAddress, testFromID := createTestInput(t)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	toAccAddress, err := sdkTypes.AccAddressFromBech32(testToAddress)
	require.Nil(t, err)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		To                       string
		IdentityID               string
	}
	tests := []struct {
		name    string
		fields  fields
		want    sdkTypes.Msg
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, testToAddress, testFromID.AsString()}, &Message{fromAccAddress.String(), toAccAddress.String(), testFromID}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				To:                       tt.fields.To,
				IdentityID:               tt.fields.IdentityID,
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
	commonTransactionRequest, testToAddress, testFromID := createTestInput(t)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		To                       string
		IdentityID               string
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve with nil", fields{}, args{codec.NewLegacyAmino()}},
		{"+ve", fields{commonTransactionRequest, testToAddress, testFromID.AsString()}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				To:                       tt.fields.To,
				IdentityID:               tt.fields.IdentityID,
			}
			tr.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	commonTransactionRequest, testToAddress, testFromID := createTestInput(t)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		To                       string
		IdentityID               string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve with nil", fields{}, true},
		{"+ve", fields{commonTransactionRequest, testToAddress, testFromID.AsString()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				To:                       tt.fields.To,
				IdentityID:               tt.fields.IdentityID,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
