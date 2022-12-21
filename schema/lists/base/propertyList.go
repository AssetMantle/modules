package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/base"
	"sort"
)

var _ lists.PropertyList = (*PropertyList)(nil)

func (propertyList *PropertyList) GetProperty(propertyID ids.PropertyID) properties.Property {
	if i, found := propertyList.Search(base.NewEmptyMesaPropertyFromID(propertyID)); found {
		return propertyList.GetList()[i]
	}

	return nil
}
func (propertyList *PropertyList) GetList() []properties.Property {
	Properties := make([]properties.Property, len(propertyList.PropertyList))
	for i, listable := range propertyList.PropertyList {
		Properties[i] = listable
	}
	return Properties
}
func (propertyList *PropertyList) Search(property properties.Property) (index int, found bool) {
	index = sort.Search(
		len(propertyList.PropertyList),
		func(i int) bool {
			return propertyList.PropertyList[i].Compare(property) >= 0
		},
	)

	if index < len(propertyList.PropertyList) && propertyList.PropertyList[index].Compare(property) == 0 {
		return index, true
	}

	return index, false
}
func (propertyList *PropertyList) GetPropertyIDList() lists.IDList {
	propertyIDList := NewIDList()
	for _, property := range propertyList.GetList() {
		propertyIDList = propertyIDList.Add(property.GetID())
	}
	return propertyIDList
}
func (propertyList *PropertyList) Add(properties ...properties.Property) lists.PropertyList {
	updatedList := propertyList
	for _, listable := range properties {
		if index, found := updatedList.Search(listable); !found {
			updatedList.PropertyList = append(updatedList.PropertyList, listable.(*base.AnyProperty))
			copy(updatedList.PropertyList[index+1:], updatedList.PropertyList[index:])
			updatedList.PropertyList[index] = listable.(*base.AnyProperty)
		}
	}
	return updatedList
}
func (propertyList *PropertyList) Remove(properties ...properties.Property) lists.PropertyList {
	updatedList := propertyList

	for _, listable := range properties {
		if index, found := updatedList.Search(listable); found {
			updatedList.PropertyList = append(updatedList.PropertyList[:index], updatedList.PropertyList[index+1:]...)
		}
	}

	return updatedList
}
func (propertyList *PropertyList) Mutate(properties ...properties.Property) lists.PropertyList {
	updatedList := propertyList

	for _, listable := range properties {
		if index, found := updatedList.Search(listable); found {
			updatedList.PropertyList[index] = listable.(*base.AnyProperty)
		}
	}

	return updatedList
}
func (propertyList *PropertyList) ScrubData() lists.PropertyList {
	newPropertyList := NewPropertyList()
	for _, listable := range propertyList.PropertyList {
		if property := listable; property.IsMeta() {
			newPropertyList = newPropertyList.Add(property.Impl.(properties.Property).ScrubData())
		} else {
			newPropertyList = newPropertyList.Add(property)
		}
	}
	return newPropertyList
}

func NewPropertyList(properties ...properties.Property) lists.PropertyList {
	var propertyList []*base.AnyProperty

	for _, dataVal := range properties {
		propertyList = append(propertyList, dataVal.(*base.AnyProperty))
	}
	return &PropertyList{PropertyList: propertyList}
}
