/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type QueryKeeper interface {
	Enquire(sdkTypes.Context, QueryRequest) QueryResponse
	Keeper
}
