// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import "github.com/AssetMantle/schema/x/helpers"

type simulator struct {
}

var _ helpers.Simulator = (*simulator)(nil)

func newSimulator() helpers.Simulator {
	return simulator{}
}
