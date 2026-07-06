// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulation

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	errorConstants "github.com/AssetMantle/modules/helpers/constants"
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

func TestSafeOperation_PanicReturnsError(t *testing.T) {
	panickingOp := func(_ *rand.Rand, _ *baseapp.BaseApp, _ sdkTypes.Context, _ []simulationTypes.Account, _ string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		panic("test panic")
	}

	safeOp := SafeOperation("assets/define", panickingOp)
	msg, futures, err := safeOp(rand.New(rand.NewSource(1)), nil, sdkTypes.Context{}, nil, "test-chain")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "test panic")
	assert.Contains(t, err.Error(), "assets/define")
	assert.Nil(t, futures)
	assert.False(t, msg.OK)
}

func TestSafeOperation_PassesThrough(t *testing.T) {
	expectedMsg := simulationTypes.NoOpMsg("test", "test", "success")
	normalOp := func(_ *rand.Rand, _ *baseapp.BaseApp, _ sdkTypes.Context, _ []simulationTypes.Account, _ string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		return expectedMsg, nil, nil
	}

	safeOp := SafeOperation("test/route", normalOp)
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

func TestIsBusinessRejection(t *testing.T) {
	testCases := []struct {
		name string
		err  error
		want bool
	}{
		{"gate disabled", errorConstants.NotAuthorized.Wrapf("minting is not enabled"), true},
		{"split exhausted", errorConstants.InsufficientBalance.Wrapf("1 is less then 2"), true},
		{"random collision", errorConstants.EntityAlreadyExists.Wrapf("name identity with ID x already exists"), true},
		{"missing random state", errorConstants.EntityNotFound.Wrapf("split with ID x not found"), true},
		{"bank cannot afford bond", sdkErrors.ErrInsufficientFunds.Wrapf("spendable balance too low"), true},
		{"property cap exceeded", errorConstants.InvalidRequest.Wrapf("total property count 30 exceeds maximum 12"), true},
		{"sim db miss", ErrSimDBMiss, true},
		{"bond below minimum is a bug", errorConstants.InvalidRequest.Wrapf("bound amount is less than min allowed 10"), false},
		{"generic error is a bug", fmt.Errorf("boom"), false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.want, IsBusinessRejection(testCase.err))
		})
	}
}

func TestRejectionOrError(t *testing.T) {
	t.Run("business rejection becomes NoOp with nil error", func(t *testing.T) {
		msg, futures, err := RejectionOrError("assets", "mint", errorConstants.NotAuthorized.Wrapf("minting is not enabled"))
		require.NoError(t, err)
		assert.Nil(t, futures)
		assert.False(t, msg.OK)
		assert.Contains(t, msg.Comment, "minting is not enabled")
	})

	t.Run("delivery bug propagates the error", func(t *testing.T) {
		bug := fmt.Errorf("no matching transaction for message")
		msg, futures, err := RejectionOrError("assets", "mint", bug)
		require.Error(t, err)
		assert.Equal(t, bug, err)
		assert.Nil(t, futures)
		assert.False(t, msg.OK)
	})
}
