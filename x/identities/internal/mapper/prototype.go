// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mapper

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/identities/internal/key"
	"github.com/AssetMantle/modules/x/identities/internal/mappable"
)

func Prototype() helpers.Mapper {
	return baseHelpers.NewMapper(key.Prototype, mappable.Prototype)
}
