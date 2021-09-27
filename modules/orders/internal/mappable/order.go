/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables" //nolint:typecheck
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	baseTypes "github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
	"strconv"
)

var _ mappables.Order = (*Order)(nil)

func (order Order) GetStructReference() codec.ProtoMarshaler {
	return &order
}
func (order Order) GetID() types.ID {
	return &order.ID
}
func (order Order) GetClassificationID() types.ID {
	return key.ReadClassificationID(&order.ID)
}
func (order Order) GetImmutableProperties() types.Properties {
	return order.HasImmutables.GetImmutableProperties()
}

func (order Order) GenerateHashID() types.ID {
	return order.HasImmutables.GenerateHashID()
}

func (order Order) GetMutableProperties() types.Properties {
	return order.HasMutables.GetMutableProperties()
}

func (order Order) Mutate(propertyList ...types.Property) traits.HasMutables {
	return order.HasMutables.Mutate(propertyList...)
}
func (order Order) GetRateID() types.ID {
	return key.ReadRateID(&order.ID)
}
func (order Order) GetCreationID() types.ID {
	return key.ReadCreationID(&order.ID)
}
func (order Order) GetMakerOwnableID() types.ID {
	return key.ReadMakerOwnableID(&order.ID)
}
func (order Order) GetTakerOwnableID() types.ID {
	return key.ReadTakerOwnableID(&order.ID)
}
func (order Order) GetMakerID() types.ID {
	return key.ReadMakerID(&order.ID)
}
func (order Order) GetCreation() types.MetaProperty {
	heightValue, Error := strconv.ParseInt(key.ReadCreationID(&order.ID).String(), 10, 64)
	if Error != nil {
		return base.NewMetaProperty(base.NewID(properties.Creation), base.NewMetaFact(base.NewHeightData(base.NewHeight(0))))
	}

	return base.NewMetaProperty(base.NewID(properties.Creation), base.NewMetaFact(base.NewHeightData(base.NewHeight(heightValue))))
}
func (order Order) GetExchangeRate() types.MetaProperty {
	decValue, Error := sdkTypes.NewDecFromStr(key.ReadRateID(&order.ID).String())
	if Error != nil {
		return base.NewMetaProperty(base.NewID(properties.ExchangeRate), base.NewMetaFact(base.NewDecData(sdkTypes.ZeroDec())))
	}

	return base.NewMetaProperty(base.NewID(properties.ExchangeRate), base.NewMetaFact(base.NewDecData(decValue)))
}
func (order Order) GetTakerID() types.Property {
	if takerID := order.HasImmutables.GetImmutableProperties().Get(base.NewID(properties.TakerID)); takerID != nil {
		return takerID
	} else if takerID := order.HasMutables.GetMutableProperties().Get(base.NewID(properties.TakerID)); takerID != nil {
		return takerID
	} else {
		data, _ := base.ReadIDData("")
		return base.NewProperty(base.NewID(properties.TakerID), base.NewFact(data))
	}
}
func (order Order) GetExpiry() types.Property {
	if property := order.HasImmutables.GetImmutableProperties().Get(base.NewID(properties.Expiry)); property != nil {
		return property
	} else if property := order.HasMutables.GetMutableProperties().Get(base.NewID(properties.Expiry)); property != nil {
		return property
	} else {
		return base.NewProperty(base.NewID(properties.Expiry), base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	}
}
func (order Order) GetMakerOwnableSplit() types.Property {
	if split := order.HasImmutables.GetImmutableProperties().Get(base.NewID(properties.MakerOwnableSplit)); split != nil {
		return split
	} else if split := order.HasMutables.GetMutableProperties().Get(base.NewID(properties.MakerOwnableSplit)); split != nil {
		return split
	} else {
		data, _ := base.ReadDecData("")
		return base.NewProperty(base.NewID(properties.MakerOwnableSplit), base.NewFact(data))
	}
}
func (order Order) GetKey() helpers.Key {
	return key.FromID(&order.ID)
}
func (Order) RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, module.Name, Order{})
}

func NewOrder(orderID types.ID, immutableProperties types.Properties, mutableProperties types.Properties) mappables.Order {
	return &Order{
		ID:            *base.NewID(orderID.String()),
		HasImmutables: baseTraits.HasImmutables{Properties: *baseTypes.NewProperties(immutableProperties.GetList()...)},
		HasMutables:   baseTraits.HasMutables{Properties: *baseTypes.NewProperties(mutableProperties.GetList()...)},
	}
}
