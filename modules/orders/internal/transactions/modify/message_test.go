// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package modify

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
	"github.com/AssetMantle/modules/utilities/transaction"
)

var (
	expiresInHeight = base.NewHeight(60).(*base.Height)
	testMessage     = newMessage(fromAccAddress, testFromID, testOrderID, takerOwnableSplit, makerOwnableSplit, expiresInHeight, mutableMetaProperties.(*baseLists.PropertyList), mutableProperties)
)

type fields struct {
	From                  string
	FromID                *baseIDs.IdentityID
	OrderID               *baseIDs.OrderID
	MakerOwnableSplit     sdkTypes.Dec
	TakerOwnableSplit     sdkTypes.Dec
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
		{"+ve", args{testMessage}, &Message{fromAccAddress.String(), testFromID, testOrderID, takerOwnableSplit, makerOwnableSplit, expiresInHeight, mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}},
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
		{"+ve", fields{fromAccAddress.String(), testFromID, testOrderID, takerOwnableSplit, makerOwnableSplit, expiresInHeight, mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, sdkTypes.MustSortJSON(transaction.RegisterLegacyAminoCodec(messagePrototype).MustMarshalJSON(testMessage))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				OrderID:               tt.fields.OrderID,
				MakerOwnableSplit:     tt.fields.MakerOwnableSplit,
				TakerOwnableSplit:     tt.fields.TakerOwnableSplit,
				ExpiresIn:             tt.fields.ExpiresIn,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
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
		{"+ve", fields{fromAccAddress.String(), testFromID, testOrderID, takerOwnableSplit, makerOwnableSplit, expiresInHeight, mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				OrderID:               tt.fields.OrderID,
				MakerOwnableSplit:     tt.fields.MakerOwnableSplit,
				TakerOwnableSplit:     tt.fields.TakerOwnableSplit,
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
		{"+ve", fields{fromAccAddress.String(), testFromID, testOrderID, takerOwnableSplit, makerOwnableSplit, expiresInHeight, mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				OrderID:               tt.fields.OrderID,
				MakerOwnableSplit:     tt.fields.MakerOwnableSplit,
				TakerOwnableSplit:     tt.fields.TakerOwnableSplit,
				ExpiresIn:             tt.fields.ExpiresIn,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
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
		{"+ve", fields{fromAccAddress.String(), testFromID, testOrderID, takerOwnableSplit, makerOwnableSplit, expiresInHeight, mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, module.Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				OrderID:               tt.fields.OrderID,
				MakerOwnableSplit:     tt.fields.MakerOwnableSplit,
				TakerOwnableSplit:     tt.fields.TakerOwnableSplit,
				ExpiresIn:             tt.fields.ExpiresIn,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
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
		{"+ve", fields{fromAccAddress.String(), testFromID, testOrderID, takerOwnableSplit, makerOwnableSplit, expiresInHeight, mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				OrderID:               tt.fields.OrderID,
				MakerOwnableSplit:     tt.fields.MakerOwnableSplit,
				TakerOwnableSplit:     tt.fields.TakerOwnableSplit,
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
		{"+ve", fields{fromAccAddress.String(), testFromID, testOrderID, takerOwnableSplit, makerOwnableSplit, expiresInHeight, mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				OrderID:               tt.fields.OrderID,
				MakerOwnableSplit:     tt.fields.MakerOwnableSplit,
				TakerOwnableSplit:     tt.fields.TakerOwnableSplit,
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

func Test_newMessage(t *testing.T) {
	type args struct {
		from                  sdkTypes.AccAddress
		fromID                ids.IdentityID
		orderID               ids.OrderID
		takerOwnableSplit     sdkTypes.Dec
		makerOwnableSplit     sdkTypes.Dec
		expiresIn             types.Height
		mutableMetaProperties lists.PropertyList
		mutableProperties     lists.PropertyList
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{
		{"+ve", args{fromAccAddress, testFromID, testOrderID, takerOwnableSplit, makerOwnableSplit, expiresInHeight, mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}, &Message{fromAccAddress.String(), testFromID, testOrderID, takerOwnableSplit, makerOwnableSplit, expiresInHeight, mutableMetaProperties.(*baseLists.PropertyList), mutableProperties.(*baseLists.PropertyList)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMessage(tt.args.from, tt.args.fromID, tt.args.orderID, tt.args.takerOwnableSplit, tt.args.makerOwnableSplit, tt.args.expiresIn, tt.args.mutableMetaProperties.(*baseLists.PropertyList), tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
