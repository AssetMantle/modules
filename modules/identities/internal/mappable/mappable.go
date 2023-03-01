// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/documents/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

var _ helpers.Mappable = (*Mappable)(nil)

func (mappable *Mappable) ValidateBasic() error {
	return mappable.Identity.ValidateBasic()
}
func (mappable *Mappable) GetKey() helpers.Key {
	return key.NewKey(baseIDs.NewIdentityID(mappable.GetIdentity().GetClassificationID(), mappable.GetIdentity().GetImmutables()))
}
func (*Mappable) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, Mappable{})
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

func MappablesFromInterface(mappables []helpers.Mappable) []*Mappable {
	Mappables := make([]*Mappable, len(mappables))
	for index, mappable := range mappables {
		Mappables[index] = mappable.(*Mappable)
	}
	return Mappables
}

func MappablesToInterface(mappables []*Mappable) []helpers.Mappable {
	Mappables := make([]helpers.Mappable, len(mappables))
	for index, mappable := range mappables {
		Mappables[index] = mappable
	}
	return Mappables
}
