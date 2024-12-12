// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

var (
	immutables            = baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutableMetaProperties = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData())).(*baseLists.PropertyList)
	mutables              = baseQualified.NewMutables(mutableMetaProperties)
	classificationID      = baseIDs.NewClassificationID(immutables, mutables).(*baseIDs.ClassificationID)
	fromAddress           = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, _     = sdkTypes.AccAddressFromBech32(fromAddress)
	fromID                = baseIDs.NewIdentityID(classificationID, immutables).(*baseIDs.IdentityID)
)

type fields struct {
	From                 string
	FromID               *baseIDs.IdentityID
	ToID                 *baseIDs.IdentityID
	ClassificationID     *baseIDs.ClassificationID
	MaintainedProperties *baseLists.PropertyList
	CanMintAsset         bool
	CanBurnAsset         bool
	CanRenumerateAsset   bool
	CanAddMaintainer     bool
	CanRemoveMaintainer  bool
	CanMutateMaintainer  bool
}

func Test_messagePrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Message
	}{
		{"+ve", &Message{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := messagePrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("messagePrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_GetSigners(t *testing.T) {

	tests := []struct {
		name   string
		fields fields
		want   []sdkTypes.AccAddress
	}{
		{"+ve", fields{fromAccAddress.String(), fromID, fromID, classificationID, mutableMetaProperties, true, true, true, true, true, true}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:                 tt.fields.From,
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
			if got := message.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_ValidateBasic(t *testing.T) {

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"-ve", fields{}, true},
		{"+ve", fields{fromAccAddress.String(), fromID, fromID, classificationID, mutableMetaProperties, true, true, true, true, true, true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:                 tt.fields.From,
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
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_NewMessage(t *testing.T) {
	type args struct {
		from                 sdkTypes.AccAddress
		fromID               ids.IdentityID
		toID                 ids.IdentityID
		classificationID     ids.ClassificationID
		maintainedProperties lists.PropertyList
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
		want sdkTypes.Msg
	}{
		{"+ve", args{}, &Message{}},
		{"+ve", args{fromAccAddress, fromID, fromID, classificationID, mutableMetaProperties, true, true, true, true, true, true}, &Message{fromAccAddress.String(), fromID, fromID, classificationID, mutableMetaProperties, true, true, true, true, true, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessage(tt.args.from, tt.args.fromID, tt.args.toID, tt.args.classificationID, tt.args.maintainedProperties, tt.args.canMintAsset, tt.args.canBurnAsset, tt.args.canRenumerateAsset, tt.args.canAddMaintainer, tt.args.canRemoveMaintainer, tt.args.canMutateMaintainer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
