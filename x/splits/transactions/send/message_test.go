// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package send

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

type fields struct {
	From    string
	FromID  *baseIDs.IdentityID
	ToID    *baseIDs.IdentityID
	AssetID *baseIDs.AssetID
	Value   types.Int
}

func Test_messageFromInterface(t *testing.T) {
	type args struct {
		msg types.Msg
	}
	tests := []struct {
		name string
		args args
		want *Message
	}{
		{"+ve", args{NewMessage(fromAccAddress, fromID, fromID, assetID, testRate)}, &Message{fromAccAddress.String(), fromID, fromID, assetID, testRate.String()}},
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
		want   []types.AccAddress
	}{
		{"+ve", fields{fromAccAddress.String(), fromID, fromID, assetID, testRate}, []types.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				ToID:    tt.fields.ToID,
				AssetID: tt.fields.AssetID,
				Value:   tt.fields.Value.String(),
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
		{"+ve", fields{fromAccAddress.String(), fromID, fromID, assetID, testRate}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				ToID:    tt.fields.ToID,
				AssetID: tt.fields.AssetID,
				Value:   tt.fields.Value.String(),
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
		{"+ve", fields{fromAccAddress.String(), fromID, fromID, assetID, testRate}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				ToID:    tt.fields.ToID,
				AssetID: tt.fields.AssetID,
				Value:   tt.fields.Value.String(),
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
		{"+ve", fields{fromAccAddress.String(), fromID, fromID, assetID, testRate}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				ToID:    tt.fields.ToID,
				AssetID: tt.fields.AssetID,
				Value:   tt.fields.Value.String(),
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_NewMessage(t *testing.T) {
	type args struct {
		from    types.AccAddress
		fromID  ids.IdentityID
		toID    ids.IdentityID
		assetID ids.AssetID
		value   types.Int
	}
	tests := []struct {
		name string
		args args
		want types.Msg
	}{
		{"+ve", args{fromAccAddress, fromID, fromID, assetID, testRate}, &Message{fromAccAddress.String(), fromID, fromID, assetID, testRate.String()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessage(tt.args.from, tt.args.fromID, tt.args.toID, tt.args.assetID, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
