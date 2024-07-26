// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package put

import (
	"encoding/json"
	"fmt"
	"github.com/AssetMantle/modules/utilities/rest"
	"reflect"
	"testing"

	baseDocuments "github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/AssetMantle/schema/types/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
)

var (
	fromAddress                   = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, _             = sdkTypes.AccAddressFromBech32(fromAddress)
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
	testClassificationID          = baseIDs.NewClassificationID(immutables, mutables).(*baseIDs.ClassificationID)
	testFromID                    = baseIDs.NewIdentityID(testClassificationID, immutables).(*baseIDs.IdentityID)
	makerAssetID                  = baseDocuments.NewCoinAsset("makerAssetID").GetCoinAssetID().(*baseIDs.AssetID)
	takerAssetID                  = baseDocuments.NewCoinAsset("takerAssetID").GetCoinAssetID().(*baseIDs.AssetID)
	commonTransactionRequest      = rest.PrototypeCommonTransactionRequest()
	expiryHeight                  = int64(60)
	makerSplit                    = sdkTypes.NewInt(60)
	takerSplit                    = sdkTypes.NewInt(60)
)

func Test_newTransactionRequest(t *testing.T) {
	type args struct {
		commonTransactionRequest rest.CommonTransactionRequest
		FromID                   string
		MakerAssetID             string
		TakerAssetID             string
		MakerSplit               string
		TakerSplit               string
		ExpiryHeight             int64
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve", args{commonTransactionRequest, testFromID.AsString(), makerAssetID.AsString(), takerAssetID.AsString(), makerSplit.String(), takerSplit.String(), expiryHeight}, newTransactionRequest(commonTransactionRequest, testFromID.AsString(), makerAssetID.AsString(), takerAssetID.AsString(), makerSplit.String(), takerSplit.String(), expiryHeight)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.commonTransactionRequest, tt.args.FromID, tt.args.MakerAssetID, tt.args.TakerAssetID, tt.args.MakerSplit, tt.args.TakerSplit, tt.args.ExpiryHeight); !reflect.DeepEqual(got, tt.want) {
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
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromIdentityID, constants.ClassificationID, constants.TakerID, constants.MakerAssetID, constants.TakerAssetID, constants.ExpiresIn, constants.MakerSplit, constants.TakerSplit, constants.ImmutableMetaProperties, constants.ImmutableProperties, constants.MutableMetaProperties, constants.MutableProperties})

	viper.Set(constants.FromIdentityID.GetName(), testFromID.AsString())
	viper.Set(constants.ClassificationID.GetName(), testClassificationID.AsString())
	viper.Set(constants.TakerID.GetName(), testFromID.AsString())
	viper.Set(constants.MakerAssetID.GetName(), makerAssetID.AsString())
	viper.Set(constants.TakerAssetID.GetName(), takerAssetID.AsString())
	viper.Set(constants.ExpiresIn.GetName(), expiryHeight)
	viper.Set(constants.MakerSplit.GetName(), makerSplit.String())
	viper.Set(constants.TakerSplit.GetName(), takerSplit.String())
	viper.Set(constants.ImmutableMetaProperties.GetName(), immutableMetaPropertiesString)
	viper.Set(constants.ImmutableProperties.GetName(), immutablePropertiesString)
	viper.Set(constants.MutableMetaProperties.GetName(), mutableMetaPropertiesString)
	viper.Set(constants.MutableProperties.GetName(), mutablePropertiesString)
	type fields struct {
		commonTransactionRequest rest.CommonTransactionRequest
		FromID                   string
		MakerAssetID             string
		TakerAssetID             string
		MakerSplit               string
		TakerSplit               string
		ExpiryHeight             int64
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
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), makerAssetID.AsString(), takerAssetID.AsString(), makerSplit.String(), takerSplit.String(), expiryHeight}, args{cliCommand, client.Context{}.WithCodec(baseHelpers.CodecPrototype())}, newTransactionRequest(commonTransactionRequest, testFromID.AsString(), makerAssetID.AsString(), takerAssetID.AsString(), makerSplit.String(), takerSplit.String(), expiryHeight), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
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
	jsonMessage, err := json.Marshal(newTransactionRequest(commonTransactionRequest, testFromID.AsString(), makerAssetID.AsString(), takerAssetID.AsString(), makerSplit.String(), takerSplit.String(), expiryHeight))
	require.NoError(t, err)
	type fields struct {
		commonTransactionRequest rest.CommonTransactionRequest
		FromID                   string
		MakerAssetID             string
		TakerAssetID             string
		MakerSplit               string
		TakerSplit               string
		ExpiryHeight             int64
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
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), makerAssetID.AsString(), takerAssetID.AsString(), makerSplit.String(), takerSplit.String(), expiryHeight}, args{jsonMessage}, newTransactionRequest(commonTransactionRequest, testFromID.AsString(), makerAssetID.AsString(), takerAssetID.AsString(), makerSplit.String(), takerSplit.String(), expiryHeight), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				tt.fields.commonTransactionRequest,
				tt.fields.FromID,
				tt.fields.MakerAssetID,
				tt.fields.TakerAssetID,
				tt.fields.MakerSplit,
				tt.fields.TakerSplit,
				tt.fields.ExpiryHeight,
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
		commonTransactionRequest rest.CommonTransactionRequest
		FromID                   string
		MakerAssetID             string
		TakerAssetID             string
		MakerSplit               string
		TakerSplit               string
		ExpiryHeight             int64
	}
	tests := []struct {
		name   string
		fields fields
		want   rest.CommonTransactionRequest
	}{
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), makerAssetID.AsString(), takerAssetID.AsString(), makerSplit.String(), takerSplit.String(), expiryHeight}, commonTransactionRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				tt.fields.commonTransactionRequest,
				tt.fields.FromID,
				tt.fields.MakerAssetID,
				tt.fields.TakerAssetID,
				tt.fields.MakerSplit,
				tt.fields.TakerSplit,
				tt.fields.ExpiryHeight,
			}
			if got := transactionRequest.GetCommonTransactionRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommonTransactionRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	type fields struct {
		commonTransactionRequest rest.CommonTransactionRequest
		FromID                   string
		MakerAssetID             string
		TakerAssetID             string
		MakerSplit               string
		TakerSplit               string
		ExpiryHeight             int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    sdkTypes.Msg
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), makerAssetID.AsString(), takerAssetID.AsString(), makerSplit.String(), takerSplit.String(), expiryHeight}, NewMessage(fromAccAddress, testFromID, makerAssetID, takerAssetID, makerSplit, takerSplit, base.NewHeight(60)), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				tt.fields.commonTransactionRequest,
				tt.fields.FromID,
				tt.fields.MakerAssetID,
				tt.fields.TakerAssetID,
				tt.fields.MakerSplit,
				tt.fields.TakerSplit,
				tt.fields.ExpiryHeight,
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
		commonTransactionRequest rest.CommonTransactionRequest
		FromID                   string
		MakerAssetID             string
		TakerAssetID             string
		MakerSplit               string
		TakerSplit               string
		ExpiryHeight             int64
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), makerAssetID.AsString(), takerAssetID.AsString(), makerSplit.String(), takerSplit.String(), expiryHeight}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				tt.fields.commonTransactionRequest,
				tt.fields.FromID,
				tt.fields.MakerAssetID,
				tt.fields.TakerAssetID,
				tt.fields.MakerSplit,
				tt.fields.TakerSplit,
				tt.fields.ExpiryHeight,
			}
			tr.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	type fields struct {
		commonTransactionRequest rest.CommonTransactionRequest
		FromID                   string
		MakerAssetID             string
		TakerAssetID             string
		MakerSplit               string
		TakerSplit               string
		ExpiryHeight             int64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), makerAssetID.AsString(), takerAssetID.AsString(), makerSplit.String(), takerSplit.String(), expiryHeight}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				tt.fields.commonTransactionRequest,
				tt.fields.FromID,
				tt.fields.MakerAssetID,
				tt.fields.TakerAssetID,
				tt.fields.MakerSplit,
				tt.fields.TakerSplit,
				tt.fields.ExpiryHeight,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
