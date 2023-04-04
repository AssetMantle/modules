// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/schema/x/helpers"
)

var _ helpers.Message = (*TestMessage)(nil)

func NewTestMessage(addr sdkTypes.AccAddress, id string) sdkTypes.Msg {
	return &TestMessage{
		From: addr,
		ID:   id,
	}
}

func (m *TestMessage) RegisterInterface(registry types.InterfaceRegistry) {
	// TODO implement me
	panic("implement me")
}

func (m *TestMessage) GenerateOnSuccessEvents() sdkTypes.Events {
	// TODO implement me
	panic("implement me")
}

func (message TestMessage) Route() string { return "TestMessage" }
func (message TestMessage) Type() string  { return "TestMessage" }
func (message TestMessage) GetSignBytes() []byte {
	bz, err := json.Marshal(message.From)
	if err != nil {
		panic(err)
	}

	return sdkTypes.MustSortJSON(bz)
}
func (message TestMessage) ValidateBasic() error { return nil }
func (message TestMessage) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.From}
}
func (message TestMessage) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	legacyAmino.RegisterConcrete(TestMessage{}, "test/TestMessage", nil)
}

func TestMessagePrototype() helpers.Message {
	return &TestMessage{}
}
