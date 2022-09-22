// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"github.com/AssetMantle/modules/modules/identities/internal/module"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/utilities/transaction"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func createTestInput(t *testing.T) (ids.IdentityID, ids.IdentityID, ids.ClassificationID, sdkTypes.AccAddress, lists.PropertyList) {
	immutables := baseQualified.NewImmutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)
	testToID := baseIDs.NewIdentityID(testClassificationID, immutables)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	maintainedProperty := "maintainedProperty:S|maintainedProperty"
	maintainedProperties, err := utilities.ReadProperties(maintainedProperty)
	return testFromID, testToID, testClassificationID, fromAccAddress, maintainedProperties
}

func Test_messageFromInterface(t *testing.T) {
	testFromID, testToID, testClassificationID, fromAccAddress, maintainedProperties := createTestInput(t)
	type args struct {
		msg sdkTypes.Msg
	}
	tests := []struct {
		name string
		args args
		want message
	}{

		{"+ve", args{message{fromAccAddress, testFromID, testToID, testClassificationID, maintainedProperties, false, false, false}}, message{fromAccAddress, testFromID, testToID, testClassificationID, maintainedProperties, false, false, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := messageFromInterface(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("messageFromInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_messagePrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Message
	}{

		{"+ve", message{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := messagePrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("messagePrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_GetSignBytes(t *testing.T) {
	testFromID, testToID, testClassificationID, fromAccAddress, maintainedProperties := createTestInput(t)
	type fields struct {
		From                 sdkTypes.AccAddress
		FromID               ids.IdentityID
		ToID                 ids.IdentityID
		ClassificationID     ids.ClassificationID
		MaintainedProperties lists.PropertyList
		AddMaintainer        bool
		RemoveMaintainer     bool
		MutateMaintainer     bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{

		{"+ve", fields{fromAccAddress, testFromID, testToID, testClassificationID, maintainedProperties, false, false, false}, sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(message{fromAccAddress, testFromID, testToID, testClassificationID, maintainedProperties, false, false, false}))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                 tt.fields.From,
				FromID:               tt.fields.FromID,
				ToID:                 tt.fields.ToID,
				ClassificationID:     tt.fields.ClassificationID,
				MaintainedProperties: tt.fields.MaintainedProperties,
				AddMaintainer:        tt.fields.AddMaintainer,
				RemoveMaintainer:     tt.fields.RemoveMaintainer,
				MutateMaintainer:     tt.fields.MutateMaintainer,
			}
			if got := message.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSignBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_GetSigners(t *testing.T) {
	testFromID, testToID, testClassificationID, fromAccAddress, maintainedProperties := createTestInput(t)
	type fields struct {
		From                 sdkTypes.AccAddress
		FromID               ids.IdentityID
		ToID                 ids.IdentityID
		ClassificationID     ids.ClassificationID
		MaintainedProperties lists.PropertyList
		AddMaintainer        bool
		RemoveMaintainer     bool
		MutateMaintainer     bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []sdkTypes.AccAddress
	}{

		{"+ve", fields{fromAccAddress, testFromID, testToID, testClassificationID, maintainedProperties, false, false, false}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                 tt.fields.From,
				FromID:               tt.fields.FromID,
				ToID:                 tt.fields.ToID,
				ClassificationID:     tt.fields.ClassificationID,
				MaintainedProperties: tt.fields.MaintainedProperties,
				AddMaintainer:        tt.fields.AddMaintainer,
				RemoveMaintainer:     tt.fields.RemoveMaintainer,
				MutateMaintainer:     tt.fields.MutateMaintainer,
			}
			if got := message.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_RegisterCodec(t *testing.T) {
	testFromID, testToID, testClassificationID, fromAccAddress, maintainedProperties := createTestInput(t)
	type fields struct {
		From                 sdkTypes.AccAddress
		FromID               ids.IdentityID
		ToID                 ids.IdentityID
		ClassificationID     ids.ClassificationID
		MaintainedProperties lists.PropertyList
		AddMaintainer        bool
		RemoveMaintainer     bool
		MutateMaintainer     bool
	}
	type args struct {
		codec *codec.Codec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{

		{"+ve", fields{fromAccAddress, testFromID, testToID, testClassificationID, maintainedProperties, false, false, false}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := message{
				From:                 tt.fields.From,
				FromID:               tt.fields.FromID,
				ToID:                 tt.fields.ToID,
				ClassificationID:     tt.fields.ClassificationID,
				MaintainedProperties: tt.fields.MaintainedProperties,
				AddMaintainer:        tt.fields.AddMaintainer,
				RemoveMaintainer:     tt.fields.RemoveMaintainer,
				MutateMaintainer:     tt.fields.MutateMaintainer,
			}
			me.RegisterCodec(tt.args.codec)
		})
	}
}

func Test_message_Route(t *testing.T) {
	testFromID, testToID, testClassificationID, fromAccAddress, maintainedProperties := createTestInput(t)
	type fields struct {
		From                 sdkTypes.AccAddress
		FromID               ids.IdentityID
		ToID                 ids.IdentityID
		ClassificationID     ids.ClassificationID
		MaintainedProperties lists.PropertyList
		AddMaintainer        bool
		RemoveMaintainer     bool
		MutateMaintainer     bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{

		{"+ve", fields{fromAccAddress, testFromID, testToID, testClassificationID, maintainedProperties, false, false, false}, module.Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                 tt.fields.From,
				FromID:               tt.fields.FromID,
				ToID:                 tt.fields.ToID,
				ClassificationID:     tt.fields.ClassificationID,
				MaintainedProperties: tt.fields.MaintainedProperties,
				AddMaintainer:        tt.fields.AddMaintainer,
				RemoveMaintainer:     tt.fields.RemoveMaintainer,
				MutateMaintainer:     tt.fields.MutateMaintainer,
			}
			if got := message.Route(); got != tt.want {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_Type(t *testing.T) {
	testFromID, testToID, testClassificationID, fromAccAddress, maintainedProperties := createTestInput(t)
	type fields struct {
		From                 sdkTypes.AccAddress
		FromID               ids.IdentityID
		ToID                 ids.IdentityID
		ClassificationID     ids.ClassificationID
		MaintainedProperties lists.PropertyList
		AddMaintainer        bool
		RemoveMaintainer     bool
		MutateMaintainer     bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{

		{"+ve", fields{fromAccAddress, testFromID, testToID, testClassificationID, maintainedProperties, false, false, false}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                 tt.fields.From,
				FromID:               tt.fields.FromID,
				ToID:                 tt.fields.ToID,
				ClassificationID:     tt.fields.ClassificationID,
				MaintainedProperties: tt.fields.MaintainedProperties,
				AddMaintainer:        tt.fields.AddMaintainer,
				RemoveMaintainer:     tt.fields.RemoveMaintainer,
				MutateMaintainer:     tt.fields.MutateMaintainer,
			}
			if got := message.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_ValidateBasic(t *testing.T) {
	testFromID, testToID, testClassificationID, fromAccAddress, maintainedProperties := createTestInput(t)
	type fields struct {
		From                 sdkTypes.AccAddress
		FromID               ids.IdentityID
		ToID                 ids.IdentityID
		ClassificationID     ids.ClassificationID
		MaintainedProperties lists.PropertyList
		AddMaintainer        bool
		RemoveMaintainer     bool
		MutateMaintainer     bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{

		{"+ve", fields{fromAccAddress, testFromID, testToID, testClassificationID, maintainedProperties, false, false, false}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                 tt.fields.From,
				FromID:               tt.fields.FromID,
				ToID:                 tt.fields.ToID,
				ClassificationID:     tt.fields.ClassificationID,
				MaintainedProperties: tt.fields.MaintainedProperties,
				AddMaintainer:        tt.fields.AddMaintainer,
				RemoveMaintainer:     tt.fields.RemoveMaintainer,
				MutateMaintainer:     tt.fields.MutateMaintainer,
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newMessage(t *testing.T) {
	testFromID, testToID, testClassificationID, fromAccAddress, maintainedProperties := createTestInput(t)
	type args struct {
		from                 sdkTypes.AccAddress
		fromID               ids.IdentityID
		toID                 ids.IdentityID
		classificationID     ids.ClassificationID
		maintainedProperties lists.PropertyList
		addMaintainer        bool
		removeMaintainer     bool
		mutateMaintainer     bool
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{
		// TODO: Add test cases.
		{"+ve", args{fromAccAddress, testFromID, testToID, testClassificationID, maintainedProperties, false, false, false}, message{fromAccAddress, testFromID, testToID, testClassificationID, maintainedProperties, false, false, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMessage(tt.args.from, tt.args.fromID, tt.args.toID, tt.args.classificationID, tt.args.maintainedProperties, tt.args.addMaintainer, tt.args.removeMaintainer, tt.args.mutateMaintainer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
