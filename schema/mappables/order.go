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
	GetMakerID() types.ID
	GetTakerID() types.ID
	GetMakerAssetAmount() sdkTypes.Dec
	GetMakerAssetData() types.ID
	GetTakerAssetAmount() sdkTypes.Dec
	GetTakerAssetData() types.ID
	GetSalt() types.Height
	SetTakerID(types.ID) Order
}
