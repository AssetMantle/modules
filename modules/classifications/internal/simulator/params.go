/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package simulator

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/x/simulation"
)

func (simulator) ParamChangeList(_ *rand.Rand) []simulation.ParamChange {
	return []simulation.ParamChange{
		//simulation.NewSimParamChange(module.Name, dummy.ID.String(),
		//	func(r *rand.Rand) string {
		//		// a, _ := json.Marshal(base.NewDecData(sdk.NewDec(1)))
		//		// a, _ := common.Codec.MarshalJSON(base.NewParameter(dummy.ID, dummy.DefaultData, func(interface{}) error { return nil }).GetData())
		//
		//		// return sdk.NewDec(1).String()
		//		// return base.NewParameter(dummy.ID, dummy.DefaultData, func(interface{}) error { return nil }).GetData().String()
		//
		//		//a, _ := base.NewDecData(sdk.NewDec(23)).MarshalJSON()
		//		//return fmt.Sprintf("\"%s\"", a)
		//
		//		return fmt.Sprintf("\"%v\"", base.NewDecData(sdk.NewDecWithPrec(int64(r.Intn(199)), 2)))
		//		//return fmt.Sprintf("\"%s\"", sdk.NewDecWithPrec(int64(r.Intn(99)), 2))
		//	}),
	}
}
