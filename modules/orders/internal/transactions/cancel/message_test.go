// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package cancel

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/orders/internal/module"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	"github.com/AssetMantle/modules/utilities/transaction"
)

func CreateTestInputForMessages(t *testing.T) (ids.OrderID, ids.IdentityID, sdkTypes.AccAddress, sdkTypes.Msg) {
	// testOrderID := baseIDs.NewStringID("orderID")
	immutables := baseQualified.NewImmutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	// mutables2 := baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID"), baseData.NewStringData(""))))

	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)
	testMakerOwnableID := baseIDs.NewOwnableID(baseIDs.NewStringID("makerOwnableID"))
	testTakerOwnableID := baseIDs.NewOwnableID(baseIDs.NewStringID("takerOwnableID"))
	rate := sdkTypes.SmallestDec()
	creationHeight := baseTypes.NewHeight(0)
	// takerOwnableID ids.OwnableID, rate types.Dec, creationHeight types.Height, makerID ids.IdentityID, imutables qualified.Immutables

	testOrderID := baseIDs.NewOrderID(testClassificationID, testMakerOwnableID, testTakerOwnableID, rate, creationHeight, testFromID, immutables)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	testMessage := newMessage(fromAccAddress, testFromID, testOrderID)

	return testOrderID, testFromID, fromAccAddress, testMessage
}

func Test_messageFromInterface(t *testing.T) {
	testOrderID, testFromID, fromAccAddress, testMessage := CreateTestInputForMessages(t)
	type args struct {
		msg sdkTypes.Msg
	}
	tests := []struct {
		name string
		args args
		want message
	}{
		// TODO: Add test cases.
		{"+ve", args{testMessage}, message{fromAccAddress, testFromID, testOrderID}},
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
	testOrderID, testFromID, fromAccAddress, testMessage := CreateTestInputForMessages(t)

	type fields struct {
		From    sdkTypes.AccAddress
		FromID  ids.IdentityID
		OrderID ids.OrderID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testOrderID}, sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(testMessage))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				OrderID: tt.fields.OrderID,
			}
			if got := message.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSignBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_GetSigners(t *testing.T) {
	testOrderID, testFromID, fromAccAddress, _ := CreateTestInputForMessages(t)

	type fields struct {
		From    sdkTypes.AccAddress
		FromID  ids.IdentityID
		OrderID ids.OrderID
	}
	tests := []struct {
		name   string
		fields fields
		want   []sdkTypes.AccAddress
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testOrderID}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				OrderID: tt.fields.OrderID,
			}
			if got := message.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_RegisterCodec(t *testing.T) {
	testOrderID, testFromID, fromAccAddress, _ := CreateTestInputForMessages(t)

	type fields struct {
		From    sdkTypes.AccAddress
		FromID  ids.IdentityID
		OrderID ids.OrderID
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
		{"+ve", fields{fromAccAddress, testFromID, testOrderID}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				OrderID: tt.fields.OrderID,
			}
			me.RegisterCodec(tt.args.codec)
		})
	}
}

func Test_message_Route(t *testing.T) {
	testOrderID, testFromID, fromAccAddress, _ := CreateTestInputForMessages(t)

	type fields struct {
		From    sdkTypes.AccAddress
		FromID  ids.IdentityID
		OrderID ids.OrderID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testOrderID}, module.Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				OrderID: tt.fields.OrderID,
			}
			if got := message.Route(); got != tt.want {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_Type(t *testing.T) {
	testOrderID, testFromID, fromAccAddress, _ := CreateTestInputForMessages(t)

	type fields struct {
		From    sdkTypes.AccAddress
		FromID  ids.IdentityID
		OrderID ids.OrderID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testOrderID}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				OrderID: tt.fields.OrderID,
			}
			if got := message.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_ValidateBasic(t *testing.T) {
	testOrderID, testFromID, fromAccAddress, _ := CreateTestInputForMessages(t)

	type fields struct {
		From    sdkTypes.AccAddress
		FromID  ids.IdentityID
		OrderID ids.OrderID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{fromAccAddress, testFromID, testOrderID}, false},
		{"-ve", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				OrderID: tt.fields.OrderID,
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newMessage(t *testing.T) {
	testOrderID, testFromID, fromAccAddress, _ := CreateTestInputForMessages(t)

	type args struct {
		from    sdkTypes.AccAddress
		fromID  ids.IdentityID
		orderID ids.OrderID
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{
		// TODO: Add test cases.
		{"+ve", args{fromAccAddress, testFromID, testOrderID}, message{fromAccAddress, testFromID, testOrderID}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMessage(tt.args.from, tt.args.fromID, tt.args.orderID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
