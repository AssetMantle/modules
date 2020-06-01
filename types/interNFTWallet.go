package types

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type InterNFTWallet interface {
	NFTWallet

	String() string

	AccAddress() sdkTypes.AccAddress
	NFTID() ID
}
