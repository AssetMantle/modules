// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

import (
	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
	"github.com/AssetMantle/modules/utilities/transaction"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"reflect"
	"testing"
)

var (
	expiresInHeight = base.NewHeight(60)
	testMessage     = newMessage(fromAccAddress, testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)
)

func Test_messageFromInterface(t *testing.T) {
	type args struct {
		msg sdkTypes.Msg
	}
	tests := []struct {
		name string
		args args
		want message
	}{
		// TODO: Add test cases.
		{"+ve", args{testMessage}, message{fromAccAddress, testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}},
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
		// TODO: Add test cases.
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
		From                    sdkTypes.AccAddress
		FromID                  ids.IdentityID
		ClassificationID        ids.ClassificationID
		TakerID                 ids.IdentityID
		MakerOwnableID          ids.OwnableID
		TakerOwnableID          ids.OwnableID
		ExpiresIn               types.Height
		MakerOwnableSplit       sdkTypes.Dec
		TakerOwnableSplit       sdkTypes.Dec
		ImmutableMetaProperties lists.PropertyList
		ImmutableProperties     lists.PropertyList
		MutableMetaProperties   lists.PropertyList
		MutableProperties       lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(testMessage))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
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
			if got := message.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSignBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_GetSigners(t *testing.T) {
	type fields struct {
		From                    sdkTypes.AccAddress
		FromID                  ids.IdentityID
		ClassificationID        ids.ClassificationID
		TakerID                 ids.IdentityID
		MakerOwnableID          ids.OwnableID
		TakerOwnableID          ids.OwnableID
		ExpiresIn               types.Height
		MakerOwnableSplit       sdkTypes.Dec
		TakerOwnableSplit       sdkTypes.Dec
		ImmutableMetaProperties lists.PropertyList
		ImmutableProperties     lists.PropertyList
		MutableMetaProperties   lists.PropertyList
		MutableProperties       lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   []sdkTypes.AccAddress
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
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
	type fields struct {
		From                    sdkTypes.AccAddress
		FromID                  ids.IdentityID
		ClassificationID        ids.ClassificationID
		TakerID                 ids.IdentityID
		MakerOwnableID          ids.OwnableID
		TakerOwnableID          ids.OwnableID
		ExpiresIn               types.Height
		MakerOwnableSplit       sdkTypes.Dec
		TakerOwnableSplit       sdkTypes.Dec
		ImmutableMetaProperties lists.PropertyList
		ImmutableProperties     lists.PropertyList
		MutableMetaProperties   lists.PropertyList
		MutableProperties       lists.PropertyList
	}
	type args struct {
		codec *codec.Codec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := message{
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
			me.RegisterCodec(tt.args.codec)
		})
	}
}

func Test_message_Route(t *testing.T) {
	type fields struct {
		From                    sdkTypes.AccAddress
		FromID                  ids.IdentityID
		ClassificationID        ids.ClassificationID
		TakerID                 ids.IdentityID
		MakerOwnableID          ids.OwnableID
		TakerOwnableID          ids.OwnableID
		ExpiresIn               types.Height
		MakerOwnableSplit       sdkTypes.Dec
		TakerOwnableSplit       sdkTypes.Dec
		ImmutableMetaProperties lists.PropertyList
		ImmutableProperties     lists.PropertyList
		MutableMetaProperties   lists.PropertyList
		MutableProperties       lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, module.Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
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
			if got := message.Route(); got != tt.want {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_Type(t *testing.T) {
	type fields struct {
		From                    sdkTypes.AccAddress
		FromID                  ids.IdentityID
		ClassificationID        ids.ClassificationID
		TakerID                 ids.IdentityID
		MakerOwnableID          ids.OwnableID
		TakerOwnableID          ids.OwnableID
		ExpiresIn               types.Height
		MakerOwnableSplit       sdkTypes.Dec
		TakerOwnableSplit       sdkTypes.Dec
		ImmutableMetaProperties lists.PropertyList
		ImmutableProperties     lists.PropertyList
		MutableMetaProperties   lists.PropertyList
		MutableProperties       lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
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
	type fields struct {
		From                    sdkTypes.AccAddress
		FromID                  ids.IdentityID
		ClassificationID        ids.ClassificationID
		TakerID                 ids.IdentityID
		MakerOwnableID          ids.OwnableID
		TakerOwnableID          ids.OwnableID
		ExpiresIn               types.Height
		MakerOwnableSplit       sdkTypes.Dec
		TakerOwnableSplit       sdkTypes.Dec
		ImmutableMetaProperties lists.PropertyList
		ImmutableProperties     lists.PropertyList
		MutableMetaProperties   lists.PropertyList
		MutableProperties       lists.PropertyList
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, false},
		{"-ve for nil", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
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

func Test_newMessage(t *testing.T) {
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
		// TODO: Add test cases.
		{"+ve", args{fromAccAddress, testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, message{fromAccAddress, testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMessage(tt.args.from, tt.args.fromID, tt.args.classificationID, tt.args.takerID, tt.args.makerOwnableID, tt.args.takerOwnableID, tt.args.expiresIn, tt.args.makerOwnableSplit, tt.args.takerOwnableSplit, tt.args.immutableMetaProperties, tt.args.immutableProperties, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
