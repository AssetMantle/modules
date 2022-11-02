// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mutate

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/identities/internal/module"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/utilities/transaction"
)

func createTestInputForMessage(t *testing.T) (types.AccAddress, ids.IdentityID, ids.IdentityID, lists.PropertyList, lists.PropertyList) {
	testFrom, err := types.AccAddressFromBech32("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	require.Nil(t, err)
	immutables := baseQualified.NewImmutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)
	testIdentityID := baseIDs.NewIdentityID(testClassificationID, immutables)
	metaProperty := baseProperties.NewMetaProperty(baseIDs.NewStringID("id"), baseData.NewStringData("Data"))
	mesaProperty := baseProperties.NewMesaProperty(baseIDs.NewStringID("id1"), baseData.NewStringData("Data1"))
	testMutableMetaProperties := base.NewMetaPropertyList([]properties.MetaProperty{metaProperty}...)
	testMutableProperties := base.NewPropertyList([]properties.Property{mesaProperty}...)

	return testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties
}

func Test_messageFromInterface(t *testing.T) {
	testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties := createTestInputForMessage(t)
	type args struct {
		msg types.Msg
	}
	tests := []struct {
		name string
		args args
		want message
	}{
		// TODO: Add test cases.
		{"+ve with nil", args{}, message{}},
		{"+ve", args{message{testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}}, message{testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}},
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
	testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties := createTestInputForMessage(t)
	type fields struct {
		From                  types.AccAddress
		FromID                ids.IdentityID
		IdentityID            ids.IdentityID
		MutableMetaProperties lists.PropertyList
		MutableProperties     lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
		{"+ve", fields{testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}, types.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(message{testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				IdentityID:            tt.fields.IdentityID,
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
	testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties := createTestInputForMessage(t)
	type fields struct {
		From                  types.AccAddress
		FromID                ids.IdentityID
		IdentityID            ids.IdentityID
		MutableMetaProperties lists.PropertyList
		MutableProperties     lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   []types.AccAddress
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{}, []types.AccAddress{nil}},
		{"+ve", fields{testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}, []types.AccAddress{testFrom}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				IdentityID:            tt.fields.IdentityID,
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
	testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties := createTestInputForMessage(t)
	type fields struct {
		From                  types.AccAddress
		FromID                ids.IdentityID
		IdentityID            ids.IdentityID
		MutableMetaProperties lists.PropertyList
		MutableProperties     lists.PropertyList
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
		{"+ve with nil", fields{}, args{codec.New()}},
		{"+ve", fields{testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				IdentityID:            tt.fields.IdentityID,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
			}
			me.RegisterCodec(tt.args.codec)
		})
	}
}

func Test_message_Route(t *testing.T) {
	testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties := createTestInputForMessage(t)
	type fields struct {
		From                  types.AccAddress
		FromID                ids.IdentityID
		IdentityID            ids.IdentityID
		MutableMetaProperties lists.PropertyList
		MutableProperties     lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{}, module.Name},
		{"+ve", fields{testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}, module.Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				IdentityID:            tt.fields.IdentityID,
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
	testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties := createTestInputForMessage(t)
	type fields struct {
		From                  types.AccAddress
		FromID                ids.IdentityID
		IdentityID            ids.IdentityID
		MutableMetaProperties lists.PropertyList
		MutableProperties     lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{}, Transaction.GetName()},
		{"+ve", fields{testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				IdentityID:            tt.fields.IdentityID,
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
	testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties := createTestInputForMessage(t)
	type fields struct {
		From                  types.AccAddress
		FromID                ids.IdentityID
		IdentityID            ids.IdentityID
		MutableMetaProperties lists.PropertyList
		MutableProperties     lists.PropertyList
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{}, true},
		{"+ve", fields{testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				IdentityID:            tt.fields.IdentityID,
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
	testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties := createTestInputForMessage(t)
	type args struct {
		from                  types.AccAddress
		fromID                ids.IdentityID
		identityID            ids.IdentityID
		mutableMetaProperties lists.PropertyList
		mutableProperties     lists.PropertyList
	}
	tests := []struct {
		name string
		args args
		want types.Msg
	}{
		// TODO: Add test cases.
		{"+ve with nil", args{}, message{}},
		{"+ve", args{testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}, message{testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMessage(tt.args.from, tt.args.fromID, tt.args.identityID, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
