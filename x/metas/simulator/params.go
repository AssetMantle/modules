// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"
	"strconv"

	"github.com/AssetMantle/modules/utilities/random"

	"github.com/AssetMantle/modules/x/metas/common"
	"github.com/AssetMantle/modules/x/metas/module"
	"github.com/AssetMantle/modules/x/metas/parameters/revealEnabled"

	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

func (simulator) ParamChangeList(_ *rand.Rand) []simulationTypes.ParamChange {
	return []simulationTypes.ParamChange{
		simulation.NewSimParamChange(module.Name,
			string(revealEnabled.Parameter.GetMetaProperty().GetID().Bytes()),
			func(r *rand.Rand) string {
				bytes, err := common.LegacyAmino.MarshalJSON(strconv.FormatBool(random.GenerateRandomBool()))
				if err != nil {
					panic(err)
				}
				return string(bytes)
			}),
	}
}
