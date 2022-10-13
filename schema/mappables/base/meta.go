package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/mappables"
)

type meta struct {
	data.Data
}

var _ mappables.Meta = (*meta)(nil)

func (meta meta) GetData() data.Data { return meta.Data }

func NewMeta(data data.Data) mappables.Meta {
	return meta{
		Data: data,
	}
}
