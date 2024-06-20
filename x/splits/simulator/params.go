// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/x/assets/parameters/wrap_allowed_coins"
	"github.com/AssetMantle/modules/x/splits/constants"
)

func (simulator) ParamChangeList(_ *rand.Rand) []simulationTypes.LegacyParamChange {
	return []simulationTypes.LegacyParamChange{
		simulation.NewSimLegacyParamChange(constants.ModuleName,
			string(wrap_allowed_coins.Parameter.GetMetaProperty().GetID().Bytes()),
			func(r *rand.Rand) string {
				bytes, err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(base.GenerateRandomCoinListString(5))
				if err != nil {
					panic(err)
				}
				return string(bytes)
			}),
	}
}
