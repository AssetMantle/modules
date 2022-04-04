// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"strconv"

	"github.com/AssetMantle/modules/schema/mappables/qualified"
	qualified2 "github.com/AssetMantle/modules/schema/qualified/base"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/ids"
	"github.com/AssetMantle/modules/constants/properties"
	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type order struct {
	qualified.Document //nolint:govet
}

var _ mappables.Order = (*order)(nil)

// TODO use get property
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
func (order order) GetCreation() types.MetaProperty {
	heightValue, Error := strconv.ParseInt(key.ReadCreationID(order.ID).String(), 10, 64)
	if Error != nil {
		return base.NewMetaProperty(ids.CreationProperty, base.NewHeightData(base.NewHeight(0)))
	}

	return base.NewMetaProperty(ids.CreationProperty, base.NewHeightData(base.NewHeight(heightValue)))
}
func (order order) GetExchangeRate() types.MetaProperty {
	decValue, Error := sdkTypes.NewDecFromStr(key.ReadRateID(order.ID).String())
	if Error != nil {
		return base.NewMetaProperty(ids.ExchangeRateProperty, base.NewDecData(sdkTypes.ZeroDec()))
	}

	return base.NewMetaProperty(ids.ExchangeRateProperty, base.NewDecData(decValue))
}
func (order order) GetTakerID() types.Property {
	if takerID := order.HasImmutables.GetImmutableProperties().Get(ids.TakerIDProperty); takerID != nil {
		return takerID
	} else if takerID := order.HasMutables.GetMutableProperties().Get(ids.TakerIDProperty); takerID != nil {
		return takerID
	} else {

		return properties.TakerID
	}
}
func (order order) GetExpiry() types.Property {
	if expiry := order.GetProperty(ids.ExpiryProperty); expiry != nil {
		return expiry
	}

	return properties.Expiry
}
func (order order) GetMakerOwnableSplit() types.Property {
	if makerOwnableSplit := order.GetProperty(ids.MakerOwnableSplitProperty); makerOwnableSplit != nil {
		return makerOwnableSplit
	}

	return properties.MakerOwnableSplit
}
func (order order) GetKey() helpers.Key {
	return key.FromID(order.ID)
}
func (order) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, order{})
}

func NewOrder(orderID types.ID, immutableProperties types.Properties, mutableProperties types.Properties) mappables.Order {
	return order{
		Document: qualified.Document{
			ID:               orderID,
			ClassificationID: key.ReadClassificationID(orderID),
			HasImmutables:    qualified2.HasImmutables{Properties: immutableProperties},
			HasMutables:      qualified2.HasMutables{Properties: mutableProperties},
		},
	}
}
