/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package simulator

import (
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"math/rand"
)

func (simulator) ParamChangeList(_ *rand.Rand) []simulation.ParamChange {
	return []simulation.ParamChange{}
}
