// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/AssetMantle/modules/helpers"
	codecUtilities "github.com/AssetMantle/schema/codec/utilities"
	"github.com/AssetMantle/schema/documents"
	"github.com/AssetMantle/schema/documents/base"
	"github.com/cosmos/cosmos-sdk/codec"
)

var _ helpers.Mappable = (*Mappable)(nil)

func (mappable *Mappable) ValidateBasic() error {
	return mappable.Asset.ValidateBasic()
}
func (*Mappable) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, Mappable{})
}

func Prototype() helpers.Mappable {
	return &Mappable{}
}

func NewMappable(asset documents.Asset) helpers.Mappable {
	return &Mappable{
		Asset: asset.Get().(*base.Document),
	}
}

func GetAsset(mappable helpers.Mappable) documents.Asset {
	return base.NewAssetFromDocument(mappable.(*Mappable).Asset)
}

func MappablesFromInterface(mappables []helpers.Mappable) []*Mappable {
	Mappables := make([]*Mappable, len(mappables))
	for index, mappable := range mappables {
		Mappables[index] = mappable.(*Mappable)
	}
	return Mappables
}

func MappablesToInterface(mappables []*Mappable) []helpers.Mappable {
	Mappables := make([]helpers.Mappable, len(mappables))
	for index, mappable := range mappables {
		Mappables[index] = mappable
	}
	return Mappables
}
