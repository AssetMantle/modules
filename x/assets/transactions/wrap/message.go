// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrap

import (
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/assets/constants"
	codecUtilities "github.com/AssetMantle/schema/codec/utilities"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

var _ helpers.Message = (*Message)(nil)

func (message *Message) ValidateBasic() error {
	if _, err := sdkTypes.AccAddressFromBech32(message.From); err != nil {
		return errorConstants.InvalidMessage.Wrapf("invalid from address %s", err.Error())
	}
	if err := message.FromID.ValidateBasic(); err != nil {
		return errorConstants.InvalidMessage.Wrapf("invalid from id %s", err.Error())
	}
	if err := message.Coins.Validate(); err != nil {
		return errorConstants.InvalidMessage.Wrapf("invalid coins %s", err.Error())
	}
	if message.Coins.Len() > constants.MaxListLength {
		return errorConstants.InvalidMessage.Wrapf("number of coins in message: %d exceeds maximum allowed: %d", message.Coins.Len(), constants.MaxListLength)
	}
	return nil
}
func (message *Message) GetSigners() []sdkTypes.AccAddress {
	from, err := sdkTypes.AccAddressFromBech32(message.From)
	if err != nil {
		panic(err)
	}
	return []sdkTypes.AccAddress{from}
}
func (*Message) RegisterLegacyAminoCodec(legacyAmino *sdkCodec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, Message{})
}
func (message *Message) RegisterInterface(interfaceRegistry types.InterfaceRegistry) {
	interfaceRegistry.RegisterImplementations((*sdkTypes.Msg)(nil), message)
}

func messagePrototype() helpers.Message {
	return &Message{}
}

func NewMessage(from sdkTypes.AccAddress, fromID ids.IdentityID, coins sdkTypes.Coins) sdkTypes.Msg {
	return &Message{
		From:   from.String(),
		FromID: fromID.(*baseIDs.IdentityID),
		Coins:  coins,
	}
}
