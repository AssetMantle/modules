// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"
	"strconv"

	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/utilities/random"
	"github.com/AssetMantle/modules/x/metas/constants"
	"github.com/AssetMantle/modules/x/metas/parameters/reveal_enabled"
)

func (simulator) ParamChangeList(_ *rand.Rand) []simulationTypes.LegacyParamChange {
	return []simulationTypes.LegacyParamChange{
		simulation.NewSimLegacyParamChange(constants.ModuleName,
			string(reveal_enabled.Parameter.GetMetaProperty().GetID().Bytes()),
			func(r *rand.Rand) string {
				bytes, err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(strconv.FormatBool(random.GenerateRandomBool()))
				if err != nil {
					panic(err)
				}
				return string(bytes)
			}),
	}
}
