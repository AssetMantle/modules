/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappables

import "github.com/persistenceOne/persistenceSDK/schema/traits"

type Meta interface {
	//TODO return to interface
	Get() string
	traits.Mappable
}
