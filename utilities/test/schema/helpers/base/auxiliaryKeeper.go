// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"context"

	"github.com/AssetMantle/schema/x/helpers"
)

type testAuxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*testAuxiliaryKeeper)(nil)

func (t testAuxiliaryKeeper) Help(_ context.Context, _ helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	return nil, nil
}

func (t testAuxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return testAuxiliaryKeeper{mapper: mapper}
}

func TestAuxiliaryKeeperPrototype() helpers.AuxiliaryKeeper {
	return testAuxiliaryKeeper{}
}
