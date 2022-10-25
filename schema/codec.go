// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package schema

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
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
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*error)(nil), nil)

	data.RegisterCodec(codec)
	baseData.RegisterCodec(codec)

	errors.RegisterCodec(codec)
	baseErrors.RegisterCodec(codec)

	helpers.RegisterCodec(codec)

	ids.RegisterCodec(codec)
	baseIDs.RegisterCodec(codec)

	lists.RegisterCodec(codec)
	baseLists.RegisterCodec(codec)

	parameters.RegisterCodec(codec)
	baseParameters.RegisterCodec(codec)

	properties.RegisterCodec(codec)
	baseProperties.RegisterCodec(codec)

	qualified.RegisterCodec(codec)
	baseQualified.RegisterCodec(codec)

	traits.RegisterCodec(codec)

	types.RegisterCodec(codec)
	baseTypes.RegisterCodec(codec)
}
