/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	abci "github.com/tendermint/tendermint/abci/types"
)

type block struct {
	mapper     helpers.Mapper
	parameters helpers.Parameters
}

var _ helpers.Block = (*block)(nil)

func (b block) Begin(_ sdkTypes.Context, _ abci.RequestBeginBlock) {

}

func (b block) End(_ sdkTypes.Context, _ abci.RequestEndBlock) {

}

func (b block) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, _ ...interface{}) helpers.Block {
	return block{mapper, parameters}
}

func TestBlockPrototype() helpers.Block {
	return block{}
}
