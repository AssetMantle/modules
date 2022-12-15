// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package schema

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	baseDocuments "github.com/AssetMantle/modules/schema/documents/base"
	"github.com/AssetMantle/modules/schema/errors"
	baseErrors "github.com/AssetMantle/modules/schema/errors/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/parameters"
	baseParameters "github.com/AssetMantle/modules/schema/parameters/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/traits"
	typesSchema "github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {

	data.RegisterLegacyAminoCodec(legacyAmino)
	baseData.RegisterLegacyAminoCodec(legacyAmino)

	documents.RegisterLegacyAminoCodec(legacyAmino)
	baseDocuments.RegisterLegacyAminoCodec(legacyAmino)

	errors.RegisterLegacyAminoCodec(legacyAmino)
	baseErrors.RegisterLegacyAminoCodec(legacyAmino)

	helpers.RegisterLegacyAminoCodec(legacyAmino)

	ids.RegisterLegacyAminoCodec(legacyAmino)
	baseIDs.RegisterLegacyAminoCodec(legacyAmino)

	lists.RegisterLegacyAminoCodec(legacyAmino)
	baseLists.RegisterLegacyAminoCodec(legacyAmino)

	parameters.RegisterLegacyAminoCodec(legacyAmino)
	baseParameters.RegisterLegacyAminoCodec(legacyAmino)

	properties.RegisterLegacyAminoCodec(legacyAmino)
	baseProperties.RegisterLegacyAminoCodec(legacyAmino)

	qualified.RegisterLegacyAminoCodec(legacyAmino)
	baseQualified.RegisterLegacyAminoCodec(legacyAmino)

	traits.RegisterLegacyAminoCodec(legacyAmino)

	typesSchema.RegisterLegacyAminoCodec(legacyAmino)
	baseTypes.RegisterLegacyAminoCodec(legacyAmino)
}
