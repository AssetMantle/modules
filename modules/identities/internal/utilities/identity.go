package utilities

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/mappables"
)

// TODO Implement
func IsProvisioned(identity mappables.Identity, accAddress sdkTypes.AccAddress) bool {
	// _, found := identity.GetAuthentication().GetDataID().Search(baseData.NewAccAddressData(accAddress).GetID())
	// return found
	panic("implement")
}

// TODO Implement
func ProvisionAddress(identity mappables.Identity, accAddress sdkTypes.AccAddress) mappables.Identity {
	// identity.Document = identity.Document.Mutate(
	// 	baseProperties.NewPropertyWithDataID(
	// 		identity.GetAuthentication().GetID(),
	// 		base.NewListDataID(identity.GetAuthentication().GetDataID().Add(baseData.NewAccAddressData(accAddress).GetID())))).(baseQualified.Document)
	// return identity
	panic("implement")
}

// TODO Implement
func UnprovisionAddress(identity mappables.Identity, accAddress sdkTypes.AccAddress) mappables.Identity {
	// identity.Document = identity.Document.Mutate(
	// 	baseProperties.NewPropertyWithDataID(
	// 		identity.GetAuthentication().GetID(),
	// 		base.NewListDataID(identity.GetAuthentication().GetDataID().Remove(baseData.NewAccAddressData(accAddress).GetID())))).(baseQualified.Document)
	// return identity
	panic("implement")
}
