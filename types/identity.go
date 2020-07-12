package types

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type Identity interface {
	GetID() ID
	GetAddressList() []sdkTypes.AccAddress
	GetDeletedAddressList() []sdkTypes.AccAddress

	AddAddress(sdkTypes.AccAddress) Identity
	DeleteAddress(sdkTypes.AccAddress) Identity

	IsActive(sdkTypes.AccAddress) bool
}

type identity struct {
	ID                 ID
	AddressList        []sdkTypes.AccAddress
	DeletedAddressList []sdkTypes.AccAddress
}

var _ Identity = (*identity)(nil)

func (identity identity) GetID() ID                             { return identity.ID }
func (identity identity) GetAddressList() []sdkTypes.AccAddress { return identity.AddressList }
func (identity identity) GetDeletedAddressList() []sdkTypes.AccAddress {
	return identity.DeletedAddressList
}

func (identity identity) AddAddress(accAddress sdkTypes.AccAddress) Identity {
	identity.AddressList = append(identity.AddressList, accAddress)
	return identity
}
func (identity identity) DeleteAddress(accAddress sdkTypes.AccAddress) Identity {
	for i, activeAddress := range identity.AddressList {
		if activeAddress.Equals(accAddress) {
			identity.AddressList = append(identity.AddressList[:i], identity.AddressList[i+1:]...)
			identity.DeletedAddressList = append(identity.DeletedAddressList, accAddress)
		}
	}
	return identity
}

func (identity identity) IsActive(accAddress sdkTypes.AccAddress) bool {
	for _, activeAddress := range identity.AddressList {
		if activeAddress.Equals(accAddress) {
			return true
		}
	}
	return false
}
