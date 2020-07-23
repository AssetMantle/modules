package utilities

type Auxiliary interface {
	GetName() string
	GetKeeper() AuxiliaryKeeper
	InitializeKeeper(Mapper)
}
