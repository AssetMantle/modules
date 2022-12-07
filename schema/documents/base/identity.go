package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/types"
)

//type identity struct {
//	documents.Document
//}

var _ documents.Identity = (*Identity)(nil)

func (identity *Identity) GenerateHashID() ids.HashID {
	return identity.Document.GenerateHashID()
}

func (identity *Identity) GetClassificationID() ids.ClassificationID {
	return identity.Document.GetClassificationID()
}

func (identity *Identity) GetProperty(id ids.PropertyID) properties.Property {
	return identity.Document.GetProperty(id)
}

func (identity *Identity) GetImmutables() qualified.Immutables {
	return identity.Document.GetImmutables()
}

func (identity *Identity) GetMutables() qualified.Mutables {
	return identity.Document.GetMutables()
}

func (identity *Identity) Mutate(property ...properties.Property) documents.Document {
	return identity.Document.Mutate(property...)
}

func (identity *Identity) GetExpiry() types.Height {
	if property := identity.Document.GetProperty(constants.ExpiryHeightProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().(data.HeightData).Get()
	}

	return constants.ExpiryHeightProperty.GetData().(data.HeightData).Get()
}
func (identity *Identity) GetAuthentication() lists.DataList {
	if property := identity.Document.GetProperty(constants.AuthenticationProperty.GetID()); property != nil && property.IsMeta() {
		return base.NewDataList(property.(properties.MetaProperty).GetData().(data.ListData).Get()...)
	}

	return base.NewDataList(constants.AuthenticationProperty.GetData().(data.ListData).Get()...)
}
func (identity *Identity) IsProvisioned(accAddress sdkTypes.AccAddress) bool {
	_, isProvisioned := identity.GetAuthentication().Search(baseData.NewAccAddressData(accAddress))
	return isProvisioned
}
func (identity *Identity) ProvisionAddress(accAddresses ...sdkTypes.AccAddress) documents.Identity {
	identity.Document = identity.Document.Mutate(baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), baseData.NewListData(identity.GetAuthentication().Add(accAddressesToData(accAddresses...)...))))
	return identity
}
func (identity *Identity) UnprovisionAddress(accAddresses ...sdkTypes.AccAddress) documents.Identity {
	identity.Document = identity.Document.Mutate(baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), baseData.NewListData(identity.GetAuthentication().Remove(accAddressesToData(accAddresses...)...))))
	return identity
}
func accAddressesToData(accAddresses ...sdkTypes.AccAddress) []data.Data {
	accAddressData := make([]data.Data, len(accAddresses))
	for i, accAddress := range accAddresses {
		accAddressData[i] = baseData.NewAccAddressData(accAddress)
	}
	return accAddressData
}

func NewIdentity(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) documents.Identity {
	return &Identity{
		Document: NewDocument(classificationID, immutables, mutables),
	}
}
