package base

import "github.com/persistenceOne/persistenceSDK/schema/helpers"

type auxiliary struct {
	moduleName       string
	name             string
	route            string
	auxiliaryKeeper  helpers.AuxiliaryKeeper
	initializeKeeper func(helpers.Mapper) helpers.AuxiliaryKeeper
}

var _ helpers.Auxiliary = (*auxiliary)(nil)

func (auxiliary auxiliary) GetName() string                    { return auxiliary.name }
func (auxiliary auxiliary) GetKeeper() helpers.AuxiliaryKeeper { return auxiliary.auxiliaryKeeper }
func (auxiliary auxiliary) InitializeKeeper(mapper helpers.Mapper) {
	auxiliary.auxiliaryKeeper = auxiliary.initializeKeeper(mapper)
}

func NewAuxiliary(moduleName string, name string, route string, initializeKeeper func(helpers.Mapper) helpers.AuxiliaryKeeper) helpers.Auxiliary {
	return auxiliary{
		moduleName:       moduleName,
		name:             name,
		route:            route,
		initializeKeeper: initializeKeeper,
	}
}
