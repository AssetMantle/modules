// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

//type mappable struct {
//	data.Data
//}
//
//var _ helpers.Mappable = (*Mappable)(nil)
//
//func (mappable Mappable) GetKey() helpers.Key {
//	return key.NewKey(base.GenerateDataID(mappable.Impl.(*Mappable_Data).Data))
//}
//func (Mappable) RegisterCodec(codec *codec.LegacyAmino) {
//	schema.RegisterModuleConcrete(codec, Mappable{})
//}
//
//func (m *Mappable) RegisterInterfaces(registry types.InterfaceRegistry) {
//	registry.RegisterInterface("mappable", (*helpers.Mappable)(nil), &Mappable{})
//}

func NewMappable(data data.Data) helpers.Mappable {
	return baseHelpers.NewDataMappable(data)
}

func Prototype() helpers.Mappable {
	return baseHelpers.DataMappablePrototype()
}
