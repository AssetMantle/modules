// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"encoding/json"
	"math/rand"
	"testing"

	assetsSimulator "github.com/AssetMantle/modules/x/assets/simulator"
	identitiesSimulator "github.com/AssetMantle/modules/x/identities/simulator"
	"github.com/AssetMantle/modules/x/splits/constants"
	"github.com/cosmos/cosmos-sdk/types/module"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupGenesisState(t *testing.T) module.SimulationState {
	t.Helper()
	r := rand.New(rand.NewSource(42))
	accs := simulationTypes.RandomAccounts(r, 5)
	simState := module.SimulationState{
		Rand:      r,
		Accounts:  accs,
		GenState:  make(map[string]json.RawMessage),
		AppParams: make(simulationTypes.AppParams),
	}
	assetsSimulator.Prototype().RandomizedGenesisState(&simState)
	identitiesSimulator.Prototype().RandomizedGenesisState(&simState)
	return simState
}

func TestRandomizedGenesisState(t *testing.T) {
	simState := setupGenesisState(t)

	s := newSimulator()
	s.RandomizedGenesisState(&simState)

	genesisBytes := simState.GenState[constants.ModuleName]
	require.NotNil(t, genesisBytes)
	assert.Greater(t, len(genesisBytes), 0)
}

func TestWeightedOperations(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	simState := module.SimulationState{
		Rand:      r,
		AppParams: make(simulationTypes.AppParams),
	}

	s := newSimulator()
	ops := s.WeightedOperations(simState, nil)

	// Splits only has govern tx — nil is correct.
	// Module logic is exercised transitively by other modules' operations.
	require.Nil(t, ops)
}

func TestParamChangeList(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	s := newSimulator()
	changes := s.ParamChangeList(r)

	require.Len(t, changes, 1)
	assert.Equal(t, constants.ModuleName, changes[0].Subspace())
	simValue := changes[0].SimValue()(r)
	assert.NotEmpty(t, simValue)
}

func TestProposalMessages(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	simState := module.SimulationState{
		Rand:      r,
		AppParams: make(simulationTypes.AppParams),
	}

	s := newSimulator()
	msgs := s.ProposalMessages(simState)

	require.Len(t, msgs, 1)
}
