package types

type Properties interface {
	Get(ID) Property

	PropertyList() []Property
	Add(Property) error
	Remove(Property) error
	Mutate(Property) error
}
type BasePropertyList []BaseProperty
type BaseProperties struct {
	BasePropertyList BasePropertyList
}

var _ Properties = (*BaseProperties)(nil)

func (baseProperties BaseProperties) Get(id ID) Property {
	for _, property := range baseProperties.BasePropertyList {
		if property.ID().Compare(id) == 0 {
			return property
		}
	}
	return nil
}
func (baseProperties *BaseProperties) PropertyList() []Property {
	var propertyList []Property
	for _, baseProperty := range baseProperties.BasePropertyList {
		propertyList = append(propertyList, &baseProperty)
	}
	return propertyList
}
func (baseProperties *BaseProperties) Add(property Property) error {
	propertyList := baseProperties.BasePropertyList
	for i, oldProperty := range propertyList {
		if oldProperty.ID().Compare(property.ID()) < 0 {
			propertyList = append(append(propertyList[:i], BaseProperty{
				BaseID: BaseID{
					IDString: property.ID().String(),
				},
				BaseFact: BaseFact{
					BaseBytes:      property.Fact().Bytes(),
					BaseSignatures: BaseSignaturesFromInterface(property.Fact().Signatures()),
				},
			}), propertyList[i+1:]...)
		}
	}
	return nil
}
func (baseProperties *BaseProperties) Remove(property Property) error {
	propertyList := baseProperties.BasePropertyList
	for i, oldProperty := range propertyList {
		if oldProperty.ID().Compare(property.ID()) == 0 {
			propertyList = append(propertyList[:i], propertyList[i+1:]...)
		}
	}
	return nil
}
func (baseProperties *BaseProperties) Mutate(property Property) error {
	basePropertyList := baseProperties.BasePropertyList
	for i, oldProperty := range basePropertyList {
		if oldProperty.ID().Compare(property.ID()) == 0 {
			basePropertyList[i] = BaseProperty{
				BaseID: BaseID{
					IDString: property.ID().String(),
				},
				BaseFact: BaseFact{
					BaseBytes:      property.Fact().Bytes(),
					BaseSignatures: BaseSignaturesFromInterface(property.Fact().Signatures()),
				},
			}
		}
	}
	return nil
}
func BasePropertiesFromInterface(properties Properties) BaseProperties {
	var basePropertyList []BaseProperty
	for _, property := range properties.PropertyList() {
		basePropertyList = append(basePropertyList, BaseProperty{
			BaseID: BaseID{IDString: property.ID().String()},
			BaseFact: BaseFact{
				BaseBytes:      property.Fact().Bytes(),
				BaseSignatures: BaseSignaturesFromInterface(property.Fact().Signatures()),
			},
		})
	}
	return BaseProperties{BasePropertyList: basePropertyList}
}
