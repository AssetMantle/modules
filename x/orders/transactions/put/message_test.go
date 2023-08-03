// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package put

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/lists"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	"github.com/AssetMantle/schema/go/types"
	"github.com/AssetMantle/schema/go/types/base"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

var (
	expiresInHeight = base.NewHeight(60).(*base.Height)
	testMessage     = NewMessage(fromAccAddress, testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties.(*baseLists.PropertyList), immutableProperties, mutableMetaProperties, mutableProperties)
)

type fields struct {
	From                    string
	FromID                  *baseIDs.IdentityID
	ClassificationID        *baseIDs.ClassificationID
	TakerID                 *baseIDs.IdentityID
	MakerOwnableID          *baseIDs.AnyOwnableID
	TakerOwnableID          *baseIDs.AnyOwnableID
	ExpiresIn               *base.Height
	MakerOwnableSplit       sdkTypes.Dec
	TakerOwnableSplit       sdkTypes.Dec
	ImmutableMetaProperties *baseLists.PropertyList
	ImmutableProperties     *baseLists.PropertyList
	MutableMetaProperties   *baseLists.PropertyList
	MutableProperties       *baseLists.PropertyList
}

func Test_messageFromInterface(t *testing.T) {
	type args struct {
		msg sdkTypes.Msg
	}
	tests := []struct {
		name string
		args args
		want *Message
	}{
		{"+ve", args{testMessage}, &Message{fromAccAddress.String(), testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties.(*baseLists.PropertyList), immutableProperties.(*baseLists.PropertyList), mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}},
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
		{"+ve", fields{fromAccAddress.String(), testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties.(*baseLists.PropertyList), immutableProperties.(*baseLists.PropertyList), mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:                    tt.fields.From,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
				TakerID:                 tt.fields.TakerID,
				MakerOwnableID:          tt.fields.MakerOwnableID,
				TakerOwnableID:          tt.fields.TakerOwnableID,
				ExpiresIn:               tt.fields.ExpiresIn,
				MakerOwnableSplit:       tt.fields.MakerOwnableSplit,
				TakerOwnableSplit:       tt.fields.TakerOwnableSplit,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
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
		{"+ve", fields{fromAccAddress.String(), testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties.(*baseLists.PropertyList), immutableProperties.(*baseLists.PropertyList), mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Message{
				From:                    tt.fields.From,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
				TakerID:                 tt.fields.TakerID,
				MakerOwnableID:          tt.fields.MakerOwnableID,
				TakerOwnableID:          tt.fields.TakerOwnableID,
				ExpiresIn:               tt.fields.ExpiresIn,
				MakerOwnableSplit:       tt.fields.MakerOwnableSplit,
				TakerOwnableSplit:       tt.fields.TakerOwnableSplit,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
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
		{"+ve", fields{fromAccAddress.String(), testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties.(*baseLists.PropertyList), immutableProperties.(*baseLists.PropertyList), mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:                    tt.fields.From,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
				TakerID:                 tt.fields.TakerID,
				MakerOwnableID:          tt.fields.MakerOwnableID,
				TakerOwnableID:          tt.fields.TakerOwnableID,
				ExpiresIn:               tt.fields.ExpiresIn,
				MakerOwnableSplit:       tt.fields.MakerOwnableSplit,
				TakerOwnableSplit:       tt.fields.TakerOwnableSplit,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
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
		{"+ve", fields{fromAccAddress.String(), testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties.(*baseLists.PropertyList), immutableProperties.(*baseLists.PropertyList), mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, false},
		{"-ve for nil", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:                    tt.fields.From,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
				TakerID:                 tt.fields.TakerID,
				MakerOwnableID:          tt.fields.MakerOwnableID,
				TakerOwnableID:          tt.fields.TakerOwnableID,
				ExpiresIn:               tt.fields.ExpiresIn,
				MakerOwnableSplit:       tt.fields.MakerOwnableSplit,
				TakerOwnableSplit:       tt.fields.TakerOwnableSplit,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_NewMessage(t *testing.T) {
	type args struct {
		from                    sdkTypes.AccAddress
		fromID                  ids.IdentityID
		classificationID        ids.ClassificationID
		takerID                 ids.IdentityID
		makerOwnableID          ids.OwnableID
		takerOwnableID          ids.OwnableID
		expiresIn               types.Height
		makerOwnableSplit       sdkTypes.Dec
		takerOwnableSplit       sdkTypes.Dec
		immutableMetaProperties lists.PropertyList
		immutableProperties     lists.PropertyList
		mutableMetaProperties   lists.PropertyList
		mutableProperties       lists.PropertyList
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{
		{"+ve", args{fromAccAddress, testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties.(*baseLists.PropertyList), immutableProperties.(*baseLists.PropertyList), mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, &Message{fromAccAddress.String(), testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties.(*baseLists.PropertyList), immutableProperties.(*baseLists.PropertyList), mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessage(tt.args.from, tt.args.fromID, tt.args.classificationID, tt.args.takerID, tt.args.makerOwnableID, tt.args.takerOwnableID, tt.args.expiresIn, tt.args.makerOwnableSplit, tt.args.takerOwnableSplit, tt.args.immutableMetaProperties.(*baseLists.PropertyList), tt.args.immutableProperties, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}