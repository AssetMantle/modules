/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type testAuxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*testAuxiliaryKeeper)(nil)

func (t testAuxiliaryKeeper) Help(_ sdkTypes.Context, _ helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	return nil
}

func (t testAuxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return testAuxiliaryKeeper{mapper: mapper}
}

func TestAuxiliaryKeeperPrototype() helpers.AuxiliaryKeeper {
	return testAuxiliaryKeeper{}
}
