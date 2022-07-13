// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/metas/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/traits"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type metaID struct {
	Type ids.ID
	Hash ids.ID
}

var _ ids.MetaID = (*metaID)(nil)
var _ helpers.Key = (*metaID)(nil)

func (metaID metaID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, metaID.Type.Bytes()...)
	Bytes = append(Bytes, metaID.Hash.Bytes()...)

	return Bytes
}
func (metaID metaID) String() string {
	return stringUtilities.JoinIDStrings(metaID.Type.String(), metaID.Hash.String())
}
func (metaID metaID) Compare(listable traits.Listable) int {
	return bytes.Compare(metaID.Bytes(), metaIDFromInterface(listable).Bytes())
}
func (metaID metaID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(metaID.Bytes())
}
func (metaID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, metaID{})
}
func (metaID metaID) IsPartial() bool {
	return len(metaID.Hash.Bytes()) == 0
}
func (metaID metaID) Equals(key helpers.Key) bool {
	return metaID.Compare(metaIDFromInterface(key)) == 0
}

func NewMetaID(Type ids.ID, hash ids.ID) ids.MetaID {
	return metaID{
		Type: Type,
		Hash: hash,
	}
}
