// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package modify

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/AssetMantle/schema/types/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
)

var (
	fromAddress                   = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, _             = types.AccAddressFromBech32(fromAddress)
	immutableMetaPropertiesString = "defaultImmutableMeta1:S|defaultImmutableMeta1"
	immutablePropertiesString     = "defaultMutableMeta1:S|defaultMutableMeta1"
	mutableMetaPropertiesString   = "defaultMutableMeta1:S|defaultMutableMeta1"
	mutablePropertiesString       = "defaultMutable1:S|defaultMutable1"
	immutableMetaProperties, _    = baseLists.NewPropertyList().FromMetaPropertiesString(immutableMetaPropertiesString)
	immutableProperties, _        = baseLists.NewPropertyList().FromMetaPropertiesString(immutablePropertiesString)
	mutableMetaProperties, _      = baseLists.NewPropertyList().FromMetaPropertiesString(mutableMetaPropertiesString)
	mutableProperties, _          = baseLists.NewPropertyList().FromMetaPropertiesString(mutablePropertiesString)
	immutables                    = baseQualified.NewImmutables(immutableProperties)
	mutables                      = baseQualified.NewMutables(mutableProperties)

	testClassificationID     = baseIDs.NewClassificationID(immutables, mutables).(*baseIDs.ClassificationID)
	testFromID               = baseIDs.NewIdentityID(testClassificationID, immutables).(*baseIDs.IdentityID)
	testOrderID              = baseIDs.NewOrderID(testClassificationID, immutables).(*baseIDs.OrderID)
	commonTransactionRequest = baseHelpers.PrototypeCommonTransactionRequest()
	expiresIn                = int64(60)
	makerSplit               = types.NewInt(60)
	takerSplit               = types.NewInt(60)
)

func Test_newTransactionRequest(t *testing.T) {
	type args struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		fromID                   string
		orderID                  string
		takerSplit               string
		makerSplit               string
		expiresIn                int64
		mutableMetaProperties    string
		mutableProperties        string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve", args{commonTransactionRequest, testFromID.AsString(), testOrderID.AsString(), takerSplit.String(), makerSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString}, transactionRequest{commonTransactionRequest, testFromID.AsString(), testOrderID.AsString(), takerSplit.String(), makerSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.commonTransactionRequest, tt.args.fromID, tt.args.orderID, tt.args.takerSplit, tt.args.makerSplit, tt.args.expiresIn, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
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
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromIdentityID, constants.OrderID, constants.TakerSplit, constants.MakerSplit, constants.ExpiresIn, constants.MutableMetaProperties, constants.MutableProperties})

	viper.Set(constants.FromIdentityID.GetName(), testFromID.AsString())
	viper.Set(constants.OrderID.GetName(), testOrderID.AsString())
	viper.Set(constants.TakerSplit.GetName(), takerSplit.String())
	viper.Set(constants.MakerSplit.GetName(), makerSplit.String())
	viper.Set(constants.ExpiresIn.GetName(), expiresIn)
	viper.Set(constants.MutableMetaProperties.GetName(), mutableMetaPropertiesString)
	viper.Set(constants.MutableProperties.GetName(), mutablePropertiesString)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		OrderID                  string
		TakerSplit               string
		MakerSplit               string
		ExpiresIn                int64
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
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), testOrderID.AsString(), takerSplit.String(), makerSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString}, args{cliCommand, client.Context{}.WithCodec(baseHelpers.CodecPrototype())}, transactionRequest{commonTransactionRequest, testFromID.AsString(), testOrderID.AsString(), takerSplit.String(), makerSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				OrderID:                  tt.fields.OrderID,
				TakerSplit:               tt.fields.TakerSplit,
				MakerSplit:               tt.fields.MakerSplit,
				ExpiresIn:                tt.fields.ExpiresIn,
				MutableMetaProperties:    tt.fields.MutableMetaProperties,
				MutableProperties:        tt.fields.MutableProperties,
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
	jsonMessage, err := json.Marshal(newTransactionRequest(commonTransactionRequest, testFromID.AsString(), testOrderID.AsString(), takerSplit.String(), makerSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString))
	require.NoError(t, err)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		OrderID                  string
		TakerSplit               string
		MakerSplit               string
		ExpiresIn                int64
		MutableMetaProperties    string
		MutableProperties        string
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
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), testOrderID.AsString(), takerSplit.String(), makerSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString}, args{jsonMessage}, newTransactionRequest(commonTransactionRequest, testFromID.AsString(), testOrderID.AsString(), takerSplit.String(), makerSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				OrderID:                  tt.fields.OrderID,
				TakerSplit:               tt.fields.TakerSplit,
				MakerSplit:               tt.fields.MakerSplit,
				ExpiresIn:                tt.fields.ExpiresIn,
				MutableMetaProperties:    tt.fields.MutableMetaProperties,
				MutableProperties:        tt.fields.MutableProperties,
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
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		OrderID                  string
		TakerSplit               string
		MakerSplit               string
		ExpiresIn                int64
		MutableMetaProperties    string
		MutableProperties        string
	}
	tests := []struct {
		name   string
		fields fields
		want   helpers.CommonTransactionRequest
	}{
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), testOrderID.AsString(), takerSplit.String(), makerSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString}, commonTransactionRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				OrderID:                  tt.fields.OrderID,
				TakerSplit:               tt.fields.TakerSplit,
				MakerSplit:               tt.fields.MakerSplit,
				ExpiresIn:                tt.fields.ExpiresIn,
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
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		OrderID                  string
		TakerSplit               string
		MakerSplit               string
		ExpiresIn                int64
		MutableMetaProperties    string
		MutableProperties        string
	}
	tests := []struct {
		name    string
		fields  fields
		want    types.Msg
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), testOrderID.AsString(), takerSplit.String(), makerSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString}, NewMessage(fromAccAddress, testFromID, testOrderID, takerSplit, makerSplit, base.NewHeight(60), mutableMetaProperties, mutableProperties), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				OrderID:                  tt.fields.OrderID,
				TakerSplit:               tt.fields.TakerSplit,
				MakerSplit:               tt.fields.MakerSplit,
				ExpiresIn:                tt.fields.ExpiresIn,
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

func Test_transactionRequest_RegisterCodec(t *testing.T) {
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		OrderID                  string
		TakerSplit               string
		MakerSplit               string
		ExpiresIn                int64
		MutableMetaProperties    string
		MutableProperties        string
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), testOrderID.AsString(), takerSplit.String(), makerSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				OrderID:                  tt.fields.OrderID,
				TakerSplit:               tt.fields.TakerSplit,
				MakerSplit:               tt.fields.MakerSplit,
				ExpiresIn:                tt.fields.ExpiresIn,
				MutableMetaProperties:    tt.fields.MutableMetaProperties,
				MutableProperties:        tt.fields.MutableProperties,
			}
			tr.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		OrderID                  string
		TakerSplit               string
		MakerSplit               string
		ExpiresIn                int64
		MutableMetaProperties    string
		MutableProperties        string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), testOrderID.AsString(), takerSplit.String(), makerSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				OrderID:                  tt.fields.OrderID,
				TakerSplit:               tt.fields.TakerSplit,
				MakerSplit:               tt.fields.MakerSplit,
				ExpiresIn:                tt.fields.ExpiresIn,
				MutableMetaProperties:    tt.fields.MutableMetaProperties,
				MutableProperties:        tt.fields.MutableProperties,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
