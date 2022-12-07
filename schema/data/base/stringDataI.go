package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.StringData = (*StringDataI)(nil)

func (s *StringDataI) GetID() ids.DataID {
	return s.Impl.(data.StringData).GetID()
}
func (s *StringDataI) Bytes() []byte {
	return s.Impl.(data.StringData).Bytes()
}
func (s *StringDataI) GetType() ids.StringID {
	return s.Impl.(data.StringData).GetType()
}
func (s *StringDataI) ZeroValue() data.Data {
	return s.Impl.(data.StringData).ZeroValue()
}
func (s *StringDataI) GenerateHashID() ids.HashID {
	return s.Impl.(data.StringData).GenerateHashID()
}
func (s *StringDataI) Compare(listable traits.Listable) int {
	return s.Impl.(data.StringData).Compare(listable)
}
func (s *StringDataI) Get() string {
	return s.Impl.(data.StringData).Get()
}
