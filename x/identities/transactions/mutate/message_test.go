// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mutate

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/x/identities/module"

	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/lists"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	"github.com/AssetMantle/schema/go/properties"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/utilities/transaction"
)

type fields struct {
	From                  string
	FromID                *baseIDs.IdentityID
	IdentityID            *baseIDs.IdentityID
	MutableMetaProperties *baseLists.PropertyList
	MutableProperties     *baseLists.PropertyList
}

func createTestInputForMessage(t *testing.T) (types.AccAddress, *baseIDs.IdentityID, *baseIDs.IdentityID, *baseLists.PropertyList, *baseLists.PropertyList) {
	testFrom, err := types.AccAddressFromBech32("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	require.Nil(t, err)
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)
	testIdentityID := baseIDs.NewIdentityID(testClassificationID, immutables)
	metaProperty := baseProperties.NewMetaProperty(baseIDs.NewStringID("id"), baseData.NewStringData("Data"))
	mesaProperty := baseProperties.NewMesaProperty(baseIDs.NewStringID("id1"), baseData.NewStringData("Data1"))
	testMutableMetaProperties := baseLists.NewPropertyList([]properties.Property{metaProperty}...)
	testMutableProperties := baseLists.NewPropertyList([]properties.Property{mesaProperty}...)

	return testFrom, testFromID.(*baseIDs.IdentityID), testIdentityID.(*baseIDs.IdentityID), testMutableMetaProperties.(*baseLists.PropertyList), testMutableProperties.(*baseLists.PropertyList)
}

func Test_messageFromInterface(t *testing.T) {
	testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties := createTestInputForMessage(t)
	type args struct {
		msg types.Msg
	}
	tests := []struct {
		name string
		args args
		want *Message
	}{
		{"+ve with nil", args{}, &Message{}},
		{"+ve", args{&Message{testFrom.String(), testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}}, &Message{testFrom.String(), testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}},
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
	testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties := createTestInputForMessage(t)

	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve", fields{testFrom.String(), testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}, types.MustSortJSON(transaction.RegisterLegacyAminoCodec(messagePrototype).MustMarshalJSON(&Message{testFrom.String(), testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
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

	tests := []struct {
		name   string
		fields fields
		want   []types.AccAddress
	}{
		{"+ve with nil", fields{}, []types.AccAddress{nil}},
		{"+ve", fields{testFrom.String(), testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}, []types.AccAddress{testFrom}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
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

	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve with nil", fields{}, args{codec.NewLegacyAmino()}},
		{"+ve", fields{testFrom.String(), testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				IdentityID:            tt.fields.IdentityID,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
			}
			me.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_message_Route(t *testing.T) {
	testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties := createTestInputForMessage(t)

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve with nil", fields{}, module.Name},
		{"+ve", fields{testFrom.String(), testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}, module.Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
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

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve with nil", fields{}, Transaction.GetName()},
		{"+ve", fields{testFrom.String(), testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
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

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve with nil", fields{}, true},
		{"+ve", fields{testFrom.String(), testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
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
		{"+ve with nil", args{}, &Message{}},
		{"+ve", args{testFrom, testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}, &Message{testFrom.String(), testFromID, testIdentityID, testMutableMetaProperties, testMutableProperties}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMessage(tt.args.from, tt.args.fromID, tt.args.identityID, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
