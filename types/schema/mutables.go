package schema

type Mutables interface {
	Get() Properties
	GetMaintainersID() ID
}

type mutables struct {
	Properties    Properties
	MaintainersID ID
}

var _ Mutables = (*mutables)(nil)

func (mutables mutables) Get() Properties {
	return mutables.Properties
}
func (mutables mutables) GetMaintainersID() ID {
	return mutables.MaintainersID
}
func NewMutables(properties Properties, maintainersID ID) Mutables {
	return mutables{
		Properties:    properties,
		MaintainersID: maintainersID,
	}
}
