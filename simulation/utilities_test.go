// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulation

import (
	"fmt"
	"math/rand"
	"testing"

	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomAddresses(t *testing.T) {
	t.Run("valid rand produces AccAddress slice", func(t *testing.T) {
		got := GenerateRandomAddresses(rand.New(rand.NewSource(7)))
		assert.IsType(t, []sdkTypes.AccAddress{}, got)
	})

	t.Run("nil rand panics", func(t *testing.T) {
		assert.Panics(t, func() { GenerateRandomAddresses(nil) })
	})
}

func TestRandomBool(t *testing.T) {
	t.Run("valid rand produces bool", func(t *testing.T) {
		got := RandomBool(rand.New(rand.NewSource(7)))
		assert.IsType(t, false, got)
	})

	t.Run("nil rand panics", func(t *testing.T) {
		assert.Panics(t, func() { RandomBool(nil) })
	})
}

func TestSafeOperation_RecoversPanic(t *testing.T) {
	panickingOp := func(_ *rand.Rand, _ *baseapp.BaseApp, _ sdkTypes.Context, _ []simulationTypes.Account, _ string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		panic("test panic")
	}

	safeOp := SafeOperation(panickingOp)
	msg, futures, err := safeOp(rand.New(rand.NewSource(1)), nil, sdkTypes.Context{}, nil, "test-chain")

	require.NoError(t, err)
	assert.Nil(t, futures)
	assert.False(t, msg.OK)
	assert.Contains(t, msg.Comment, "test panic")
}

func TestSafeOperation_PassesThrough(t *testing.T) {
	expectedMsg := simulationTypes.NoOpMsg("test", "test", "success")
	normalOp := func(_ *rand.Rand, _ *baseapp.BaseApp, _ sdkTypes.Context, _ []simulationTypes.Account, _ string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		return expectedMsg, nil, nil
	}

	safeOp := SafeOperation(normalOp)
	msg, futures, err := safeOp(rand.New(rand.NewSource(1)), nil, sdkTypes.Context{}, nil, "test-chain")

	require.NoError(t, err)
	assert.Nil(t, futures)
	assert.False(t, msg.OK) // NoOpMsg has OK=false
	assert.Equal(t, "success", msg.Comment)
}

func TestExecuteMessage_NilModule(t *testing.T) {
	_, err := ExecuteMessage(sdkTypes.Context{}, nil, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "nil module or message")
}

func TestExecuteMessage_NilMessage(t *testing.T) {
	_, err := ExecuteMessage(sdkTypes.Context{}, nil, nil)
	require.Error(t, err)
}

func TestCalculateBondAmount_NonNil(t *testing.T) {
	immutables := baseQualified.NewImmutables(nil)
	mutables := baseQualified.NewMutables(nil)

	result := CalculateBondAmount(immutables, mutables)
	require.NotNil(t, result)
	assert.NotNil(t, result.Get())
}

func TestGetGenesisProperties_NonNil(t *testing.T) {
	// Reset global state for test isolation
	Immutables = &baseQualified.Immutables{}
	Mutables = &baseQualified.Mutables{}

	r := rand.New(rand.NewSource(42))
	imm, mut := GetGenesisProperties(r)

	require.NotNil(t, imm)
	require.NotNil(t, mut)
	assert.NotNil(t, imm.GetImmutablePropertyList())
	assert.NotNil(t, mut.GetMutablePropertyList())
}

func TestGetGenesisProperties_Idempotent(t *testing.T) {
	// Reset global state
	Immutables = &baseQualified.Immutables{}
	Mutables = &baseQualified.Mutables{}

	r := rand.New(rand.NewSource(42))
	imm1, mut1 := GetGenesisProperties(r)
	imm2, mut2 := GetGenesisProperties(r)

	// Second call returns cached values (PropertyList already non-nil)
	assert.Equal(t, fmt.Sprintf("%v", imm1), fmt.Sprintf("%v", imm2))
	assert.Equal(t, fmt.Sprintf("%v", mut1), fmt.Sprintf("%v", mut2))
}
