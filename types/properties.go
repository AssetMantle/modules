package types

type Properties interface {
	Get(ID) Property

	AddProperty(Property) error
	RemoveProperty(Property) error
	MutateProperty(Property) error
}

type BaseProperties struct {
	PropertyList []Property
}

var _ Properties = (*BaseProperties)(nil)

func (BaseProperties BaseProperties) Get(ID) Property                {}
func (BaseProperties *BaseProperties) AddProperty(Property) error    {}
func (BaseProperties *BaseProperties) RemoveProperty(Property) error {}
func (BaseProperties *BaseProperties) MutateProperty(Property) error {}
