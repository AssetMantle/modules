// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"fmt"
	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/spf13/viper"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/stretchr/testify/assert"
)

var (
	commonTransactionRequest    = helpers.PrototypeCommonTransactionRequest()
	mutableMetaPropertiesString = "testMutableMeta1:S|mutableMeta"
	mutableMetaProperties1      = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("testMutableMeta1"), baseData.NewStringData("mutableMeta")))
)

func Test_newTransactionRequest(t *testing.T) {
	type args struct {
		commonTransactionRequest helpers.CommonTransactionRequest
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
		{"valid", args{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), fmt.Sprint(mutableMetaProperties), true, true, true, true, true, true}, transactionRequest{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), fmt.Sprint(mutableMetaProperties), true, true, true, true, true, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newTransactionRequest(tt.args.commonTransactionRequest, tt.args.fromID, tt.args.toID, tt.args.classificationID, tt.args.maintainedProperties, tt.args.canMintAsset, tt.args.canBurnAsset, tt.args.canRenumerateAsset, tt.args.canAddMaintainer, tt.args.canRemoveMaintainer, tt.args.canMutateMaintainer)
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
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.ToIdentityID, constants.FromIdentityID, constants.ClassificationID, constants.MaintainedProperties, constants.CanMintAsset, constants.CanBurnAsset, constants.CanRenumerateAsset, constants.CanAddMaintainer, constants.CanRemoveMaintainer, constants.CanMutateMaintainer})

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
		commonTransactionRequest helpers.CommonTransactionRequest
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
		{"valid", fields{}, args{cliCommand, client.Context{}.WithCodec(baseHelpers.CodecPrototype())}, transactionRequest{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), mutableMetaPropertiesString, true, true, true, true, true, true}, false},
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
			assert.Equal(t, fmt.Sprint(tt.want), fmt.Sprint(got), "FromCLI()")
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
		want   helpers.CommonTransactionRequest
	}{
		{"valid", fields{commonTransactionRequest, fromID.AsString(), fromID.AsString(), classificationID.AsString(), fmt.Sprint(mutableMetaProperties), true, true, true, true, true, true}, commonTransactionRequest},
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
			got := transactionRequest.GetCommonTransactionRequest()
			assert.Equal(t, tt.want, got, "GetCommonTransactionRequest()")
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
		{"valid", fields{commonTransactionRequest.SetFrom(fromAddress), fromID.AsString(), fromID.AsString(), classificationID.AsString(), mutableMetaPropertiesString, true, true, true, true, true, true}, NewMessage(fromAccAddress, fromID, fromID, classificationID, mutableMetaProperties1, true, true, true, true, true, true), false},
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
			assert.Equal(t, tt.want, got, "MakeMsg() got")
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
		{"nil inputs", fields{}, true},
		{"valid", fields{commonTransactionRequest.SetFrom(fromAddress), fromID.AsString(), fromID.AsString(), classificationID.AsString(), mutableMetaPropertiesString, true, true, true, true, true, true}, false},
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
