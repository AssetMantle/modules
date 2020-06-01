package types

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type AssetWallet interface {
	NFTWallet

	String() string

	AccAddress() sdkTypes.AccAddress
	NFTID() ID
}
