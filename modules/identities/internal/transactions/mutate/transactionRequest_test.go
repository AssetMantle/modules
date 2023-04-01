// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mutate

import (
	"encoding/json"
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

func createTestInput(t *testing.T) (*codec.LegacyAmino, helpers.CLICommand, client.Context, string, string, lists.PropertyList, lists.PropertyList, string, sdkTypes.AccAddress, rest.BaseReq) {
	var legacyAmino = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()

	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromID, constants.IdentityID, constants.MutableMetaProperties, constants.MutableProperties})

	mutableMetaPropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutablePropertiesString := "defaultMutable1:S|defaultMutable1"
	mutableMetaProperties, err := utilities.ReadMetaPropertyList(mutableMetaPropertiesString)
	require.Equal(t, nil, err)
	mutableProperties, err := utilities.ReadMetaPropertyList(mutablePropertiesString)
	require.Equal(t, nil, err)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}

	return legacyAmino, cliCommand, baseHelpers.TestClientContext, mutableMetaPropertiesString, mutablePropertiesString, mutableMetaProperties, mutableProperties, fromAddress, fromAccAddress, testBaseReq
}

func Test_newTransactionRequest(t *testing.T) {
	_, _, _, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, testBaseReq := createTestInput(t)
	type args struct {
		baseReq               rest.BaseReq
		fromID                string
		identityID            string
		mutableMetaProperties string
		mutableProperties     string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve", args{testBaseReq, "fromID", "identityID", mutableMetaPropertiesString, mutablePropertiesString}, transactionRequest{testBaseReq, "fromID", "identityID", mutableMetaPropertiesString, mutablePropertiesString}},
		{"-ve with nil", args{}, transactionRequest{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.baseReq, tt.args.fromID, tt.args.identityID, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
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
	_, cliCommand, context, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, testBaseReq := createTestInput(t)
	type fields struct {
		BaseReq               rest.BaseReq
		FromID                string
		IdentityID            string
		MutableMetaProperties string
		MutableProperties     string
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
		{"+ve", fields{testBaseReq, "fromID", "identityID", mutableMetaPropertiesString, mutablePropertiesString}, args{cliCommand: cliCommand, context: context}, transactionRequest{cliCommand.ReadBaseReq(context), cliCommand.ReadString(constants.FromID), cliCommand.ReadString(constants.IdentityID), cliCommand.ReadString(constants.MutableMetaProperties), cliCommand.ReadString(constants.MutableProperties)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:               tt.fields.BaseReq,
				FromID:                tt.fields.FromID,
				IdentityID:            tt.fields.IdentityID,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
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
	_, _, _, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, testBaseReq := createTestInput(t)
	jsonMessage, _ := json.Marshal(transactionRequest{testBaseReq, "fromID", "identityID", mutableMetaPropertiesString, mutablePropertiesString})

	type fields struct {
		BaseReq               rest.BaseReq
		FromID                string
		IdentityID            string
		MutableMetaProperties string
		MutableProperties     string
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
		{"+ve", fields{testBaseReq, "fromID", "identityID", mutableMetaPropertiesString, mutablePropertiesString}, args{jsonMessage}, transactionRequest{testBaseReq, "fromID", "identityID", mutableMetaPropertiesString, mutablePropertiesString}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:               tt.fields.BaseReq,
				FromID:                tt.fields.FromID,
				IdentityID:            tt.fields.IdentityID,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
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
	_, _, context, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, testBaseReq := createTestInput(t)

	type fields struct {
		BaseReq               rest.BaseReq
		FromID                string
		IdentityID            string
		MutableMetaProperties string
		MutableProperties     string
	}
	tests := []struct {
		name   string
		fields fields
		want   rest.BaseReq
	}{
		{"+ve", fields{}, rest.BaseReq{From: context.GetFromAddress().String(), ChainID: context.ChainID, Simulate: context.Simulate}},
		{"+ve", fields{testBaseReq, "fromID", "identityID", mutableMetaPropertiesString, mutablePropertiesString}, testBaseReq},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:               tt.fields.BaseReq,
				FromID:                tt.fields.FromID,
				IdentityID:            tt.fields.IdentityID,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
			}
			if got := transactionRequest.GetBaseReq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBaseReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	_, _, _, mutableMetaPropertiesString, mutablePropertiesString, mutableMetaProperties, mutableProperties, _, fromAccAddress, testBaseReq := createTestInput(t)
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)
	type fields struct {
		BaseReq               rest.BaseReq
		FromID                string
		IdentityID            string
		MutableMetaProperties string
		MutableProperties     string
	}
	tests := []struct {
		name    string
		fields  fields
		want    sdkTypes.Msg
		wantErr bool
	}{
		{"+ve", fields{testBaseReq, testFromID.AsString(), testFromID.AsString(), mutableMetaPropertiesString, mutablePropertiesString}, newMessage(fromAccAddress, testFromID, testFromID, mutableMetaProperties, mutableProperties.ScrubData()), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:               tt.fields.BaseReq,
				FromID:                tt.fields.FromID,
				IdentityID:            tt.fields.IdentityID,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
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
	_, _, _, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, testBaseReq := createTestInput(t)

	type fields struct {
		BaseReq               rest.BaseReq
		FromID                string
		IdentityID            string
		MutableMetaProperties string
		MutableProperties     string
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testBaseReq, "fromID", "identityID", mutableMetaPropertiesString, mutablePropertiesString}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				BaseReq:               tt.fields.BaseReq,
				FromID:                tt.fields.FromID,
				IdentityID:            tt.fields.IdentityID,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
			}
			tr.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	_, _, _, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, testBaseReq := createTestInput(t)

	type fields struct {
		BaseReq               rest.BaseReq
		FromID                string
		IdentityID            string
		MutableMetaProperties string
		MutableProperties     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{testBaseReq, "fromID", "identityID", mutableMetaPropertiesString, mutablePropertiesString}, false},
		{"+ve with nil", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:               tt.fields.BaseReq,
				FromID:                tt.fields.FromID,
				IdentityID:            tt.fields.IdentityID,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
