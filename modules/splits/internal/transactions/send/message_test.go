// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package send

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/splits/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/utilities/transaction"
)

func Test_messageFromInterface(t *testing.T) {
	type args struct {
		msg types.Msg
	}
	tests := []struct {
		name string
		args args
		want message
	}{
		{"+ve", args{newMessage(fromAccAddress, fromID, fromID, ownableID, testRate)}, message{fromAccAddress, fromID, fromID, ownableID, testRate}},
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
	type fields struct {
		From      types.AccAddress
		FromID    ids.IdentityID
		ToID      ids.IdentityID
		OwnableID ids.OwnableID
		Value     types.Dec
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve", fields{fromAccAddress, fromID, fromID, ownableID, testRate}, types.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(message{fromAccAddress, fromID, fromID, ownableID, testRate}))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:      tt.fields.From,
				FromID:    tt.fields.FromID,
				ToID:      tt.fields.ToID,
				OwnableID: tt.fields.OwnableID,
				Value:     tt.fields.Value,
			}
			if got := message.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSignBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_GetSigners(t *testing.T) {
	type fields struct {
		From      types.AccAddress
		FromID    ids.IdentityID
		ToID      ids.IdentityID
		OwnableID ids.OwnableID
		Value     types.Dec
	}
	tests := []struct {
		name   string
		fields fields
		want   []types.AccAddress
	}{
		{"+ve", fields{fromAccAddress, fromID, fromID, ownableID, testRate}, []types.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:      tt.fields.From,
				FromID:    tt.fields.FromID,
				ToID:      tt.fields.ToID,
				OwnableID: tt.fields.OwnableID,
				Value:     tt.fields.Value,
			}
			if got := message.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_RegisterCodec(t *testing.T) {
	type fields struct {
		From      types.AccAddress
		FromID    ids.IdentityID
		ToID      ids.IdentityID
		OwnableID ids.OwnableID
		Value     types.Dec
	}
	type args struct {
		codec *codec.Codec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{fromAccAddress, fromID, fromID, ownableID, testRate}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := message{
				From:      tt.fields.From,
				FromID:    tt.fields.FromID,
				ToID:      tt.fields.ToID,
				OwnableID: tt.fields.OwnableID,
				Value:     tt.fields.Value,
			}
			me.RegisterCodec(tt.args.codec)
		})
	}
}

func Test_message_Route(t *testing.T) {
	type fields struct {
		From      types.AccAddress
		FromID    ids.IdentityID
		ToID      ids.IdentityID
		OwnableID ids.OwnableID
		Value     types.Dec
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{fromAccAddress, fromID, fromID, ownableID, testRate}, module.Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:      tt.fields.From,
				FromID:    tt.fields.FromID,
				ToID:      tt.fields.ToID,
				OwnableID: tt.fields.OwnableID,
				Value:     tt.fields.Value,
			}
			if got := message.Route(); got != tt.want {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_Type(t *testing.T) {
	type fields struct {
		From      types.AccAddress
		FromID    ids.IdentityID
		ToID      ids.IdentityID
		OwnableID ids.OwnableID
		Value     types.Dec
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{fromAccAddress, fromID, fromID, ownableID, testRate}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:      tt.fields.From,
				FromID:    tt.fields.FromID,
				ToID:      tt.fields.ToID,
				OwnableID: tt.fields.OwnableID,
				Value:     tt.fields.Value,
			}
			if got := message.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_ValidateBasic(t *testing.T) {
	type fields struct {
		From      types.AccAddress
		FromID    ids.IdentityID
		ToID      ids.IdentityID
		OwnableID ids.OwnableID
		Value     types.Dec
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{fromAccAddress, fromID, fromID, ownableID, testRate}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:      tt.fields.From,
				FromID:    tt.fields.FromID,
				ToID:      tt.fields.ToID,
				OwnableID: tt.fields.OwnableID,
				Value:     tt.fields.Value,
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newMessage(t *testing.T) {
	type args struct {
		from      types.AccAddress
		fromID    ids.IdentityID
		toID      ids.IdentityID
		ownableID ids.OwnableID
		value     types.Dec
	}
	tests := []struct {
		name string
		args args
		want types.Msg
	}{
		{"+ve", args{fromAccAddress, fromID, fromID, ownableID, testRate}, message{fromAccAddress, fromID, fromID, ownableID, testRate}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMessage(tt.args.from, tt.args.fromID, tt.args.toID, tt.args.ownableID, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
