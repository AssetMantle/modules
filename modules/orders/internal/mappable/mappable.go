// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/documents/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

var _ helpers.Mappable = (*Mappable)(nil)

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
