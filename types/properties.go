package types

type Properties interface {
	Get(ID) Property

	PropertyList() []Property
	Add(Property) Properties
	Remove(Property) Properties
	Mutate(Property) Properties
}

type BaseProperties struct {
	BasePropertyList []BaseProperty
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
func (baseProperties BaseProperties) PropertyList() []Property {
	var propertyList []Property
	for _, baseProperty := range baseProperties.BasePropertyList {
		propertyList = append(propertyList, baseProperty)
	}
	return propertyList
}
func (baseProperties BaseProperties) Add(property Property) Properties {
	basePropertyList := baseProperties.BasePropertyList
	for i, oldProperty := range basePropertyList {
		if oldProperty.ID().Compare(property.ID()) < 0 {
			basePropertyList = append(append(basePropertyList[:i], BaseProperty{
				BaseID: BaseID{
					IDString: property.ID().String(),
				},
				BaseFact: BaseFact{
					BaseString:     property.Fact().String(),
					BaseSignatures: BaseSignaturesFromInterface(property.Fact().Signatures()),
				},
			}), basePropertyList[i+1:]...)
		}
	}
	return BaseProperties{BasePropertyList: basePropertyList}
}
func (baseProperties BaseProperties) Remove(property Property) Properties {
	basePropertyList := baseProperties.BasePropertyList
	for i, oldProperty := range basePropertyList {
		if oldProperty.ID().Compare(property.ID()) == 0 {
			basePropertyList = append(basePropertyList[:i], basePropertyList[i+1:]...)
		}
	}
	return BaseProperties{BasePropertyList: basePropertyList}
}
func (baseProperties BaseProperties) Mutate(property Property) Properties {
	basePropertyList := baseProperties.BasePropertyList
	for i, oldProperty := range basePropertyList {
		if oldProperty.ID().Compare(property.ID()) == 0 {
			basePropertyList[i] = BaseProperty{
				BaseID: BaseID{
					IDString: property.ID().String(),
				},
				BaseFact: BaseFact{
					BaseString:     property.Fact().String(),
					BaseSignatures: BaseSignaturesFromInterface(property.Fact().Signatures()),
				},
			}
		}
	}
	return BaseProperties{BasePropertyList: basePropertyList}
}
func BasePropertiesFromInterface(properties Properties) BaseProperties {
	var basePropertyList []BaseProperty
	for _, property := range properties.PropertyList() {
		basePropertyList = append(basePropertyList, BaseProperty{
			BaseID: BaseID{IDString: property.ID().String()},
			BaseFact: BaseFact{
				BaseString:     property.Fact().String(),
				BaseSignatures: BaseSignaturesFromInterface(property.Fact().Signatures()),
			},
		})
	}
	return BaseProperties{BasePropertyList: basePropertyList}
}
