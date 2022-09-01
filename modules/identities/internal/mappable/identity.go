// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/types"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type identity struct {
	qualified.Document
}

var _ mappables.Identity = (*identity)(nil)

func (identity identity) GetExpiry() types.Height {
	if property := identity.Document.GetProperty(constants.ExpiryHeightProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().(data.HeightData).Get()
	}

	return constants.ExpiryHeightProperty.GetData().(data.HeightData).Get()
}
func (identity identity) GetAuthentication() lists.DataList {
	if property := identity.Document.GetProperty(constants.AuthenticationProperty.GetID()); property != nil && property.IsMeta() {
		return base.NewDataList(property.(properties.MetaProperty).GetData().(data.ListData).Get()...)
	}

	return base.NewDataList(constants.AuthenticationProperty.GetData().(data.ListData).Get()...)
}
func (identity identity) IsProvisioned(accAddress sdkTypes.AccAddress) bool {
	_, isProvisioned := identity.GetAuthentication().Search(baseData.NewAccAddressData(accAddress))
	return isProvisioned
}
func (identity identity) ProvisionAddress(accAddresses ...sdkTypes.AccAddress) mappables.Identity {
	identity.Document = identity.Document.Mutate(baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), baseData.NewListData(identity.GetAuthentication().Add(accAddressesToData(accAddresses...)...))))
	return identity
}
func (identity identity) UnprovisionAddress(accAddresses ...sdkTypes.AccAddress) mappables.Identity {
	identity.Document = identity.Document.Mutate(baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), baseData.NewListData(identity.GetAuthentication().Remove(accAddressesToData(accAddresses...)...))))
	return identity
}
func (identity identity) GetKey() helpers.Key {
	return key.NewKey(baseIDs.NewIdentityID(identity.Document.GetClassificationID(), identity.Document.GetImmutables()))
}
func accAddressesToData(accAddresses ...sdkTypes.AccAddress) []data.Data {
	accAddressData := make([]data.Data, len(accAddresses))
	for i, accAddress := range accAddresses {
		accAddressData[i] = baseData.NewAccAddressData(accAddress)
	}
	return accAddressData
}
func (identity) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, identity{})
}

func NewIdentity(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) mappables.Identity {
	return identity{Document: baseQualified.NewDocument(classificationID, immutables, mutables)}
}

func Prototype() helpers.Mappable {
	return identity{}
}
