// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package block

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	"github.com/AssetMantle/modules/schema/helpers"
)

type block struct {
	mapper     helpers.Mapper
	parameters helpers.Parameters
}

var _ helpers.Block = (*block)(nil)

func (block block) Begin(_ sdkTypes.Context, _ abciTypes.RequestBeginBlock) {

}

func (block block) End(_ sdkTypes.Context, _ abciTypes.RequestEndBlock) {

}

func Prototype() helpers.Block {
	return block{}
}

func (block block) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, _ ...interface{}) helpers.Block {
	block.mapper, block.parameters = mapper, parameters
	return block
}
