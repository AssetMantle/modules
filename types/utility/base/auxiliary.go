package base

import "github.com/persistenceOne/persistenceSDK/types/utility"

type auxiliary struct {
	moduleName       string
	name             string
	route            string
	auxiliaryKeeper  utility.AuxiliaryKeeper
	initializeKeeper func(utility.Mapper) utility.AuxiliaryKeeper
}

var _ utility.Auxiliary = (*auxiliary)(nil)

func (auxiliary auxiliary) GetName() string                    { return auxiliary.name }
func (auxiliary auxiliary) GetKeeper() utility.AuxiliaryKeeper { return auxiliary.auxiliaryKeeper }
func (auxiliary auxiliary) InitializeKeeper(mapper utility.Mapper) {
	auxiliary.auxiliaryKeeper = auxiliary.initializeKeeper(mapper)
}

func NewAuxiliary(moduleName string, name string, route string, initializeKeeper func(utility.Mapper) utility.AuxiliaryKeeper) utility.Auxiliary {
	return auxiliary{
		moduleName:       moduleName,
		name:             name,
		route:            route,
		initializeKeeper: initializeKeeper,
	}
}
