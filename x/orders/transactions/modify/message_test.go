// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package modify

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/types"
	"github.com/AssetMantle/schema/types/base"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

var (
	expiresInHeight = base.NewHeight(60).(*base.Height)
	testMessage     = NewMessage(fromAccAddress, testFromID, testOrderID, takerSplit, makerSplit, expiresInHeight, mutableMetaProperties.(*baseLists.PropertyList), mutableProperties)
)

type fields struct {
	From                  string
	FromID                *baseIDs.IdentityID
	OrderID               *baseIDs.OrderID
	MakerSplit            sdkTypes.Int
	TakerSplit            sdkTypes.Int
	ExpiresIn             *base.Height
	MutableMetaProperties *baseLists.PropertyList
	MutableProperties     *baseLists.PropertyList
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
		{"+ve", args{testMessage}, &Message{fromAccAddress.String(), testFromID, testOrderID, takerSplit.String(), makerSplit.String(), expiresInHeight, mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}},
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
		{"+ve", fields{fromAccAddress.String(), testFromID, testOrderID, takerSplit, makerSplit, expiresInHeight, mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				OrderID:               tt.fields.OrderID,
				MakerSplit:            tt.fields.MakerSplit.String(),
				TakerSplit:            tt.fields.TakerSplit.String(),
				ExpiresIn:             tt.fields.ExpiresIn,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
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
		{"+ve", fields{fromAccAddress.String(), testFromID, testOrderID, takerSplit, makerSplit, expiresInHeight, mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				OrderID:               tt.fields.OrderID,
				MakerSplit:            tt.fields.MakerSplit.String(),
				TakerSplit:            tt.fields.TakerSplit.String(),
				ExpiresIn:             tt.fields.ExpiresIn,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
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
		{"+ve", fields{fromAccAddress.String(), testFromID, testOrderID, takerSplit, makerSplit, expiresInHeight, mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, Transaction.GetServicePath()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				OrderID:               tt.fields.OrderID,
				MakerSplit:            tt.fields.MakerSplit.String(),
				TakerSplit:            tt.fields.TakerSplit.String(),
				ExpiresIn:             tt.fields.ExpiresIn,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
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
		{"-ve for nil", fields{}, true},
		{"+ve", fields{fromAccAddress.String(), testFromID, testOrderID, takerSplit, makerSplit, expiresInHeight, mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				OrderID:               tt.fields.OrderID,
				MakerSplit:            tt.fields.MakerSplit.String(),
				TakerSplit:            tt.fields.TakerSplit.String(),
				ExpiresIn:             tt.fields.ExpiresIn,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_NewMessage(t *testing.T) {
	type args struct {
		from                  sdkTypes.AccAddress
		fromID                ids.IdentityID
		orderID               ids.OrderID
		takerSplit            sdkTypes.Int
		makerSplit            sdkTypes.Int
		expiresIn             types.Height
		mutableMetaProperties lists.PropertyList
		mutableProperties     lists.PropertyList
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{
		{"+ve", args{fromAccAddress, testFromID, testOrderID, takerSplit, makerSplit, expiresInHeight, mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, &Message{fromAccAddress.String(), testFromID, testOrderID, takerSplit.String(), makerSplit.String(), expiresInHeight, mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessage(tt.args.from, tt.args.fromID, tt.args.orderID, tt.args.takerSplit, tt.args.makerSplit, tt.args.expiresIn, tt.args.mutableMetaProperties.(*baseLists.PropertyList), tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
