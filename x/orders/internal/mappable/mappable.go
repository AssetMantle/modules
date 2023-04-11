// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/AssetMantle/schema/x/documents"
	"github.com/AssetMantle/schema/x/documents/base"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/helpers"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
	"github.com/AssetMantle/modules/x/orders/internal/key"
)

var _ helpers.Mappable = (*Mappable)(nil)

func (mappable *Mappable) ValidateBasic() error {
	return mappable.Order.ValidateBasic()
}
func (mappable *Mappable) GetKey() helpers.Key {
	return key.NewKey(baseIDs.NewOrderID(mappable.Order.GetClassificationID(), mappable.Order.GetImmutables()))
}
func (*Mappable) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, Mappable{})
}

func NewMappable(order documents.Order) helpers.Mappable {
	return &Mappable{Order: order.Get().(*base.Document)}
}

func Prototype() helpers.Mappable {
	return &Mappable{}
}

func GetOrder(mappable helpers.Mappable) documents.Order {
	return base.NewOrderFromDocument(mappable.(*Mappable).Order)
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
