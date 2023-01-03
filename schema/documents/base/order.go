package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/types"
)

type order struct {
	documents.Document
}

var _ documents.Order = (*order)(nil)

func (order order) GetExchangeRate() sdkTypes.Dec {
	if property := order.GetProperty(constants.ExchangeRateProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.DecData).Get()
	}
	return constants.ExchangeRateProperty.GetData().Get().(data.DecData).Get()
}
func (order order) GetCreationHeight() types.Height {
	if property := order.GetProperty(constants.CreationHeightProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.HeightData).Get()
	}
	return constants.CreationHeightProperty.GetData().Get().(data.HeightData).Get()
}
func (order order) GetMakerOwnableID() ids.OwnableID {
	if property := order.GetProperty(constants.MakerOwnableIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.IDData).Get().(ids.OwnableID)
	}
	return constants.MakerOwnableIDProperty.GetData().Get().(data.IDData).Get().(ids.OwnableID)
}
func (order order) GetTakerOwnableID() ids.OwnableID {
	if property := order.GetProperty(constants.TakerOwnableIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.IDData).Get().(ids.OwnableID)
	}
	return constants.TakerOwnableIDProperty.GetData().Get().(data.IDData).Get().(ids.OwnableID)
}
func (order order) GetMakerID() ids.IdentityID {
	if property := order.GetProperty(constants.MakerIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.IDData).Get().(ids.IdentityID)
	}
	return constants.MakerIDProperty.GetData().Get().(data.IDData).Get().(ids.IdentityID)
}
func (order order) GetTakerID() ids.IdentityID {
	if property := order.GetProperty(constants.TakerIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.IDData).Get().(ids.IdentityID)
	}
	return constants.TakerIDProperty.GetData().Get().(data.IDData).Get().(ids.IdentityID)
}
func (order order) GetExpiryHeight() types.Height {
	if property := order.GetProperty(constants.ExpiryHeightProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.HeightData).Get()
	}
	return constants.ExpiryHeightProperty.GetData().Get().(data.HeightData).Get()
}
func (order order) GetMakerOwnableSplit() sdkTypes.Dec {
	if property := order.GetProperty(constants.MakerOwnableSplitProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.DecData).Get()
	}
	return constants.MakerOwnableSplitProperty.GetData().Get().(data.DecData).Get()
}

func NewOrder(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) documents.Order {
	return order{Document: NewDocument(classificationID, immutables, mutables)}
}

func NewOrderFromDocument(document documents.Document) documents.Order {
	return order{Document: document}
}
