// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package immediate

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	typesSchema "github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

var _ helpers.Message = (*Message)(nil)

func (message *Message) Type() string { return Transaction.GetName() }
func (message *Message) ValidateBasic() error {
	if _, err := sdkTypes.AccAddressFromBech32(message.From); err != nil {
		return err
	}
	if err := message.FromID.ValidateBasic(); err != nil {
		return err
	}
	if err := message.ClassificationID.ValidateBasic(); err != nil {
		return err
	}
	if err := message.TakerID.ValidateBasic(); err != nil {
		return err
	}
	if err := message.MakerOwnableID.ValidateBasic(); err != nil {
		return err
	}
	if err := message.TakerOwnableID.ValidateBasic(); err != nil {
		return err
	}
	if err := message.ExpiresIn.ValidateBasic(); err != nil {
		return err
	}
	if err := message.ImmutableMetaProperties.ValidateBasic(); err != nil {
		return err
	}
	if err := message.MutableMetaProperties.ValidateBasic(); err != nil {
		return err
	}
	if err := message.ImmutableProperties.ValidateBasic(); err != nil {
		return err
	}
	if err := message.MutableProperties.ValidateBasic(); err != nil {
		return err
	}
	if _, err := sdkTypes.NewDecFromStr(message.MakerOwnableSplit); err != nil {
		return err
	} else if _, err := sdkTypes.NewDecFromStr(message.TakerOwnableSplit); err != nil {
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
func (*Message) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, Message{})
}
func (message *Message) RegisterInterface(interfaceRegistry types.InterfaceRegistry) {
	interfaceRegistry.RegisterImplementations((*sdkTypes.Msg)(nil), message)
}

func messageFromInterface(msg sdkTypes.Msg) *Message {
	switch value := msg.(type) {
	case *Message:
		return value
	default:
		return &Message{}
	}
}
func messagePrototype() helpers.Message {
	return &Message{}
}
func newMessage(from sdkTypes.AccAddress, fromID ids.IdentityID, classificationID ids.ClassificationID, takerID ids.IdentityID, makerOwnableID ids.AnyOwnableID, takerOwnableID ids.AnyOwnableID, expiresIn typesSchema.Height, makerOwnableSplit sdkTypes.Dec, takerOwnableSplit sdkTypes.Dec, immutableMetaProperties lists.PropertyList, immutableProperties lists.PropertyList, mutableMetaProperties lists.PropertyList, mutableProperties lists.PropertyList) sdkTypes.Msg {
	return &Message{
		From:                    from.String(),
		FromID:                  fromID.(*baseIDs.IdentityID),
		ClassificationID:        classificationID.(*baseIDs.ClassificationID),
		TakerID:                 takerID.(*baseIDs.IdentityID),
		MakerOwnableID:          makerOwnableID.(*baseIDs.AnyOwnableID),
		TakerOwnableID:          takerOwnableID.(*baseIDs.AnyOwnableID),
		ExpiresIn:               expiresIn.(*baseTypes.Height),
		MakerOwnableSplit:       makerOwnableSplit.String(),
		TakerOwnableSplit:       takerOwnableSplit.String(),
		ImmutableMetaProperties: immutableMetaProperties.(*baseLists.PropertyList),
		ImmutableProperties:     immutableProperties.(*baseLists.PropertyList),
		MutableMetaProperties:   mutableMetaProperties.(*baseLists.PropertyList),
		MutableProperties:       mutableProperties.(*baseLists.PropertyList),
	}
}
