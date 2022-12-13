/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

var _ mappables.Meta = (*Meta)(nil)

func (meta Meta) GetStructReference() codec.ProtoMarshaler {
	return &meta
}
func (meta Meta) GetData() types.Data { return &meta.Data }
func (meta Meta) GetID() types.ID     { return &meta.ID }
func (meta Meta) GetKey() helpers.Key {
	return key.FromID(meta.GetID())
}
func (Meta) RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, module.Name, Meta{})
}

func NewMeta(data types.Data) *Meta {
	return &Meta{
		ID:   *base.NewID(key.GenerateMetaID(data).String()),
		Data: *base.NewData(data),
	}
}
