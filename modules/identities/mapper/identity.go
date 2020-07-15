package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types"
)

type identity struct {
	ID                 types.ID
	AddressList        []sdkTypes.AccAddress
	DeletedAddressList []sdkTypes.AccAddress
	Immutables         types.Immutables
	Mutables           types.Mutables
}

var _ types.InterIdentity = (*identity)(nil)

func (identity identity) GetID() types.ID { return identity.ID }
func (identity identity) GetChainID() types.ID {
	return identityIDFromInterface(identity.ID).ChainID
}

func (identity identity) GetClassificationID() types.ID {
	return identityIDFromInterface(identity.ID).ClassificationID
}
func (identity identity) GetAddressList() []sdkTypes.AccAddress { return identity.AddressList }
func (identity identity) GetDeletedAddressList() []sdkTypes.AccAddress {
	return identity.DeletedAddressList
}

func (identity identity) AddAddress(accAddress sdkTypes.AccAddress) types.InterIdentity {
	identity.AddressList = append(identity.AddressList, accAddress)
	return identity
}
func (identity identity) DeleteAddress(accAddress sdkTypes.AccAddress) types.InterIdentity {
	for i, activeAddress := range identity.AddressList {
		if activeAddress.Equals(accAddress) {
			identity.AddressList = append(identity.AddressList[:i], identity.AddressList[i+1:]...)
			identity.DeletedAddressList = append(identity.DeletedAddressList, accAddress)
		}
	}
	return identity
}
func (identity identity) GetImmutables() types.Immutables { return identity.Immutables }
func (identity identity) GetMutables() types.Mutables     { return identity.Mutables }
func (identity identity) IsActive(accAddress sdkTypes.AccAddress) bool {
	for _, activeAddress := range identity.AddressList {
		if activeAddress.Equals(accAddress) {
			return true
		}
	}
	return false
}
func NewIdentity(id types.ID, addressList []sdkTypes.AccAddress, deletedAddressList []sdkTypes.AccAddress, immutables types.Immutables, mutables types.Mutables) types.InterIdentity {
	return identity{
		ID:                 id,
		AddressList:        addressList,
		DeletedAddressList: deletedAddressList,
		Immutables:         immutables,
		Mutables:           mutables,
	}
}
