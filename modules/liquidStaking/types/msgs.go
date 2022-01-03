package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgAddChain = "liquidStaking/AddChain"
)

var _ sdk.Msg = &MsgAddChain{}

func NewMsgAddChain(fromAddr sdk.AccAddress, amount sdk.Coin) *MsgAddChain {
	return &MsgAddChain{FromAddress: fromAddr.String(), Amount: amount}
}

// Route Implements Msg.
func (msg MsgAddChain) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgAddChain) Type() string { return TypeMsgAddChain }

// ValidateBasic Implements Msg.
func (msg MsgAddChain) ValidateBasic() error {
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgAddChain) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners Implements Msg.
func (msg MsgAddChain) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
