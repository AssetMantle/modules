// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"
	"strconv"

	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/utilities/random"
	"github.com/AssetMantle/modules/x/assets/module"
	"github.com/AssetMantle/modules/x/assets/parameters/burnEnabled"
	"github.com/AssetMantle/modules/x/assets/parameters/mintEnabled"
	"github.com/AssetMantle/modules/x/assets/parameters/renumerateEnabled"
)

func (simulator) ParamChangeList(_ *rand.Rand) []simulationTypes.ParamChange {
	return []simulationTypes.ParamChange{
		simulation.NewSimParamChange(module.Name,
			string(mintEnabled.Parameter.GetMetaProperty().GetID().Bytes()),
			func(r *rand.Rand) string {
				bytes, err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(strconv.FormatBool(random.GenerateRandomBool()))
				if err != nil {
					panic(err)
				}
				return string(bytes)
			}),
		simulation.NewSimParamChange(module.Name,
			string(burnEnabled.Parameter.GetMetaProperty().GetID().Bytes()),
			func(r *rand.Rand) string {
				bytes, err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(strconv.FormatBool(random.GenerateRandomBool()))
				if err != nil {
					panic(err)
				}
				return string(bytes)
			}),
		simulation.NewSimParamChange(module.Name,
			string(renumerateEnabled.Parameter.GetMetaProperty().GetID().Bytes()),
			func(r *rand.Rand) string {
				bytes, err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(strconv.FormatBool(random.GenerateRandomBool()))
				if err != nil {
					panic(err)
				}
				return string(bytes)
			}),
	}
}
