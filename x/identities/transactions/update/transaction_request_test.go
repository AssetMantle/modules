// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package update

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
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

func createTestInput(t *testing.T) (*codec.LegacyAmino, helpers.CLICommand, client.Context, string, string, lists.PropertyList, lists.PropertyList, string, sdkTypes.AccAddress, helpers.CommonTransactionRequest) {
	var legacyAmino = baseHelpers.CodecPrototype().GetLegacyAmino()

	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromIdentityID, constants.IdentityID, constants.MutableMetaProperties, constants.MutableProperties})

	mutableMetaPropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutablePropertiesString := "defaultMutable1:S|defaultMutable1"
	mutableMetaProperties, err := baseLists.NewPropertyList().FromMetaPropertiesString(mutableMetaPropertiesString)
	require.Equal(t, nil, err)
	mutableProperties, err := baseLists.NewPropertyList().FromMetaPropertiesString(mutablePropertiesString)
	require.Equal(t, nil, err)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	commonTransactionRequest := helpers.PrototypeCommonTransactionRequest()

	return legacyAmino, cliCommand, client.Context{}.WithCodec(baseHelpers.CodecPrototype()), mutableMetaPropertiesString, mutablePropertiesString, mutableMetaProperties, mutableProperties, fromAddress, fromAccAddress, commonTransactionRequest
}

func Test_newTransactionRequest(t *testing.T) {
	_, _, _, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, commonTransactionRequest := createTestInput(t)
	type args struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		fromID                   string
		identityID               string
		mutableMetaProperties    string
		mutableProperties        string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve", args{commonTransactionRequest, "fromID", "identityID", mutableMetaPropertiesString, mutablePropertiesString}, transactionRequest{commonTransactionRequest, "fromID", "identityID", mutableMetaPropertiesString, mutablePropertiesString}},
		{"-ve with nil", args{}, transactionRequest{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.commonTransactionRequest, tt.args.fromID, tt.args.identityID, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
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
	_, cliCommand, context, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, commonTransactionRequest := createTestInput(t)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		IdentityID               string
		MutableMetaProperties    string
		MutableProperties        string
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
		{"+ve", fields{commonTransactionRequest, "fromID", "identityID", mutableMetaPropertiesString, mutablePropertiesString}, args{cliCommand: cliCommand, context: context}, transactionRequest{cliCommand.ReadCommonTransactionRequest(context), cliCommand.ReadString(constants.FromIdentityID), cliCommand.ReadString(constants.IdentityID), cliCommand.ReadString(constants.MutableMetaProperties), cliCommand.ReadString(constants.MutableProperties)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				IdentityID:               tt.fields.IdentityID,
				MutableMetaProperties:    tt.fields.MutableMetaProperties,
				MutableProperties:        tt.fields.MutableProperties,
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

func Test_transactionRequest_GetBaseReq(t *testing.T) {
	_, _, _, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, commonTransactionRequest := createTestInput(t)

	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		IdentityID               string
		MutableMetaProperties    string
		MutableProperties        string
	}
	tests := []struct {
		name   string
		fields fields
		want   helpers.CommonTransactionRequest
	}{
		{"+ve", fields{}, helpers.PrototypeCommonTransactionRequest()},
		{"+ve", fields{commonTransactionRequest, "fromID", "identityID", mutableMetaPropertiesString, mutablePropertiesString}, commonTransactionRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				IdentityID:               tt.fields.IdentityID,
				MutableMetaProperties:    tt.fields.MutableMetaProperties,
				MutableProperties:        tt.fields.MutableProperties,
			}
			if got := transactionRequest.GetCommonTransactionRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommonTransactionRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	_, _, _, mutableMetaPropertiesString, mutablePropertiesString, mutableMetaProperties, mutableProperties, _, fromAccAddress, commonTransactionRequest := createTestInput(t)
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		IdentityID               string
		MutableMetaProperties    string
		MutableProperties        string
	}
	tests := []struct {
		name    string
		fields  fields
		want    sdkTypes.Msg
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), testFromID.AsString(), mutableMetaPropertiesString, mutablePropertiesString}, NewMessage(fromAccAddress, testFromID, testFromID, mutableMetaProperties, mutableProperties.ScrubData()), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				IdentityID:               tt.fields.IdentityID,
				MutableMetaProperties:    tt.fields.MutableMetaProperties,
				MutableProperties:        tt.fields.MutableProperties,
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

func Test_transactionRequest_Validate(t *testing.T) {
	_, _, _, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, commonTransactionRequest := createTestInput(t)

	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		IdentityID               string
		MutableMetaProperties    string
		MutableProperties        string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, "fromID", "identityID", mutableMetaPropertiesString, mutablePropertiesString}, false},
		{"+ve with nil", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				IdentityID:               tt.fields.IdentityID,
				MutableMetaProperties:    tt.fields.MutableMetaProperties,
				MutableProperties:        tt.fields.MutableProperties,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
