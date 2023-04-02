package base

import (
	"sort"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ lists.PropertyList = (*PropertyList)(nil)

func (propertyList *PropertyList) ValidateBasic() error {
	for _, property := range propertyList.Properties {
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
	Properties := make([]properties.AnyProperty, len(propertyList.Properties))
	for i, listable := range propertyList.Properties {
		Properties[i] = listable
	}
	return Properties
}
func (propertyList *PropertyList) Search(property properties.Property) (index int, found bool) {
	index = sort.Search(
		len(propertyList.Properties),
		func(i int) bool {
			return propertyList.Properties[i].Compare(property) >= 0
		},
	)

	if index < len(propertyList.Properties) && propertyList.Properties[index].Compare(property) == 0 {
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
			updatedList.Properties = append(updatedList.Properties, listable.ToAnyProperty().(*base.AnyProperty))
			copy(updatedList.Properties[index+1:], updatedList.Properties[index:])
			updatedList.Properties[index] = listable.ToAnyProperty().(*base.AnyProperty)
		} else {
			updatedList.Properties[index] = listable.ToAnyProperty().(*base.AnyProperty)
		}
	}
	return &updatedList
}
func (propertyList PropertyList) Remove(properties ...properties.Property) lists.PropertyList {
	updatedList := propertyList

	for _, listable := range properties {
		if index, found := updatedList.Search(listable); found {
			updatedList.Properties = append(updatedList.Properties[:index], updatedList.Properties[index+1:]...)
		}
	}

	return &updatedList
}

// TODO: Check if this is required
func (propertyList PropertyList) Mutate(properties ...properties.Property) lists.PropertyList {
	updatedList := propertyList

	for _, listable := range properties {
		if index, found := updatedList.Search(listable); found {
			updatedList.Properties[index] = listable.ToAnyProperty().(*base.AnyProperty)
		}
	}

	return &updatedList
}
func (propertyList *PropertyList) ScrubData() lists.PropertyList {
	newPropertyList := NewPropertyList()
	for _, listable := range propertyList.Properties {
		if property := listable; property.IsMeta() {
			newPropertyList = newPropertyList.Add(property.Get().(properties.MetaProperty).ScrubData())
		} else {
			newPropertyList = newPropertyList.Add(property)
		}
	}
	return newPropertyList
}
func (propertyList *PropertyList) sort() lists.PropertyList {
	sort.Slice(propertyList.Properties, func(i, j int) bool {
		return propertyList.Properties[i].Compare(propertyList.Properties[j]) <= 0
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
