// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/metas/internal/key"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type mappable struct {
	mappables.Meta
}

var _ helpers.Mappable = (*mappable)(nil)

func (mappable mappable) GetKey() helpers.Key {
	return key.NewKey(base.NewMetaID(mappable.Meta.GetData().GetType(), mappable.GetData().GenerateHashID()))
}
func (mappable) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, mappable{})
}

func NewMappable(meta mappables.Meta) helpers.Mappable {
	return mappable{Meta: meta}
}

func Prototype() helpers.Mappable {
	return mappable{}
}
