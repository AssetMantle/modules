// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"github.com/AssetMantle/modules/utilities/random"
	"math/rand"
	"strconv"

	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/AssetMantle/modules/x/maintainers/internal/common"
	"github.com/AssetMantle/modules/x/maintainers/internal/module"
	"github.com/AssetMantle/modules/x/maintainers/internal/parameters/deputizeAllowed"
)

func (simulator) ParamChangeList(_ *rand.Rand) []simulationTypes.ParamChange {
	return []simulationTypes.ParamChange{
		simulation.NewSimParamChange(module.Name,
			string(deputizeAllowed.Parameter.GetMetaProperty().GetID().Bytes()),
			func(r *rand.Rand) string {
				bytes, _ := common.LegacyAmino.MarshalJSON(strconv.FormatBool(random.GenerateRandomBool()))
				return string(bytes)
			}),
	}
}
