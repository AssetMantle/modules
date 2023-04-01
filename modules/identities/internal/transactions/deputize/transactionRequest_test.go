// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
)

func Test_newTransactionRequest(t *testing.T) {

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	maintainedProperty := "maintainedProperties:S|maintainedProperties"
	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	type args struct {
		baseReq              rest.BaseReq
		fromID               string
		toID                 string
		classificationID     string
		maintainedProperties string
		CanMintAsset         bool `json:"canMintAsset"`
		CanBurnAsset         bool `json:"canBurnAsset"`
		CanRenumerateAsset   bool `json:"canRenumerateAsset"`
		CanAddMaintainer     bool `json:"canAddMaintainer"`
		CanRemoveMaintainer  bool `json:"canRemoveMaintainer"`
		CanMutateMaintainer  bool `json:"canMutateMaintainer"`
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve", args{testBaseReq, "fromID", "toID", "classificationID", maintainedProperty, false, false, false, false, false, false}, transactionRequest{testBaseReq, "fromID", "toID", "classificationID", maintainedProperty, false, false, false, false, false, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.baseReq, tt.args.fromID, tt.args.toID, tt.args.classificationID, tt.args.maintainedProperties, tt.args.CanMintAsset, tt.args.CanBurnAsset, tt.args.CanRenumerateAsset, tt.args.CanAddMaintainer, tt.args.CanRemoveMaintainer, tt.args.CanMutateMaintainer); !reflect.DeepEqual(got, tt.want) {
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
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromID, constants.ToID, constants.ClassificationID, constants.MaintainedProperties, constants.CanMintAsset, constants.CanBurnAsset, constants.CanRenumerateAsset, constants.CanAddMaintainer, constants.CanRemoveMaintainer, constants.CanMutateMaintainer})

	type fields struct {
		BaseReq              rest.BaseReq
		FromID               string
		ToID                 string
		ClassificationID     string
		MaintainedProperties string
		CanMintAsset         bool `json:"canMintAsset"`
		CanBurnAsset         bool `json:"canBurnAsset"`
		CanRenumerateAsset   bool `json:"canRenumerateAsset"`
		CanAddMaintainer     bool `json:"canAddMaintainer"`
		CanRemoveMaintainer  bool `json:"canRemoveMaintainer"`
		CanMutateMaintainer  bool `json:"canMutateMaintainer"`
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
		{"+ve", fields{BaseReq: rest.BaseReq{From: baseHelpers.TestClientContext.GetFromAddress().String(), ChainID: baseHelpers.TestClientContext.ChainID, Simulate: baseHelpers.TestClientContext.Simulate}, FromID: "", ToID: "", ClassificationID: "", MaintainedProperties: "", CanMintAsset: false, CanBurnAsset: false, CanRenumerateAsset: false, CanAddMaintainer: false, CanRemoveMaintainer: false, CanMutateMaintainer: false}, args{cliCommand, baseHelpers.TestClientContext}, transactionRequest{BaseReq: rest.BaseReq{From: baseHelpers.TestClientContext.GetFromAddress().String(), ChainID: baseHelpers.TestClientContext.ChainID, Simulate: baseHelpers.TestClientContext.Simulate}, FromID: "", ToID: "", ClassificationID: "", MaintainedProperties: "", CanAddMaintainer: false, CanRemoveMaintainer: false, CanMutateMaintainer: false}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:              tt.fields.BaseReq,
				FromID:               tt.fields.FromID,
				ToID:                 tt.fields.ToID,
				ClassificationID:     tt.fields.ClassificationID,
				MaintainedProperties: tt.fields.MaintainedProperties,
				CanMintAsset:         tt.fields.CanMintAsset,
				CanBurnAsset:         tt.fields.CanBurnAsset,
				CanRenumerateAsset:   tt.fields.CanRenumerateAsset,
				CanAddMaintainer:     tt.fields.CanAddMaintainer,
				CanRemoveMaintainer:  tt.fields.CanRemoveMaintainer,
				CanMutateMaintainer:  tt.fields.CanMutateMaintainer,
			}
			got, err := transactionRequest.FromCLI(tt.args.cliCommand, tt.args.context)
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

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	maintainedProperty := "maintainedProperties:S|maintainedProperties"
	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	jsonMessage, _ := json.Marshal(newTransactionRequest(testBaseReq, "fromID", "toID", "classificationID", maintainedProperty, false, false, false, false, false, false))
	type fields struct {
		BaseReq              rest.BaseReq
		FromID               string
		ToID                 string
		ClassificationID     string
		MaintainedProperties string
		CanMintAsset         bool `json:"canMintAsset"`
		CanBurnAsset         bool `json:"canBurnAsset"`
		CanRenumerateAsset   bool `json:"canRenumerateAsset"`
		CanAddMaintainer     bool `json:"canAddMaintainer"`
		CanRemoveMaintainer  bool `json:"canRemoveMaintainer"`
		CanMutateMaintainer  bool `json:"canMutateMaintainer"`
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
		{"+ve", fields{testBaseReq, "fromID", "toID", "classificationID", maintainedProperty, false, false, false, false, false, false}, args{jsonMessage}, transactionRequest{testBaseReq, "fromID", "toID", "classificationID", maintainedProperty, false, false, false, false, false, false}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:              tt.fields.BaseReq,
				FromID:               tt.fields.FromID,
				ToID:                 tt.fields.ToID,
				ClassificationID:     tt.fields.ClassificationID,
				MaintainedProperties: tt.fields.MaintainedProperties,
				CanMintAsset:         tt.fields.CanMintAsset,
				CanBurnAsset:         tt.fields.CanBurnAsset,
				CanRenumerateAsset:   tt.fields.CanRenumerateAsset,
				CanAddMaintainer:     tt.fields.CanAddMaintainer,
				CanRemoveMaintainer:  tt.fields.CanRemoveMaintainer,
				CanMutateMaintainer:  tt.fields.CanMutateMaintainer,
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

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	maintainedProperty := "maintainedProperties:S|maintainedProperties"
	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	type fields struct {
		BaseReq              rest.BaseReq
		FromID               string
		ToID                 string
		ClassificationID     string
		MaintainedProperties string
		CanMintAsset         bool `json:"canMintAsset"`
		CanBurnAsset         bool `json:"canBurnAsset"`
		CanRenumerateAsset   bool `json:"canRenumerateAsset"`
		CanAddMaintainer     bool `json:"canAddMaintainer"`
		CanRemoveMaintainer  bool `json:"canRemoveMaintainer"`
		CanMutateMaintainer  bool `json:"canMutateMaintainer"`
	}
	tests := []struct {
		name   string
		fields fields
		want   rest.BaseReq
	}{
		{"+ve", fields{testBaseReq, "fromID", "toID", "classificationID", maintainedProperty, false, false, false, false, false, false}, testBaseReq},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:              tt.fields.BaseReq,
				FromID:               tt.fields.FromID,
				ToID:                 tt.fields.ToID,
				ClassificationID:     tt.fields.ClassificationID,
				MaintainedProperties: tt.fields.MaintainedProperties,
				CanMintAsset:         tt.fields.CanMintAsset,
				CanBurnAsset:         tt.fields.CanBurnAsset,
				CanRenumerateAsset:   tt.fields.CanRenumerateAsset,
				CanAddMaintainer:     tt.fields.CanAddMaintainer,
				CanRemoveMaintainer:  tt.fields.CanRemoveMaintainer,
				CanMutateMaintainer:  tt.fields.CanMutateMaintainer,
			}
			if got := transactionRequest.GetBaseReq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBaseReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {

	testFromID, testToID, testClassificationID, fromAccAddress, maintainedProperties := createTestInput(t)

	testBaseReq := rest.BaseReq{From: "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c", ChainID: "test", Fees: sdkTypes.NewCoins()}
	type fields struct {
		BaseReq              rest.BaseReq
		FromID               string
		ToID                 string
		ClassificationID     string
		MaintainedProperties string
		CanMintAsset         bool `json:"canMintAsset"`
		CanBurnAsset         bool `json:"canBurnAsset"`
		CanRenumerateAsset   bool `json:"canRenumerateAsset"`
		CanAddMaintainer     bool `json:"canAddMaintainer"`
		CanRemoveMaintainer  bool `json:"canRemoveMaintainer"`
		CanMutateMaintainer  bool `json:"canMutateMaintainer"`
	}
	tests := []struct {
		name    string
		fields  fields
		want    sdkTypes.Msg
		wantErr bool
	}{
		{"+ve", fields{testBaseReq, testFromID.AsString(), testFromID.AsString(), testClassificationID.AsString(), "maintainedProperty:S|maintainedProperty", false, false, false, false, false, false}, newMessage(fromAccAddress, testFromID, testToID, testClassificationID, maintainedProperties, false, false, false, false, false, false), false}, // TODO: issue==> getting MetaDataError that is not expected
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:              tt.fields.BaseReq,
				FromID:               tt.fields.FromID,
				ToID:                 tt.fields.ToID,
				ClassificationID:     tt.fields.ClassificationID,
				MaintainedProperties: tt.fields.MaintainedProperties,
				CanMintAsset:         tt.fields.CanMintAsset,
				CanBurnAsset:         tt.fields.CanBurnAsset,
				CanRenumerateAsset:   tt.fields.CanRenumerateAsset,
				CanAddMaintainer:     tt.fields.CanAddMaintainer,
				CanRemoveMaintainer:  tt.fields.CanRemoveMaintainer,
				CanMutateMaintainer:  tt.fields.CanMutateMaintainer,
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

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	maintainedProperty := "maintainedProperties:S|maintainedProperties"
	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	type fields struct {
		BaseReq              rest.BaseReq
		FromID               string
		ToID                 string
		ClassificationID     string
		MaintainedProperties string
		CanMintAsset         bool `json:"canMintAsset"`
		CanBurnAsset         bool `json:"canBurnAsset"`
		CanRenumerateAsset   bool `json:"canRenumerateAsset"`
		CanAddMaintainer     bool `json:"canAddMaintainer"`
		CanRemoveMaintainer  bool `json:"canRemoveMaintainer"`
		CanMutateMaintainer  bool `json:"canMutateMaintainer"`
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testBaseReq, "fromID", "toID", "classificationID", maintainedProperty, false, false, false, false, false, false}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				BaseReq:              tt.fields.BaseReq,
				FromID:               tt.fields.FromID,
				ToID:                 tt.fields.ToID,
				ClassificationID:     tt.fields.ClassificationID,
				MaintainedProperties: tt.fields.MaintainedProperties,
				CanMintAsset:         tt.fields.CanMintAsset,
				CanBurnAsset:         tt.fields.CanBurnAsset,
				CanRenumerateAsset:   tt.fields.CanRenumerateAsset,
				CanAddMaintainer:     tt.fields.CanAddMaintainer,
				CanRemoveMaintainer:  tt.fields.CanRemoveMaintainer,
				CanMutateMaintainer:  tt.fields.CanMutateMaintainer,
			}
			tr.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	maintainedProperty := "maintainedProperties:S|maintainedProperties"
	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	type fields struct {
		BaseReq              rest.BaseReq
		FromID               string
		ToID                 string
		ClassificationID     string
		MaintainedProperties string
		CanMintAsset         bool `json:"canMintAsset"`
		CanBurnAsset         bool `json:"canBurnAsset"`
		CanRenumerateAsset   bool `json:"canRenumerateAsset"`
		CanAddMaintainer     bool `json:"canAddMaintainer"`
		CanRemoveMaintainer  bool `json:"canRemoveMaintainer"`
		CanMutateMaintainer  bool `json:"canMutateMaintainer"`
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{testBaseReq, "fromID", "toID", "classificationID", maintainedProperty, false, false, false, false, false, false}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:              tt.fields.BaseReq,
				FromID:               tt.fields.FromID,
				ToID:                 tt.fields.ToID,
				ClassificationID:     tt.fields.ClassificationID,
				MaintainedProperties: tt.fields.MaintainedProperties,
				CanMintAsset:         tt.fields.CanMintAsset,
				CanBurnAsset:         tt.fields.CanBurnAsset,
				CanRenumerateAsset:   tt.fields.CanRenumerateAsset,
				CanAddMaintainer:     tt.fields.CanAddMaintainer,
				CanRemoveMaintainer:  tt.fields.CanRemoveMaintainer,
				CanMutateMaintainer:  tt.fields.CanMutateMaintainer,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
