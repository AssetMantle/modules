// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package documents

import (
	"github.com/AssetMantle/modules/schema/capabilities"
)

type Asset interface {
	capabilities.Burnable
	capabilities.Lockable
	capabilities.Splittable

	Document
}
