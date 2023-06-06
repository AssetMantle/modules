// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"
	"strconv"

	"github.com/AssetMantle/modules/x/maintainers/parameters/deputizeAllowed"

	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/utilities/random"
	"github.com/AssetMantle/modules/x/maintainers/constants"
)

func (simulator) ParamChangeList(_ *rand.Rand) []simulationTypes.ParamChange {
	return []simulationTypes.ParamChange{
		simulation.NewSimParamChange(constants.ModuleName,
			string(deputizeAllowed.Parameter.GetMetaProperty().GetID().Bytes()),
			func(r *rand.Rand) string {
				bytes, _ := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(strconv.FormatBool(random.GenerateRandomBool()))
				return string(bytes)
			}),
	}
}
