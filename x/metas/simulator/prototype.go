// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import "github.com/AssetMantle/modules/helpers"

func Prototype() helpers.Simulator {
	return newSimulator()
}
