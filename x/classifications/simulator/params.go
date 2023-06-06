// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math"
	"math/rand"

	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/constants"
	"github.com/AssetMantle/modules/x/classifications/parameters/bondRate"
	"github.com/AssetMantle/modules/x/classifications/parameters/maxPropertyCount"
)

func (simulator) ParamChangeList(_ *rand.Rand) []simulationTypes.ParamChange {
	return []simulationTypes.ParamChange{
		simulation.NewSimParamChange(constants.ModuleName,
			string(bondRate.Parameter.GetMetaProperty().GetID().Bytes()),
			func(r *rand.Rand) string {
				bytes, err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(rand.Intn(math.MaxInt))
				if err != nil {
					panic(err)
				}
				return string(bytes)
			},
		),
		simulation.NewSimParamChange(constants.ModuleName,
			string(maxPropertyCount.Parameter.GetMetaProperty().GetID().Bytes()),
			func(r *rand.Rand) string {
				bytes, err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(rand.Intn(1000) + 22)
				if err != nil {
					panic(err)
				}
				return string(bytes)
			},
		),
	}
}
