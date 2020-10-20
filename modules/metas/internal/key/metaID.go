/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type metaID struct {
	HashID types.ID `json:"hashID" valid:"required~required field hashID missing"`
}

var _ types.ID = (*metaID)(nil)
var _ helpers.Key = (*metaID)(nil)

func (MetaID metaID) Bytes() []byte {
	return MetaID.HashID.Bytes()
}
func (MetaID metaID) String() string {
	return MetaID.HashID.String()
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

func New(id types.ID) helpers.Key {
	return metaIDFromInterface(id)
}

func NewMetaID(hashID types.ID) types.ID {
	return metaID{
		HashID: hashID,
	}
}
