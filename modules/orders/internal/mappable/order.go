// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/types"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type order struct {
	qualified.Document
}

var _ mappables.Order = (*order)(nil)

func (order order) GetExchangeRate() sdkTypes.Dec {
	if property := order.GetProperty(constants.ExchangeRatePropertyID); property != nil && property.IsMeta() && property.GetType().Compare(dataConstants.DecDataID) == 0 {
		return property.(properties.MetaProperty).GetData().(data.DecData).Get()
	}
	return constants.ExchangeRateProperty.GetData().(data.DecData).Get()
}
func (order order) GetCreationHeight() types.Height {
	// TODO compare type with propertyID's type
	if property := order.GetProperty(constants.CreationHeightPropertyID); property != nil && property.IsMeta() && property.GetType().Compare(dataConstants.HeightDataID) == 0 {
		return property.(properties.MetaProperty).GetData().(data.HeightData).Get()
	}
	return constants.CreationHeightProperty.GetData().(data.HeightData).Get()
}
func (order order) GetMakerOwnableID() ids.OwnableID {
	if property := order.GetProperty(constants.MakerOwnableIDPropertyID); property != nil && property.IsMeta() && property.GetType().Compare(constants.MakerOwnableIDPropertyID.GetType()) == 0 {
		return property.(properties.MetaProperty).GetData().(data.IDData).Get().(ids.OwnableID)
	}
	return constants.MakerOwnableIDProperty.GetData().(data.IDData).Get().(ids.OwnableID)
}
func (order order) GetTakerOwnableID() ids.OwnableID {
	if property := order.GetProperty(constants.TakerOwnableIDPropertyID); property != nil && property.IsMeta() && property.GetType().Compare(constants.TakerOwnableIDPropertyID.GetType()) == 0 {
		return property.(properties.MetaProperty).GetData().(data.IDData).Get().(ids.OwnableID)
	}
	return constants.TakerOwnableIDProperty.GetData().(data.IDData).Get().(ids.OwnableID)
}
func (order order) GetMakerID() ids.IdentityID {
	if property := order.GetProperty(constants.MakerIDPropertyID); property != nil && property.IsMeta() && property.GetType().Compare(constants.MakerIDPropertyID.GetType()) == 0 {
		return property.(properties.MetaProperty).GetData().(data.IDData).Get().(ids.IdentityID)
	}
	return constants.MakerIDProperty.GetData().(data.IDData).Get().(ids.IdentityID)
}
func (order order) GetTakerID() ids.IdentityID {
	if property := order.GetProperty(constants.TakerIDPropertyID); property != nil && property.IsMeta() && property.GetType().Compare(constants.TakerIDPropertyID.GetType()) == 0 {
		return property.(properties.MetaProperty).GetData().(data.IDData).Get().(ids.IdentityID)
	}
	return constants.TakerIDProperty.GetData().(data.IDData).Get().(ids.IdentityID)
}
func (order order) GetExpiryHeight() types.Height {
	if property := order.GetProperty(constants.ExpiryHeightPropertyID); property != nil && property.IsMeta() && property.GetType().Compare(constants.ExpiryHeightPropertyID.GetType()) == 0 {
		return property.(properties.MetaProperty).GetData().(data.HeightData).Get()
	}
	return constants.ExpiryHeightProperty.GetData().(data.HeightData).Get()
}
func (order order) GetMakerOwnableSplit() sdkTypes.Dec {
	if property := order.GetProperty(constants.MakerOwnableSplitPropertyID); property != nil && property.IsMeta() && property.GetType().Compare(constants.MakerOwnableSplitPropertyID.GetType()) == 0 {
		return property.(properties.MetaProperty).GetData().(data.DecData).Get()
	}
	return constants.MakerOwnableSplitProperty.GetData().(data.DecData).Get()
}
func (order order) GetKey() helpers.Key {
	return key.NewKey(baseIDs.NewOrderID(order.GetClassificationID(), order.GetMakerOwnableID(), order.GetTakerOwnableID(), order.GetExchangeRate(), order.GetCreationHeight(), order.GetMakerID(), order.GetImmutables()))
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
