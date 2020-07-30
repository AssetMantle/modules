package mappables

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Order interface {
	GetID() types.ID
	traits.InterChain
	traits.Burnable
	traits.Lockable
	traits.HasImmutables
	traits.Mappable
	GetMakerAddress() sdkTypes.AccAddress
	GetTakerAddress() sdkTypes.AccAddress
	GetMakerAssetAmount() sdkTypes.Dec
	GetMakerAssetData() interface{}
	GetTakerAssetAmount() sdkTypes.Dec
	GetTakerAssetData() interface{}
	GetSalt() types.Height
	SetTakerAddress(sdkTypes.AccAddress) Order
}
