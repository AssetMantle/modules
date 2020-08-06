/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import "github.com/persistenceOne/persistenceSDK/schema/helpers"

type auxiliary struct {
	moduleName       string
	name             string
	route            string
	auxiliaryKeeper  helpers.AuxiliaryKeeper
	initializeKeeper func(helpers.Mapper, []interface{}) helpers.AuxiliaryKeeper
}

var _ helpers.Auxiliary = (*auxiliary)(nil)

func (auxiliary auxiliary) GetName() string                    { return auxiliary.name }
func (auxiliary auxiliary) GetKeeper() helpers.AuxiliaryKeeper { return auxiliary.auxiliaryKeeper }
func (auxiliary *auxiliary) InitializeKeeper(mapper helpers.Mapper, auxiliaryKeepers ...interface{}) {
	auxiliary.auxiliaryKeeper = auxiliary.initializeKeeper(mapper, auxiliaryKeepers)
}
func NewAuxiliary(moduleName string, name string, route string, initializeKeeper func(helpers.Mapper, []interface{}) helpers.AuxiliaryKeeper) helpers.Auxiliary {
	return &auxiliary{
		moduleName:       moduleName,
		name:             name,
		route:            route,
		initializeKeeper: initializeKeeper,
	}
}
