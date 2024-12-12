// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrap

import (
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/cometbft/cometbft/crypto/ed25519"
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

var (
	testAddress = types.AccAddress(ed25519.GenPrivKey().PubKey().Address()).String()
	testID      = baseIDs.PrototypeIdentityID().(*baseIDs.IdentityID)
	testCoins   = types.NewCoins(types.NewCoin(denom, types.NewInt(100)))
)

const (
	denom = "stake"
)

func TestMessage_ValidateBasic(t *testing.T) {
	tests := []struct {
		name    string
		message *Message
		wantErr helpers.Error
	}{
		{
			"valid message",
			&Message{
				From:   testAddress,
				FromID: testID,
				Coins:  testCoins,
			},
			nil,
		},
		{
			"multiple coins",
			&Message{
				From:   testAddress,
				FromID: testID,
				Coins:  types.Coins{types.Coin{Denom: denom + "1", Amount: types.NewInt(100)}, types.Coin{Denom: denom + "2", Amount: types.NewInt(100)}},
			},
			nil,
		},
		{
			"duplicate denom",
			&Message{
				From:   testAddress,
				FromID: testID,
				Coins:  types.Coins{types.Coin{Denom: denom, Amount: types.NewInt(100)}, types.Coin{Denom: denom, Amount: types.NewInt(100)}},
			},
			constants.InvalidMessage,
		},
		{
			"too many coins",
			&Message{
				From:   testAddress,
				FromID: testID,
				Coins: types.Coins{
					types.Coin{Denom: denom, Amount: types.NewInt(100)},
					types.Coin{Denom: denom + "01", Amount: types.NewInt(100)},
					types.Coin{Denom: denom + "02", Amount: types.NewInt(100)},
					types.Coin{Denom: denom + "03", Amount: types.NewInt(100)},
					types.Coin{Denom: denom + "04", Amount: types.NewInt(100)},
					types.Coin{Denom: denom + "05", Amount: types.NewInt(100)},
					types.Coin{Denom: denom + "06", Amount: types.NewInt(100)},
					types.Coin{Denom: denom + "07", Amount: types.NewInt(100)},
					types.Coin{Denom: denom + "08", Amount: types.NewInt(100)},
					types.Coin{Denom: denom + "09", Amount: types.NewInt(100)},
					types.Coin{Denom: denom + "10", Amount: types.NewInt(100)},
				},
			},
			constants.InvalidMessage,
		},
		{
			"unsorted coins",
			&Message{
				From:   testAddress,
				FromID: testID,
				Coins: types.Coins{
					types.Coin{Denom: denom + "1", Amount: types.NewInt(100)},
					types.Coin{Denom: denom, Amount: types.NewInt(100)},
				},
			},
			constants.InvalidMessage,
		},
		{
			"empty message",
			&Message{},
			constants.InvalidMessage,
		},
		{
			"invalid from address",
			&Message{
				From:   "invalid",
				FromID: testID,
				Coins:  testCoins,
			},
			constants.InvalidMessage,
		},
		{
			name: "invalid from id",
			message: &Message{
				From:   testAddress,
				FromID: &baseIDs.IdentityID{HashID: &baseIDs.HashID{IDBytes: []byte{1, 2, 3, 4}}},
				Coins:  testCoins,
			},
			wantErr: constants.InvalidMessage,
		},
		{
			name: "invalid coin amount",
			message: &Message{
				From:   testAddress,
				FromID: testID,
				Coins:  types.Coins{types.Coin{Denom: denom, Amount: types.NewInt(-100)}},
			},
			wantErr: constants.InvalidMessage,
		},
		{
			name: "invalid coin denom",
			message: &Message{
				From:   testAddress,
				FromID: testID,
				Coins:  types.Coins{types.Coin{Denom: "", Amount: types.NewInt(100)}},
			},
			wantErr: constants.InvalidMessage,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.message.ValidateBasic()

			if err != nil && tt.wantErr == nil || err == nil && tt.wantErr != nil || err != nil && tt.wantErr != nil && !tt.wantErr.Is(err) {
				t.Errorf("\n got: \n %v \n want: \n %v", err, tt.wantErr)
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

func Test_message_GetSigners(t *testing.T) {
	type fields struct {
		From   string
		FromID *baseIDs.IdentityID
		Coins  types.Coins
	}
	tests := []struct {
		name   string
		fields fields
		want   []types.AccAddress
	}{
		{"+ve", fields{fromAccAddress.String(), fromID, coins}, []types.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:   tt.fields.From,
				FromID: tt.fields.FromID,
				Coins:  tt.fields.Coins,
			}
			if got := message.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NewMessage(t *testing.T) {
	type args struct {
		from   types.AccAddress
		fromID ids.IdentityID
		coins  types.Coins
	}
	tests := []struct {
		name string
		args args
		want types.Msg
	}{
		{"+ve", args{fromAccAddress, fromID, coins}, &Message{fromAccAddress.String(), fromID, coins}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessage(tt.args.from, tt.args.fromID, tt.args.coins); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
