/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type metaID struct {
	TypeID types.ID `json:"typeID" valid:"required~required field typeID missing"`
	HashID types.ID `json:"hashID" valid:"required~required field hashID missing"`
}

var _ types.ID = (*metaID)(nil)
var _ helpers.Key = (*metaID)(nil)

func (metaID metaID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, metaID.TypeID.Bytes()...)
	Bytes = append(Bytes, metaID.HashID.Bytes()...)

	return Bytes
}
func (metaID metaID) String() string {
	var values []string
	values = append(values, metaID.TypeID.String())
	values = append(values, metaID.HashID.String())

	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}
func (metaID metaID) Equals(id types.ID) bool {
	return bytes.Equal(metaID.Bytes(), id.Bytes())
}
func (metaID metaID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(metaID.Bytes())
}
func (metaID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, metaID{})
}
func (metaID metaID) IsPartial() bool {
	return len(metaID.HashID.Bytes()) == 0
}
func (metaID metaID) Matches(key helpers.Key) bool {
	return metaID.Equals(metaIDFromInterface(key))
}

func NewMetaID(typeID types.ID, hashID types.ID) types.ID {
	return metaID{
		TypeID: typeID,
		HashID: hashID,
	}
}
