// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package provision

import (
	"github.com/AssetMantle/modules/modules/identities/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/utilities/transaction"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func createInputForMessage(t *testing.T) (ids.ID, string, sdkTypes.AccAddress, string, sdkTypes.AccAddress, sdkTypes.Msg) {
	testIdentityID := baseIDs.NewID("identityID")

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	toAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	toAccAddress, err := sdkTypes.AccAddressFromBech32(toAddress)
	require.Nil(t, err)

	testMessage := newMessage(fromAccAddress, toAccAddress, testIdentityID)

	return testIdentityID, fromAddress, fromAccAddress, toAddress, toAccAddress, testMessage
}

func Test_messageFromInterface(t *testing.T) {
	testIdentityID, _, fromAccAddress, _, toAccAddress, testMessage := createInputForMessage(t)
	type args struct {
		msg sdkTypes.Msg
	}
	tests := []struct {
		name string
		args args
		want message
	}{
		// TODO: Add test cases.
		{"+ve", args{testMessage}, message{fromAccAddress, toAccAddress, testIdentityID}},
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
	testIdentityID, _, fromAccAddress, _, toAccAddress, testMessage := createInputForMessage(t)
	type fields struct {
		From       sdkTypes.AccAddress
		To         sdkTypes.AccAddress
		IdentityID ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, toAccAddress, testIdentityID}, sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(testMessage))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:       tt.fields.From,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			if got := message.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSignBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_GetSigners(t *testing.T) {
	testIdentityID, _, fromAccAddress, _, toAccAddress, _ := createInputForMessage(t)

	type fields struct {
		From       sdkTypes.AccAddress
		To         sdkTypes.AccAddress
		IdentityID ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   []sdkTypes.AccAddress
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, toAccAddress, testIdentityID}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:       tt.fields.From,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			if got := message.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_RegisterCodec(t *testing.T) {
	testIdentityID, _, fromAccAddress, _, toAccAddress, _ := createInputForMessage(t)

	type fields struct {
		From       sdkTypes.AccAddress
		To         sdkTypes.AccAddress
		IdentityID ids.ID
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
		{"+ve", fields{fromAccAddress, toAccAddress, testIdentityID}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := message{
				From:       tt.fields.From,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			me.RegisterCodec(tt.args.codec)
		})
	}
}

func Test_message_Route(t *testing.T) {
	testIdentityID, _, fromAccAddress, _, toAccAddress, _ := createInputForMessage(t)

	type fields struct {
		From       sdkTypes.AccAddress
		To         sdkTypes.AccAddress
		IdentityID ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, toAccAddress, testIdentityID}, module.Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:       tt.fields.From,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			if got := message.Route(); got != tt.want {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_Type(t *testing.T) {
	testIdentityID, _, fromAccAddress, _, toAccAddress, _ := createInputForMessage(t)

	type fields struct {
		From       sdkTypes.AccAddress
		To         sdkTypes.AccAddress
		IdentityID ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, toAccAddress, testIdentityID}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:       tt.fields.From,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			if got := message.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_ValidateBasic(t *testing.T) {
	testIdentityID, _, fromAccAddress, _, toAccAddress, _ := createInputForMessage(t)

	type fields struct {
		From       sdkTypes.AccAddress
		To         sdkTypes.AccAddress
		IdentityID ids.ID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, toAccAddress, testIdentityID}, false},
		{"+ve", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:       tt.fields.From,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newMessage(t *testing.T) {
	testIdentityID, _, fromAccAddress, _, toAccAddress, _ := createInputForMessage(t)

	type args struct {
		from       sdkTypes.AccAddress
		to         sdkTypes.AccAddress
		identityID ids.ID
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{
		// TODO: Add test cases.
		{"+ve", args{fromAccAddress, toAccAddress, testIdentityID}, message{fromAccAddress, toAccAddress, testIdentityID}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMessage(tt.args.from, tt.args.to, tt.args.identityID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
