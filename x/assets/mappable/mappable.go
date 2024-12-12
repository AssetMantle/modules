// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/schema/documents"
	"github.com/AssetMantle/schema/documents/base"
)

var _ helpers.Mappable = (*Mappable)(nil)

func (mappable *Mappable) ValidateBasic() error {
	return mappable.Asset.ValidateBasic()
}

func Prototype() helpers.Mappable {
	return &Mappable{}
}

func NewMappable(asset documents.Asset) helpers.Mappable {
	return &Mappable{
		Asset: asset.Get().(*base.Document),
	}
}

func GetAsset(mappable helpers.Mappable) documents.Asset {
	return base.NewAssetFromDocument(mappable.(*Mappable).Asset)
}
