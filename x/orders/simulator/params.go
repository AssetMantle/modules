// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math"
	"math/rand"

	"github.com/AssetMantle/modules/x/orders/common"
	"github.com/AssetMantle/modules/x/orders/module"
	"github.com/AssetMantle/modules/x/orders/parameters/maxOrderLife"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

func (simulator) ParamChangeList(_ *rand.Rand) []simulationTypes.ParamChange {
	return []simulationTypes.ParamChange{
		simulation.NewSimParamChange(module.Name,
			string(maxOrderLife.Parameter.GetMetaProperty().GetID().Bytes()),
			func(r *rand.Rand) string {
				bytes, err := common.LegacyAmino.MarshalJSON(rand.Intn(math.MaxInt))
				if err != nil {
					panic(err)
				}
				return string(bytes)
			}),
	}
}
