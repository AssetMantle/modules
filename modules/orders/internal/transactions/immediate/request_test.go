// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package immediate

import (
	"encoding/json"
	"fmt"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/types/base"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
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
	testClassificationID          = baseIDs.NewClassificationID(immutables, mutables)
	testFromID                    = baseIDs.NewIdentityID(testClassificationID, immutables)
	makerOwnableID                = baseIDs.NewOwnableID(baseIDs.NewStringID("makerownableid"))
	takerOwnableID                = baseIDs.NewOwnableID(baseIDs.NewStringID("takerownableid"))
	testBaseRequest               = rest.BaseReq{From: fromAddress, ChainID: "test", Fees: types.NewCoins()}
	expiresIn                     = int64(60)
	makerOwnableSplit             = types.NewDec(60)
	takerOwnableSplit             = types.NewDec(60)
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
		// TODO: Add test cases.
		{"+ve", args{testBaseRequest, testFromID.String(), testClassificationID.String(), testFromID.String(), makerOwnableID.String(), takerOwnableID.String(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, newTransactionRequest(testBaseRequest, testFromID.String(), testClassificationID.String(), testFromID.String(), makerOwnableID.String(), takerOwnableID.String(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString)},
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
		// TODO: Add test cases.
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
		cliContext context.CLIContext
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.TransactionRequest
		wantErr bool
	}{
		// TODO: Add test cases.
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
			got, err := transactionRequest.FromCLI(tt.args.cliCommand, tt.args.cliContext)
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
	jsonMessage, err := json.Marshal(newTransactionRequest(testBaseRequest, testFromID.String(), testClassificationID.String(), testFromID.String(), makerOwnableID.String(), takerOwnableID.String(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString))
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
		// TODO: Add test cases.
		{"+ve", fields{testBaseRequest, testFromID.String(), testClassificationID.String(), testFromID.String(), makerOwnableID.String(), takerOwnableID.String(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, args{jsonMessage}, newTransactionRequest(testBaseRequest, testFromID.String(), testClassificationID.String(), testFromID.String(), makerOwnableID.String(), takerOwnableID.String(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString), false},
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
		// TODO: Add test cases.
		{"+ve", fields{testBaseRequest, testFromID.String(), testClassificationID.String(), testFromID.String(), makerOwnableID.String(), takerOwnableID.String(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, testBaseRequest},
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
		want    types.Msg
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{testBaseRequest, testFromID.String(), testClassificationID.String(), testFromID.String(), makerOwnableID.String(), takerOwnableID.String(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, newMessage(fromAccAddress, testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, base.NewHeight(60), makerOwnableSplit, takerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties), false},
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
		codec *codec.Codec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"+ve", fields{testBaseRequest, testFromID.String(), testClassificationID.String(), testFromID.String(), makerOwnableID.String(), takerOwnableID.String(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, args{codec.New()}},
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
			tr.RegisterCodec(tt.args.codec)
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
		// TODO: Add test cases.
		{"+ve", fields{testBaseRequest, testFromID.String(), testClassificationID.String(), testFromID.String(), makerOwnableID.String(), takerOwnableID.String(), expiresIn, makerOwnableSplit.String(), takerOwnableSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, false},
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
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
