// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unprovision

import (
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
)

func createInputForMessage(t *testing.T) (*baseIDs.IdentityID, string, types.AccAddress, types.Msg, helpers.CommonTransactionRequest) {
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

	commonTransactionRequest := helpers.PrototypeCommonTransactionRequest()

	testMessage := NewMessage(fromAccAddress, toAccAddress, testIdentityID)

	return testIdentityID.(*baseIDs.IdentityID), toAddress, toAccAddress, testMessage, commonTransactionRequest
}

func Test_newTransactionRequest(t *testing.T) {
	testIdentityID, toAddress, _, _, commonTransactionRequest := createInputForMessage(t)
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
		{"nil inputs", args{}, transactionRequest{}},
		{"valid", args{commonTransactionRequest, toAddress, testIdentityID.AsString()}, transactionRequest{commonTransactionRequest, toAddress, testIdentityID.AsString()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newTransactionRequest(tt.args.commonTransactionRequest, tt.args.to, tt.args.identityID)
			assert.Equal(t, tt.want, got, "newTransactionRequest()")
		})
	}
}

func Test_requestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.TransactionRequest
	}{
		{"valid", transactionRequest{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := requestPrototype()
			assert.Equal(t, tt.want, got, "requestPrototype()")
		})
	}
}

func Test_transactionRequest_FromCLI(t *testing.T) {
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.To, constants.IdentityID})

	testIdentityID, toAddress, _, _, commonTransactionRequest := createInputForMessage(t)
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
		{"valid", fields{commonTransactionRequest, toAddress, testIdentityID.AsString()}, args{cliCommand, client.Context{}.WithCodec(baseHelpers.CodecPrototype())}, transactionRequest{cliCommand.ReadCommonTransactionRequest(client.Context{}.WithCodec(baseHelpers.CodecPrototype())), cliCommand.ReadString(constants.To), cliCommand.ReadString(constants.IdentityID)}, false},
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
			assert.Equal(t, tt.want, got, "FromCLI() got")
		})
	}
}

func Test_transactionRequest_GetBaseReq(t *testing.T) {
	testIdentityID, toAddress, _, _, commonTransactionRequest := createInputForMessage(t)
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
		{"valid", fields{commonTransactionRequest, toAddress, testIdentityID.AsString()}, commonTransactionRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				To:                       tt.fields.To,
				IdentityID:               tt.fields.IdentityID,
			}
			got := transactionRequest.GetCommonTransactionRequest()
			assert.Equal(t, tt.want, got, "GetCommonTransactionRequest()")
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	testIdentityID, toAddress, _, _, commonTransactionRequest := createInputForMessage(t)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	_, err := types.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		To                       string
		IdentityID               string
	}
	tests := []struct {
		name    string
		fields  fields
		want    types.Msg
		wantErr bool
	}{
		{"valid", fields{commonTransactionRequest, toAddress, testIdentityID.AsString()}, nil, true},
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
			assert.Equal(t, tt.want, got, "MakeMsg() got")
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	testIdentityID, toAddress, _, _, commonTransactionRequest := createInputForMessage(t)
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
		{"nil inputs", fields{}, true},
		{"valid", fields{commonTransactionRequest, toAddress, testIdentityID.AsString()}, true},
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
