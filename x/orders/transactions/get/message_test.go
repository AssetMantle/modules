// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package get

import (
	"cosmossdk.io/math"
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
)

var (
	testMessage = NewMessage(fromAccAddress, testFromID, testOrderID)
)

type fields struct {
	From       string
	FromID     *baseIDs.IdentityID
	TakerSplit math.Int
	OrderID    *baseIDs.OrderID
}

func Test_messagePrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Message
	}{
		{"valid", &Message{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := messagePrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("messagePrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_GetSigners(t *testing.T) {

	tests := []struct {
		name   string
		fields fields
		want   []types.AccAddress
	}{
		{"valid", fields{fromAccAddress.String(), testFromID, takerSplit, testOrderID}, []types.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
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

func Test_message_ValidateBasic(t *testing.T) {
	t.Run("nil inputs error", func(t *testing.T) {
		message := &Message{}
		err := message.ValidateBasic()
		require.Error(t, err)
	})
	t.Run("valid", func(t *testing.T) {
		message := &Message{
			From:    fromAccAddress.String(),
			FromID:  testFromID,
			OrderID: testOrderID,
		}
		err := message.ValidateBasic()
		require.NoError(t, err)
	})
}

func Test_NewMessage(t *testing.T) {
	type args struct {
		from       types.AccAddress
		fromID     ids.IdentityID
		takerSplit math.Int
		orderID    ids.OrderID
	}
	tests := []struct {
		name string
		args args
		want types.Msg
	}{
		{"valid", args{fromAccAddress, testFromID, takerSplit, testOrderID}, &Message{fromAccAddress.String(), testFromID, testOrderID}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessage(tt.args.from, tt.args.fromID, tt.args.orderID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
