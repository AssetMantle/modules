package types

type Properties interface {
	Get(ID) Property

	Add(Property) error
	Remove(Property) error
	Mutate(Property) error
}

type BaseProperties struct {
	PropertyList []Property
}

var _ Properties = (*BaseProperties)(nil)

func (baseProperties BaseProperties) Get(id ID) Property {
	for _, property := range baseProperties.PropertyList {
		if property.ID().Compare(id) == 0 {
			return property
		}
	}
	return nil
}
func (baseProperties *BaseProperties) Add(property Property) error {
	propertyList := baseProperties.PropertyList
	for i, oldProperty := range propertyList {
		if oldProperty.ID().Compare(property.ID()) < 0 {
			propertyList = append(append(propertyList[:i], property), propertyList[i+1:]...)
		}
	}
	return nil
}
func (baseProperties *BaseProperties) Remove(property Property) error {
	propertyList := baseProperties.PropertyList
	for i, oldProperty := range propertyList {
		if oldProperty.ID().Compare(property.ID()) == 0 {
			propertyList = append(propertyList[:i], propertyList[i+1:]...)
		}
	}
	return nil
}
func (baseProperties *BaseProperties) Mutate(property Property) error {
	propertyList := baseProperties.PropertyList
	for i, oldProperty := range propertyList {
		if oldProperty.ID().Compare(property.ID()) == 0 {
			propertyList[i] = property
		}
	}
	return nil
}
