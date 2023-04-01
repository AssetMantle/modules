// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"context"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/AssetMantle/modules/schema/helpers"
)

type block struct {
	mapper           helpers.Mapper
	parameterManager helpers.ParameterManager
}

var _ helpers.Block = (*block)(nil)

func (b block) Begin(_ context.Context, _ abci.RequestBeginBlock) {

}

func (b block) End(_ context.Context, _ abci.RequestEndBlock) {

}

func (b block) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, _ ...interface{}) helpers.Block {
	return block{mapper, parameterManager}
}

func TestBlockPrototype() helpers.Block {
	return block{}
}
