// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/schema/documents"
	"github.com/AssetMantle/schema/documents/base"
)

var _ helpers.Mappable = (*Mappable)(nil)

func (mappable *Mappable) ValidateBasic() error {
	return mappable.Order.ValidateBasic()
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
