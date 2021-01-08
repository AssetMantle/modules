/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type NFTWallet interface {
	GetAccAddress() sdkTypes.AccAddress
	GetNFTID() ID
}
