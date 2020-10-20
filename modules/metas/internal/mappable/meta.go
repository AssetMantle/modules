/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
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

type meta struct {
	Data types.Data `json:"data" valid:"required field data missing"`
}

var _ mappables.Meta = (*meta)(nil)

func (meta meta) GetData() types.Data { return meta.Data }
func (meta meta) GetID() types.ID     { return base.NewID(meta.Data.GenerateHash()) }
func (meta meta) GetKey() helpers.Key {
	return key.New(meta.GetID())
}
func (meta) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, meta{})
}

func NewMeta(data types.Data) mappables.Meta {
	return meta{
		Data: data,
	}
}
