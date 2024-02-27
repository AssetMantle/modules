// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unwrap

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

type fields struct {
	From   string
	FromID *baseIDs.IdentityID
	Coins  sdkTypes.Coins
}

func Test_messageFromInterface(t *testing.T) {
	type args struct {
		msg helpers.Message
	}
	tests := []struct {
		name string
		args args
		want *Message
	}{
		{"+ve", args{NewMessage(fromAccAddress, fromID, coins).(*Message)}, &Message{fromAccAddress.String(), fromID, coins}},
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
		{"+ve", fields{fromAccAddress.String(), fromID, coins}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:   tt.fields.From,
				FromID: tt.fields.FromID,
				Coins:  tt.fields.Coins,
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
		{"+ve", fields{fromAccAddress.String(), fromID, coins}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Message{
				From:   tt.fields.From,
				FromID: tt.fields.FromID,
				Coins:  tt.fields.Coins,
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
		{"+ve", fields{fromAccAddress.String(), fromID, coins}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:   tt.fields.From,
				FromID: tt.fields.FromID,
				Coins:  tt.fields.Coins,
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
		{"+ve", fields{fromAccAddress.String(), fromID, coins}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:   tt.fields.From,
				FromID: tt.fields.FromID,
				Coins:  tt.fields.Coins,
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_NewMessage(t *testing.T) {
	type args struct {
		from   sdkTypes.AccAddress
		fromID ids.IdentityID
		coins  sdkTypes.Coins
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{
		{"+ve", args{fromAccAddress, fromID, coins}, &Message{fromAccAddress.String(), fromID, coins}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessage(tt.args.from, tt.args.fromID, tt.args.coins); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
