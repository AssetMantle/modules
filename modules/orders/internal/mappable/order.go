/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type order struct {
	ID         types.ID         `json:"key" valid:"required~required field key missing"`
	Immutables types.Immutables `json:"immutables" valid:"required field immutables missing"`
	Mutables   types.Mutables   `json:"mutables" valid:"required~required field mutables missing"`
}

var _ mappables.Order = (*order)(nil)

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
func (order order) GetTakerID() types.Property {
	if takerID := order.Immutables.Get().Get(base.NewID(properties.TakerID)); takerID != nil {
		return takerID
	} else if takerID := order.Mutables.Get().Get(base.NewID(properties.TakerID)); takerID != nil {
		return takerID
	} else {
		data, _ := base.ReadIDData("")
		return base.NewProperty(base.NewID(properties.TakerID), base.NewFact(data))
	}
}
func (order order) GetCreation() types.Property {
	heightValue, Error := strconv.ParseInt(key.ReadCreationID(order.ID).String(), 10, 64)
	if Error != nil {
		return base.NewMetaProperty(base.NewID(properties.MakerOwnableSplit), base.NewMetaFact(base.NewHeightData(base.NewHeight(0))))
	}

	return base.NewMetaProperty(base.NewID(properties.MakerOwnableSplit), base.NewMetaFact(base.NewHeightData(base.NewHeight(heightValue))))
}
func (order order) GetExpiry() types.Property {
	if expiry := order.Immutables.Get().Get(base.NewID(properties.Expiry)); expiry != nil {
		return expiry
	} else if creation := order.Mutables.Get().Get(base.NewID(properties.Expiry)); creation != nil {
		return creation
	} else {
		data, _ := base.ReadHeightData("")
		return base.NewProperty(base.NewID(properties.Expiry), base.NewFact(data))
	}
}
func (order order) GetMakerOwnableSplit() types.Property {
	if split := order.Immutables.Get().Get(base.NewID(properties.MakerOwnableSplit)); split != nil {
		return split
	} else if split := order.Mutables.Get().Get(base.NewID(properties.MakerOwnableSplit)); split != nil {
		return split
	} else {
		data, _ := base.ReadDecData("")
		return base.NewProperty(base.NewID(properties.MakerOwnableSplit), base.NewFact(data))
	}
}
func (order order) GetExchangeRate() types.Property {
	decValue, Error := sdkTypes.NewDecFromStr(key.ReadRateID(order.ID).String())
	if Error != nil {
		return base.NewMetaProperty(base.NewID(properties.ExchangeRate), base.NewMetaFact(base.NewDecData(sdkTypes.ZeroDec())))
	}

	return base.NewMetaProperty(base.NewID(properties.ExchangeRate), base.NewMetaFact(base.NewDecData(decValue)))
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
func (order order) GetKey() helpers.Key {
	return key.New(order.ID)
}
func (order) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, order{})
}

func NewOrder(orderID types.ID, immutables types.Immutables, mutables types.Mutables) mappables.Order {
	return order{
		ID:         orderID,
		Immutables: immutables,
		Mutables:   mutables,
	}
}
