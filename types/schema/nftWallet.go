package schema

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type NFTWallet interface {
	GetAccAddress() sdkTypes.AccAddress
	GetNFTID() ID
}
