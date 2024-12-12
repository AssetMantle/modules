// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/schema/documents"
	"github.com/AssetMantle/schema/documents/base"
)

var _ helpers.Mappable = (*Mappable)(nil)

func (mappable *Mappable) ValidateBasic() error {
	return mappable.Classification.ValidateBasic()
}
func NewMappable(classification documents.Classification) helpers.Mappable {
	return &Mappable{
		Classification: classification.Get().(*base.Document),
	}
}

func Prototype() helpers.Mappable {
	return &Mappable{}
}

func GetClassification(mappable helpers.Mappable) documents.Classification {
	return base.NewClassificationFromDocument(mappable.(*Mappable).Classification)
}
