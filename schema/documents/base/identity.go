package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/types"
)

type identity struct {
	documents.Document
}

var _ documents.Identity = (*identity)(nil)

func (identity identity) GetExpiry() types.Height {
	if property := identity.Document.GetProperty(constants.ExpiryHeightProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().(data.HeightData).Get()
	}

	return constants.ExpiryHeightProperty.GetData().(data.HeightData).Get()
}
func (identity identity) GetAuthentication() lists.AnyDataList {
	var dataList []data.Data

	if property := identity.Document.GetProperty(constants.AuthenticationProperty.GetID()); property != nil && property.IsMeta() {
		for _, anyData := range property.(properties.MetaProperty).GetData().(data.ListData).Get() {
			dataList = append(dataList, anyData.(*baseData.AnyData).Impl.(data.Data))
		}
	} else {
		for _, anyData := range constants.AuthenticationProperty.GetData().(data.ListData).Get() {
			dataList = append(dataList, anyData.(*baseData.AnyData).Impl.(data.Data))
		}
	}
	return baseData.NewDataList(dataList...)
}
func (identity identity) IsProvisioned(accAddress sdkTypes.AccAddress) bool {
	_, isProvisioned := identity.GetAuthentication().Search(baseData.NewAccAddressData(accAddress))
	return isProvisioned
}
func (identity identity) ProvisionAddress(accAddresses ...sdkTypes.AccAddress) documents.Identity {
	var accAddressList []data.AnyData
	for _, address := range accAddressesToData(accAddresses...) {
		accAddressList = append(accAddressList, address.ToAnyData())
	}
	identity.Document = identity.Document.Mutate(baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), baseData.NewListData(identity.GetAuthentication().Add(accAddressList...))))
	return identity
}
func (identity identity) UnprovisionAddress(accAddresses ...sdkTypes.AccAddress) documents.Identity {
	var accAddressList []data.AnyData
	for _, address := range accAddressesToData(accAddresses...) {
		accAddressList = append(accAddressList, address.ToAnyData())
	}
	identity.Document = identity.Document.Mutate(baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), baseData.NewListData(identity.GetAuthentication().Remove(accAddressList...))))
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
	return identity{Document: NewDocument(classificationID, immutables, mutables)}
}
