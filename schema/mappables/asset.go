// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappables

import (
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

type Asset interface {
	traits.Burnable
	traits.Lockable
	traits.Splittable

	qualified.Document
	helpers.Mappable
}
