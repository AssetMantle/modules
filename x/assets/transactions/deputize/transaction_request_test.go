// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"encoding/json"
	"fmt"
	"github.com/AssetMantle/modules/utilities/rest"
	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/spf13/viper"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
)

var (
	commonTransactionRequest    = rest.PrototypeCommonTransactionRequest()
	mutableMetaPropertiesString = "testMutableMeta1:S|mutableMeta"
	mutableMetaProperties1      = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("testMutableMeta1"), baseData.NewStringData("mutableMeta")))
)

func Test_newTransactionRequest(t *testing.T) {
	type args struct {
		commonTransactionRequest rest.CommonTransactionRequest
		fromID                   string
		toID                     string
		classificationID         string
		maintainedProperties     string
		canMintAsset             bool
		canBurnAsset             bool
		canRenumerateAsset       bool
		canAddMaintainer         bool
		canRemoveMaintainer      bool
		canMutateMaintainer      bool
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve", args{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), fmt.Sprint(mutableMetaProperties), true, true, true, true, true, true}, transactionRequest{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), fmt.Sprint(mutableMetaProperties), true, true, true, true, true, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.commonTransactionRequest, tt.args.fromID, tt.args.toID, tt.args.classificationID, tt.args.maintainedProperties, tt.args.canMintAsset, tt.args.canBurnAsset, tt.args.canRenumerateAsset, tt.args.canAddMaintainer, tt.args.canRemoveMaintainer, tt.args.canMutateMaintainer); !reflect.DeepEqual(got, tt.want) {
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
	cliCommand := base.NewCLICommand("", "", "", []helpers.CLIFlag{constants.ToIdentityID, constants.FromIdentityID, constants.ClassificationID, constants.MaintainedProperties, constants.CanMintAsset, constants.CanBurnAsset, constants.CanRenumerateAsset, constants.CanAddMaintainer, constants.CanRemoveMaintainer, constants.CanMutateMaintainer})

	viper.Set(constants.ToIdentityID.GetName(), fromID.AsString())
	viper.Set(constants.FromIdentityID.GetName(), fromID.AsString())
	viper.Set(constants.ClassificationID.GetName(), classificationID.AsString())
	viper.Set(constants.MaintainedProperties.GetName(), mutableMetaPropertiesString)
	viper.Set(constants.CanMintAsset.GetName(), true)
	viper.Set(constants.CanBurnAsset.GetName(), true)
	viper.Set(constants.CanRenumerateAsset.GetName(), true)
	viper.Set(constants.CanAddMaintainer.GetName(), true)
	viper.Set(constants.CanRemoveMaintainer.GetName(), true)
	viper.Set(constants.CanMutateMaintainer.GetName(), true)
	type fields struct {
		commonTransactionRequest rest.CommonTransactionRequest
		FromID                   string
		ToID                     string
		ClassificationID         string
		MaintainedProperties     string
		CanMintAsset             bool
		CanBurnAsset             bool
		CanRenumerateAsset       bool
		CanAddMaintainer         bool
		CanRemoveMaintainer      bool
		CanMutateMaintainer      bool
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
		{"+ve", fields{}, args{cliCommand, client.Context{}.WithCodec(base.CodecPrototype())}, transactionRequest{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), mutableMetaPropertiesString, true, true, true, true, true, true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
				MaintainedProperties:     tt.fields.MaintainedProperties,
				CanMintAsset:             tt.fields.CanMintAsset,
				CanBurnAsset:             tt.fields.CanBurnAsset,
				CanRenumerateAsset:       tt.fields.CanRenumerateAsset,
				CanAddMaintainer:         tt.fields.CanAddMaintainer,
				CanRemoveMaintainer:      tt.fields.CanRemoveMaintainer,
				CanMutateMaintainer:      tt.fields.CanMutateMaintainer,
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
		commonTransactionRequest rest.CommonTransactionRequest
		FromID                   string
		ToID                     string
		ClassificationID         string
		MaintainedProperties     string
		CanMintAsset             bool
		CanBurnAsset             bool
		CanRenumerateAsset       bool
		CanAddMaintainer         bool
		CanRemoveMaintainer      bool
		CanMutateMaintainer      bool
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
		{"+ve", fields{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), fmt.Sprint(mutableMetaProperties), true, true, true, true, true, true}, args{sdkTypes.MustSortJSON(base.CodecPrototype().MustMarshalJSON(&Message{fromAccAddress.String(), fromID, fromID, classificationID, mutableMetaProperties, true, true, true, true, true, true}))}, transactionRequest{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), fmt.Sprint(mutableMetaProperties), true, true, true, true, true, true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
				MaintainedProperties:     tt.fields.MaintainedProperties,
				CanMintAsset:             tt.fields.CanMintAsset,
				CanBurnAsset:             tt.fields.CanBurnAsset,
				CanRenumerateAsset:       tt.fields.CanRenumerateAsset,
				CanAddMaintainer:         tt.fields.CanAddMaintainer,
				CanRemoveMaintainer:      tt.fields.CanRemoveMaintainer,
				CanMutateMaintainer:      tt.fields.CanMutateMaintainer,
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
		ToID                     string
		ClassificationID         string
		MaintainedProperties     string
		CanMintAsset             bool
		CanBurnAsset             bool
		CanRenumerateAsset       bool
		CanAddMaintainer         bool
		CanRemoveMaintainer      bool
		CanMutateMaintainer      bool
	}
	tests := []struct {
		name   string
		fields fields
		want   rest.CommonTransactionRequest
	}{
		{"+ve", fields{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), fmt.Sprint(mutableMetaProperties), true, true, true, true, true, true}, commonTransactionRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
				MaintainedProperties:     tt.fields.MaintainedProperties,
				CanMintAsset:             tt.fields.CanMintAsset,
				CanBurnAsset:             tt.fields.CanBurnAsset,
				CanRenumerateAsset:       tt.fields.CanRenumerateAsset,
				CanAddMaintainer:         tt.fields.CanAddMaintainer,
				CanRemoveMaintainer:      tt.fields.CanRemoveMaintainer,
				CanMutateMaintainer:      tt.fields.CanMutateMaintainer,
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
		ToID                     string
		ClassificationID         string
		MaintainedProperties     string
		CanMintAsset             bool
		CanBurnAsset             bool
		CanRenumerateAsset       bool
		CanAddMaintainer         bool
		CanRemoveMaintainer      bool
		CanMutateMaintainer      bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    sdkTypes.Msg
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), mutableMetaPropertiesString, true, true, true, true, true, true}, NewMessage(fromAccAddress, fromID, fromID, classificationID, mutableMetaProperties1, true, true, true, true, true, true), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
				MaintainedProperties:     tt.fields.MaintainedProperties,
				CanMintAsset:             tt.fields.CanMintAsset,
				CanBurnAsset:             tt.fields.CanBurnAsset,
				CanRenumerateAsset:       tt.fields.CanRenumerateAsset,
				CanAddMaintainer:         tt.fields.CanAddMaintainer,
				CanRemoveMaintainer:      tt.fields.CanRemoveMaintainer,
				CanMutateMaintainer:      tt.fields.CanMutateMaintainer,
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
		commonTransactionRequest rest.CommonTransactionRequest
		FromID                   string
		ToID                     string
		ClassificationID         string
		MaintainedProperties     string
		CanMintAsset             bool
		CanBurnAsset             bool
		CanRenumerateAsset       bool
		CanAddMaintainer         bool
		CanRemoveMaintainer      bool
		CanMutateMaintainer      bool
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), fmt.Sprint(mutableMetaProperties), true, true, true, true, true, true}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
				MaintainedProperties:     tt.fields.MaintainedProperties,
				CanMintAsset:             tt.fields.CanMintAsset,
				CanBurnAsset:             tt.fields.CanBurnAsset,
				CanRenumerateAsset:       tt.fields.CanRenumerateAsset,
				CanAddMaintainer:         tt.fields.CanAddMaintainer,
				CanRemoveMaintainer:      tt.fields.CanRemoveMaintainer,
				CanMutateMaintainer:      tt.fields.CanMutateMaintainer,
			}
			tr.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	type fields struct {
		commonTransactionRequest rest.CommonTransactionRequest
		FromID                   string
		ToID                     string
		ClassificationID         string
		MaintainedProperties     string
		CanMintAsset             bool
		CanBurnAsset             bool
		CanRenumerateAsset       bool
		CanAddMaintainer         bool
		CanRemoveMaintainer      bool
		CanMutateMaintainer      bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve with nil", fields{}, true},
		{"+ve", fields{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), fmt.Sprint(mutableMetaProperties), true, true, true, true, true, true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
				MaintainedProperties:     tt.fields.MaintainedProperties,
				CanMintAsset:             tt.fields.CanMintAsset,
				CanBurnAsset:             tt.fields.CanBurnAsset,
				CanRenumerateAsset:       tt.fields.CanRenumerateAsset,
				CanAddMaintainer:         tt.fields.CanAddMaintainer,
				CanRemoveMaintainer:      tt.fields.CanRemoveMaintainer,
				CanMutateMaintainer:      tt.fields.CanMutateMaintainer,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
