package utility

type Auxiliary interface {
	GetName() string
	GetKeeper() AuxiliaryKeeper
	InitializeKeeper(Mapper)
}

type auxiliary struct {
	moduleName       string
	name             string
	route            string
	auxiliaryKeeper  AuxiliaryKeeper
	initializeKeeper func(Mapper) AuxiliaryKeeper
}

var _ Auxiliary = (*auxiliary)(nil)

func (auxiliary auxiliary) GetName() string            { return auxiliary.name }
func (auxiliary auxiliary) GetKeeper() AuxiliaryKeeper { return auxiliary.auxiliaryKeeper }
func (auxiliary auxiliary) InitializeKeeper(mapper Mapper) {
	auxiliary.auxiliaryKeeper = auxiliary.initializeKeeper(mapper)
}

func NewAuxiliary(moduleName string, name string, route string, initializeKeeper func(Mapper) AuxiliaryKeeper) Auxiliary {
	return auxiliary{
		moduleName:       moduleName,
		name:             name,
		route:            route,
		initializeKeeper: initializeKeeper,
	}
}
