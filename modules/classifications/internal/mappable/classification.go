// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/modules/classifications/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/mappables"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/types"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type classification struct {
	baseQualified.Document //nolint:govet
}

var _ mappables.Classification = (*classification)(nil)

func (classification classification) GetKey() helpers.Key {
	return key.FromID(classification.ID)
}
func (classification) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, module.Name, classification{})
}

func NewClassification(id types.ID, immutableProperties lists.PropertyList, mutableProperties lists.PropertyList) mappables.Classification {
	return classification{
		Document: baseQualified.Document{
			ID:         id,
			Immutables: baseQualified.Immutables{PropertyList: immutableProperties},
			Mutables:   baseQualified.Mutables{Properties: mutableProperties},
		},
	}
}
