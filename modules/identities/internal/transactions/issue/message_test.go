// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package issue

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/identities/internal/module"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/utilities/transaction"
)

func createTestInput(t *testing.T) (ids.IdentityID, ids.ClassificationID, string, sdkTypes.AccAddress, string, sdkTypes.AccAddress, lists.MetaPropertyList, lists.PropertyList, lists.MetaPropertyList, lists.PropertyList) {
	immutables := baseQualified.NewImmutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)

	const fromAddress = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	const toAddress = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	var toAccAddress sdkTypes.AccAddress
	toAccAddress, err = sdkTypes.AccAddressFromBech32(toAddress)
	require.Nil(t, err)

	var immutableMetaProperties lists.MetaPropertyList
	immutableMetaProperties, err = utilities.ReadMetaPropertyList("defaultImmutableMeta1:S|defaultImmutableMeta1")
	require.Equal(t, nil, err)

	var immutableProperties lists.PropertyList
	immutableProperties, err = utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)

	var mutableMetaProperties lists.MetaPropertyList
	mutableMetaProperties, err = utilities.ReadMetaPropertyList("defaultMutableMeta1:S|defaultMutableMeta1")
	require.Equal(t, nil, err)

	var mutableProperties lists.PropertyList
	mutableProperties, err = utilities.ReadProperties("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, err)

	return testFromID, testClassificationID, fromAddress, fromAccAddress, toAddress, toAccAddress, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties
}

func Test_messageFromInterface(t *testing.T) {
	testFromID, testClassificationID, _, fromAccAddress, _, toAccAddress, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties := createTestInput(t)
	type args struct {
		msg sdkTypes.Msg
	}
	tests := []struct {
		name string
		args args
		want message
	}{

		{"+ve", args{newMessage(fromAccAddress, toAccAddress, testFromID, testClassificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)}, message{fromAccAddress, toAccAddress, testFromID, testClassificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}},
		{"+ve with nil", args{}, message{}},
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
	testFromID, testClassificationID, _, fromAccAddress, _, toAccAddress, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties := createTestInput(t)
	type fields struct {
		From                    sdkTypes.AccAddress
		To                      sdkTypes.AccAddress
		FromID                  ids.IdentityID
		ClassificationID        ids.ClassificationID
		ImmutableMetaProperties lists.MetaPropertyList
		ImmutableProperties     lists.PropertyList
		MutableMetaProperties   lists.MetaPropertyList
		MutableProperties       lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{

		{"+ve", fields{fromAccAddress, toAccAddress, testFromID, testClassificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(message{fromAccAddress, toAccAddress, testFromID, testClassificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                    tt.fields.From,
				To:                      tt.fields.To,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
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
	testFromID, testClassificationID, _, fromAccAddress, _, toAccAddress, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties := createTestInput(t)
	type fields struct {
		From                    sdkTypes.AccAddress
		To                      sdkTypes.AccAddress
		FromID                  ids.IdentityID
		ClassificationID        ids.ClassificationID
		ImmutableMetaProperties lists.MetaPropertyList
		ImmutableProperties     lists.PropertyList
		MutableMetaProperties   lists.MetaPropertyList
		MutableProperties       lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   []sdkTypes.AccAddress
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, toAccAddress, testFromID, testClassificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                    tt.fields.From,
				To:                      tt.fields.To,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
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
	testFromID, testClassificationID, _, fromAccAddress, _, toAccAddress, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties := createTestInput(t)
	type fields struct {
		From                    sdkTypes.AccAddress
		To                      sdkTypes.AccAddress
		FromID                  ids.IdentityID
		ClassificationID        ids.ClassificationID
		ImmutableMetaProperties lists.MetaPropertyList
		ImmutableProperties     lists.PropertyList
		MutableMetaProperties   lists.MetaPropertyList
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
		{"+ve", fields{fromAccAddress, toAccAddress, testFromID, testClassificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := message{
				From:                    tt.fields.From,
				To:                      tt.fields.To,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
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
	testFromID, testClassificationID, _, fromAccAddress, _, toAccAddress, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties := createTestInput(t)
	type fields struct {
		From                    sdkTypes.AccAddress
		To                      sdkTypes.AccAddress
		FromID                  ids.IdentityID
		ClassificationID        ids.ClassificationID
		ImmutableMetaProperties lists.MetaPropertyList
		ImmutableProperties     lists.PropertyList
		MutableMetaProperties   lists.MetaPropertyList
		MutableProperties       lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, toAccAddress, testFromID, testClassificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, module.Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                    tt.fields.From,
				To:                      tt.fields.To,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
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
	testFromID, testClassificationID, _, fromAccAddress, _, toAccAddress, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties := createTestInput(t)
	type fields struct {
		From                    sdkTypes.AccAddress
		To                      sdkTypes.AccAddress
		FromID                  ids.IdentityID
		ClassificationID        ids.ClassificationID
		ImmutableMetaProperties lists.MetaPropertyList
		ImmutableProperties     lists.PropertyList
		MutableMetaProperties   lists.MetaPropertyList
		MutableProperties       lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, toAccAddress, testFromID, testClassificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                    tt.fields.From,
				To:                      tt.fields.To,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
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
	testFromID, testClassificationID, _, fromAccAddress, _, toAccAddress, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties := createTestInput(t)
	type fields struct {
		From                    sdkTypes.AccAddress
		To                      sdkTypes.AccAddress
		FromID                  ids.IdentityID
		ClassificationID        ids.ClassificationID
		ImmutableMetaProperties lists.MetaPropertyList
		ImmutableProperties     lists.PropertyList
		MutableMetaProperties   lists.MetaPropertyList
		MutableProperties       lists.PropertyList
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, toAccAddress, testFromID, testClassificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, false},
		{"-ve with nil", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                    tt.fields.From,
				To:                      tt.fields.To,
				FromID:                  tt.fields.FromID,
				ClassificationID:        tt.fields.ClassificationID,
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
	testFromID, testClassificationID, _, fromAccAddress, _, toAccAddress, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties := createTestInput(t)
	type args struct {
		from                    sdkTypes.AccAddress
		to                      sdkTypes.AccAddress
		fromID                  ids.IdentityID
		classificationID        ids.ClassificationID
		immutableMetaProperties lists.MetaPropertyList
		immutableProperties     lists.PropertyList
		mutableMetaProperties   lists.MetaPropertyList
		mutableProperties       lists.PropertyList
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{
		// TODO: Add test cases.
		{"+ve", args{fromAccAddress, toAccAddress, testFromID, testClassificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties}, newMessage(fromAccAddress, toAccAddress, testFromID, testClassificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)},
		{"-ve with nil", args{}, message{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMessage(tt.args.from, tt.args.to, tt.args.fromID, tt.args.classificationID, tt.args.immutableMetaProperties, tt.args.immutableProperties, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
