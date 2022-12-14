// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/metas/module/key"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids/base"
)

//type mappable struct {
//	data.Data
//}

var _ helpers.Mappable = (*Mappable)(nil)

func (mappable Mappable) GetKey() helpers.Key {
	return key.NewKey(base.GenerateDataID(mappable.Data))
}
func (Mappable) RegisterCodec(codec *codec.LegacyAmino) {
	schema.RegisterModuleConcrete(codec, Mappable{})
}

func NewMappable(data data.Data) helpers.Mappable {
	return &Mappable{Data: data.(*baseData.Data)}
}

func Prototype() helpers.Mappable {
	return &Mappable{}
}
