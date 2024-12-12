// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"math"
	"math/rand"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/orders/constants"
	"github.com/AssetMantle/modules/x/orders/parameters/max_order_life"
)

func (simulator) ParamChangeList(_ *rand.Rand) []simulationTypes.LegacyParamChange {
	return []simulationTypes.LegacyParamChange{
		simulation.NewSimLegacyParamChange(constants.ModuleName,
			string(max_order_life.Parameter.GetMetaProperty().GetID().Bytes()),
			func(r *rand.Rand) string {
				bytes, err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(rand.Intn(math.MaxInt))
				if err != nil {
					panic(err)
				}
				return string(bytes)
			}),
	}
}
