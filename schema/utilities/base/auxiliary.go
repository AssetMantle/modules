package base

import "github.com/persistenceOne/persistenceSDK/schema/utilities"

type auxiliary struct {
	moduleName       string
	name             string
	route            string
	auxiliaryKeeper  utilities.AuxiliaryKeeper
	initializeKeeper func(utilities.Mapper) utilities.AuxiliaryKeeper
}

var _ utilities.Auxiliary = (*auxiliary)(nil)

func (auxiliary auxiliary) GetName() string                      { return auxiliary.name }
func (auxiliary auxiliary) GetKeeper() utilities.AuxiliaryKeeper { return auxiliary.auxiliaryKeeper }
func (auxiliary auxiliary) InitializeKeeper(mapper utilities.Mapper) {
	auxiliary.auxiliaryKeeper = auxiliary.initializeKeeper(mapper)
}

func NewAuxiliary(moduleName string, name string, route string, initializeKeeper func(utilities.Mapper) utilities.AuxiliaryKeeper) utilities.Auxiliary {
	return auxiliary{
		moduleName:       moduleName,
		name:             name,
		route:            route,
		initializeKeeper: initializeKeeper,
	}
}
