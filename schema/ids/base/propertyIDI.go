package base

import (
	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids"

	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type propertyIDI ids.PropertyID

var _ ids2.PropertyID = (*propertyIDI)(nil)

func (propertyIDI *propertyIDI) GetKey() ids2.StringID {
	return propertyIDI.Impl.(ids2.PropertyID).GetKey()
}
func (propertyIDI *propertyIDI) GetType() ids2.StringID {
	return propertyIDI.Impl.(ids2.PropertyID).GetType()
}
func (propertyIDI *propertyIDI) Compare(listable traits.Listable) int {
	return propertyIDI.Impl.(ids2.PropertyID).Compare(listable)
}
func (propertyIDI *propertyIDI) String() string {
	return propertyIDI.Impl.(ids2.PropertyID).String()
}
func (propertyIDI *propertyIDI) Bytes() []byte {
	return propertyIDI.Impl.(ids2.PropertyID).Bytes()
}
func (propertyIDI *propertyIDI) IsPropertyID() {
}
