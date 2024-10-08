// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/schema/documents"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
)

var _ helpers.Mappable = (*Mappable)(nil)

func (mappable *Mappable) ValidateBasic() error {
	return mappable.Maintainer.ValidateBasic()
}

func NewMappable(maintainer documents.Maintainer) helpers.Mappable {
	return &Mappable{Maintainer: maintainer.Get().(*baseDocuments.Document)}
}

func Prototype() helpers.Mappable {
	return &Mappable{}
}

func GetMaintainer(mappable helpers.Mappable) documents.Maintainer {
	return baseDocuments.NewMaintainerFromDocument(mappable.(*Mappable).Maintainer)
}
