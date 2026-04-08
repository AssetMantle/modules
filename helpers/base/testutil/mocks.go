// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package testutil

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/AssetMantle/modules/helpers"
)

// MockAuxiliary implements helpers.Auxiliary using testify/mock.
type MockAuxiliary struct {
	mock.Mock
}

var _ helpers.Auxiliary = (*MockAuxiliary)(nil)

func (m *MockAuxiliary) GetName() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockAuxiliary) GetKeeper() helpers.AuxiliaryKeeper {
	args := m.Called()
	return args.Get(0).(helpers.AuxiliaryKeeper)
}

func (m *MockAuxiliary) Initialize(_ helpers.Mapper, _ helpers.ParameterManager, _ ...interface{}) helpers.Auxiliary {
	panic("MockAuxiliary.Initialize should not be called — use NewMockAuxiliaryPair instead")
}

// MockAuxiliaryKeeper implements helpers.AuxiliaryKeeper using testify/mock.
type MockAuxiliaryKeeper struct {
	mock.Mock
}

var _ helpers.AuxiliaryKeeper = (*MockAuxiliaryKeeper)(nil)

func (m *MockAuxiliaryKeeper) Help(ctx context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(helpers.AuxiliaryResponse), args.Error(1)
}

func (m *MockAuxiliaryKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, i []interface{}) helpers.Keeper {
	args := m.Called(mapper, parameterManager, i)
	return args.Get(0).(helpers.Keeper)
}

// NewMockAuxiliaryPair creates a MockAuxiliary and MockAuxiliaryKeeper pair
// with GetKeeper already wired. Returns both so callers can set expectations
// on the keeper.
func NewMockAuxiliaryPair() (*MockAuxiliary, *MockAuxiliaryKeeper) {
	keeper := new(MockAuxiliaryKeeper)
	auxiliary := new(MockAuxiliary)
	auxiliary.On("GetKeeper").Return(keeper)
	return auxiliary, keeper
}
