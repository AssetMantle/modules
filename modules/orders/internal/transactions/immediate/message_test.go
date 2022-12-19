// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package immediate

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
	"github.com/AssetMantle/modules/utilities/transaction"
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
		{"+ve", fields{fromAccAddress, testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, sdkTypes.MustSortJSON(transaction.RegisterLegacyAminoCodec(messagePrototype).MustMarshalJSON(testMessage))},
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
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{fromAccAddress, testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, args{codec.NewLegacyAmino()}},
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
			me.RegisterLegacyAminoCodec(tt.args.legacyAmino)
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
		{"+ve", fields{fromAccAddress, testFromID, testClassificationID, testFromID, makerOwnableID, takerOwnableID, expiresInHeight, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, false},
		{"-ve", fields{}, true},
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