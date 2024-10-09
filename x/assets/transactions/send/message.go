// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package send

import (
	"cosmossdk.io/math"
	"github.com/AssetMantle/modules/helpers"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var _ helpers.Message = (*Message)(nil)

func (message *Message) GetValueAsInt() (math.Int, error) {
	value, ok := sdkTypes.NewIntFromString(message.Value)
	if !ok {
		return sdkTypes.ZeroInt(), errorConstants.IncorrectFormat.Wrapf("send value %s is not a valid integer", message.Value)
	} else if value.IsNegative() {
		return sdkTypes.ZeroInt(), errorConstants.InvalidParameter.Wrapf("send value is negative %s", message.Value)
	}

	return value, nil
}

func (message *Message) ValidateBasic() error {
	if _, err := sdkTypes.AccAddressFromBech32(message.From); err != nil {
		return errorConstants.InvalidMessage.Wrapf("invalid from address %s", err.Error())
	}
	if err := message.FromID.ValidateBasic(); err != nil {
		return errorConstants.InvalidMessage.Wrapf("invalid from id %s", err.Error())
	}
	if err := message.ToID.ValidateBasic(); err != nil {
		return errorConstants.InvalidMessage.Wrapf("invalid to id %s", err.Error())
	}
	if err := message.AssetID.ValidateBasic(); err != nil {
		return errorConstants.InvalidMessage.Wrapf("invalid asset id %s", err.Error())
	}
	if _, err := message.GetValueAsInt(); err != nil {
		return errorConstants.InvalidMessage.Wrapf("invalid value %s", err.Error())
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
func (message *Message) RegisterInterface(interfaceRegistry types.InterfaceRegistry) {
	interfaceRegistry.RegisterImplementations((*sdkTypes.Msg)(nil), message)
}

func messagePrototype() helpers.Message {
	return &Message{}
}
func NewMessage(from sdkTypes.AccAddress, fromID ids.IdentityID, toID ids.IdentityID, assetID ids.AssetID, value math.Int) sdkTypes.Msg {
	return &Message{
		From:    from.String(),
		FromID:  fromID.(*baseIDs.IdentityID),
		ToID:    toID.(*baseIDs.IdentityID),
		AssetID: assetID.(*baseIDs.AssetID),
		Value:   value.String(),
	}
}
