/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

type Block interface {
	Begin(context sdkTypes.Context, beginBlockRequest abciTypes.RequestBeginBlock)
	End(context sdkTypes.Context, endBlockRequest abciTypes.RequestEndBlock)
	Initialize(mapper Mapper, parameters Parameters, auxiliaries []Auxiliary) Block
}
