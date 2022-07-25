// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type classification struct {
	qualified.Document
}

var _ mappables.Classification = (*classification)(nil)

func (classification classification) GetKey() helpers.Key {
	return key.FromID(classification.ID)
}
func (classification) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, classification{})
}

func NewClassification(immutableProperties lists.PropertyList, mutableProperties lists.PropertyList) mappables.Classification {
	return classification{
		document: baseQualified.New{
			ID: key.NewClassificationID(immutableProperties, mutableProperties),
			// TODO Add classificationID
			Immutables: baseQualified.immutables{PropertyList: immutableProperties},
			Mutables:   baseQualified.mutables{PropertyList: mutableProperties},
		},
	}
}
