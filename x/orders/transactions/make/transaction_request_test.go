// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

import (
	"cosmossdk.io/math"
	"fmt"
	"testing"

	baseDocuments "github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
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
	makerAssetID                  = baseDocuments.NewCoinAsset("makerAssetID").GetCoinAssetID()
	takerAssetID                  = baseDocuments.NewCoinAsset("takerAssetID").GetCoinAssetID()
	commonTransactionRequest      = helpers.PrototypeCommonTransactionRequest()
	expiresIn                     = int64(60)
	makerSplit                    = math.NewInt(60)
	takerSplit                    = math.NewInt(60)
)

func Test_newTransactionRequest(t *testing.T) {
	type args struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		fromID                   string
		classificationID         string
		takerID                  string
		makerAssetID             string
		takerAssetID             string
		expiresIn                int64
		makerSplit               string
		takerSplit               string
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
		{"valid", args{commonTransactionRequest, testFromID.AsString(), testClassificationID.AsString(), testFromID.AsString(), makerAssetID.AsString(), takerAssetID.AsString(), expiresIn, makerSplit.String(), takerSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, newTransactionRequest(commonTransactionRequest, testFromID.AsString(), testClassificationID.AsString(), testFromID.AsString(), makerAssetID.AsString(), takerAssetID.AsString(), expiresIn, makerSplit.String(), takerSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newTransactionRequest(tt.args.commonTransactionRequest, tt.args.fromID, tt.args.classificationID, tt.args.takerID, tt.args.makerAssetID, tt.args.takerAssetID, tt.args.expiresIn, tt.args.makerSplit, tt.args.takerSplit, tt.args.immutableMetaProperties, tt.args.immutableProperties, tt.args.mutableMetaProperties, tt.args.mutableProperties)
			assert.Equal(t, tt.want, got, "newTransactionRequest()")
		})
	}
}

func Test_requestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.TransactionRequest
	}{
		{"valid", transactionRequest{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := requestPrototype()
			assert.Equal(t, tt.want, got, "requestPrototype()")
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
	viper.Set(constants.ExpiresIn.GetName(), expiresIn)
	viper.Set(constants.MakerSplit.GetName(), makerSplit.String())
	viper.Set(constants.TakerSplit.GetName(), takerSplit.String())
	viper.Set(constants.ImmutableMetaProperties.GetName(), immutableMetaPropertiesString)
	viper.Set(constants.ImmutableProperties.GetName(), immutablePropertiesString)
	viper.Set(constants.MutableMetaProperties.GetName(), mutableMetaPropertiesString)
	viper.Set(constants.MutableProperties.GetName(), mutablePropertiesString)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ClassificationID         string
		TakerID                  string
		MakerAssetID             string
		TakerAssetID             string
		ExpiresIn                int64
		MakerSplit               string
		TakerSplit               string
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
		{"valid", fields{commonTransactionRequest, testFromID.AsString(), testClassificationID.AsString(), testFromID.AsString(), makerAssetID.AsString(), takerAssetID.AsString(), expiresIn, makerSplit.String(), takerSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, args{cliCommand, client.Context{}.WithCodec(baseHelpers.CodecPrototype())}, newTransactionRequest(commonTransactionRequest, testFromID.AsString(), testClassificationID.AsString(), testFromID.AsString(), makerAssetID.AsString(), takerAssetID.AsString(), expiresIn, makerSplit.String(), takerSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ClassificationID:         tt.fields.ClassificationID,
				TakerID:                  tt.fields.TakerID,
				MakerAssetID:             tt.fields.MakerAssetID,
				TakerAssetID:             tt.fields.TakerAssetID,
				ExpiresIn:                tt.fields.ExpiresIn,
				MakerSplit:               tt.fields.MakerSplit,
				TakerSplit:               tt.fields.TakerSplit,
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
			assert.Equal(t, fmt.Sprint(tt.want), fmt.Sprint(got), "FromCLI()")
		})
	}
}

func Test_transactionRequest_GetBaseReq(t *testing.T) {
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ClassificationID         string
		TakerID                  string
		MakerAssetID             string
		TakerAssetID             string
		ExpiresIn                int64
		MakerSplit               string
		TakerSplit               string
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
		{"valid", fields{commonTransactionRequest, testFromID.AsString(), testClassificationID.AsString(), testFromID.AsString(), makerAssetID.AsString(), takerAssetID.AsString(), expiresIn, makerSplit.String(), takerSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, commonTransactionRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ClassificationID:         tt.fields.ClassificationID,
				TakerID:                  tt.fields.TakerID,
				MakerAssetID:             tt.fields.MakerAssetID,
				TakerAssetID:             tt.fields.TakerAssetID,
				ExpiresIn:                tt.fields.ExpiresIn,
				MakerSplit:               tt.fields.MakerSplit,
				TakerSplit:               tt.fields.TakerSplit,
				ImmutableMetaProperties:  tt.fields.ImmutableMetaProperties,
				ImmutableProperties:      tt.fields.ImmutableProperties,
				MutableMetaProperties:    tt.fields.MutableMetaProperties,
				MutableProperties:        tt.fields.MutableProperties,
			}
			got := transactionRequest.GetCommonTransactionRequest()
			assert.Equal(t, tt.want, got, "GetCommonTransactionRequest()")
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ClassificationID         string
		TakerID                  string
		MakerAssetID             string
		TakerAssetID             string
		ExpiresIn                int64
		MakerSplit               string
		TakerSplit               string
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
		{"valid", fields{commonTransactionRequest, testFromID.AsString(), testClassificationID.AsString(), testFromID.AsString(), makerAssetID.AsString(), takerAssetID.AsString(), expiresIn, makerSplit.String(), takerSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ClassificationID:         tt.fields.ClassificationID,
				TakerID:                  tt.fields.TakerID,
				MakerAssetID:             tt.fields.MakerAssetID,
				TakerAssetID:             tt.fields.TakerAssetID,
				ExpiresIn:                tt.fields.ExpiresIn,
				MakerSplit:               tt.fields.MakerSplit,
				TakerSplit:               tt.fields.TakerSplit,
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
			assert.Equal(t, fmt.Sprint(tt.want), fmt.Sprint(got), "MakeMsg()")
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ClassificationID         string
		TakerID                  string
		MakerAssetID             string
		TakerAssetID             string
		ExpiresIn                int64
		MakerSplit               string
		TakerSplit               string
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
		{"valid", fields{commonTransactionRequest, testFromID.AsString(), testClassificationID.AsString(), testFromID.AsString(), makerAssetID.AsString(), takerAssetID.AsString(), expiresIn, makerSplit.String(), takerSplit.String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ClassificationID:         tt.fields.ClassificationID,
				TakerID:                  tt.fields.TakerID,
				MakerAssetID:             tt.fields.MakerAssetID,
				TakerAssetID:             tt.fields.TakerAssetID,
				ExpiresIn:                tt.fields.ExpiresIn,
				MakerSplit:               tt.fields.MakerSplit,
				TakerSplit:               tt.fields.TakerSplit,
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
