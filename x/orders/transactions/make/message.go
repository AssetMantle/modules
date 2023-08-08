// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

import (
	codecUtilities "github.com/AssetMantle/schema/go/codec/utilities"
	"github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/lists"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	typesSchema "github.com/AssetMantle/schema/go/types"
	baseTypes "github.com/AssetMantle/schema/go/types/base"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
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
	if err := message.MakerAssetID.ValidateBasic(); err != nil {
		return err
	}
	if err := message.TakerAssetID.ValidateBasic(); err != nil {
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
	if _, ok := sdkTypes.NewIntFromString(message.MakerSplit); !ok {
		return constants.IncorrectFormat.Wrapf("maker split %s is not a valid integer", message.MakerSplit)
	} else if _, ok := sdkTypes.NewIntFromString(message.TakerSplit); !ok {
		return constants.IncorrectFormat.Wrapf("taker split %s is not a valid integer", message.TakerSplit)
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
func NewMessage(from sdkTypes.AccAddress, fromID ids.IdentityID, classificationID ids.ClassificationID, takerID ids.IdentityID, makerAssetID ids.AssetID, takerAssetID ids.AssetID, expiresIn typesSchema.Height, makerSplit sdkTypes.Int, takerSplit sdkTypes.Int, immutableMetaProperties lists.PropertyList, immutableProperties lists.PropertyList, mutableMetaProperties lists.PropertyList, mutableProperties lists.PropertyList) sdkTypes.Msg {

	return &Message{
		From:                    from.String(),
		FromID:                  fromID.(*baseIDs.IdentityID),
		ClassificationID:        classificationID.(*baseIDs.ClassificationID),
		TakerID:                 takerID.(*baseIDs.IdentityID),
		MakerAssetID:            makerAssetID.(*baseIDs.AssetID),
		TakerAssetID:            takerAssetID.(*baseIDs.AssetID),
		ExpiresIn:               expiresIn.(*baseTypes.Height),
		MakerSplit:              makerSplit.String(),
		TakerSplit:              takerSplit.String(),
		ImmutableMetaProperties: immutableMetaProperties.(*baseLists.PropertyList),
		ImmutableProperties:     immutableProperties.(*baseLists.PropertyList),
		MutableMetaProperties:   mutableMetaProperties.(*baseLists.PropertyList),
		MutableProperties:       mutableProperties.(*baseLists.PropertyList),
	}
}
