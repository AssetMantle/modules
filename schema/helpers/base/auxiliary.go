// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import "github.com/AssetMantle/modules/schema/helpers"

type auxiliary struct {
	name            string
	auxiliaryKeeper helpers.AuxiliaryKeeper
	keeperPrototype func() helpers.AuxiliaryKeeper
}

var _ helpers.Auxiliary = (*auxiliary)(nil)

func (auxiliary auxiliary) GetName() string                    { return auxiliary.name }
func (auxiliary auxiliary) GetKeeper() helpers.AuxiliaryKeeper { return auxiliary.auxiliaryKeeper }
func (auxiliary auxiliary) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaryKeepers ...interface{}) helpers.Auxiliary {
	auxiliary.auxiliaryKeeper = auxiliary.keeperPrototype().Initialize(mapper, parameterManager, auxiliaryKeepers).(helpers.AuxiliaryKeeper)
	return auxiliary
}
func NewAuxiliary(name string, keeperPrototype func() helpers.AuxiliaryKeeper) helpers.Auxiliary {
	return auxiliary{
		name:            name,
		keeperPrototype: keeperPrototype,
	}
}
