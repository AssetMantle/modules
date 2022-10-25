// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package quash

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
	"github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/utilities/transaction"
)

func createInputForMessage(t *testing.T) (ids.IdentityID, ids.IdentityID, sdkTypes.AccAddress, sdkTypes.Msg) {
	immutables := baseQualified.NewImmutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)
	testIdentityID := baseIDs.NewIdentityID(testClassificationID, immutables)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	testMessage := newMessage(fromAccAddress, testFromID, testIdentityID)

	return testIdentityID, testFromID, fromAccAddress, testMessage
}

func Test_messageFromInterface(t *testing.T) {
	testIdentityID, testFromID, fromAccAddress, testMessage := createInputForMessage(t)
	type args struct {
		msg sdkTypes.Msg
	}
	tests := []struct {
		name string
		args args
		want message
	}{
		// TODO: Add test cases.
		{"+ve", args{testMessage}, message{fromAccAddress, testFromID, testIdentityID}},
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
	testIdentityID, testFromID, fromAccAddress, testMessage := createInputForMessage(t)
	type fields struct {
		From       sdkTypes.AccAddress
		FromID     ids.IdentityID
		IdentityID ids.IdentityID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testIdentityID}, sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(testMessage))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:       tt.fields.From,
				FromID:     tt.fields.FromID,
				IdentityID: tt.fields.IdentityID,
			}
			if got := message.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSignBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_GetSigners(t *testing.T) {
	testIdentityID, testFromID, fromAccAddress, _ := createInputForMessage(t)

	type fields struct {
		From       sdkTypes.AccAddress
		FromID     ids.IdentityID
		IdentityID ids.IdentityID
	}
	tests := []struct {
		name   string
		fields fields
		want   []sdkTypes.AccAddress
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testIdentityID}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:       tt.fields.From,
				FromID:     tt.fields.FromID,
				IdentityID: tt.fields.IdentityID,
			}
			if got := message.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_RegisterCodec(t *testing.T) {
	testIdentityID, testFromID, fromAccAddress, _ := createInputForMessage(t)
	type fields struct {
		From       sdkTypes.AccAddress
		FromID     ids.IdentityID
		IdentityID ids.IdentityID
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
		{"+ve", fields{fromAccAddress, testFromID, testIdentityID}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := message{
				From:       tt.fields.From,
				FromID:     tt.fields.FromID,
				IdentityID: tt.fields.IdentityID,
			}
			me.RegisterCodec(tt.args.codec)
		})
	}
}

func Test_message_Route(t *testing.T) {
	testIdentityID, testFromID, fromAccAddress, _ := createInputForMessage(t)
	type fields struct {
		From       sdkTypes.AccAddress
		FromID     ids.IdentityID
		IdentityID ids.IdentityID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testIdentityID}, module.Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:       tt.fields.From,
				FromID:     tt.fields.FromID,
				IdentityID: tt.fields.IdentityID,
			}
			if got := message.Route(); got != tt.want {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_Type(t *testing.T) {
	testIdentityID, testFromID, fromAccAddress, _ := createInputForMessage(t)
	type fields struct {
		From       sdkTypes.AccAddress
		FromID     ids.IdentityID
		IdentityID ids.IdentityID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testIdentityID}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:       tt.fields.From,
				FromID:     tt.fields.FromID,
				IdentityID: tt.fields.IdentityID,
			}
			if got := message.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_ValidateBasic(t *testing.T) {
	testIdentityID, testFromID, fromAccAddress, _ := createInputForMessage(t)
	type fields struct {
		From       sdkTypes.AccAddress
		FromID     ids.IdentityID
		IdentityID ids.IdentityID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testIdentityID}, false},
		{"-ve", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:       tt.fields.From,
				FromID:     tt.fields.FromID,
				IdentityID: tt.fields.IdentityID,
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newMessage(t *testing.T) {
	testIdentityID, testFromID, fromAccAddress, _ := createInputForMessage(t)
	type args struct {
		from       sdkTypes.AccAddress
		fromID     ids.IdentityID
		identityID ids.IdentityID
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{
		// TODO: Add test cases.
		{"+ve", args{fromAccAddress, testFromID, testIdentityID}, message{fromAccAddress, testFromID, testIdentityID}},
		{"-ve", args{}, message{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMessage(tt.args.from, tt.args.fromID, tt.args.identityID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
