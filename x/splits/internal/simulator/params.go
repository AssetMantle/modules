// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"github.com/AssetMantle/modules/simulation/schema/types/base"
	"math/rand"

	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/AssetMantle/modules/x/splits/internal/common"
	"github.com/AssetMantle/modules/x/splits/internal/module"
	"github.com/AssetMantle/modules/x/splits/internal/parameters/wrapAllowedCoins"
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
