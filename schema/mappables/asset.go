// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappables

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
)

type Asset interface {
	traits.Burnable
	traits.Lockable
	traits.Splittable

	Document
	helpers.Mappable
}
