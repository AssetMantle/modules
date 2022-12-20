package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.ID = (*AnyID)(nil)

func (m *AnyID) Compare(listable traits.Listable) int {
	return m.Impl.(ids.ID).Compare(listable)
}

func (m *AnyID) Bytes() []byte {
	return m.Impl.(ids.ID).Bytes()
}
