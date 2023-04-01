// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/AssetMantle/modules/modules/assets/internal/common"
	"github.com/AssetMantle/modules/modules/assets/internal/module"
	"github.com/AssetMantle/modules/modules/assets/internal/parameters/mintEnabled"
	"github.com/AssetMantle/modules/schema/data/base"
)

func (simulator) ParamChangeList(_ *rand.Rand) []simulationTypes.ParamChange {
	return []simulationTypes.ParamChange{
		simulation.NewSimParamChange(module.Name,
			mintEnabled.ID.AsString(),
			func(r *rand.Rand) string {
				bytes, err := common.LegacyAmino.MarshalJSON(mintEnabled.Parameter.Mutate(base.NewDecData(sdkTypes.NewDecWithPrec(int64(r.Intn(99)), 2))))
				if err != nil {
					panic(err)
				}
				return string(bytes)
			}),
	}
}
