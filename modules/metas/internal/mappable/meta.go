// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/metas/internal/key"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type meta struct {
	data.Data
}

var _ mappables.Meta = (*meta)(nil)

func (meta meta) GetData() data.Data { return meta.Data }
func (meta meta) GetKey() helpers.Key {
	return key.NewKey(base.NewMetaID(meta.Data.GetType(), meta.GetData().GenerateHashID()))
}
func (meta) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, meta{})
}

func NewMeta(data data.Data) mappables.Meta {
	return meta{
		Data: data,
	}
}

func Prototype() helpers.Mappable {
	return meta{}
}
