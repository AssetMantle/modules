package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/base"
	"sort"
)

var _ lists.PropertyList = (*List_PropertyList)(nil)

func (propertyList *List_PropertyList) GetProperty(propertyID ids.ID) properties.Property {
	if i, found := propertyList.Search(base.NewEmptyMesaPropertyFromID(propertyID)); found {
		return propertyList.GetList()[i]
	}

	return nil
}
func (propertyList *List_PropertyList) Search(property properties.Property) (index int, found bool) {
	index = sort.Search(
		len(propertyList.PropertyList.List),
		func(i int) bool {
			return propertyList.PropertyList.List[i].Compare(property) >= 0
		},
	)

	if index < len(propertyList.PropertyList.List) && propertyList.PropertyList.List[index].Compare(property) == 0 {
		return index, true
	}

	return index, false
}
func (propertyList *List_PropertyList) GetList() []properties.Property {
	Properties := make([]properties.Property, len(propertyList.PropertyList.List))
	for i, listable := range propertyList.PropertyList.List {
		Properties[i] = listable
	}
	return Properties
}
func (propertyList *List_PropertyList) GetPropertyIDList() lists.List {
	propertyIDList := NewIDList()
	for _, property := range propertyList.GetList() {
		propertyIDList = propertyIDList.Add(property.GetID())
	}
	return propertyIDList
}
func (propertyList *List_PropertyList) Add(properties ...properties.Property) lists.List {
	updatedList := propertyList
	for _, listable := range properties {
		if index, found := updatedList.Search(listable); !found {
			updatedList.PropertyList.List = append(updatedList.PropertyList.List, listable.(*base.Property))
			copy(updatedList.PropertyList.List[index+1:], updatedList.PropertyList.List[index:])
			updatedList.PropertyList.List[index] = listable.(*base.Property)
		}
	}
	return &List{
		Impl: updatedList,
	}
}
func (propertyList *List_PropertyList) Remove(properties ...properties.Property) lists.List {
	updatedList := propertyList

	for _, listable := range properties {
		if index, found := updatedList.Search(listable); found {
			updatedList.PropertyList.List = append(updatedList.PropertyList.List[:index], updatedList.PropertyList.List[index+1:]...)
		}
	}

	return &List{
		Impl: updatedList,
	}
}
func (propertyList *List_PropertyList) Mutate(properties ...properties.Property) lists.List {
	updatedList := propertyList

	for _, listable := range properties {
		if index, found := updatedList.Search(listable); found {
			updatedList.PropertyList.List[index] = listable.(*base.Property)
		}
	}

	return &List{
		Impl: updatedList,
	}
}
func (propertyList *List_PropertyList) ScrubData() lists.List {
	newPropertyList := NewPropertyList()
	for _, listable := range propertyList.PropertyList.List {
		if property := listable; property.IsMeta() {
			newPropertyList = newPropertyList.Add(property.Impl.(properties.Property).ScrubData())
		} else {
			newPropertyList = newPropertyList.Add(property)
		}
	}
	return newPropertyList
}

func NewPropertyList(properties ...properties.Property) lists.List {
	var propertyList []*base.Property

	for _, dataVal := range properties {
		propertyList = append(propertyList, dataVal.(*base.Property))
	}
	sort.Slice(propertyList, func(i, j int) bool {
		return propertyList[i].Compare(propertyList[j]) <= 0
	})
	return &List{
		Impl: &List_PropertyList{
			PropertyList: &PropertyList{
				List: propertyList,
			},
		},
	}
}
