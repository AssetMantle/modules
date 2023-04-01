// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package issue

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

func createTestInputForRequest(t *testing.T) (*codec.LegacyAmino, helpers.CLICommand, client.Context, string, string, string, string, lists.PropertyList, lists.PropertyList, lists.PropertyList, lists.PropertyList, string, sdkTypes.AccAddress, string, sdkTypes.AccAddress, rest.BaseReq) {
	var legacyAmino = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()

	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromID, constants.To, constants.ClassificationID, constants.ImmutableMetaProperties, constants.ImmutableProperties, constants.MutableMetaProperties, constants.MutableProperties})

	immutableMetaPropertiesString := "defaultImmutableMeta1:S|defaultImmutableMeta1"
	immutablePropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutableMetaPropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutablePropertiesString := "defaultMutable1:S|defaultMutable1"
	immutableMetaProperties, err := utilities.ReadMetaPropertyList(immutableMetaPropertiesString)
	require.Equal(t, nil, err)

	var immutableProperties lists.PropertyList
	immutableProperties, err = utilities.ReadMetaPropertyList(immutablePropertiesString)
	require.Equal(t, nil, err)

	var mutableMetaProperties lists.PropertyList
	mutableMetaProperties, err = utilities.ReadMetaPropertyList(mutableMetaPropertiesString)
	require.Equal(t, nil, err)

	var mutableProperties lists.PropertyList
	mutableProperties, err = utilities.ReadMetaPropertyList(mutablePropertiesString)
	require.Equal(t, nil, err)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	var fromAccAddress sdkTypes.AccAddress
	fromAccAddress, err = sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	toAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	var toAccAddress sdkTypes.AccAddress
	toAccAddress, err = sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}

	return legacyAmino, cliCommand, baseHelpers.TestClientContext, immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties, fromAddress, fromAccAddress, toAddress, toAccAddress, testBaseReq
}

func Test_newTransactionRequest(t *testing.T) {
	_, _, _, immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, _, _, toAddress, _, testBaseReq := createTestInputForRequest(t)
	type args struct {
		baseReq                 rest.BaseReq
		to                      string
		fromID                  string
		classificationID        string
		immutableMetaProperties string
		immutableProperties     string
		mutableMetaProperties   string
		mutableProperties       string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{

		{"+ve", args{testBaseReq, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, transactionRequest{testBaseReq, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.baseReq, tt.args.to, tt.args.fromID, tt.args.classificationID, tt.args.immutableMetaProperties, tt.args.immutableProperties, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
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
	_, cliCommand, context, immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, _, _, toAddress, _, testBaseReq := createTestInputForRequest(t)
	type fields struct {
		BaseReq                 rest.BaseReq
		To                      string
		FromID                  string
		ClassificationID        string
		ImmutableMetaProperties string
		ImmutableProperties     string
		MutableMetaProperties   string
		MutableProperties       string
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
		{"+ve", fields{testBaseReq, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, args{cliCommand, baseHelpers.TestClientContext}, transactionRequest{cliCommand.ReadBaseReq(context), cliCommand.ReadString(constants.To), cliCommand.ReadString(constants.FromID), cliCommand.ReadString(constants.ClassificationID), cliCommand.ReadString(constants.ImmutableMetaProperties), cliCommand.ReadString(constants.ImmutableProperties), cliCommand.ReadString(constants.MutableMetaProperties), cliCommand.ReadString(constants.MutableProperties)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				To:                      tt.fields.To,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
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
	_, _, _, immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, _, _, toAddress, _, testBaseReq := createTestInputForRequest(t)
	jsonMessage, err := json.Marshal(newTransactionRequest(testBaseReq, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString))
	require.Equal(t, nil, err)
	jsonMessage1, err := json.Marshal(newTransactionRequest(testBaseReq, toAddress, "", "", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString))
	require.Equal(t, nil, err)
	type fields struct {
		BaseReq                 rest.BaseReq
		To                      string
		FromID                  string
		ClassificationID        string
		ImmutableMetaProperties string
		ImmutableProperties     string
		MutableMetaProperties   string
		MutableProperties       string
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

		{"+ve", fields{testBaseReq, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, args{jsonMessage}, transactionRequest{testBaseReq, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, false},
		{"+ve", fields{testBaseReq, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, args{jsonMessage1}, transactionRequest{testBaseReq, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				To:                      tt.fields.To,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
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
	_, _, _, immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, _, _, toAddress, _, testBaseReq := createTestInputForRequest(t)
	type fields struct {
		BaseReq                 rest.BaseReq
		To                      string
		FromID                  string
		ClassificationID        string
		ImmutableMetaProperties string
		ImmutableProperties     string
		MutableMetaProperties   string
		MutableProperties       string
	}
	tests := []struct {
		name   string
		fields fields
		want   rest.BaseReq
	}{

		{"+ve", fields{testBaseReq, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, testBaseReq},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				To:                      tt.fields.To,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
			}
			if got := transactionRequest.GetBaseReq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBaseReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	_, _, _, immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties, _, fromAccAddress, toAddress, toAccAddress, testBaseReq := createTestInputForRequest(t)
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)
	type fields struct {
		BaseReq                 rest.BaseReq
		To                      string
		FromID                  string
		ClassificationID        string
		ImmutableMetaProperties string
		ImmutableProperties     string
		MutableMetaProperties   string
		MutableProperties       string
	}
	tests := []struct {
		name    string
		fields  fields
		want    sdkTypes.Msg
		wantErr bool
	}{
		{"+ve", fields{testBaseReq, toAddress, "9UNIA3_tulK2vRE0nSmsHKNzhDxoCBHI4z8XXfLO1FM=", "9UNIA3_tulK2vRE0nSmsHKNzhDxoCBHI4z8XXfLO1FM=", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, newMessage(fromAccAddress, toAccAddress, testFromID, testClassificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties), false}, // TODO: issue==> getting MetaDataError that is not expected
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				To:                      tt.fields.To,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
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
	_, _, _, immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, _, _, toAddress, _, testBaseReq := createTestInputForRequest(t)
	type fields struct {
		BaseReq                 rest.BaseReq
		To                      string
		FromID                  string
		ClassificationID        string
		ImmutableMetaProperties string
		ImmutableProperties     string
		MutableMetaProperties   string
		MutableProperties       string
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testBaseReq, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				To:                      tt.fields.To,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
			}
			tr.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	_, _, _, immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, _, _, toAddress, _, testBaseReq := createTestInputForRequest(t)
	testBaseReq1 := rest.BaseReq{}

	type fields struct {
		BaseReq                 rest.BaseReq
		To                      string
		FromID                  string
		ClassificationID        string
		ImmutableMetaProperties string
		ImmutableProperties     string
		MutableMetaProperties   string
		MutableProperties       string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{testBaseReq, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, false},
		{"+ve", fields{testBaseReq1, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, false},
		{"-ve", fields{ClassificationID: ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				To:                      tt.fields.To,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
