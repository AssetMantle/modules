package mappables

import (
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type InterNFTWallet interface {
	types.NFTWallet
	traits.Mappable
}
