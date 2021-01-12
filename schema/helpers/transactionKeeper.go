/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type TransactionKeeper interface {
	Transact(sdkTypes.Context, sdkTypes.Msg) TransactionResponse
	Keeper
}
