package base

import (
	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids2.PropertyID = (*PropertyIDI)(nil)

func (propertyIDI *PropertyIDI) GetKey() ids2.StringID {
	return propertyIDI.Impl.(ids2.PropertyID).GetKey()
}
func (propertyIDI *PropertyIDI) GetType() ids2.StringID {
	return propertyIDI.Impl.(ids2.PropertyID).GetType()
}
func (propertyIDI *PropertyIDI) Compare(listable traits.Listable) int {
	return propertyIDI.Impl.(ids2.PropertyID).Compare(listable)
}
func (propertyIDI *PropertyIDI) Bytes() []byte {
	return propertyIDI.Impl.(ids2.PropertyID).Bytes()
}
func (propertyIDI *PropertyIDI) IsPropertyID() {
}
