// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mapper

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/assets/record"
)

func Prototype() helpers.Mapper {
	return baseHelpers.NewMapper(record.Prototype)
}
