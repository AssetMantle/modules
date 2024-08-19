// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package send

import (
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/documents/base"
	"github.com/cometbft/cometbft/crypto/ed25519"
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

const (
	denom = "stake"
)

var (
	testAddress = types.AccAddress(ed25519.GenPrivKey().PubKey().Address()).String()
	testID      = baseIDs.PrototypeIdentityID().(*baseIDs.IdentityID)
	coinAsset   = base.NewCoinAsset(denom)
	coinAssetID = coinAsset.GetCoinAssetID().(*baseIDs.AssetID)
	testValue   = types.NewInt(100).String()
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
				From:    testAddress,
				FromID:  testID,
				ToID:    testID,
				AssetID: coinAssetID,
				Value:   testValue,
			},
			nil,
		},
		{
			"empty message",
			&Message{},
			errorConstants.InvalidMessage,
		},
		{
			"invalid from address",
			&Message{
				From:    "invalid",
				FromID:  testID,
				ToID:    testID,
				AssetID: coinAssetID,
				Value:   testValue,
			},
			errorConstants.InvalidMessage,
		},
		{
			"invalid from id",
			&Message{
				From:    testAddress,
				FromID:  &baseIDs.IdentityID{HashID: &baseIDs.HashID{IDBytes: []byte{1, 2, 3, 4}}},
				ToID:    testID,
				AssetID: coinAssetID,
				Value:   testValue,
			},
			errorConstants.InvalidMessage,
		},
		{
			"invalid to id",
			&Message{
				From:    testAddress,
				FromID:  testID,
				ToID:    &baseIDs.IdentityID{HashID: &baseIDs.HashID{IDBytes: []byte{1, 2, 3, 4}}},
				AssetID: coinAssetID,
				Value:   testValue,
			},
			errorConstants.InvalidMessage,
		},
		{
			"invalid asset id",
			&Message{
				From:    testAddress,
				FromID:  testID,
				ToID:    testID,
				AssetID: &baseIDs.AssetID{HashID: &baseIDs.HashID{IDBytes: []byte{1, 2, 3, 4}}},
				Value:   testValue,
			},
			errorConstants.InvalidMessage,
		},
		{
			"invalid value",
			&Message{
				From:    testAddress,
				FromID:  testID,
				ToID:    testID,
				AssetID: coinAssetID,
				Value:   "invalid",
			},
			errorConstants.InvalidMessage,
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

func Test_messageFromInterface(t *testing.T) {
	type args struct {
		msg types.Msg
	}
	tests := []struct {
		name string
		args args
		want *Message
	}{
		{"+ve", args{NewMessage(fromAccAddress, fromID, fromID, assetID, testRate)}, &Message{fromAccAddress.String(), fromID, fromID, assetID, testRate.String()}},
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

func Test_message_GetSigners(t *testing.T) {
	type fields struct {
		From    string
		FromID  *baseIDs.IdentityID
		ToID    *baseIDs.IdentityID
		AssetID *baseIDs.AssetID
		Value   types.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   []types.AccAddress
	}{
		{"+ve", fields{fromAccAddress.String(), fromID, fromID, assetID, testRate}, []types.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				ToID:    tt.fields.ToID,
				AssetID: tt.fields.AssetID,
				Value:   tt.fields.Value.String(),
			}
			if got := message.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_RegisterCodec(t *testing.T) {
	type fields struct {
		From    string
		FromID  *baseIDs.IdentityID
		ToID    *baseIDs.IdentityID
		AssetID *baseIDs.AssetID
		Value   types.Int
	}
	type args struct {
		legacyAmino *sdkCodec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{fromAccAddress.String(), fromID, fromID, assetID, testRate}, args{codec.GetLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				ToID:    tt.fields.ToID,
				AssetID: tt.fields.AssetID,
				Value:   tt.fields.Value.String(),
			}
			me.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_message_Type(t *testing.T) {
	type fields struct {
		From    string
		FromID  *baseIDs.IdentityID
		ToID    *baseIDs.IdentityID
		AssetID *baseIDs.AssetID
		Value   types.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{fromAccAddress.String(), fromID, fromID, assetID, testRate}, Transaction.GetServicePath()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				ToID:    tt.fields.ToID,
				AssetID: tt.fields.AssetID,
				Value:   tt.fields.Value.String(),
			}
			if got := message.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NewMessage(t *testing.T) {
	type args struct {
		from    types.AccAddress
		fromID  ids.IdentityID
		toID    ids.IdentityID
		assetID ids.AssetID
		value   types.Int
	}
	tests := []struct {
		name string
		args args
		want types.Msg
	}{
		{"+ve", args{fromAccAddress, fromID, fromID, assetID, testRate}, &Message{fromAccAddress.String(), fromID, fromID, assetID, testRate.String()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessage(tt.args.from, tt.args.fromID, tt.args.toID, tt.args.assetID, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
