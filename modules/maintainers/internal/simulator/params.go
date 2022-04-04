// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/AssetMantle/modules/modules/maintainers/internal/common"
	"github.com/AssetMantle/modules/modules/maintainers/internal/module"
	"github.com/AssetMantle/modules/modules/maintainers/internal/parameters/dummy"
	"github.com/AssetMantle/modules/schema/types/base"
)

func (simulator) ParamChangeList(_ *rand.Rand) []simulation.ParamChange {
	return []simulation.ParamChange{
		simulation.NewSimParamChange(module.Name,
			dummy.ID.String(),
			func(r *rand.Rand) string {
				bytes, err := common.Codec.MarshalJSON(dummy.Parameter.Mutate(base.NewDecData(sdkTypes.NewDecWithPrec(int64(r.Intn(99)), 2))).GetData())
				if err != nil {
					panic(err)
				}
				return string(bytes)
			}),
	}
}
