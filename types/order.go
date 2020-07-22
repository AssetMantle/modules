package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Order interface {
	NFT
	InterChain
	Burnable
	Lockable
	HasImmutables
	HasMutables
	GetMakerAddress() sdkTypes.AccAddress
	GetTakerAddress() sdkTypes.AccAddress
	//GetSenderAddress() sdkTypes.AccAddress
	//GetFeeRecipientAddress() sdkTypes.AccAddress
	GetMakerAssetAmount() sdkTypes.Dec
	GetMakerAssetData() ID
	//GetMakerFee() sdkTypes.Dec
	//GetMakerFeeAssetData() ID
	GetTakerAssetAmount() sdkTypes.Dec
	GetTakerAssetData() ID
	//GetTakerFee() sdkTypes.Dec
	//GetTakerFeeAssetData() ID
	//GetExpirationTime() Height
	GetSalt() Height
	SetTakerAddress(sdkTypes.AccAddress) Order
}
