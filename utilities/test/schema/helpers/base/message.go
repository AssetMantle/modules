package base

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

// msg type for testing
type testMsg struct {
	From sdkTypes.AccAddress
	ID   string
}

var _ helpers.Message = (*testMsg)(nil)

func NewTestMsg(addr sdkTypes.AccAddress, id string) *testMsg {
	return &testMsg{
		From: addr,
		ID:   id,
	}
}
func (msg testMsg) Route() string { return "testMsg" }
func (msg testMsg) Type() string  { return "Test message" }
func (msg testMsg) GetSignBytes() []byte {
	bz, err := json.Marshal(msg.From)
	if err != nil {
		panic(err)
	}
	return sdkTypes.MustSortJSON(bz)
}
func (msg testMsg) ValidateBasic() error { return nil }
func (msg testMsg) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{msg.From}
}
func (msg testMsg) RegisterCodec(_ *codec.Codec) {
}

func TestMessagePrototype() helpers.Message {
	return testMsg{}
}
