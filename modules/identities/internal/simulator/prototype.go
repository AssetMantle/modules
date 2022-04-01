// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import "github.com/persistenceOne/persistenceSDK/schema/helpers"

func Prototype() helpers.Simulator {
	return newSimulator()
}
