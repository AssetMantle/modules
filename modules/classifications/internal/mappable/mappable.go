// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/helpers"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type mappable struct {
	documents.Classification
}

var _ helpers.Mappable = (*mappable)(nil)

func (classification mappable) GetKey() helpers.Key {
	return key.NewKey(classification.GetClassificationID())
}
func (mappable) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, mappable{})
}

func NewMappable(classification documents.Classification) helpers.Mappable {
	return mappable{
		Classification: classification,
	}
}

func Prototype() helpers.Mappable {
	return mappable{}
}
