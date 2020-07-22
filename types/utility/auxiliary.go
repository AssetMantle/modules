package utility

type Auxiliary interface {
	GetName() string
	GetKeeper() AuxiliaryKeeper
	InitializeKeeper(Mapper)
}
