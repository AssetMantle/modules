// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/schema/data"
	baseData "github.com/AssetMantle/schema/data/base"
)

var _ helpers.Mappable = (*Mappable)(nil)

func (mappable *Mappable) ValidateBasic() error {
	return mappable.Data.ValidateBasic()
}

func NewMappable(data data.Data) helpers.Mappable {
	return &Mappable{Data: data.ToAnyData().(*baseData.AnyData)}
}

func Prototype() helpers.Mappable {
	return &Mappable{}
}

func GetData(mappable helpers.Mappable) data.Data {
	return mappable.(*Mappable).Data
}
