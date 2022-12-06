package base

import (
	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids2.StringID = (*StringIDI)(nil)

func (stringIDI *StringIDI) Compare(listable traits.Listable) int {
	return stringIDI.Impl.(ids2.StringID).Compare(listable)
}

func (stringIDI *StringIDI) Bytes() []byte {
	return stringIDI.Impl.(ids2.StringID).Bytes()
}

func (stringIDI *StringIDI) IsStringID() {
}
