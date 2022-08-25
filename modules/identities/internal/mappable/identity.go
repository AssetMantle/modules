// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables"
	propertiesSchema "github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type identity struct {
	qualified.Document
}

var _ mappables.Identity = (*identity)(nil)

func (identity identity) GetExpiry() propertiesSchema.Property {
	if property := identity.Document.GetProperty(constants.ExpiryHeightProperty.GetID()); property != nil {
		return property
	}

	return constants.ExpiryHeightProperty
}
func (identity identity) GetAuthentication() propertiesSchema.Property {
	if property := identity.Document.GetProperty(constants.AuthenticationProperty.GetID()); property != nil {
		return property
	}

	return constants.AuthenticationProperty
}
func (identity identity) GetKey() helpers.Key {
	return key.NewKey(baseIDs.NewIdentityID(identity.Document.GetClassificationID(), identity.Document.GetImmutables()))
}
func (identity) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, identity{})
}

func NewIdentity(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) mappables.Identity {
	return identity{Document: baseQualified.NewDocument(classificationID, immutables, mutables)}
}

func Prototype() helpers.Mappable {
	return identity{}
}
