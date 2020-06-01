package types

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type NFTWallet interface {
	AccAddress() sdkTypes.AccAddress
	NFTID() ID
}
