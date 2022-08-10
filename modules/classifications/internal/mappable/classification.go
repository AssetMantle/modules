// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids/base"
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
	return key.NewKey(classification.GetClassificationID())
}
func (classification) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, classification{})
}

func NewClassification(immutables qualified.Immutables, mutables qualified.Mutables) mappables.Classification {
	return classification{
		Document: baseQualified.NewDocument(base.NewClassificationID(immutables, mutables), immutables, mutables),
	}
}

func Prototype() helpers.Mappable {
	return classification{}
}
