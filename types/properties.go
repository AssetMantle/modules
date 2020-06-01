package types

type Properties interface {
	String() string

	Property(ID) Property

	AddProperty(Property) error
	RemoveProperty(Property) error
	MutateProperty(Property) error
}

type BaseProperties struct {
	PropertyList []Property
}

var _ Properties = (*BaseProperties)(nil)

func (BaseProperties BaseProperties) String() string                 {}
func (BaseProperties BaseProperties) Property(ID) Property           {}
func (BaseProperties *BaseProperties) AddProperty(Property) error    {}
func (BaseProperties *BaseProperties) RemoveProperty(Property) error {}
func (BaseProperties *BaseProperties) MutateProperty(Property) error {}
