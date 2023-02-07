package base

import (
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.ID = (*ID)(nil)

func (id *ID) Compare(listable traits.Listable) int {
	return id.Impl.(ids.ID).Compare(listable)
}
func (id *ID) Bytes() []byte {
	return id.Impl.(ids.ID).Bytes()
}

func idFromInterface(i interface{}) ids.ID {
	switch value := i.(type) {
	case ids.ID:
		return value
	default:
		panic(constants.MetaDataError)
	}
}
