// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerate

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/x/data/base"
	"github.com/AssetMantle/schema/x/helpers"
	"github.com/AssetMantle/schema/x/ids"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	baseLists "github.com/AssetMantle/schema/x/lists/base"
	baseProperties "github.com/AssetMantle/schema/x/properties/base"
	baseQualified "github.com/AssetMantle/schema/x/qualified/base"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/assets/internal/module"
	"github.com/AssetMantle/modules/utilities/transaction"
)

var (
	immutableMetaProperties = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))).(*baseLists.PropertyList)
	immutables              = baseQualified.NewImmutables(immutableMetaProperties)
	mutableMetaProperties   = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData())).(*baseLists.PropertyList)
	mutables                = baseQualified.NewMutables(mutableMetaProperties)
	classificationID        = baseIDs.NewClassificationID(immutables, mutables).(*baseIDs.ClassificationID)
	fromAddress             = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, _       = sdkTypes.AccAddressFromBech32(fromAddress)
	fromID                  = baseIDs.NewIdentityID(classificationID, immutables).(*baseIDs.IdentityID)
	testAssetID             = baseIDs.NewAssetID(classificationID, immutables).(*baseIDs.AssetID)
)

type fields struct {
	From    string
	FromID  *baseIDs.IdentityID
	AssetID *baseIDs.AssetID
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
		{"+ve", args{newMessage(fromAccAddress, fromID, testAssetID)}, newMessage(fromAccAddress, fromID, testAssetID).(*Message)},
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

func Test_message_GetSignBytes(t *testing.T) {

	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve", fields{fromAccAddress.String(), fromID, testAssetID}, sdkTypes.MustSortJSON(transaction.RegisterLegacyAminoCodec(messagePrototype).MustMarshalJSON(newMessage(fromAccAddress, fromID, testAssetID)))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				AssetID: tt.fields.AssetID,
			}
			if got := message.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSignBytes() = %v, want %v", got, tt.want)
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
		{"+ve", fields{fromAccAddress.String(), fromID, testAssetID}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				AssetID: tt.fields.AssetID,
			}
			if got := message.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_RegisterCodec(t *testing.T) {

	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{fromAccAddress.String(), fromID, testAssetID}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				AssetID: tt.fields.AssetID,
			}
			me.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_message_Route(t *testing.T) {

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{fromAccAddress.String(), fromID, testAssetID}, module.Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				AssetID: tt.fields.AssetID,
			}
			if got := message.Route(); got != tt.want {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_Type(t *testing.T) {

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{fromAccAddress.String(), fromID, testAssetID}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				AssetID: tt.fields.AssetID,
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
		{"+ve", fields{fromAccAddress.String(), fromID, testAssetID}, false},
		{"+ve with Nil", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				AssetID: tt.fields.AssetID,
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newMessage(t *testing.T) {
	type args struct {
		from    sdkTypes.AccAddress
		fromID  ids.IdentityID
		assetID ids.AssetID
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{
		{"+ve", args{fromAccAddress, fromID, testAssetID}, &Message{fromAccAddress.String(), fromID, testAssetID}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMessage(tt.args.from, tt.args.fromID, tt.args.assetID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
