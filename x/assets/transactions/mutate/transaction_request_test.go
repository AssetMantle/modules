// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mutate

import (
	"fmt"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/spf13/viper"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
)

var (
	commonTransactionRequest    = helpers.PrototypeCommonTransactionRequest()
	mutableMetaPropertiesString = "testMutableMeta1:S|mutableMeta"
	mutableMetaProperties1      = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("testMutableMeta1"), baseData.NewStringData("mutableMeta")))
	mutablePropertiesString     = "testMutable1:S|mutable"
	mutableProperties1          = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("testMutable1"), baseData.NewStringData("mutable")))
)

func Test_newTransactionRequest(t *testing.T) {
	type args struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		fromID                   string
		assetID                  string
		mutableMetaProperties    string
		mutableProperties        string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve", args{commonTransactionRequest, fromID.AsString(), testAssetID.AsString(), mutableMetaPropertiesString, mutablePropertiesString}, transactionRequest{commonTransactionRequest, fromID.AsString(), testAssetID.AsString(), mutableMetaPropertiesString, mutablePropertiesString}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.commonTransactionRequest, tt.args.fromID, tt.args.assetID, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
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
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.AssetID, constants.FromIdentityID, constants.MutableMetaProperties, constants.MutableProperties})

	viper.Set(constants.AssetID.GetName(), testAssetID.AsString())
	viper.Set(constants.FromIdentityID.GetName(), fromID.AsString())
	viper.Set(constants.MutableMetaProperties.GetName(), mutableMetaPropertiesString)
	viper.Set(constants.MutableProperties.GetName(), mutablePropertiesString)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		AssetID                  string
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
		{"+ve", fields{commonTransactionRequest, fromID.AsString(), testAssetID.AsString(), mutableMetaPropertiesString, mutablePropertiesString}, args{cliCommand, client.Context{}.WithCodec(baseHelpers.CodecPrototype())}, transactionRequest{commonTransactionRequest, fromID.AsString(), testAssetID.AsString(), mutableMetaPropertiesString, mutablePropertiesString}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				AssetID:                  tt.fields.AssetID,
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
func Test_transactionRequest_GetBaseReq(t *testing.T) {
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		AssetID                  string
		MutableMetaProperties    string
		MutableProperties        string
	}
	tests := []struct {
		name   string
		fields fields
		want   helpers.CommonTransactionRequest
	}{
		{"+ve", fields{commonTransactionRequest, fromID.AsString(), testAssetID.AsString(), mutableMetaPropertiesString, mutablePropertiesString}, commonTransactionRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				AssetID:                  tt.fields.AssetID,
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
		AssetID                  string
		MutableMetaProperties    string
		MutableProperties        string
	}
	tests := []struct {
		name    string
		fields  fields
		want    sdkTypes.Msg
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, fromID.AsString(), testAssetID.AsString(), mutableMetaPropertiesString, mutablePropertiesString}, NewMessage(fromAccAddress, fromID, testAssetID, mutableMetaProperties1, mutableProperties1.ScrubData()), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				AssetID:                  tt.fields.AssetID,
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
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		AssetID                  string
		MutableMetaProperties    string
		MutableProperties        string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, fromID.AsString(), testAssetID.AsString(), mutableMetaPropertiesString, mutablePropertiesString}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				AssetID:                  tt.fields.AssetID,
				MutableMetaProperties:    tt.fields.MutableMetaProperties,
				MutableProperties:        tt.fields.MutableProperties,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
