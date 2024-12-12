// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/schema/types"
	baseTypes "github.com/AssetMantle/schema/types/base"
)

var _ helpers.Mappable = (*Mappable)(nil)

func (mappable *Mappable) ValidateBasic() error {
	return mappable.Split.ValidateBasic()
}

func NewMappable(split types.Split) helpers.Mappable {
	return &Mappable{Split: split.(*baseTypes.Split)}
}

func Prototype() helpers.Mappable {
	return &Mappable{}
}

func GetSplit(mappable helpers.Mappable) types.Split {
	if mappable == nil || mappable.(*Mappable).Split == nil {
		return nil
	}
	return mappable.(*Mappable).Split
}
