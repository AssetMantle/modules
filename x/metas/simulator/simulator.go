// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import "github.com/AssetMantle/modules/helpers"

type simulator struct{}

var _ helpers.Simulator = (*simulator)(nil)

func newSimulator() helpers.Simulator {
	return simulator{}
}
