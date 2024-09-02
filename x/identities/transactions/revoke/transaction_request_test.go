// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	"encoding/json"
	"fmt"
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

func createTestInput(t *testing.T) (helpers.CommonTransactionRequest, string, *baseIDs.IdentityID, *baseIDs.IdentityID, *baseIDs.ClassificationID) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)
	testToID := baseIDs.NewIdentityID(testClassificationID, immutables)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	commonTransactionRequest := helpers.PrototypeCommonTransactionRequest().SetFrom(fromAddress)
	testToAddress := "cosmos1vx8knpllrj7n963p9ttd80w47kpacrhuts497x"
	return commonTransactionRequest, testToAddress, testFromID.(*baseIDs.IdentityID), testToID.(*baseIDs.IdentityID), testClassificationID.(*baseIDs.ClassificationID)
}

func Test_newTransactionRequest(t *testing.T) {
	commonTransactionRequest, _, testFromID, testToID, testClassificationID := createTestInput(t)
	type args struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		fromID                   string
		toID                     string
		classificationID         string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve with nil", args{}, transactionRequest{}},
		{"+ve", args{commonTransactionRequest, testFromID.AsString(), testToID.AsString(), testClassificationID.AsString()}, transactionRequest{commonTransactionRequest, testFromID.AsString(), testToID.AsString(), testClassificationID.AsString()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.commonTransactionRequest, tt.args.fromID, tt.args.toID, tt.args.classificationID); !reflect.DeepEqual(got, tt.want) {
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
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromIdentityID, constants.ToIdentityID, constants.ClassificationID})

	commonTransactionRequest, _, testFromID, testToID, testClassificationID := createTestInput(t)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ToID                     string
		ClassificationID         string
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
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), testToID.AsString(), testClassificationID.AsString()}, args{cliCommand, client.Context{}.WithCodec(baseHelpers.CodecPrototype())}, transactionRequest{cliCommand.ReadCommonTransactionRequest(client.Context{}.WithCodec(baseHelpers.CodecPrototype())), cliCommand.ReadString(constants.FromIdentityID), cliCommand.ReadString(constants.ToIdentityID), cliCommand.ReadString(constants.ClassificationID)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
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
	commonTransactionRequest, fromAddress, testFromID, testToID, testClassificationID := createTestInput(t)
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ToID                     string
		ClassificationID         string
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
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), testToID.AsString(), testClassificationID.AsString()}, args{sdkTypes.MustSortJSON(baseHelpers.CodecPrototype().MustMarshalJSON(&Message{fromAccAddress.String(), testFromID, testToID, testClassificationID}))}, transactionRequest{commonTransactionRequest, testFromID.AsString(), testToID.AsString(), testClassificationID.AsString()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
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
	commonTransactionRequest, _, testFromID, testToID, testClassificationID := createTestInput(t)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ToID                     string
		ClassificationID         string
	}
	tests := []struct {
		name   string
		fields fields
		want   helpers.CommonTransactionRequest
	}{
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), testToID.AsString(), testClassificationID.AsString()}, commonTransactionRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
			}
			if got := transactionRequest.GetCommonTransactionRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommonTransactionRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	commonTransactionRequest, fromAccAddress, testFromID, testToID, testClassificationID := createTestInput(t)
	testFromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAccAddress)
	require.Nil(t, err)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ToID                     string
		ClassificationID         string
	}
	tests := []struct {
		name    string
		fields  fields
		want    sdkTypes.Msg
		wantErr bool
	}{
		// TODO: Type & Data same but Not matching
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), testToID.AsString(), testClassificationID.AsString()}, &Message{testFromAccAddress.String(), testFromID, testToID, testClassificationID}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
			}
			got, err := transactionRequest.MakeMsg()
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("MakeMsg() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_RegisterCodec(t *testing.T) {
	commonTransactionRequest, _, testFromID, testToID, testClassificationID := createTestInput(t)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ToID                     string
		ClassificationID         string
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
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), testToID.AsString(), testClassificationID.AsString()}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
			}
			tr.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	commonTransactionRequest, _, testFromID, testToID, testClassificationID := createTestInput(t)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ToID                     string
		ClassificationID         string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve with nil", fields{}, true},
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), testToID.AsString(), testClassificationID.AsString()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
