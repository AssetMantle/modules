// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	"github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/x/splits/common"
	"github.com/AssetMantle/modules/x/splits/module"
	"github.com/AssetMantle/modules/x/splits/parameters/wrapAllowedCoins"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

func (simulator) ParamChangeList(_ *rand.Rand) []simulationTypes.ParamChange {
	return []simulationTypes.ParamChange{
		simulation.NewSimParamChange(module.Name,
			string(wrapAllowedCoins.Parameter.GetMetaProperty().GetID().Bytes()),
			func(r *rand.Rand) string {
				bytes, err := common.LegacyAmino.MarshalJSON(base.GenerateRandomCoinListString(5))
				if err != nil {
					panic(err)
				}
				return string(bytes)
			}),
	}
}
