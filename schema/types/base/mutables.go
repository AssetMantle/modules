package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type mutables struct {
	Properties    types.Properties
	MaintainersID types.ID
}

var _ types.Mutables = (*mutables)(nil)

func (mutables mutables) Get() types.Properties {
	return mutables.Properties
}
func (mutables mutables) GetMaintainersID() types.ID {
	return mutables.MaintainersID
}
func NewMutables(properties types.Properties, maintainersID types.ID) types.Mutables {
	return mutables{
		Properties:    properties,
		MaintainersID: maintainersID,
	}
}
