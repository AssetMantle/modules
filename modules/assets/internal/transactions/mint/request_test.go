// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint

import (
	"encoding/json"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/utilities/transaction"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"reflect"
	"testing"
)

var (
	testBaseRequest               = rest.BaseReq{From: fromAddress, ChainID: "test", Fees: types.NewCoins()}
	immutableMetaPropertiesString = "testImmutableMeta1:S|immutableMeta"
	immutableMetaProperties1      = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("testImmutableMeta1"), baseData.NewStringData("immutableMeta")))
	immutablePropertiesString     = "testImmutable1:S|immutable"
	immutableProperties1          = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("testImmutable1"), baseData.NewStringData("immutable")))
	mutableMetaPropertiesString   = "testMutableMeta1:S|mutableMeta"
	mutableMetaProperties1        = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("testMutableMeta1"), baseData.NewStringData("mutableMeta")))
	mutablePropertiesString       = "testMutable1:S|mutable"
	mutableProperties1            = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("testMutable1"), baseData.NewStringData("mutable")))
)

func Test_newTransactionRequest(t *testing.T) {
	type args struct {
		baseReq                 rest.BaseReq
		fromID                  string
		toID                    string
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
		// TODO: Add test cases.
		{"+ve", args{testBaseRequest, fromID.String(), fromID.String(), classificationID.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, transactionRequest{testBaseRequest, fromID.String(), fromID.String(), classificationID.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.baseReq, tt.args.fromID, tt.args.toID, tt.args.classificationID, tt.args.immutableMetaProperties, tt.args.immutableProperties, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
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
		ToID                    string
		ClassificationID        string
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
				ToID:                    tt.fields.ToID,
				ClassificationID:        tt.fields.ClassificationID,
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
	type fields struct {
		BaseReq                 rest.BaseReq
		FromID                  string
		ToID                    string
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
		// TODO: Add test cases.
		{"+ve", fields{testBaseRequest, fromID.String(), fromID.String(), classificationID.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, args{types.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(message{fromAccAddress, fromID, fromID, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}))}, newTransactionRequest(testBaseRequest, fromID.String(), fromID.String(), classificationID.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				FromID:                  tt.fields.FromID,
				ToID:                    tt.fields.ToID,
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
	type fields struct {
		BaseReq                 rest.BaseReq
		FromID                  string
		ToID                    string
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
		// TODO: Add test cases.
		{"+ve", fields{testBaseRequest, fromID.String(), fromID.String(), classificationID.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, testBaseRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				FromID:                  tt.fields.FromID,
				ToID:                    tt.fields.ToID,
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
	type fields struct {
		BaseReq                 rest.BaseReq
		FromID                  string
		ToID                    string
		ClassificationID        string
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
		{"+ve", fields{testBaseRequest, fromID.String(), fromID.String(), classificationID.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, newMessage(fromAccAddress, fromID, fromID, classificationID, immutableMetaProperties1, immutableProperties1.ScrubData(), mutableMetaProperties1, mutableProperties1.ScrubData()), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				FromID:                  tt.fields.FromID,
				ToID:                    tt.fields.ToID,
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeMsg() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_RegisterCodec(t *testing.T) {
	type fields struct {
		BaseReq                 rest.BaseReq
		FromID                  string
		ToID                    string
		ClassificationID        string
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
		{"+ve", fields{testBaseRequest, fromID.String(), fromID.String(), classificationID.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				FromID:                  tt.fields.FromID,
				ToID:                    tt.fields.ToID,
				ClassificationID:        tt.fields.ClassificationID,
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
		ToID                    string
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
		// TODO: Add test cases.
		{"+ve", fields{testBaseRequest, fromID.String(), fromID.String(), classificationID.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				FromID:                  tt.fields.FromID,
				ToID:                    tt.fields.ToID,
				ClassificationID:        tt.fields.ClassificationID,
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
