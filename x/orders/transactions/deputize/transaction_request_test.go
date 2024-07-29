// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
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
	fromAddress       = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, _ = types.AccAddressFromBech32(fromAddress)
	immutables        = baseQualified.NewImmutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables          = baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))

	testClassificationID     = baseIDs.NewClassificationID(immutables, mutables)
	testFromID               = baseIDs.NewIdentityID(testClassificationID, immutables)
	maintainedPropertyString = "maintainedProperty:S|maintainedProperty"
	maintainedProperties, _  = base.NewPropertyList().FromMetaPropertiesString(maintainedPropertyString)
	commonTransactionRequest = baseHelpers.PrototypeCommonTransactionRequest()
)

func Test_newTransactionRequest(t *testing.T) {
	type args struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		fromID                   string
		toID                     string
		classificationID         string
		maintainedProperties     string
		CanMakeOrder             bool `json:"canMakeOrder"`
		CanCancelOrder           bool `json:"canCancelOrder"`
		CanAddMaintainer         bool `json:"canAddMaintainer"`
		CanRemoveMaintainer      bool `json:"canRemoveMaintainer"`
		CanMutateMaintainer      bool `json:"canMutateMaintainer"`
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve", args{commonTransactionRequest, testFromID.AsString(), testFromID.AsString(), testClassificationID.AsString(), maintainedPropertyString, true, true, true, true, true}, newTransactionRequest(commonTransactionRequest, testFromID.AsString(), testFromID.AsString(), testClassificationID.AsString(), maintainedPropertyString, true, true, true, true, true)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.commonTransactionRequest, tt.args.fromID, tt.args.toID, tt.args.classificationID, tt.args.maintainedProperties, tt.args.CanMakeOrder, tt.args.CanCancelOrder, tt.args.CanAddMaintainer, tt.args.CanRemoveMaintainer, tt.args.CanMutateMaintainer); !reflect.DeepEqual(got, tt.want) {
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
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromIdentityID, constants.ToIdentityID, constants.ClassificationID, constants.MaintainedProperties, constants.CanMintAsset, constants.CanBurnAsset, constants.CanRenumerateAsset, constants.CanAddMaintainer, constants.CanRemoveMaintainer, constants.CanMutateMaintainer})

	viper.Set(constants.FromIdentityID.GetName(), testFromID.AsString())
	viper.Set(constants.ToIdentityID.GetName(), testFromID.AsString())
	viper.Set(constants.ClassificationID.GetName(), testClassificationID.AsString())
	viper.Set(constants.MaintainedProperties.GetName(), maintainedPropertyString)
	viper.Set(constants.CanMintAsset.GetName(), true)
	viper.Set(constants.CanBurnAsset.GetName(), true)
	viper.Set(constants.CanRenumerateAsset.GetName(), true)
	viper.Set(constants.CanAddMaintainer.GetName(), true)
	viper.Set(constants.CanRemoveMaintainer.GetName(), true)
	viper.Set(constants.CanMutateMaintainer.GetName(), true)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ToID                     string
		ClassificationID         string
		MaintainedProperties     string
		CanMakeOrder             bool `json:"canMakeOrder"`
		CanCancelOrder           bool `json:"canCancelOrder"`
		CanAddMaintainer         bool `json:"canAddMaintainer"`
		CanRemoveMaintainer      bool `json:"canRemoveMaintainer"`
		CanMutateMaintainer      bool `json:"canMutateMaintainer"`
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
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), testFromID.AsString(), testClassificationID.AsString(), maintainedPropertyString, true, true, true, true, true}, args{cliCommand, client.Context{}.WithCodec(baseHelpers.CodecPrototype())}, newTransactionRequest(commonTransactionRequest, testFromID.AsString(), testFromID.AsString(), testClassificationID.AsString(), maintainedPropertyString, true, true, true, true, true), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
				MaintainedProperties:     tt.fields.MaintainedProperties,
				CanMakeOrder:             tt.fields.CanMakeOrder,
				CanCancelOrder:           tt.fields.CanCancelOrder,
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
	jsonMessage, err := json.Marshal(newTransactionRequest(commonTransactionRequest, testFromID.AsString(), testFromID.AsString(), testClassificationID.AsString(), maintainedPropertyString, true, true, true, true, true))
	require.NoError(t, err)
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ToID                     string
		ClassificationID         string
		MaintainedProperties     string
		CanMakeOrder             bool `json:"canMakeOrder"`
		CanCancelOrder           bool `json:"canCancelOrder"`
		CanAddMaintainer         bool `json:"canAddMaintainer"`
		CanRemoveMaintainer      bool `json:"canRemoveMaintainer"`
		CanMutateMaintainer      bool `json:"canMutateMaintainer"`
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
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), testFromID.AsString(), testClassificationID.AsString(), maintainedPropertyString, true, true, true, true, true}, args{jsonMessage}, newTransactionRequest(commonTransactionRequest, testFromID.AsString(), testFromID.AsString(), testClassificationID.AsString(), maintainedPropertyString, true, true, true, true, true), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
				MaintainedProperties:     tt.fields.MaintainedProperties,
				CanMakeOrder:             tt.fields.CanMakeOrder,
				CanCancelOrder:           tt.fields.CanCancelOrder,
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
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ToID                     string
		ClassificationID         string
		MaintainedProperties     string
		CanMakeOrder             bool `json:"canMakeOrder"`
		CanCancelOrder           bool `json:"canCancelOrder"`
		CanAddMaintainer         bool `json:"canAddMaintainer"`
		CanRemoveMaintainer      bool `json:"canRemoveMaintainer"`
		CanMutateMaintainer      bool `json:"canMutateMaintainer"`
	}
	tests := []struct {
		name   string
		fields fields
		want   helpers.CommonTransactionRequest
	}{
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), testFromID.AsString(), testClassificationID.AsString(), maintainedPropertyString, true, true, true, true, true}, commonTransactionRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
				MaintainedProperties:     tt.fields.MaintainedProperties,
				CanMakeOrder:             tt.fields.CanMakeOrder,
				CanCancelOrder:           tt.fields.CanCancelOrder,
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
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ToID                     string
		ClassificationID         string
		MaintainedProperties     string
		CanMakeOrder             bool `json:"canMakeOrder"`
		CanCancelOrder           bool `json:"canCancelOrder"`
		CanAddMaintainer         bool `json:"canAddMaintainer"`
		CanRemoveMaintainer      bool `json:"canRemoveMaintainer"`
		CanMutateMaintainer      bool `json:"canMutateMaintainer"`
	}
	tests := []struct {
		name    string
		fields  fields
		want    types.Msg
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), testFromID.AsString(), testClassificationID.AsString(), maintainedPropertyString, true, true, true, true, true}, NewMessage(fromAccAddress, testFromID, testFromID, testClassificationID, maintainedProperties, true, true, true, true, true), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
				MaintainedProperties:     tt.fields.MaintainedProperties,
				CanMakeOrder:             tt.fields.CanMakeOrder,
				CanCancelOrder:           tt.fields.CanCancelOrder,
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
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ToID                     string
		ClassificationID         string
		MaintainedProperties     string
		CanMakeOrder             bool `json:"canMakeOrder"`
		CanCancelOrder           bool `json:"canCancelOrder"`
		CanAddMaintainer         bool `json:"canAddMaintainer"`
		CanRemoveMaintainer      bool `json:"canRemoveMaintainer"`
		CanMutateMaintainer      bool `json:"canMutateMaintainer"`
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), testFromID.AsString(), testClassificationID.AsString(), maintainedPropertyString, true, true, true, true, true}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
				MaintainedProperties:     tt.fields.MaintainedProperties,
				CanMakeOrder:             tt.fields.CanMakeOrder,
				CanCancelOrder:           tt.fields.CanCancelOrder,
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
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		ToID                     string
		ClassificationID         string
		MaintainedProperties     string
		CanMakeOrder             bool `json:"canMakeOrder"`
		CanCancelOrder           bool `json:"canCancelOrder"`
		CanAddMaintainer         bool `json:"canAddMaintainer"`
		CanRemoveMaintainer      bool `json:"canRemoveMaintainer"`
		CanMutateMaintainer      bool `json:"canMutateMaintainer"`
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), testFromID.AsString(), testClassificationID.AsString(), maintainedPropertyString, true, true, true, true, true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				ToID:                     tt.fields.ToID,
				ClassificationID:         tt.fields.ClassificationID,
				MaintainedProperties:     tt.fields.MaintainedProperties,
				CanMakeOrder:             tt.fields.CanMakeOrder,
				CanCancelOrder:           tt.fields.CanCancelOrder,
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
