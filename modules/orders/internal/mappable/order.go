/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type order struct {
	ID types.ID `json:"id" valid:"required~required field key missing"`
	traits.HasMutables
	traits.HasImmutables
}

var _ mappables.Order = (*order)(nil)

func (order order) GetID() types.ID {
	return order.ID
}
func (order order) GetClassificationID() types.ID {
	return key.ReadClassificationID(order.ID)
}
func (order order) GetRateID() types.ID {
	return key.ReadRateID(order.ID)
}
func (order order) GetCreationID() types.ID {
	return key.ReadCreationID(order.ID)
}
func (order order) GetMakerOwnableID() types.ID {
	return key.ReadMakerOwnableID(order.ID)
}
func (order order) GetTakerOwnableID() types.ID {
	return key.ReadTakerOwnableID(order.ID)
}
func (order order) GetMakerID() types.ID {
	return key.ReadMakerID(order.ID)
}
func (order order) GetCreation() types.Property {
	if creationID := order.HasImmutables.GetImmutableProperties().Get(base.NewID(properties.Creation)); creationID != nil {
		return creationID
	} else if creationID := order.HasMutables.GetMutableProperties().Get(base.NewID(properties.Creation)); creationID != nil {
		return creationID
	} else {
		data, _ := base.ReadHeightData("")
		return base.NewProperty(base.NewID(properties.Creation), base.NewFact(data))
	}
}
func (order order) GetExchangeRate() types.Property {
	if rateID := order.HasImmutables.GetImmutableProperties().Get(base.NewID(properties.ExchangeRate)); rateID != nil {
		return rateID
	} else if rateID := order.HasMutables.GetMutableProperties().Get(base.NewID(properties.ExchangeRate)); rateID != nil {
		return rateID
	} else {
		data, _ := base.ReadIDData("")
		return base.NewProperty(base.NewID(properties.ExchangeRate), base.NewFact(data))
	}
}
func (order order) GetTakerID() types.Property {
	if takerID := order.HasImmutables.GetImmutableProperties().Get(base.NewID(properties.TakerID)); takerID != nil {
		return takerID
	} else if takerID := order.HasMutables.GetMutableProperties().Get(base.NewID(properties.TakerID)); takerID != nil {
		return takerID
	} else {
		data, _ := base.ReadIDData("")
		return base.NewProperty(base.NewID(properties.TakerID), base.NewFact(data))
	}
}
func (order order) GetExpiry() types.Property {
	if property := order.HasImmutables.GetImmutableProperties().Get(base.NewID(properties.Expiry)); property != nil {
		return property
	} else if property := order.HasMutables.GetMutableProperties().Get(base.NewID(properties.Expiry)); property != nil {
		return property
	} else {
		data, _ := base.ReadHeightData("")

		return base.NewProperty(base.NewID(properties.Expiry), base.NewFact(data))
	}
}
func (order order) GetMakerOwnableSplit() types.Property {
	if split := order.HasImmutables.GetImmutableProperties().Get(base.NewID(properties.MakerOwnableSplit)); split != nil {
		return split
	} else if split := order.HasMutables.GetMutableProperties().Get(base.NewID(properties.MakerOwnableSplit)); split != nil {
		return split
	} else {
		data, _ := base.ReadDecData("")
		return base.NewProperty(base.NewID(properties.MakerOwnableSplit), base.NewFact(data))
	}
}
func (order order) GetKey() helpers.Key {
	return key.FromID(order.ID)
}
func (order) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, order{})
}

func NewOrder(orderID types.ID, immutableProperties types.Properties, mutableProperties types.Properties) mappables.Order {
	return order{
		ID:            orderID,
		HasImmutables: baseTraits.HasImmutables{Properties: immutableProperties},
		HasMutables:   baseTraits.HasMutables{Properties: mutableProperties},
	}
}
