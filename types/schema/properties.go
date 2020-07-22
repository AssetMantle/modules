package schema

type Properties interface {
	Get(ID) Property

	GetList() []Property

	Add(Property) Properties
	Remove(Property) Properties
	Mutate(Property) Properties
}

type properties struct {
	PropertyList []Property
}

var _ Properties = (*properties)(nil)

func (properties properties) Get(id ID) Property {
	for _, property := range properties.GetList() {
		if property.GetID().Compare(id) == 0 {
			return property
		}
	}
	return nil
}
func (properties properties) GetList() []Property {
	var propertyList []Property
	for _, baseProperty := range properties.PropertyList {
		propertyList = append(propertyList, baseProperty)
	}
	return propertyList
}
func (properties properties) Add(property Property) Properties {
	propertyList := properties.GetList()
	propertyList = append(propertyList, property)
	return NewProperties(propertyList)
}
func (properties properties) Remove(property Property) Properties {
	propertyList := properties.GetList()
	for i, oldProperty := range propertyList {
		if oldProperty.GetID().Compare(property.GetID()) == 0 {
			propertyList = append(propertyList[:i], propertyList[i+1:]...)
		}
	}
	return NewProperties(propertyList)
}
func (properties properties) Mutate(property Property) Properties {
	propertyList := properties.GetList()
	for i, oldProperty := range propertyList {
		if oldProperty.GetID().Compare(property.GetID()) == 0 {
			propertyList[i] = property
		}
	}
	return NewProperties(propertyList)
}
func NewProperties(propertyList []Property) Properties {
	return properties{
		PropertyList: propertyList,
	}
}
