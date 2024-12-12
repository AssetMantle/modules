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
	return mappable.Identity.ValidateBasic()
}

func NewMappable(identity documents.Identity) helpers.Mappable {
	return &Mappable{
		Identity: identity.Get().(*base.Document),
	}
}

func GetIdentity(mappable helpers.Mappable) documents.Identity {
	return base.NewIdentityFromDocument(mappable.(*Mappable).Identity)
}

func Prototype() helpers.Mappable {
	return &Mappable{}
}
