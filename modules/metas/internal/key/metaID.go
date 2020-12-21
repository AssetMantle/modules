/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
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
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
	"strings"
)

type metaID struct {
	TypeID types.ID `json:"typeID" valid:"required~required field typeID missing"`
	HashID types.ID `json:"hashID" valid:"required~required field hashID missing"`
}

var _ types.ID = (*metaID)(nil)
var _ helpers.Key = (*metaID)(nil)

func (MetaID metaID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, MetaID.TypeID.Bytes()...)
	Bytes = append(Bytes, MetaID.HashID.Bytes()...)
	return Bytes
}
func (MetaID metaID) String() string {
	var values []string
	values = append(values, MetaID.TypeID.String())
	values = append(values, MetaID.HashID.String())
	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}
func (MetaID metaID) Equals(id types.ID) bool {
	return bytes.Compare(MetaID.Bytes(), id.Bytes()) == 0
}
func (MetaID metaID) GenerateStoreKeyBytes() []byte {
	return append([]byte{0x13}, MetaID.Bytes()...)
}
func (metaID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, metaID{})
}
func (MetaID metaID) IsPartial() bool {
	if len(MetaID.HashID.Bytes()) > 0 {
		return false
	}
	return true
}
func (MetaID metaID) Matches(key helpers.Key) bool {
	switch value := key.(type) {
	case metaID:
		return bytes.Compare(MetaID.Bytes(), value.Bytes()) == 0
	default:
		return false
	}
}

func GenerateMetaID(data types.Data) types.ID {
	return metaID{
		TypeID: data.GetTypeID(),
		HashID: data.GenerateHashID(),
	}
}

func New(id types.ID) helpers.Key {
	return metaIDFromInterface(id)
}

func NewMetaID(typeID types.ID, hashID types.ID) types.ID {
	return metaID{
		TypeID: typeID,
		HashID: hashID,
	}
}
