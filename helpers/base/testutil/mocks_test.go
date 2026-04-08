// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package testutil

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/AssetMantle/modules/helpers"
)

func TestNewMockAuxiliaryPair(t *testing.T) {
	aux, keeper := NewMockAuxiliaryPair()

	// GetKeeper should return the wired keeper
	assert.Same(t, keeper, aux.GetKeeper())

	// set up a mock expectation on the keeper
	keeper.On("Help", mock.Anything, mock.Anything).
		Return(new(helpers.AuxiliaryResponse), nil).Once()

	resp, err := keeper.Help(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	keeper.AssertExpectations(t)
	aux.AssertExpectations(t)
}

func TestNewMockAuxiliaryPair_multipleAuxiliaries(t *testing.T) {
	auth, authKeeper := NewMockAuxiliaryPair()
	bond, bondKeeper := NewMockAuxiliaryPair()

	// each pair is independent
	assert.NotSame(t, authKeeper, bondKeeper)
	assert.Same(t, authKeeper, auth.GetKeeper())
	assert.Same(t, bondKeeper, bond.GetKeeper())
}

func TestNewNamedMockAuxiliaryPair(t *testing.T) {
	aux, keeper := NewNamedMockAuxiliaryPair("authenticate")

	assert.Equal(t, "authenticate", aux.GetName())
	assert.Same(t, keeper, aux.GetKeeper())

	// set up a mock expectation on the keeper
	keeper.On("Help", mock.Anything, mock.Anything).
		Return(new(helpers.AuxiliaryResponse), nil).Once()

	resp, err := keeper.Help(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	keeper.AssertExpectations(t)
	aux.AssertExpectations(t)
}

func TestNewNamedMockAuxiliaryPair_differentNames(t *testing.T) {
	auth, _ := NewNamedMockAuxiliaryPair("authenticate")
	bond, _ := NewNamedMockAuxiliaryPair("bond")

	assert.Equal(t, "authenticate", auth.GetName())
	assert.Equal(t, "bond", bond.GetName())
}
