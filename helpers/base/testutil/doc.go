// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

// Package testutil provides shared test infrastructure for AssetMantle module tests.
//
// Use NewTestContext for simple store+context setup.
// Use NewTestContextWithBankAuth for tests that need auth+bank keepers.
// Use NewMockAuxiliaryPair for mock auxiliary+keeper pairs.
//
// Conventions:
//   - Test names: lowercase descriptive (e.g. "valid message", "empty address")
//   - Use require for setup, assert for assertions
//   - Always call mock.AssertExpectations(t) on mocks
//   - No package-level var blocks with side effects
//   - No reflect.DeepEqual — use assert.Equal or type-specific checks
package testutil
