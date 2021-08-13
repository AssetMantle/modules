package base

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

func (accAddress AccAddress) String() string {
	err := sdkTypes.VerifyAddressFormat(accAddress.GetValue())
	if err != nil {
		return ""
	}
	return sdkTypes.AccAddress(accAddress.GetValue()).String()
}

func (accAddress AccAddress) GetBytes() []byte {
	return accAddress.GetValue()
}

func (accAddress AccAddress) AsSDKTypesAccAddress() sdkTypes.AccAddress {
	return accAddress.GetValue()
}

func NewAccAddressFromString(address string) AccAddress {
	accAddress, err := sdkTypes.AccAddressFromBech32(address)
	if err != nil {
		return AccAddress{}
	}
	return AccAddress{Value: accAddress}
}

func NewAccAddressFromSDKTypesAccAddress(address sdkTypes.AccAddress) AccAddress {
	return AccAddress{Value: address}
}
