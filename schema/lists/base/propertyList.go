package base

import (
	"sort"

	"github.com/AssetMantle/modules/schema/traits"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/base"
)

var _ lists.PropertyList = (*PropertyList)(nil)

func (propertyList *PropertyList) ValidateBasic() error {
	for _, property := range propertyList.PropertyList {
		if err := property.ValidateBasic(); err != nil {
			return err
		}
	}
	return nil
}
func (propertyList *PropertyList) GetProperty(propertyID ids.PropertyID) properties.AnyProperty {
	if i, found := propertyList.Search(base.NewEmptyMesaPropertyFromID(propertyID)); found {
		return propertyList.GetList()[i]
	}

	return nil
}
func (propertyList *PropertyList) GetList() []properties.AnyProperty {
	Properties := make([]properties.AnyProperty, len(propertyList.PropertyList))
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
func (propertyList PropertyList) Add(properties ...properties.Property) lists.PropertyList {
	updatedList := propertyList
	for _, listable := range properties {
		if index, found := updatedList.Search(listable); !found {
			updatedList.PropertyList = append(updatedList.PropertyList, listable.ToAnyProperty().(*base.AnyProperty))
			copy(updatedList.PropertyList[index+1:], updatedList.PropertyList[index:])
			updatedList.PropertyList[index] = listable.ToAnyProperty().(*base.AnyProperty)
		} else {
			updatedList.PropertyList[index] = listable.ToAnyProperty().(*base.AnyProperty)
		}
	}
	return &updatedList
}
func (propertyList PropertyList) Remove(properties ...properties.Property) lists.PropertyList {
	updatedList := propertyList

	for _, listable := range properties {
		if index, found := updatedList.Search(listable); found {
			updatedList.PropertyList = append(updatedList.PropertyList[:index], updatedList.PropertyList[index+1:]...)
		}
	}

	return &updatedList
}

//TODO: Check if this is required
func (propertyList PropertyList) Mutate(properties ...properties.Property) lists.PropertyList {
	updatedList := propertyList

	for _, listable := range properties {
		if index, found := updatedList.Search(listable); found {
			updatedList.PropertyList[index] = listable.ToAnyProperty().(*base.AnyProperty)
		}
	}

	return &updatedList
}
func (propertyList *PropertyList) ScrubData() lists.PropertyList {
	newPropertyList := NewPropertyList()
	for _, listable := range propertyList.PropertyList {
		if property := listable; property.IsMeta() {
			newPropertyList = newPropertyList.Add(property.Get().(properties.MetaProperty).ScrubData())
		} else {
			newPropertyList = newPropertyList.Add(property)
		}
	}
	return newPropertyList
}
func (propertyList *PropertyList) sort() lists.PropertyList {
	sort.Slice(propertyList.PropertyList, func(i, j int) bool {
		return propertyList.PropertyList[i].Compare(propertyList.PropertyList[j]) <= 0
	})

	return propertyList
}

// TODO: Evaluate need
func propertiesToListables(properties ...properties.Property) []traits.Listable {
	listables := make([]traits.Listable, len(properties))
	for i, property := range properties {
		listables[i] = property
	}
	return listables
}

func NewPropertyList(properties ...properties.Property) lists.PropertyList {
	return (&PropertyList{}).Add(properties...)
}
