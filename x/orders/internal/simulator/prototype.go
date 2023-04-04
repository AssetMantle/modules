// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import "github.com/AssetMantle/schema/x/helpers"

func Prototype() helpers.Simulator {
	return newSimulator()
}
