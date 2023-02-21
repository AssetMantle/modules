// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"

	abciTypes "github.com/tendermint/tendermint/abci/types"
)

type Block interface {
	Begin(context.Context, abciTypes.RequestBeginBlock)
	End(context.Context, abciTypes.RequestEndBlock)
	Initialize(Mapper, ParameterManager, ...interface{}) Block
}
