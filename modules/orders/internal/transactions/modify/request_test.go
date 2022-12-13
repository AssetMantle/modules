// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package modify

import (
	"encoding/json"
	"fmt"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/types/base"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

var (
	fromAddress                   = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, _             = types.AccAddressFromBech32(fromAddress)
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

	testClassificationID = baseIDs.NewClassificationID(immutables, mutables)
	testFromID           = baseIDs.NewIdentityID(testClassificationID, immutables)
	testOrderID          = baseIDs.NewOrderID(testClassificationID, immutables)
	testBaseRequest      = rest.BaseReq{From: fromAddress, ChainID: "test", Fees: types.NewCoins()}
	expiresIn            = int64(60)
	makerOwnableSplit    = types.NewDec(60)
	takerOwnableSplit    = types.NewDec(60)
)

func Test_newTransactionRequest(t *testing.T) {
	type args struct {
		baseReq               rest.BaseReq
		fromID                string
		orderID               string
		takerOwnableSplit     string
		makerOwnableSplit     string
		expiresIn             int64
		mutableMetaProperties string
		mutableProperties     string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve", args{testBaseRequest, testFromID.String(), testOrderID.String(), takerOwnableSplit.String(), makerOwnableSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString}, transactionRequest{testBaseRequest, testFromID.String(), testOrderID.String(), takerOwnableSplit.String(), makerOwnableSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.baseReq, tt.args.fromID, tt.args.orderID, tt.args.takerOwnableSplit, tt.args.makerOwnableSplit, tt.args.expiresIn, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
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
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromID, constants.OrderID, constants.TakerOwnableSplit, constants.MakerOwnableSplit, constants.ExpiresIn, constants.MutableMetaProperties, constants.MutableProperties})
	cliContext := context.NewCLIContext().WithCodec(codec.New()).WithFromAddress(fromAccAddress).WithChainID("test")
	viper.Set(constants.FromID.GetName(), testFromID.String())
	viper.Set(constants.OrderID.GetName(), testOrderID.String())
	viper.Set(constants.TakerOwnableSplit.GetName(), takerOwnableSplit.String())
	viper.Set(constants.MakerOwnableSplit.GetName(), makerOwnableSplit.String())
	viper.Set(constants.ExpiresIn.GetName(), expiresIn)
	viper.Set(constants.MutableMetaProperties.GetName(), mutableMetaPropertiesString)
	viper.Set(constants.MutableProperties.GetName(), mutablePropertiesString)
	type fields struct {
		BaseReq               rest.BaseReq
		FromID                string
		OrderID               string
		TakerOwnableSplit     string
		MakerOwnableSplit     string
		ExpiresIn             int64
		MutableMetaProperties string
		MutableProperties     string
	}
	type args struct {
		cliCommand helpers.CLICommand
		cliContext context.CLIContext
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.TransactionRequest
		wantErr bool
	}{
		{"+ve", fields{testBaseRequest, testFromID.String(), testOrderID.String(), takerOwnableSplit.String(), makerOwnableSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString}, args{cliCommand, cliContext}, transactionRequest{testBaseRequest, testFromID.String(), testOrderID.String(), takerOwnableSplit.String(), makerOwnableSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:               tt.fields.BaseReq,
				FromID:                tt.fields.FromID,
				OrderID:               tt.fields.OrderID,
				TakerOwnableSplit:     tt.fields.TakerOwnableSplit,
				MakerOwnableSplit:     tt.fields.MakerOwnableSplit,
				ExpiresIn:             tt.fields.ExpiresIn,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
			}
			got, err := transactionRequest.FromCLI(tt.args.cliCommand, tt.args.cliContext)
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
	jsonMessage, err := json.Marshal(newTransactionRequest(testBaseRequest, testFromID.String(), testOrderID.String(), takerOwnableSplit.String(), makerOwnableSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString))
	require.NoError(t, err)
	type fields struct {
		BaseReq               rest.BaseReq
		FromID                string
		OrderID               string
		TakerOwnableSplit     string
		MakerOwnableSplit     string
		ExpiresIn             int64
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
		{"+ve", fields{testBaseRequest, testFromID.String(), testOrderID.String(), takerOwnableSplit.String(), makerOwnableSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString}, args{jsonMessage}, newTransactionRequest(testBaseRequest, testFromID.String(), testOrderID.String(), takerOwnableSplit.String(), makerOwnableSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:               tt.fields.BaseReq,
				FromID:                tt.fields.FromID,
				OrderID:               tt.fields.OrderID,
				TakerOwnableSplit:     tt.fields.TakerOwnableSplit,
				MakerOwnableSplit:     tt.fields.MakerOwnableSplit,
				ExpiresIn:             tt.fields.ExpiresIn,
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
	type fields struct {
		BaseReq               rest.BaseReq
		FromID                string
		OrderID               string
		TakerOwnableSplit     string
		MakerOwnableSplit     string
		ExpiresIn             int64
		MutableMetaProperties string
		MutableProperties     string
	}
	tests := []struct {
		name   string
		fields fields
		want   rest.BaseReq
	}{
		{"+ve", fields{testBaseRequest, testFromID.String(), testOrderID.String(), takerOwnableSplit.String(), makerOwnableSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString}, testBaseRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:               tt.fields.BaseReq,
				FromID:                tt.fields.FromID,
				OrderID:               tt.fields.OrderID,
				TakerOwnableSplit:     tt.fields.TakerOwnableSplit,
				MakerOwnableSplit:     tt.fields.MakerOwnableSplit,
				ExpiresIn:             tt.fields.ExpiresIn,
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
	type fields struct {
		BaseReq               rest.BaseReq
		FromID                string
		OrderID               string
		TakerOwnableSplit     string
		MakerOwnableSplit     string
		ExpiresIn             int64
		MutableMetaProperties string
		MutableProperties     string
	}
	tests := []struct {
		name    string
		fields  fields
		want    types.Msg
		wantErr bool
	}{
		{"+ve", fields{testBaseRequest, testFromID.String(), testOrderID.String(), takerOwnableSplit.String(), makerOwnableSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString}, newMessage(fromAccAddress, testFromID, testOrderID, takerOwnableSplit, makerOwnableSplit, base.NewHeight(60), mutableMetaProperties, mutableProperties), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:               tt.fields.BaseReq,
				FromID:                tt.fields.FromID,
				OrderID:               tt.fields.OrderID,
				TakerOwnableSplit:     tt.fields.TakerOwnableSplit,
				MakerOwnableSplit:     tt.fields.MakerOwnableSplit,
				ExpiresIn:             tt.fields.ExpiresIn,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
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
		BaseReq               rest.BaseReq
		FromID                string
		OrderID               string
		TakerOwnableSplit     string
		MakerOwnableSplit     string
		ExpiresIn             int64
		MutableMetaProperties string
		MutableProperties     string
	}
	type args struct {
		codec *codec.Codec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testBaseRequest, testFromID.String(), testOrderID.String(), takerOwnableSplit.String(), makerOwnableSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				BaseReq:               tt.fields.BaseReq,
				FromID:                tt.fields.FromID,
				OrderID:               tt.fields.OrderID,
				TakerOwnableSplit:     tt.fields.TakerOwnableSplit,
				MakerOwnableSplit:     tt.fields.MakerOwnableSplit,
				ExpiresIn:             tt.fields.ExpiresIn,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
			}
			tr.RegisterCodec(tt.args.codec)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	type fields struct {
		BaseReq               rest.BaseReq
		FromID                string
		OrderID               string
		TakerOwnableSplit     string
		MakerOwnableSplit     string
		ExpiresIn             int64
		MutableMetaProperties string
		MutableProperties     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{testBaseRequest, testFromID.String(), testOrderID.String(), takerOwnableSplit.String(), makerOwnableSplit.String(), expiresIn, mutableMetaPropertiesString, mutablePropertiesString}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:               tt.fields.BaseReq,
				FromID:                tt.fields.FromID,
				OrderID:               tt.fields.OrderID,
				TakerOwnableSplit:     tt.fields.TakerOwnableSplit,
				MakerOwnableSplit:     tt.fields.MakerOwnableSplit,
				ExpiresIn:             tt.fields.ExpiresIn,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
