/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import "github.com/persistenceOne/persistenceSDK/schema/helpers"

type auxiliary struct {
	name            string
	auxiliaryKeeper helpers.AuxiliaryKeeper
	keeperPrototype func() helpers.AuxiliaryKeeper
}

var _ helpers.Auxiliary = (*auxiliary)(nil)

func (auxiliary auxiliary) GetName() string                    { return auxiliary.name }
func (auxiliary auxiliary) GetKeeper() helpers.AuxiliaryKeeper { return auxiliary.auxiliaryKeeper }
func (auxiliary auxiliary) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaryKeepers ...interface{}) helpers.Auxiliary {
	auxiliary.auxiliaryKeeper = auxiliary.keeperPrototype().Initialize(mapper, parameters, auxiliaryKeepers).(helpers.AuxiliaryKeeper)
	return auxiliary
}
func NewAuxiliary(name string, keeperPrototype func() helpers.AuxiliaryKeeper) helpers.Auxiliary {
	return auxiliary{
		name:            name,
		keeperPrototype: keeperPrototype,
	}
}
