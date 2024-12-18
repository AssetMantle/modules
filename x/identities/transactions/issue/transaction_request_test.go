// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package issue

import (
	"fmt"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
)

func createTestInputForRequest(t *testing.T) (helpers.CLICommand, client.Context, string, string, string, string, lists.PropertyList, lists.PropertyList, lists.PropertyList, lists.PropertyList, string, sdkTypes.AccAddress, string, sdkTypes.AccAddress, helpers.CommonTransactionRequest) {

	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromIdentityID, constants.To, constants.ClassificationID, constants.ImmutableMetaProperties, constants.ImmutableProperties, constants.MutableMetaProperties, constants.MutableProperties})

	immutableMetaPropertiesString := "defaultImmutableMeta1:S|defaultImmutableMeta1"
	immutablePropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutableMetaPropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutablePropertiesString := "defaultMutable1:S|defaultMutable1"
	immutableMetaProperties, err := baseLists.NewPropertyList().FromMetaPropertiesString(immutableMetaPropertiesString)
	require.Equal(t, nil, err)

	var immutableProperties lists.PropertyList
	immutableProperties, err = baseLists.NewPropertyList().FromMetaPropertiesString(immutablePropertiesString)
	require.Equal(t, nil, err)

	var mutableMetaProperties lists.PropertyList
	mutableMetaProperties, err = baseLists.NewPropertyList().FromMetaPropertiesString(mutableMetaPropertiesString)
	require.Equal(t, nil, err)

	var mutableProperties lists.PropertyList
	mutableProperties, err = baseLists.NewPropertyList().FromMetaPropertiesString(mutablePropertiesString)
	require.Equal(t, nil, err)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	var fromAccAddress sdkTypes.AccAddress
	fromAccAddress, err = sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	toAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	var toAccAddress sdkTypes.AccAddress
	toAccAddress, err = sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	commonTransactionRequest := helpers.PrototypeCommonTransactionRequest()

	return cliCommand, client.Context{}.WithCodec(baseHelpers.CodecPrototype()), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties, fromAddress, fromAccAddress, toAddress, toAccAddress, commonTransactionRequest
}

func Test_newTransactionRequest(t *testing.T) {
	_, _, immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, _, _, toAddress, _, commonTransactionRequest := createTestInputForRequest(t)
	type args struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		to                       string
		fromID                   string
		classificationID         string
		immutableMetaProperties  string
		immutableProperties      string
		mutableMetaProperties    string
		mutableProperties        string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{

		{"+ve", args{commonTransactionRequest, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, transactionRequest{commonTransactionRequest, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.commonTransactionRequest, tt.args.fromID, tt.args.classificationID, tt.args.immutableMetaProperties, tt.args.immutableProperties, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
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
	cliCommand, context, immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, _, _, toAddress, _, commonTransactionRequest := createTestInputForRequest(t)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		To                       string
		FromID                   string
		ClassificationID         string
		ImmutableMetaProperties  string
		ImmutableProperties      string
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
		{"+ve", fields{commonTransactionRequest, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, args{cliCommand, client.Context{}.WithCodec(baseHelpers.CodecPrototype())}, transactionRequest{cliCommand.ReadCommonTransactionRequest(context), cliCommand.ReadString(constants.FromIdentityID), cliCommand.ReadString(constants.ClassificationID), cliCommand.ReadString(constants.ImmutableMetaProperties), cliCommand.ReadString(constants.ImmutableProperties), cliCommand.ReadString(constants.MutableMetaProperties), cliCommand.ReadString(constants.MutableProperties)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ClassificationID:         tt.fields.ClassificationID,
				ImmutableMetaProperties:  tt.fields.ImmutableMetaProperties,
				ImmutableProperties:      tt.fields.ImmutableProperties,
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
	_, _, immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, _, _, toAddress, _, commonTransactionRequest := createTestInputForRequest(t)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		To                       string
		FromID                   string
		ClassificationID         string
		ImmutableMetaProperties  string
		ImmutableProperties      string
		MutableMetaProperties    string
		MutableProperties        string
	}
	tests := []struct {
		name   string
		fields fields
		want   helpers.CommonTransactionRequest
	}{

		{"+ve", fields{commonTransactionRequest, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, commonTransactionRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ClassificationID:         tt.fields.ClassificationID,
				ImmutableMetaProperties:  tt.fields.ImmutableMetaProperties,
				ImmutableProperties:      tt.fields.ImmutableProperties,
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
	_, _, immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties, _, fromAccAddress, toAddress, _, commonTransactionRequest := createTestInputForRequest(t)
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		To                       string
		FromID                   string
		ClassificationID         string
		ImmutableMetaProperties  string
		ImmutableProperties      string
		MutableMetaProperties    string
		MutableProperties        string
	}
	tests := []struct {
		name    string
		fields  fields
		want    sdkTypes.Msg
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, toAddress, "9UNIA3_tulK2vRE0nSmsHKNzhDxoCBHI4z8XXfLO1FM=", "9UNIA3_tulK2vRE0nSmsHKNzhDxoCBHI4z8XXfLO1FM=", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, NewMessage(fromAccAddress, testFromID, testClassificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties), false}, // TODO: issue==> getting MetaDataError that is not expected
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ClassificationID:         tt.fields.ClassificationID,
				ImmutableMetaProperties:  tt.fields.ImmutableMetaProperties,
				ImmutableProperties:      tt.fields.ImmutableProperties,
				MutableMetaProperties:    tt.fields.MutableMetaProperties,
				MutableProperties:        tt.fields.MutableProperties,
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

func Test_transactionRequest_Validate(t *testing.T) {
	_, _, immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString, _, _, _, _, _, _, toAddress, _, commonTransactionRequest := createTestInputForRequest(t)
	commonTransactionRequest1 := helpers.PrototypeCommonTransactionRequest()

	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		To                       string
		FromID                   string
		ClassificationID         string
		ImmutableMetaProperties  string
		ImmutableProperties      string
		MutableMetaProperties    string
		MutableProperties        string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, false},
		{"+ve", fields{commonTransactionRequest1, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, false},
		{"-ve", fields{ClassificationID: ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ClassificationID:         tt.fields.ClassificationID,
				ImmutableMetaProperties:  tt.fields.ImmutableMetaProperties,
				ImmutableProperties:      tt.fields.ImmutableProperties,
				MutableMetaProperties:    tt.fields.MutableMetaProperties,
				MutableProperties:        tt.fields.MutableProperties,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
