// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package put

import (
	"cosmossdk.io/math"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	typesSchema "github.com/AssetMantle/schema/types"
	baseTypes "github.com/AssetMantle/schema/types/base"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var _ helpers.Message = (*Message)(nil)

func (message *Message) ValidateBasic() error {
	if _, err := sdkTypes.AccAddressFromBech32(message.From); err != nil {
		return err
	}

	if err := message.FromID.ValidateBasic(); err != nil {
		return err
	}

	if err := message.MakerAssetID.ValidateBasic(); err != nil {
		return err
	}

	if err := message.TakerAssetID.ValidateBasic(); err != nil {
		return err
	}

	if _, ok := sdkTypes.NewIntFromString(message.MakerSplit); !ok {
		return constants.IncorrectFormat.Wrapf("maker split %s is not a valid integer", message.MakerSplit)
	}

	if _, ok := sdkTypes.NewIntFromString(message.TakerSplit); !ok {
		return constants.IncorrectFormat.Wrapf("taker split %s is not a valid integer", message.TakerSplit)
	}

	if err := message.ExpiryHeight.ValidateBasic(); err != nil {
		return err
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

func NewMessage(from sdkTypes.AccAddress, fromID ids.IdentityID, makerAssetID ids.AssetID, takerAssetID ids.AssetID, makerSplit math.Int, takerSplit math.Int, expiryHeight typesSchema.Height) sdkTypes.Msg {
	return &Message{
		From:         from.String(),
		FromID:       fromID.(*baseIDs.IdentityID),
		MakerAssetID: makerAssetID.(*baseIDs.AssetID),
		TakerAssetID: takerAssetID.(*baseIDs.AssetID),
		MakerSplit:   makerSplit.String(),
		TakerSplit:   takerSplit.String(),
		ExpiryHeight: expiryHeight.(*baseTypes.Height),
	}
}
