/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package simulator

import "github.com/persistenceOne/persistenceSDK/schema/helpers"

func Prototype() helpers.Simulator {
	return newSimulator()
}
