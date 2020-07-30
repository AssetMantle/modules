package helpers

type Auxiliary interface {
	GetName() string
	GetKeeper() AuxiliaryKeeper
	InitializeKeeper(Mapper, ...interface{})
}
