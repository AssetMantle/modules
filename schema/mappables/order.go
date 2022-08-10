// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappables

import (
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
)

type Order interface {
	// TODO check if ID return type required
	GetRateID() ids.StringID
	GetCreationID() ids.StringID
	GetMakerOwnableID() ids.OwnableID
	GetTakerOwnableID() ids.OwnableID
	GetMakerID() ids.IdentityID

	// TODO check if rest of the IDs can be converted to meta properties
	GetCreation() properties.MetaProperty
	GetExchangeRate() properties.MetaProperty

	GetTakerID() properties.Property
	GetExpiry() properties.Property
	GetMakerOwnableSplit() properties.Property

	qualified.Document
	helpers.Mappable
}
