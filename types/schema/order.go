package schema

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types/trait"
)

type Order interface {
	NFT
	trait.InterChain
	trait.Burnable
	trait.Lockable
	trait.HasImmutables
	trait.HasMutables
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
