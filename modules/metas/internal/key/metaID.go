/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
	"strings"
)

var _ types.ID = (*MetaID)(nil)
var _ helpers.Key = (*MetaID)(nil)

func (metaID MetaID) GetStructReference() codec.ProtoMarshaler {
	return &metaID
}
func (metaID MetaID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, metaID.TypeID.Bytes()...)
	Bytes = append(Bytes, metaID.HashID.Bytes()...)

	return Bytes
}
func (metaID MetaID) String() string {
	var values []string
	values = append(values, metaID.TypeID.String())
	values = append(values, metaID.HashID.String())

	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}
func (metaID MetaID) Compare(id types.ID) int {
	return bytes.Compare(metaID.Bytes(), id.Bytes())
}
func (metaID MetaID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(metaID.Bytes())
}
func (MetaID) RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, module.Name, MetaID{})
}
func (metaID MetaID) IsPartial() bool {
	return len(metaID.HashID.Bytes()) == 0
}
func (metaID MetaID) Equals(key helpers.Key) bool {
	id := metaIDFromInterface(key)
	return metaID.Compare(&id) == 0
}

func NewMetaID(typeID types.ID, hashID types.ID) types.ID {
	return &MetaID{
		TypeID: *base.NewID(typeID.String()),
		HashID: *base.NewID(hashID.String()),
	}
}
