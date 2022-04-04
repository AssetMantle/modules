// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/modules/classifications/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/mappables"
	qualifiedMappables "github.com/AssetMantle/modules/schema/mappables/qualified"
	"github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/types"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type classification struct {
	qualifiedMappables.Document //nolint:govet
}

var _ mappables.Classification = (*classification)(nil)

func (classification classification) GetKey() helpers.Key {
	return key.FromID(classification.ID)
}
func (classification) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, classification{})
}

func NewClassification(id types.ID, immutableProperties types.Properties, mutableProperties types.Properties) mappables.Classification {
	return classification{
		Document: qualifiedMappables.Document{
			ID:            id,
			HasImmutables: base.HasImmutables{Properties: immutableProperties},
			HasMutables:   base.HasMutables{Properties: mutableProperties},
		},
	}
}
