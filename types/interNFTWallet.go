package types

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type InterNFTWallet interface {
	NFTWallet

	AccAddress() sdkTypes.AccAddress
	NFTID() ID
}
