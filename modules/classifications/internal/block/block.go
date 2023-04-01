// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package block

import (
	"context"

	abciTypes "github.com/tendermint/tendermint/abci/types"

	"github.com/AssetMantle/modules/schema/helpers"
)

type block struct {
	mapper           helpers.Mapper
	parameterManager helpers.ParameterManager
}

var _ helpers.Block = (*block)(nil)

func (block block) Begin(_ context.Context, _ abciTypes.RequestBeginBlock) {

}

func (block block) End(_ context.Context, _ abciTypes.RequestEndBlock) {

}

func Prototype() helpers.Block {
	return block{}
}

func (block block) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, _ ...interface{}) helpers.Block {
	block.mapper, block.parameterManager = mapper, parameterManager
	return block
}
