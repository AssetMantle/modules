package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types/schema"
)

type identity struct {
	ID                       schema.ID
	ProvisionedAddressList   []sdkTypes.AccAddress
	UnprovisionedAddressList []sdkTypes.AccAddress
	Immutables               schema.Immutables
	Mutables                 schema.Mutables
}

var _ schema.InterIdentity = (*identity)(nil)

func (identity identity) GetID() schema.ID { return identity.ID }
func (identity identity) GetChainID() schema.ID {
	return identityIDFromInterface(identity.ID).ChainID
}

func (identity identity) GetClassificationID() schema.ID {
	return identityIDFromInterface(identity.ID).ClassificationID
}
func (identity identity) GetProvisionedAddressList() []sdkTypes.AccAddress {
	return identity.ProvisionedAddressList
}
func (identity identity) GetUnprovisionedAddressList() []sdkTypes.AccAddress {
	return identity.UnprovisionedAddressList
}
func (identity identity) ProvisionAddress(accAddress sdkTypes.AccAddress) schema.InterIdentity {
	identity.ProvisionedAddressList = append(identity.ProvisionedAddressList, accAddress)
	return identity
}
func (identity identity) UnprovisionAddress(accAddress sdkTypes.AccAddress) schema.InterIdentity {
	for i, provisionedAddress := range identity.ProvisionedAddressList {
		if provisionedAddress.Equals(accAddress) {
			identity.ProvisionedAddressList = append(identity.ProvisionedAddressList[:i], identity.ProvisionedAddressList[i+1:]...)
			identity.UnprovisionedAddressList = append(identity.UnprovisionedAddressList, accAddress)
			return identity
		}
	}
	return identity
}
func (identity identity) GetImmutables() schema.Immutables { return identity.Immutables }
func (identity identity) GetMutables() schema.Mutables     { return identity.Mutables }
func (identity identity) IsProvisioned(accAddress sdkTypes.AccAddress) bool {
	for _, provisionedAddress := range identity.ProvisionedAddressList {
		if provisionedAddress.Equals(accAddress) {
			return true
		}
	}
	return false
}
func (identity identity) IsUnprovisioned(accAddress sdkTypes.AccAddress) bool {
	for _, unprovisionedAddress := range identity.UnprovisionedAddressList {
		if unprovisionedAddress.Equals(accAddress) {
			return true
		}
	}
	return false
}
func NewIdentity(identityID schema.ID, provisionedAddressList []sdkTypes.AccAddress, unprovisionedAddressList []sdkTypes.AccAddress, immutables schema.Immutables, mutables schema.Mutables) schema.InterIdentity {
	return identity{
		ID:                       identityID,
		ProvisionedAddressList:   provisionedAddressList,
		UnprovisionedAddressList: unprovisionedAddressList,
		Immutables:               immutables,
		Mutables:                 mutables,
	}
}
