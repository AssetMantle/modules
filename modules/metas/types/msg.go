package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

const (
	TypeMsgReveal = "meta_reveal"
)

var (
	_ sdk.Msg = &MsgReveal{}
)

func NewMsgReveal(from sdk.AccAddress, data types.Data) *MsgReveal {
	return &MsgReveal{
		From:     from.String(),
		MetaFact: base.NewMetaFact(data),
	}
}

func (msg MsgReveal) Route() string { return ModuleName }

// Type implements the sdk.Msg interface.
func (msg MsgReveal) Type() string { return TypeMsgReveal }

// GetSigners implements the sdk.Msg interface.
func (msg MsgReveal) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{address.Bytes()}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgReveal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgReveal) ValidateBasic() error {
	if msg.From == "" {
		return errors.EmptyFromAddress
	}

	// TODO
	return nil
}
