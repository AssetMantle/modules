/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

// msg type for testing
type TestMessage struct {
	From sdkTypes.AccAddress
	ID   string
}

var _ helpers.Message = (*TestMessage)(nil)

func NewTestMessage(addr sdkTypes.AccAddress, id string) sdkTypes.Msg {
	return TestMessage{
		From: addr,
		ID:   id,
	}
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
func (message TestMessage) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(TestMessage{}, "test/TestMessage", nil)
}

func TestMessagePrototype() helpers.Message {
	return TestMessage{}
}
