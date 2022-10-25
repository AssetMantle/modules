package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/types"
)

type meta struct {
	data.Data
}

var _ types.Meta = (*meta)(nil)

func (meta meta) GetData() data.Data { return meta.Data }

func NewMeta(data data.Data) types.Meta {
	return meta{
		Data: data,
	}
}
