package base

import (
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ lists.List = (*List)(nil)

func (m *List) Get() []traits.Listable {
	return m.Impl.(lists.List).Get()
}

func (m *List) Search(listable traits.Listable) (index int, found bool) {
	return m.Impl.(lists.List).Search(listable)
}

func (m *List) Add(listable ...traits.Listable) lists.List {
	return m.Impl.(lists.List).Add(listable...)
}

func (m *List) Remove(listable ...traits.Listable) lists.List {
	return m.Impl.(lists.List).Remove(listable...)
}

func (m *List) Mutate(listable ...traits.Listable) lists.List {
	return m.Impl.(lists.List).Mutate(listable...)
}
