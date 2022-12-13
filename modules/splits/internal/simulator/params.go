/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package simulator

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	simTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/common"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/parameters/dummy"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"math/rand"
)

func (simulator) ParamChangeList(_ *rand.Rand) []simTypes.ParamChange {
	return []simTypes.ParamChange{
		simulation.NewSimParamChange(module.Name,
			dummy.ID.String(),
			func(r *rand.Rand) string {
				bytes, Error := common.LegacyAminoCodec.MarshalJSON(dummy.Parameter.Mutate(base.NewDecData(sdkTypes.NewDecWithPrec(int64(r.Intn(99)), 2))).GetData())
				if Error != nil {
					panic(Error)
				}
				return string(bytes)
			}),
	}
}
