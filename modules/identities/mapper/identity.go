package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/entities"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type identity struct {
	ID                       types.ID              `json:"id" valid:"required~Enter the ID"`
	ProvisionedAddressList   []sdkTypes.AccAddress `json:"provisionedAddressList" valid:"required~Enter the ProvisionedAddressList"`
	UnprovisionedAddressList []sdkTypes.AccAddress `json:"unprovisionedaddressList" valid:"required~Enter the UnprovisionedAddressList"`
	Immutables               types.Immutables      `json:"immutables" valid:"required~Enter the Immutables"`
	Mutables                 types.Mutables        `json:"mutables" valid:"required~Enter the Mutables"`
}

var _ entities.InterIdentity = (*identity)(nil)

func (identity identity) GetID() types.ID { return identity.ID }
func (identity identity) GetChainID() types.ID {
	return identityIDFromInterface(identity.ID).ChainID
}

func (identity identity) GetClassificationID() types.ID {
	return identityIDFromInterface(identity.ID).ClassificationID
}
func (identity identity) GetProvisionedAddressList() []sdkTypes.AccAddress {
	return identity.ProvisionedAddressList
}
func (identity identity) GetUnprovisionedAddressList() []sdkTypes.AccAddress {
	return identity.UnprovisionedAddressList
}
func (identity identity) ProvisionAddress(accAddress sdkTypes.AccAddress) entities.InterIdentity {
	identity.ProvisionedAddressList = append(identity.ProvisionedAddressList, accAddress)
	return identity
}
func (identity identity) UnprovisionAddress(accAddress sdkTypes.AccAddress) entities.InterIdentity {
	for i, provisionedAddress := range identity.ProvisionedAddressList {
		if provisionedAddress.Equals(accAddress) {
			identity.ProvisionedAddressList = append(identity.ProvisionedAddressList[:i], identity.ProvisionedAddressList[i+1:]...)
			identity.UnprovisionedAddressList = append(identity.UnprovisionedAddressList, accAddress)
			return identity
		}
	}
	return identity
}
func (identity identity) GetImmutables() types.Immutables { return identity.Immutables }
func (identity identity) GetMutables() types.Mutables     { return identity.Mutables }
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
func NewIdentity(identityID types.ID, provisionedAddressList []sdkTypes.AccAddress, unprovisionedAddressList []sdkTypes.AccAddress, immutables types.Immutables, mutables types.Mutables) entities.InterIdentity {
	return identity{
		ID:                       identityID,
		ProvisionedAddressList:   provisionedAddressList,
		UnprovisionedAddressList: unprovisionedAddressList,
		Immutables:               immutables,
		Mutables:                 mutables,
	}
}
