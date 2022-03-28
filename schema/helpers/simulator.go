/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/types/simulation"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	simTypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

type Simulator interface {
	RandomizedGenesisState(*module.SimulationState)
	WeightedOperations(simTypes.AppParams, codec.JSONMarshaler) simulation.WeightedOperation
	WeightedProposalContentList() []simTypes.WeightedProposalContent
	ParamChangeList(*rand.Rand) []simTypes.ParamChange
}
