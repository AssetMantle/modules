// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

type Block interface {
	Begin(sdkTypes.Context, abciTypes.RequestBeginBlock)
	End(sdkTypes.Context, abciTypes.RequestEndBlock)
	Initialize(Mapper, Parameters, ...interface{}) Block
}
