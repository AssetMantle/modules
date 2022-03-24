/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

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
