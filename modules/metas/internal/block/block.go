/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package block

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	abciTypes "github.com/tendermint/tendermint/abci/types"
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

func (block block) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, _ ...interface{}) helpers.Block {
	block.mapper, block.parameters = mapper, parameters
	return block
}
