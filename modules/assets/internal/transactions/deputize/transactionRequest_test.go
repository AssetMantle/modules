// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/spf13/viper"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/utilities/transaction"
)

var (
	testBaseRequest             = rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	mutableMetaPropertiesString = "testMutableMeta1:S|mutableMeta"
	mutableMetaProperties1      = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("testMutableMeta1"), baseData.NewStringData("mutableMeta")))
)

func Test_newTransactionRequest(t *testing.T) {
	type args struct {
		baseReq              rest.BaseReq
		fromID               string
		toID                 string
		classificationID     string
		maintainedProperties string
		canMintAsset         bool
		canBurnAsset         bool
		canRenumerateAsset   bool
		canAddMaintainer     bool
		canRemoveMaintainer  bool
		canMutateMaintainer  bool
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve", args{testBaseRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), fmt.Sprint(mutableMetaProperties), true, true, true, true, true, true}, transactionRequest{testBaseRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), fmt.Sprint(mutableMetaProperties), true, true, true, true, true, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.baseReq, tt.args.fromID, tt.args.toID, tt.args.classificationID, tt.args.maintainedProperties, tt.args.canMintAsset, tt.args.canBurnAsset, tt.args.canRenumerateAsset, tt.args.canAddMaintainer, tt.args.canRemoveMaintainer, tt.args.canMutateMaintainer); !reflect.DeepEqual(got, tt.want) {
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
	cliCommand := base.NewCLICommand("", "", "", []helpers.CLIFlag{constants.ToID, constants.FromID, constants.ClassificationID, constants.MaintainedProperties, constants.CanMintAsset, constants.CanBurnAsset, constants.CanRenumerateAsset, constants.CanAddMaintainer, constants.CanRemoveMaintainer, constants.CanMutateMaintainer})

	viper.Set(constants.ToID.GetName(), fromID.AsString())
	viper.Set(constants.FromID.GetName(), fromID.AsString())
	viper.Set(constants.ClassificationID.GetName(), classificationID.AsString())
	viper.Set(constants.MaintainedProperties.GetName(), mutableMetaPropertiesString)
	viper.Set(constants.CanMintAsset.GetName(), true)
	viper.Set(constants.CanBurnAsset.GetName(), true)
	viper.Set(constants.CanRenumerateAsset.GetName(), true)
	viper.Set(constants.CanAddMaintainer.GetName(), true)
	viper.Set(constants.CanRemoveMaintainer.GetName(), true)
	viper.Set(constants.CanMutateMaintainer.GetName(), true)
	type fields struct {
		BaseReq              rest.BaseReq
		FromID               string
		ToID                 string
		ClassificationID     string
		MaintainedProperties string
		CanMintAsset         bool
		CanBurnAsset         bool
		CanRenumerateAsset   bool
		CanAddMaintainer     bool
		CanRemoveMaintainer  bool
		CanMutateMaintainer  bool
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
		{"+ve", fields{}, args{cliCommand, base.TestClientContext}, transactionRequest{testBaseRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), mutableMetaPropertiesString, true, true, true, true, true, true}, false},
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
			if !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("FromCLI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_FromJSON(t *testing.T) {
	type fields struct {
		BaseReq              rest.BaseReq
		FromID               string
		ToID                 string
		ClassificationID     string
		MaintainedProperties string
		CanMintAsset         bool
		CanBurnAsset         bool
		CanRenumerateAsset   bool
		CanAddMaintainer     bool
		CanRemoveMaintainer  bool
		CanMutateMaintainer  bool
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
		{"+ve", fields{testBaseRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), fmt.Sprint(mutableMetaProperties), true, true, true, true, true, true}, args{sdkTypes.MustSortJSON(transaction.RegisterLegacyAminoCodec(messagePrototype).MustMarshalJSON(&Message{fromAccAddress.String(), fromID, fromID, classificationID, mutableMetaProperties, true, true, true, true, true, true}))}, transactionRequest{testBaseRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), fmt.Sprint(mutableMetaProperties), true, true, true, true, true, true}, false},
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
	type fields struct {
		BaseReq              rest.BaseReq
		FromID               string
		ToID                 string
		ClassificationID     string
		MaintainedProperties string
		CanMintAsset         bool
		CanBurnAsset         bool
		CanRenumerateAsset   bool
		CanAddMaintainer     bool
		CanRemoveMaintainer  bool
		CanMutateMaintainer  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   rest.BaseReq
	}{
		{"+ve", fields{testBaseRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), fmt.Sprint(mutableMetaProperties), true, true, true, true, true, true}, testBaseRequest},
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
	type fields struct {
		BaseReq              rest.BaseReq
		FromID               string
		ToID                 string
		ClassificationID     string
		MaintainedProperties string
		CanMintAsset         bool
		CanBurnAsset         bool
		CanRenumerateAsset   bool
		CanAddMaintainer     bool
		CanRemoveMaintainer  bool
		CanMutateMaintainer  bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    sdkTypes.Msg
		wantErr bool
	}{
		{"+ve", fields{testBaseRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), mutableMetaPropertiesString, true, true, true, true, true, true}, newMessage(fromAccAddress, fromID, fromID, classificationID, mutableMetaProperties1, true, true, true, true, true, true), false},
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
	type fields struct {
		BaseReq              rest.BaseReq
		FromID               string
		ToID                 string
		ClassificationID     string
		MaintainedProperties string
		CanMintAsset         bool
		CanBurnAsset         bool
		CanRenumerateAsset   bool
		CanAddMaintainer     bool
		CanRemoveMaintainer  bool
		CanMutateMaintainer  bool
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testBaseRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), fmt.Sprint(mutableMetaProperties), true, true, true, true, true, true}, args{codec.NewLegacyAmino()}},
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
	type fields struct {
		BaseReq              rest.BaseReq
		FromID               string
		ToID                 string
		ClassificationID     string
		MaintainedProperties string
		CanMintAsset         bool
		CanBurnAsset         bool
		CanRenumerateAsset   bool
		CanAddMaintainer     bool
		CanRemoveMaintainer  bool
		CanMutateMaintainer  bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve with nil", fields{}, true},
		{"+ve", fields{testBaseRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), fmt.Sprint(mutableMetaProperties), true, true, true, true, true, true}, false},
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
