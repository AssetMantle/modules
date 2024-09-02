// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint

import (
	"encoding/json"
	"fmt"
	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/spf13/viper"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
)

var (
	commonTransactionRequest      = helpers.PrototypeCommonTransactionRequest()
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
		commonTransactionRequest helpers.CommonTransactionRequest
		fromID                   string
		toID                     string
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
		{"+ve", args{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, transactionRequest{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.commonTransactionRequest, tt.args.fromID, tt.args.toID, tt.args.classificationID, tt.args.immutableMetaProperties, tt.args.immutableProperties, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
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
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.AssetID, constants.FromIdentityID, constants.ToIdentityID, constants.ClassificationID, constants.ImmutableMetaProperties, constants.ImmutableProperties, constants.MutableMetaProperties, constants.MutableProperties})

	viper.Set(constants.FromIdentityID.GetName(), fromID.AsString())
	viper.Set(constants.ToIdentityID.GetName(), fromID.AsString())
	viper.Set(constants.ClassificationID.GetName(), classificationID.AsString())
	viper.Set(constants.ImmutableMetaProperties.GetName(), immutableMetaPropertiesString)
	viper.Set(constants.ImmutableProperties.GetName(), immutablePropertiesString)
	viper.Set(constants.MutableMetaProperties.GetName(), mutableMetaPropertiesString)
	viper.Set(constants.MutableProperties.GetName(), mutablePropertiesString)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ToID                     string
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
		{"+ve", fields{}, args{cliCommand, client.Context{}.WithCodec(baseHelpers.CodecPrototype())}, transactionRequest{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
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
			if !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("FromCLI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_FromJSON(t *testing.T) {
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ToID                     string
		ClassificationID         string
		ImmutableMetaProperties  string
		ImmutableProperties      string
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
		{"+ve", fields{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, args{types.MustSortJSON(baseHelpers.CodecPrototype().MustMarshalJSON(&Message{fromAccAddress.String(), fromID, fromID, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}))}, newTransactionRequest(commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
				ImmutableMetaProperties:  tt.fields.ImmutableMetaProperties,
				ImmutableProperties:      tt.fields.ImmutableProperties,
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
		ToID                     string
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
		{"+ve", fields{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, commonTransactionRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
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
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ToID                     string
		ClassificationID         string
		ImmutableMetaProperties  string
		ImmutableProperties      string
		MutableMetaProperties    string
		MutableProperties        string
	}
	tests := []struct {
		name    string
		fields  fields
		want    types.Msg
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, NewMessage(fromAccAddress, fromID, fromID, classificationID, immutableMetaProperties1, immutableProperties1.ScrubData(), mutableMetaProperties1, mutableProperties1.ScrubData()), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeMsg() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_RegisterCodec(t *testing.T) {
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ToID                     string
		ClassificationID         string
		ImmutableMetaProperties  string
		ImmutableProperties      string
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
		{"+ve", fields{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
				ImmutableMetaProperties:  tt.fields.ImmutableMetaProperties,
				ImmutableProperties:      tt.fields.ImmutableProperties,
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
		ToID                     string
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
		{"+ve", fields{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
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
