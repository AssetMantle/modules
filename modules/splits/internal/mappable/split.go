/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

var _ mappables.Split = (*Split)(nil)

func (split Split) GetStructReference() codec.ProtoMarshaler {
	return &split
}
func (split Split) GetID() types.ID { return split.ID }
func (split Split) GetOwnerID() types.ID {
	return key.ReadOwnerID(split.ID)
}
func (split Split) GetOwnableID() types.ID {
	return key.ReadOwnableID(split.ID)
}
func (split Split) GetValue() sdkTypes.Dec {
	return split.Value
}
func (split Split) Send(outValue sdkTypes.Dec) traits.Transactional {
	result := split.Value.Sub(outValue)
	split.Value = result
	return split
}
func (split Split) Receive(inValue sdkTypes.Dec) traits.Transactional {
	result := split.Value.Add(inValue)
	split.Value = result
	return split
}
func (split Split) CanSend(outValue sdkTypes.Dec) bool {
	return split.Value.GTE(outValue)
}
func (split Split) GetKey() helpers.Key {
	return key.FromID(split.ID)
}
func (Split) RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, module.Name, Split{})
}

func NewSplit(splitID types.ID, value sdkTypes.Dec) mappables.Split {
	return &Split{
		ID:    splitID,
		Value: value,
	}
}
