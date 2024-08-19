// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package put

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/types"
	"github.com/AssetMantle/schema/types/base"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

var (
	expiresInHeight = base.NewHeight(60).(*base.Height)
	testMessage     = NewMessage(fromAccAddress, testFromID, makerAssetID, takerAssetID, makerSplit, takerSplit, expiresInHeight)
)

type fields struct {
	From         string
	FromID       *baseIDs.IdentityID
	MakerAssetID *baseIDs.AssetID
	TakerAssetID *baseIDs.AssetID
	MakerSplit   sdkTypes.Int
	TakerSplit   sdkTypes.Int
	ExpiryHeight *base.Height
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
		{"+ve", args{testMessage}, &Message{fromAccAddress.String(), testFromID, makerAssetID, takerAssetID, makerSplit.String(), takerSplit.String(), expiresInHeight}},
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
		{"+ve", fields{fromAccAddress.String(), testFromID, makerAssetID, takerAssetID, makerSplit, takerSplit, expiresInHeight}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:   tt.fields.From,
				FromID: tt.fields.FromID,
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
		{"+ve", fields{fromAccAddress.String(), testFromID, makerAssetID, takerAssetID, makerSplit, takerSplit, expiresInHeight}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Message{
				tt.fields.From,
				tt.fields.FromID,
				tt.fields.MakerAssetID,
				tt.fields.TakerAssetID,
				tt.fields.MakerSplit.String(),
				tt.fields.TakerSplit.String(),
				tt.fields.ExpiryHeight,
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
		{"+ve", fields{fromAccAddress.String(), testFromID, makerAssetID, takerAssetID, makerSplit, takerSplit, expiresInHeight}, Transaction.GetServicePath()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				tt.fields.From,
				tt.fields.FromID,
				tt.fields.MakerAssetID,
				tt.fields.TakerAssetID,
				tt.fields.MakerSplit.String(),
				tt.fields.TakerSplit.String(),
				tt.fields.ExpiryHeight,
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
		{"+ve", fields{fromAccAddress.String(), testFromID, makerAssetID, takerAssetID, makerSplit, takerSplit, expiresInHeight}, false},
		{"-ve for nil", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				tt.fields.From,
				tt.fields.FromID,
				tt.fields.MakerAssetID,
				tt.fields.TakerAssetID,
				tt.fields.MakerSplit.String(),
				tt.fields.TakerSplit.String(),
				tt.fields.ExpiryHeight,
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_NewMessage(t *testing.T) {
	type args struct {
		from         sdkTypes.AccAddress
		fromID       ids.IdentityID
		makerAssetID ids.AssetID
		takerAssetID ids.AssetID
		makerSplit   sdkTypes.Int
		takerSplit   sdkTypes.Int
		expiryHeight types.Height
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{
		{"+ve", args{fromAccAddress, testFromID, makerAssetID, takerAssetID, makerSplit, takerSplit, expiresInHeight}, &Message{fromAccAddress.String(), testFromID, makerAssetID, takerAssetID, makerSplit.String(), takerSplit.String(), expiresInHeight}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessage(tt.args.from, tt.args.fromID, tt.args.makerAssetID, tt.args.takerAssetID, tt.args.makerSplit, tt.args.takerSplit, tt.args.expiryHeight); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
