// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/metas/internal/key"
	"github.com/AssetMantle/modules/modules/metas/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/types"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type meta struct {
	ID   types.ID   `json:"id" valid:"required field id missing"`
	Data types.Data `json:"data" valid:"required field data missing"`
}

var _ mappables.Meta = (*meta)(nil)

func (meta meta) GetData() types.Data { return meta.Data }
func (meta meta) GetID() types.ID     { return meta.ID }
func (meta meta) GetKey() helpers.Key {
	return key.FromID(meta.GetID())
}
func (meta) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, module.Name, meta{})
}

func NewMeta(data types.Data) mappables.Meta {
	return meta{
		ID:   key.GenerateMetaID(data),
		Data: data,
	}
}
