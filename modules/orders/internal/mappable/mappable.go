// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type mappable struct {
	documents.Order
}

var _ helpers.Mappable = (*mappable)(nil)

func (mappable mappable) GetKey() helpers.Key {
	return key.NewKey(baseIDs.NewOrderID(mappable.GetClassificationID(), mappable.GetMakerOwnableID(), mappable.GetTakerOwnableID(), mappable.GetExchangeRate(), mappable.GetCreationHeight(), mappable.GetMakerID(), mappable.GetImmutables()))
}
func (mappable) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, mappable{})
}

func NewMappable(order documents.Order) helpers.Mappable {
	return mappable{Order: order}
}

func Prototype() helpers.Mappable {
	return mappable{}
}
