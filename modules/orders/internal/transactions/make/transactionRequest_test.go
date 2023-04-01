// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/types/base"
)

var (
	fromAddress                   = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, _             = sdkTypes.AccAddressFromBech32(fromAddress)
	immutableMetaPropertiesString = "defaultImmutableMeta1:S|defaultImmutableMeta1"
	immutablePropertiesString     = "defaultMutableMeta1:S|defaultMutableMeta1"
	mutableMetaPropertiesString   = "defaultMutableMeta1:S|defaultMutableMeta1"
	mutablePropertiesString       = "defaultMutable1:S|defaultMutable1"
	immutableMetaProperties, _    = utilities.ReadMetaPropertyList(immutableMetaPropertiesString)
	immutableProperties, _        = utilities.ReadMetaPropertyList(immutablePropertiesString)
	mutableMetaProperties, _      = utilities.ReadMetaPropertyList(mutableMetaPropertiesString)
	mutableProperties, _          = utilities.ReadMetaPropertyList(mutablePropertiesString)
	immutables                    = baseQualified.NewImmutables(immutableProperties)
	mutables                      = baseQualified.NewMutables(mutableProperties)
	testClassificationID          = baseIDs.NewClassificationID(immutables, mutables).(*baseIDs.ClassificationID)
	testFromID                    = baseIDs.NewIdentityID(testClassificationID, immutables).(*baseIDs.IdentityID)
	makerOwnableID                = baseIDs.NewCoinID(baseIDs.NewStringID("makerownableid")).ToAnyOwnableID().(*baseIDs.AnyOwnableID)
	takerOwnableID                = baseIDs.NewCoinID(baseIDs.NewStringID("takerownableid")).ToAnyOwnableID().(*baseIDs.AnyOwnableID)
	testBaseRequest               = rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	expiresIn                     = int64(60)
	makerOwnableSplit             = sdkTypes.NewDec(60)
	takerOwnableSplit             = sdkTypes.NewDec(60)
)

func Test_newTransactionRequest(t *testing.T) {
	type args struct {
		baseReq                 rest.BaseReq
		fromID                  string
		classificationID        string
		takerID                 string
		makerOwnableID          string
		takerOwnableID          string
		expiresIn               int64
		makerOwnableSplit       string
		takerOwnableSplit       string
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
		{"+ve", args{testBaseRequest, testFromID.AsString(), testClassificationID.AsString(), testFromID.AsString(), makerOwnableID.AsString(), takerOwnableID.AsString(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, newTransactionRequest(testBaseRequest, testFromID.AsString(), testClassificationID.AsString(), testFromID.AsString(), makerOwnableID.AsString(), takerOwnableID.AsString(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.baseReq, tt.args.fromID, tt.args.classificationID, tt.args.takerID, tt.args.makerOwnableID, tt.args.takerOwnableID, tt.args.expiresIn, tt.args.makerOwnableSplit, tt.args.takerOwnableSplit, tt.args.immutableMetaProperties, tt.args.immutableProperties, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
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
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromID, constants.ClassificationID, constants.TakerID, constants.MakerOwnableID, constants.TakerOwnableID, constants.ExpiresIn, constants.MakerOwnableSplit, constants.TakerOwnableSplit, constants.ImmutableMetaProperties, constants.ImmutableProperties, constants.MutableMetaProperties, constants.MutableProperties})

	viper.Set(constants.FromID.GetName(), testFromID.AsString())
	viper.Set(constants.ClassificationID.GetName(), testClassificationID.AsString())
	viper.Set(constants.TakerID.GetName(), testFromID.AsString())
	viper.Set(constants.MakerOwnableID.GetName(), makerOwnableID.AsString())
	viper.Set(constants.TakerOwnableID.GetName(), takerOwnableID.AsString())
	viper.Set(constants.ExpiresIn.GetName(), expiresIn)
	viper.Set(constants.MakerOwnableSplit.GetName(), makerOwnableSplit.String())
	viper.Set(constants.TakerOwnableSplit.GetName(), takerOwnableSplit.String())
	viper.Set(constants.ImmutableMetaProperties.GetName(), immutableMetaPropertiesString)
	viper.Set(constants.ImmutableProperties.GetName(), immutablePropertiesString)
	viper.Set(constants.MutableMetaProperties.GetName(), mutableMetaPropertiesString)
	viper.Set(constants.MutableProperties.GetName(), mutablePropertiesString)
	type fields struct {
		BaseReq                 rest.BaseReq
		FromID                  string
		ClassificationID        string
		TakerID                 string
		MakerOwnableID          string
		TakerOwnableID          string
		ExpiresIn               int64
		MakerOwnableSplit       string
		TakerOwnableSplit       string
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
		{"+ve", fields{testBaseRequest, testFromID.AsString(), testClassificationID.AsString(), testFromID.AsString(), makerOwnableID.AsString(), takerOwnableID.AsString(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, args{cliCommand, baseHelpers.TestClientContext}, newTransactionRequest(testBaseRequest, testFromID.AsString(), testClassificationID.AsString(), testFromID.AsString(), makerOwnableID.AsString(), takerOwnableID.AsString(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
				TakerID:                 tt.fields.TakerID,
				MakerOwnableID:          tt.fields.MakerOwnableID,
				TakerOwnableID:          tt.fields.TakerOwnableID,
				ExpiresIn:               tt.fields.ExpiresIn,
				MakerOwnableSplit:       tt.fields.MakerOwnableSplit,
				TakerOwnableSplit:       tt.fields.TakerOwnableSplit,
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
			if !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("FromCLI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_FromJSON(t *testing.T) {
	jsonMessage, err := json.Marshal(newTransactionRequest(testBaseRequest, testFromID.AsString(), testClassificationID.AsString(), testFromID.AsString(), makerOwnableID.AsString(), takerOwnableID.AsString(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString))
	require.NoError(t, err)
	type fields struct {
		BaseReq                 rest.BaseReq
		FromID                  string
		ClassificationID        string
		TakerID                 string
		MakerOwnableID          string
		TakerOwnableID          string
		ExpiresIn               int64
		MakerOwnableSplit       string
		TakerOwnableSplit       string
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
		{"+ve", fields{testBaseRequest, testFromID.AsString(), testClassificationID.AsString(), testFromID.AsString(), makerOwnableID.AsString(), takerOwnableID.AsString(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, args{jsonMessage}, newTransactionRequest(testBaseRequest, testFromID.AsString(), testClassificationID.AsString(), testFromID.AsString(), makerOwnableID.AsString(), takerOwnableID.AsString(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
				TakerID:                 tt.fields.TakerID,
				MakerOwnableID:          tt.fields.MakerOwnableID,
				TakerOwnableID:          tt.fields.TakerOwnableID,
				ExpiresIn:               tt.fields.ExpiresIn,
				MakerOwnableSplit:       tt.fields.MakerOwnableSplit,
				TakerOwnableSplit:       tt.fields.TakerOwnableSplit,
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
	type fields struct {
		BaseReq                 rest.BaseReq
		FromID                  string
		ClassificationID        string
		TakerID                 string
		MakerOwnableID          string
		TakerOwnableID          string
		ExpiresIn               int64
		MakerOwnableSplit       string
		TakerOwnableSplit       string
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
		{"+ve", fields{testBaseRequest, testFromID.AsString(), testClassificationID.AsString(), testFromID.AsString(), makerOwnableID.AsString(), takerOwnableID.AsString(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, testBaseRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
				TakerID:                 tt.fields.TakerID,
				MakerOwnableID:          tt.fields.MakerOwnableID,
				TakerOwnableID:          tt.fields.TakerOwnableID,
				ExpiresIn:               tt.fields.ExpiresIn,
				MakerOwnableSplit:       tt.fields.MakerOwnableSplit,
				TakerOwnableSplit:       tt.fields.TakerOwnableSplit,
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
	type fields struct {
		BaseReq                 rest.BaseReq
		FromID                  string
		ClassificationID        string
		TakerID                 string
		MakerOwnableID          string
		TakerOwnableID          string
		ExpiresIn               int64
		MakerOwnableSplit       string
		TakerOwnableSplit       string
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
		{"+ve", fields{testBaseRequest, testFromID.AsString(), testClassificationID.AsString(), testFromID.AsString(), makerOwnableID.AsString(), takerOwnableID.AsString(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, newMessage(fromAccAddress, testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, base.NewHeight(60), makerOwnableSplit, takerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
				TakerID:                 tt.fields.TakerID,
				MakerOwnableID:          tt.fields.MakerOwnableID,
				TakerOwnableID:          tt.fields.TakerOwnableID,
				ExpiresIn:               tt.fields.ExpiresIn,
				MakerOwnableSplit:       tt.fields.MakerOwnableSplit,
				TakerOwnableSplit:       tt.fields.TakerOwnableSplit,
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
	type fields struct {
		BaseReq                 rest.BaseReq
		FromID                  string
		ClassificationID        string
		TakerID                 string
		MakerOwnableID          string
		TakerOwnableID          string
		ExpiresIn               int64
		MakerOwnableSplit       string
		TakerOwnableSplit       string
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
		{"+ve", fields{testBaseRequest, testFromID.AsString(), testClassificationID.AsString(), testFromID.AsString(), makerOwnableID.AsString(), takerOwnableID.AsString(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
				TakerID:                 tt.fields.TakerID,
				MakerOwnableID:          tt.fields.MakerOwnableID,
				TakerOwnableID:          tt.fields.TakerOwnableID,
				ExpiresIn:               tt.fields.ExpiresIn,
				MakerOwnableSplit:       tt.fields.MakerOwnableSplit,
				TakerOwnableSplit:       tt.fields.TakerOwnableSplit,
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
	type fields struct {
		BaseReq                 rest.BaseReq
		FromID                  string
		ClassificationID        string
		TakerID                 string
		MakerOwnableID          string
		TakerOwnableID          string
		ExpiresIn               int64
		MakerOwnableSplit       string
		TakerOwnableSplit       string
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
		{"+ve", fields{testBaseRequest, testFromID.AsString(), testClassificationID.AsString(), testFromID.AsString(), makerOwnableID.AsString(), takerOwnableID.AsString(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
				TakerID:                 tt.fields.TakerID,
				MakerOwnableID:          tt.fields.MakerOwnableID,
				TakerOwnableID:          tt.fields.TakerOwnableID,
				ExpiresIn:               tt.fields.ExpiresIn,
				MakerOwnableSplit:       tt.fields.MakerOwnableSplit,
				TakerOwnableSplit:       tt.fields.TakerOwnableSplit,
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
