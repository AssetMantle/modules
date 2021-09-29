package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type metaProperty struct {
	ID   types.ID   `json:"id"`
	Data types.Data `json:"data"`
}

var _ types.MetaProperty = (*metaProperty)(nil)

func (metaProperty metaProperty) GetDataID() types.ID {
	return metaProperty.Data.GetDataID()
}

func (metaProperty metaProperty) GetTypeID() types.ID {
	return metaProperty.Data.GetTypeID()
}

func (metaProperty metaProperty) GetKeyID() types.ID {
	return metaProperty.Data.GetKeyID()
}

func (metaProperty metaProperty) GetHashID() types.ID {
	return metaProperty.Data.GenerateHashID()
}

func (metaProperty metaProperty) GetData() types.Data { return metaProperty.Data }
func (metaProperty metaProperty) GetID() types.ID     { return metaProperty.ID }
func (metaProperty metaProperty) RemoveData() types.Property {
	return NewProperty(metaProperty.ID, metaProperty.Data)
}

func NewMetaProperty(id types.ID, data types.Data) types.MetaProperty {
	return metaProperty{
		ID:   id,
		Data: data,
	}
}
