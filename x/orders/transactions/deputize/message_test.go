// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
)

type fields struct {
	From                 string
	FromID               *baseIDs.IdentityID
	ToID                 *baseIDs.IdentityID
	ClassificationID     *baseIDs.ClassificationID
	MaintainedProperties *base.PropertyList
	CanMakeOrder         bool
	CanCancelOrder       bool
	CanAddMaintainer     bool
	CanRemoveMaintainer  bool
	CanMutateMaintainer  bool
}

func CreateTestInputForMessage(t *testing.T) (*baseIDs.IdentityID, *baseIDs.IdentityID, *baseIDs.ClassificationID, sdkTypes.AccAddress, *base.PropertyList, sdkTypes.Msg) {
	immutables := baseQualified.NewImmutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	mutables2 := baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID"), baseData.NewStringData("Data3"))))

	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)

	testClassificationID2 := baseIDs.NewClassificationID(immutables, mutables2)
	testToID := baseIDs.NewIdentityID(testClassificationID2, immutables)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	maintainedProperty := "maintainedProperty:S|maintainedProperty"
	maintainedProperties, err := base.NewPropertyList().FromMetaPropertiesString(maintainedProperty)
	require.Equal(t, nil, err)

	testMessage := NewMessage(fromAccAddress, testFromID, testToID, testClassificationID, maintainedProperties, true, true, true, true, true)

	return testFromID.(*baseIDs.IdentityID), testToID.(*baseIDs.IdentityID), testClassificationID.(*baseIDs.ClassificationID), fromAccAddress, maintainedProperties.(*base.PropertyList), testMessage
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
	testFromID, testToID, testClassificationID, fromAccAddress, maintainedProperties, _ := CreateTestInputForMessage(t)

	tests := []struct {
		name   string
		fields fields
		want   []sdkTypes.AccAddress
	}{
		{"+ve", fields{fromAccAddress.String(), testFromID, testToID, testClassificationID, maintainedProperties, true, true, true, true, true}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				tt.fields.From,
				tt.fields.FromID,
				tt.fields.ToID,
				tt.fields.ClassificationID,
				tt.fields.MaintainedProperties,
				tt.fields.CanMakeOrder,
				tt.fields.CanCancelOrder,
				tt.fields.CanAddMaintainer,
				tt.fields.CanRemoveMaintainer,
				tt.fields.CanMutateMaintainer,
			}
			if got := message.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_ValidateBasic(t *testing.T) {
	testFromID, testToID, testClassificationID, fromAccAddress, maintainedProperties, _ := CreateTestInputForMessage(t)

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{fromAccAddress.String(), testFromID, testToID, testClassificationID, maintainedProperties, true, true, true, true, true}, false},
		{"-ve", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				tt.fields.From,
				tt.fields.FromID,
				tt.fields.ToID,
				tt.fields.ClassificationID,
				tt.fields.MaintainedProperties,
				tt.fields.CanMakeOrder,
				tt.fields.CanCancelOrder,
				tt.fields.CanAddMaintainer,
				tt.fields.CanRemoveMaintainer,
				tt.fields.CanMutateMaintainer,
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_NewMessage(t *testing.T) {
	testFromID, testToID, testClassificationID, fromAccAddress, maintainedProperties, _ := CreateTestInputForMessage(t)

	type args struct {
		From                 sdkTypes.AccAddress
		FromID               *baseIDs.IdentityID
		ToID                 *baseIDs.IdentityID
		ClassificationID     *baseIDs.ClassificationID
		MaintainedProperties *base.PropertyList
		CanMakeOrder         bool
		CanCancelOrder       bool
		CanAddMaintainer     bool
		CanRemoveMaintainer  bool
		CanMutateMaintainer  bool
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{
		{"+ve", args{fromAccAddress, testFromID, testToID, testClassificationID, maintainedProperties, true, true, true, true, true}, &Message{fromAccAddress.String(), testFromID, testToID, testClassificationID, maintainedProperties, true, true, true, true, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessage(tt.args.From, tt.args.FromID, tt.args.ToID, tt.args.ClassificationID, tt.args.MaintainedProperties, tt.args.CanMakeOrder, tt.args.CanCancelOrder, tt.args.CanAddMaintainer, tt.args.CanRemoveMaintainer, tt.args.CanMutateMaintainer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
