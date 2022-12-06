// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/metas/module/key"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids/base"
)

type mappable struct {
	data.Data
}

var _ helpers.Mappable = (*mappable)(nil)

func (mappable mappable) GetKey() helpers.Key {
	return key.NewKey(base.NewDataID(mappable.Data))
}
func (mappable) RegisterCodec(codec *codec.LegacyAmino) {
	schema.RegisterModuleConcrete(codec, mappable{})
}

func NewMappable(data data.Data) helpers.Mappable {
	return mappable{Data: data}
}

func Prototype() helpers.Mappable {
	return mappable{}
}
