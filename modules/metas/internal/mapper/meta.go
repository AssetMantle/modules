/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type meta struct {
	Data types.Data `json:"data" valid:"required field data missing"`
}

var _ mappables.Meta = (*meta)(nil)

func (meta meta) GetData() types.Data { return meta.Data }

func (meta meta) GetID() types.ID { return base.NewID(meta.Data.GenerateHash()) }

func (meta meta) Encode() []byte {
	return packageCodec.MustMarshalBinaryBare(meta)
}

func (meta meta) Decode(bytes []byte) helpers.Mappable {
	packageCodec.MustUnmarshalBinaryBare(bytes, &meta)
	return meta
}

func metaPrototype() helpers.Mappable {
	return meta{}
}

func NewMeta(data types.Data) mappables.Meta {
	return meta{
		Data: data,
	}
}
