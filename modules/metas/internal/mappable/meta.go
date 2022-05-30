// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/metas/internal/key"
	"github.com/AssetMantle/modules/modules/metas/internal/module"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/mappables"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type meta struct {
	ID   ids.ID    `json:"id" valid:"required field id missing"`
	Data data.Data `json:"data" valid:"required field data missing"`
}

var _ mappables.Meta = (*meta)(nil)

func (meta meta) GetData() data.Data { return meta.Data }
func (meta meta) GetID() ids.ID      { return meta.ID }
func (meta meta) GetKey() helpers.Key {
	return key.FromID(meta.GetID())
}
func (meta) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, module.Name, meta{})
}

func NewMeta(data data.Data) mappables.Meta {
	return meta{
		ID:   key.GenerateMetaID(data),
		Data: data,
	}
}
