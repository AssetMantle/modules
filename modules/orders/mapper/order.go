/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type order struct {
	ID         types.ID         `json:"id" valid:"required~required field id missing"`
	Immutables types.Immutables `json:"immutables" valid:"required field immutables missing"`
	Mutables   types.Mutables   `json:"mutables" valid:"required~required field mutables missing"`
}

var _ mappables.Order = (*order)(nil)

func (order order) GetClassificationID() types.ID {
	return orderIDFromInterface(order.ID).ClassificationID
}

func (order order) GetMakerOwnableID() types.ID {
	return orderIDFromInterface(order.ID).MakerOwnableID
}

func (order order) GetTakerOwnableID() types.ID {
	return orderIDFromInterface(order.ID).TakerOwnableID
}

func (order order) GetMakerID() types.ID {
	return orderIDFromInterface(order.ID).MakerID
}

func (order order) GetTakerID() types.Property {
	if takerID := order.Immutables.Get().Get(base.NewID(constants.TakerIDProperty)); takerID != nil {
		return takerID
	} else if takerID := order.Mutables.Get().Get(base.NewID(constants.TakerIDProperty)); takerID != nil {
		return takerID
	} else {
		return base.NewProperty(base.NewID(constants.TakerIDProperty), base.NewFact(base.ReadIDData("")))
	}
}

func (order order) GetExchangeRate() types.Property {
	if takerRate := order.Immutables.Get().Get(base.NewID(constants.ExchangeRateProperty)); takerRate != nil {
		return takerRate
	} else if takerRate := order.Mutables.Get().Get(base.NewID(constants.ExchangeRateProperty)); takerRate != nil {
		return takerRate
	} else {
		return base.NewProperty(base.NewID(constants.ExchangeRateProperty), base.NewFact(base.NewDecData(sdkTypes.SmallestDec())))
	}
}

func (order order) GetCreation() types.Property {
	if creation := order.Immutables.Get().Get(base.NewID(constants.CreationProperty)); creation != nil {
		return creation
	} else if creation := order.Mutables.Get().Get(base.NewID(constants.CreationProperty)); creation != nil {
		return creation
	} else {
		return base.NewProperty(base.NewID(constants.CreationProperty), base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	}
}

func (order order) GetExpiry() types.Property {
	if expiry := order.Immutables.Get().Get(base.NewID(constants.ExpiryProperty)); expiry != nil {
		return expiry
	} else if creation := order.Mutables.Get().Get(base.NewID(constants.ExpiryProperty)); creation != nil {
		return creation
	} else {
		return base.NewProperty(base.NewID(constants.ExpiryProperty), base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	}
}

func (order order) GetMakerOwnableSplit() types.Property {
	if split := order.Immutables.Get().Get(base.NewID(constants.MakerOwnableSplitProperty)); split != nil {
		return split
	} else if split := order.Mutables.Get().Get(base.NewID(constants.MakerOwnableSplitProperty)); split != nil {
		return split
	} else {
		return base.NewProperty(base.NewID(constants.MakerOwnableSplitProperty), base.NewFact(base.NewDecData(sdkTypes.SmallestDec())))
	}
}

func (order order) GetImmutables() types.Immutables {
	return order.Immutables
}

func (order order) GetMutables() types.Mutables {
	return order.Mutables
}

func (order order) GetID() types.ID {
	return order.ID
}

func (order order) Encode() []byte {
	return packageCodec.MustMarshalBinaryBare(order)
}

func (order order) Decode(bytes []byte) traits.Mappable {
	packageCodec.MustUnmarshalBinaryBare(bytes, &order)
	return order
}

func orderPrototype() traits.Mappable {
	return order{}
}

func NewOrder(orderID types.ID, immutables types.Immutables, mutables types.Mutables) mappables.Order {
	return order{
		ID:         orderID,
		Immutables: immutables,
		Mutables:   mutables,
	}
}
