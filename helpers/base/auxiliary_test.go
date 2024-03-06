// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"

	"github.com/AssetMantle/modules/helpers"
)

func TestGetName(t *testing.T) {

	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			"Valid name",
			"sampleName",
			"sampleName",
		},
		{
			"Empty name",
			"",
			"",
		},
		{
			"Whitespace name",
			"   ",
			"   ",
		},
		{
			"Special character in name",
			"@#$-_=+!",
			"@#$-_=+!",
		},
	}

	// Create mock KeeperPrototype function
	mockKeeperPrototype := func() helpers.AuxiliaryKeeper {
		return nil
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aux := NewAuxiliary(tt.input, mockKeeperPrototype)
			if got := aux.GetName(); got != tt.expect {
				t.Errorf("GetName() = %v, want = %v", got, tt.expect)
			}
		})
	}
}

func TestGetKeeper(t *testing.T) {
	tests := []struct {
		name        string
		auxiliary   auxiliary
		expectedObj helpers.AuxiliaryKeeper
	}{
		{
			name: "Normal case",
			auxiliary: auxiliary{
				auxiliaryKeeper: &dummyAuxiliaryKeeper{},
				name:            "testCase1",
			},
			expectedObj: &dummyAuxiliaryKeeper{},
		},
		{
			name: "Nil case",
			auxiliary: auxiliary{
				auxiliaryKeeper: nil,
				name:            "testCase2",
			},
			expectedObj: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotKeeper := tt.auxiliary.GetKeeper(); gotKeeper != tt.expectedObj {
				t.Errorf("GetKeeper() = %v, want %v", gotKeeper, tt.expectedObj)
			}
		})
	}
}

func TestInitialize(t *testing.T) {
	// Define the test cases
	testCases := []struct {
		name              string
		auxiliaryName     string
		keeperPrototype   func() helpers.AuxiliaryKeeper
		mapper            helpers.Mapper
		parameterManager  helpers.ParameterManager
		auxiliaryKeepers  []interface{}
		expectedAuxiliary helpers.Auxiliary
	}{
		{
			name:             "Valid case",
			auxiliaryName:    "TestAuxiliary",
			keeperPrototype:  func() helpers.AuxiliaryKeeper { return nil },
			mapper:           nil,
			parameterManager: nil,
			auxiliaryKeepers: []interface{}{ /*...*/ },
			expectedAuxiliary: &auxiliary{ // Adjust as per actual implementation
				name:            "TestAuxiliary",
				keeperPrototype: func() helpers.AuxiliaryKeeper { return nil },
				auxiliaryKeeper: nil,
			},
		},
		// Add more test cases for all corner cases
	}

	// Run the tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Define the auxiliary
			auxiliary := NewAuxiliary(tc.auxiliaryName, tc.keeperPrototype)

			// Call initialize method
			gotAuxiliary := auxiliary.Initialize(tc.mapper, tc.parameterManager, tc.auxiliaryKeepers...)

			// Assert that the returned Auxiliary is the expected one
			assert.Equal(t, tc.expectedAuxiliary, gotAuxiliary)
		})
	}
}

func TestNewAuxiliary(t *testing.T) {
	var tests = []struct {
		name            string
		keeperPrototype func() helpers.AuxiliaryKeeper
		expectedName    string
		keeperTest      bool
	}{
		{"test1", func() helpers.AuxiliaryKeeper { return nil }, "test1", false},
		{"", func() helpers.AuxiliaryKeeper { return nil }, "", false},
		{"test3", func() helpers.AuxiliaryKeeper { return &dummyAuxiliaryKeeper{} }, "test3", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliary := NewAuxiliary(tt.name, tt.keeperPrototype)
			if gotName := auxiliary.GetName(); gotName != tt.expectedName {
				t.Errorf("NewAuxiliary().GetName() = %v, expected %v", gotName, tt.expectedName)
			}
			if gotKeeper := auxiliary.GetKeeper(); (gotKeeper != nil) != tt.keeperTest {
				t.Errorf("NewAuxiliary().GetKeeper() existence = %v, expected %v", gotKeeper != nil, tt.keeperTest)
			}
		})
	}
}

type dummyAuxiliaryKeeper struct{}

func (d *dummyAuxiliaryKeeper) Initialize(_ helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return nil
}
func (d *dummyAuxiliaryKeeper) Help(_ context.Context, _ helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	return nil, nil
}
