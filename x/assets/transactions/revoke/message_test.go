// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

var (
	immutableMetaProperties = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))
	immutables              = baseQualified.NewImmutables(immutableMetaProperties)
	mutableMetaProperties   = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData()))
	mutables                = baseQualified.NewMutables(mutableMetaProperties)
	classificationID        = baseIDs.NewClassificationID(immutables, mutables).(*baseIDs.ClassificationID)
	fromAddress             = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, _       = sdkTypes.AccAddressFromBech32(fromAddress)
	fromID                  = baseIDs.NewIdentityID(classificationID, immutables).(*baseIDs.IdentityID)
)

type fields struct {
	From             string
	FromID           *baseIDs.IdentityID
	ToID             *baseIDs.IdentityID
	ClassificationID *baseIDs.ClassificationID
}

func Test_messageFromInterface(t *testing.T) {
	type args struct {
		msg sdkTypes.Msg
	}
	tests := []struct {
		name string
		args args
		want helpers.Message
	}{
		{"+ve", args{NewMessage(fromAccAddress, fromID, fromID, classificationID)}, NewMessage(fromAccAddress, fromID, fromID, classificationID).(*Message)},
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
		{"+ve", fields{fromAccAddress.String(), fromID, fromID, classificationID}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:             tt.fields.From,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
			}
			if got := message.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_RegisterCodec(t *testing.T) {
	type args struct {
		legacyAmino *sdkCodec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{fromAccAddress.String(), fromID, fromID, classificationID}, args{sdkCodec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Message{
				From:             tt.fields.From,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
			}
			me.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_message_Type(t *testing.T) {
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{fromAccAddress.String(), fromID, fromID, classificationID}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:             tt.fields.From,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
			}
			if got := message.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
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
		{"+ve", fields{fromAccAddress.String(), fromID, fromID, classificationID}, false},
		{"-ve", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:             tt.fields.From,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_NewMessage(t *testing.T) {
	type args struct {
		from             sdkTypes.AccAddress
		fromID           ids.IdentityID
		toID             ids.IdentityID
		classificationID ids.ClassificationID
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{
		{"+ve with nil", args{}, &Message{}},
		{"+ve", args{fromAccAddress, fromID, fromID, classificationID}, NewMessage(fromAccAddress, fromID, fromID, classificationID)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessage(tt.args.from, tt.args.fromID, tt.args.toID, tt.args.classificationID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
