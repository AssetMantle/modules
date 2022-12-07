// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	baseDocuments "github.com/AssetMantle/modules/schema/documents/base"
	"github.com/AssetMantle/modules/schema/errors"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/parameters"
	baseParameters "github.com/AssetMantle/modules/schema/parameters/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/traits"
	typesSchema "github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/schema/helpers"
)

func MakeModuleCode(keyPrototype func() helpers.Key, mappablePrototype func() helpers.Mappable) *codec.LegacyAmino {
	Codec := codec.NewLegacyAmino()
	keyPrototype().RegisterCodec(Codec)
	mappablePrototype().RegisterCodec(Codec)
	RegisterCodec(Codec)
	Codec.Seal()

	return Codec
}

func MakeMessageCodec(messagePrototype func() helpers.Message) *codec.LegacyAmino {
	Codec := codec.NewLegacyAmino()
	messagePrototype().RegisterCodec(Codec)
	RegisterCodec(Codec)
	Codec.Seal()

	return Codec
}

func RegisterCodec(codec *codec.LegacyAmino) {

	data.RegisterCodec(codec)
	baseData.RegisterCodec(codec)

	documents.RegisterCodec(codec)
	baseDocuments.RegisterCodec(codec)

	errors.RegisterCodec(codec)

	helpers.RegisterCodec(codec)

	ids.RegisterCodec(codec)
	baseIDs.RegisterCodec(codec)

	//lists.RegisterCodec(codec)
	//baseLists.RegisterCodec(codec)

	parameters.RegisterCodec(codec)
	baseParameters.RegisterCodec(codec)

	properties.RegisterCodec(codec)
	baseProperties.RegisterCodec(codec)

	qualified.RegisterCodec(codec)
	baseQualified.RegisterCodec(codec)

	traits.RegisterCodec(codec)

	typesSchema.RegisterCodec(codec)
	baseTypes.RegisterCodec(codec)
}
