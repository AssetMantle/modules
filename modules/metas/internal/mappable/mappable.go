// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/AssetMantle/schema/x/data"
	baseData "github.com/AssetMantle/schema/x/data/base"
	"github.com/AssetMantle/schema/x/helpers"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/metas/internal/key"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

var _ helpers.Mappable = (*Mappable)(nil)

func (mappable *Mappable) ValidateBasic() error {
	return mappable.Data.ValidateBasic()
}
func (mappable *Mappable) GetKey() helpers.Key {
	return key.NewKey(baseIDs.GenerateDataID(mappable.Data))
}
func (*Mappable) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, Mappable{})
}

func NewMappable(data data.Data) helpers.Mappable {
	return &Mappable{Data: data.ToAnyData().(*baseData.AnyData)}
}

func Prototype() helpers.Mappable {
	return &Mappable{}
}

func GetData(mappable helpers.Mappable) data.Data {
	return mappable.(*Mappable).Data
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
