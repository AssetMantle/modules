package types

func NewGenesisState(parameters Parameters, metas []Meta) *GenesisState {
	return &GenesisState{
		Parameters: parameters,
		Mappables:  metas,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(DefaultParameters, nil)
}
