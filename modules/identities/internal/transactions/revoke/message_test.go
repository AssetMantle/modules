// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

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

func CreateTestInputForMessage(t *testing.T) (ids.ID, ids.ID, ids.ID, sdkTypes.AccAddress, sdkTypes.Msg) {
	testFromID := baseIDs.NewStringID("fromID")
	testToID := baseIDs.NewStringID("toID")
	testClassificationID := baseIDs.NewStringID("classificationID")

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	require.Equal(t, nil, err)

	testMessage := newMessage(fromAccAddress, testFromID, testToID, testClassificationID)

	return testFromID, testToID, testClassificationID, fromAccAddress, testMessage
}

func Test_messageFromInterface(t *testing.T) {
	testFromID, testToID, testClassificationID, fromAccAddress, testMessage := CreateTestInputForMessage(t)
	type args struct {
		msg sdkTypes.Msg
	}
	tests := []struct {
		name string
		args args
		want message
	}{
		// TODO: Add test cases.
		{"+ve", args{testMessage}, message{fromAccAddress, testFromID, testToID, testClassificationID}},
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
	testFromID, testToID, testClassificationID, fromAccAddress, testMessage := CreateTestInputForMessage(t)

	type fields struct {
		From             sdkTypes.AccAddress
		FromID           ids.ID
		ToID             ids.ID
		ClassificationID ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testToID, testClassificationID}, sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(testMessage))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:             tt.fields.From,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
			}
			if got := message.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSignBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_GetSigners(t *testing.T) {
	testFromID, testToID, testClassificationID, fromAccAddress, _ := CreateTestInputForMessage(t)

	type fields struct {
		From             sdkTypes.AccAddress
		FromID           ids.ID
		ToID             ids.ID
		ClassificationID ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   []sdkTypes.AccAddress
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testToID, testClassificationID}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:             tt.fields.From,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
			}
			if got := message.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_RegisterCodec(t *testing.T) {
	testFromID, testToID, testClassificationID, fromAccAddress, _ := CreateTestInputForMessage(t)

	type fields struct {
		From             sdkTypes.AccAddress
		FromID           ids.ID
		ToID             ids.ID
		ClassificationID ids.ID
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
		{"+ve", fields{fromAccAddress, testFromID, testToID, testClassificationID}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := message{
				From:             tt.fields.From,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
			}
			me.RegisterCodec(tt.args.codec)
		})
	}
}

func Test_message_Route(t *testing.T) {
	testFromID, testToID, testClassificationID, fromAccAddress, _ := CreateTestInputForMessage(t)

	type fields struct {
		From             sdkTypes.AccAddress
		FromID           ids.ID
		ToID             ids.ID
		ClassificationID ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testToID, testClassificationID}, module.Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:             tt.fields.From,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
			}
			if got := message.Route(); got != tt.want {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_Type(t *testing.T) {
	testFromID, testToID, testClassificationID, fromAccAddress, _ := CreateTestInputForMessage(t)

	type fields struct {
		From             sdkTypes.AccAddress
		FromID           ids.ID
		ToID             ids.ID
		ClassificationID ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testToID, testClassificationID}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:             tt.fields.From,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
			}
			if got := message.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_ValidateBasic(t *testing.T) {
	testFromID, testToID, testClassificationID, fromAccAddress, _ := CreateTestInputForMessage(t)

	type fields struct {
		From             sdkTypes.AccAddress
		FromID           ids.ID
		ToID             ids.ID
		ClassificationID ids.ID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testToID, testClassificationID}, false},
		{"-ve", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:             tt.fields.From,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newMessage(t *testing.T) {
	testFromID, testToID, testClassificationID, fromAccAddress, _ := CreateTestInputForMessage(t)

	type args struct {
		from             sdkTypes.AccAddress
		fromID           ids.ID
		toID             ids.ID
		classificationID ids.ID
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{
		// TODO: Add test cases.
		{"+ve", args{fromAccAddress, testFromID, testToID, testClassificationID}, message{fromAccAddress, testFromID, testToID, testClassificationID}},
		{"-ve with nil", args{}, message{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMessage(tt.args.from, tt.args.fromID, tt.args.toID, tt.args.classificationID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
