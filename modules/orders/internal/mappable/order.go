// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/orders/internal/key"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type order struct {
	qualified.Document
}

var _ mappables.Order = (*order)(nil)

func (order order) GetRateID() ids.ID {
	return key.ReadRateID(order.ID)
}
func (order order) GetCreationID() ids.ID {
	return key.ReadCreationID(order.ID)
}
func (order order) GetMakerOwnableID() ids.ID {
	return key.ReadMakerOwnableID(order.ID)
}
func (order order) GetTakerOwnableID() ids.ID {
	return key.ReadTakerOwnableID(order.ID)
}
func (order order) GetMakerID() ids.ID {
	return key.ReadMakerID(order.ID)
}
func (order order) GetCreation() properties.MetaProperty {
	heightValue, err := strconv.ParseInt(key.ReadCreationID(order.ID).String(), 10, 64)
	if err != nil {
		return base.NewMetaProperty(constants.CreationProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(0)))
	}

	return base.NewMetaProperty(constants.CreationProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(heightValue)))
}
func (order order) GetExchangeRate() properties.MetaProperty {
	decValue, err := sdkTypes.NewDecFromStr(key.ReadRateID(order.ID).String())
	if err != nil {
		return base.NewMetaProperty(constants.ExchangeRateProperty.GetKey(), baseData.NewDecData(sdkTypes.ZeroDec()))
	}

	return base.NewMetaProperty(constants.ExchangeRateProperty.GetKey(), baseData.NewDecData(decValue))
}
func (order order) GetTakerID() properties.Property {
	if takerID := order.Document.GetImmutables().GetImmutablePropertyList().GetProperty(constants.TakerIDProperty); takerID != nil {
		return takerID
	} else if takerID := order.Document.GetMutables().GetMutablePropertyList().GetProperty(constants.TakerIDProperty); takerID != nil {
		return takerID
	} else {

		return constants.TakerID
	}
}
func (order order) GetExpiry() properties.Property {
	if expiry := order.GetProperty(constants.ExpiryProperty); expiry != nil {
		return expiry
	}

	return constants.Expiry
}
func (order order) GetMakerOwnableSplit() properties.Property {
	if makerOwnableSplit := order.GetProperty(constants.MakerOwnableSplitProperty); makerOwnableSplit != nil {
		return makerOwnableSplit
	}

	return constants.MakerOwnableSplit
}
func (order order) GetKey() helpers.Key {
	return key.NewKey(baseIDs.NewOrderID(order.GetClassificationID(), order.GetMakerOwnableID(), order.GetTakerOwnableID(), order.GetRateID(), order.GetCreationID(), order.GetMakerID(), order.GetImmutables()))
}
func (order) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, order{})
}

func NewOrder(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) mappables.Order {
	return order{Document: baseQualified.NewDocument(classificationID, immutables, mutables)}
}

func Prototype() helpers.Mappable {
	return order{}
}
