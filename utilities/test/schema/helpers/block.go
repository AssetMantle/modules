// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/persistenceOne/persistenceSDK/schema/helpers"
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
