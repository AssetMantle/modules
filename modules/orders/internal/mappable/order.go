// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/modules/orders/internal/module"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/mappables"
	properties2 "github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type order struct {
	baseQualified.Document //nolint:govet
}

var _ mappables.Order = (*order)(nil)

// TODO use get property
func (order order) GetRateID() ids2.ID {
	return key.ReadRateID(order.ID)
}
func (order order) GetCreationID() ids2.ID {
	return key.ReadCreationID(order.ID)
}
func (order order) GetMakerOwnableID() ids2.ID {
	return key.ReadMakerOwnableID(order.ID)
}
func (order order) GetTakerOwnableID() ids2.ID {
	return key.ReadTakerOwnableID(order.ID)
}
func (order order) GetMakerID() ids2.ID {
	return key.ReadMakerID(order.ID)
}
func (order order) GetCreation() properties2.MetaProperty {
	heightValue, Error := strconv.ParseInt(key.ReadCreationID(order.ID).String(), 10, 64)
	if Error != nil {
		return base.NewMetaProperty(constants.CreationProperty, baseData.NewHeightData(baseTypes.NewHeight(0)))
	}

	return base.NewMetaProperty(constants.CreationProperty, baseData.NewHeightData(baseTypes.NewHeight(heightValue)))
}
func (order order) GetExchangeRate() properties2.MetaProperty {
	decValue, Error := sdkTypes.NewDecFromStr(key.ReadRateID(order.ID).String())
	if Error != nil {
		return base.NewMetaProperty(constants.ExchangeRateProperty, baseData.NewDecData(sdkTypes.ZeroDec()))
	}

	return base.NewMetaProperty(constants.ExchangeRateProperty, baseData.NewDecData(decValue))
}
func (order order) GetTakerID() properties2.Property {
	if takerID := order.Immutables.GetImmutablePropertyList().GetProperty(constants.TakerIDProperty); takerID != nil {
		return takerID
	} else if takerID := order.Mutables.GetMutablePropertyList().GetProperty(constants.TakerIDProperty); takerID != nil {
		return takerID
	} else {

		return constants.TakerID
	}
}
func (order order) GetExpiry() properties2.Property {
	if expiry := order.GetProperty(constants.ExpiryProperty); expiry != nil {
		return expiry
	}

	return constants.Expiry
}
func (order order) GetMakerOwnableSplit() properties2.Property {
	if makerOwnableSplit := order.GetProperty(constants.MakerOwnableSplitProperty); makerOwnableSplit != nil {
		return makerOwnableSplit
	}

	return constants.MakerOwnableSplit
}
func (order order) GetKey() helpers.Key {
	return key.FromID(order.ID)
}
func (order) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, module.Name, order{})
}

func NewOrder(orderID ids2.ID, immutableProperties lists.PropertyList, mutableProperties lists.PropertyList) mappables.Order {
	return order{
		Document: baseQualified.Document{
			ID:               orderID,
			ClassificationID: key.ReadClassificationID(orderID),
			Immutables:       baseQualified.Immutables{PropertyList: immutableProperties},
			Mutables:         baseQualified.Mutables{Properties: mutableProperties},
		},
	}
}
