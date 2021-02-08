/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package simulator

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/parameters/dummy"

	"github.com/cosmos/cosmos-sdk/x/simulation"
)

func (simulator) ParamChangeList(_ *rand.Rand) []simulation.ParamChange {
	return []simulation.ParamChange{
		simulation.NewSimParamChange(module.Name, string(dummy.ID.Bytes()),
			func(r *rand.Rand) string {
				//a, _ := json.Marshal(base.NewDecData(sdk.NewDec(1)))
				//a, _ := common.Codec.MarshalJSON(base.NewParameter(dummy.ID, dummy.DefaultData, func(interface{}) error { return nil }).GetData())

				return sdk.NewDec(1).String()
				//return base.NewParameter(dummy.ID, dummy.DefaultData, func(interface{}) error { return nil }).GetData().String()

				//return fmt.Sprintf("\"%s\"", base.NewDecData(sdk.NewDecWithPrec(int64(r.Intn(199)), 2)))
			}),
	}
}
