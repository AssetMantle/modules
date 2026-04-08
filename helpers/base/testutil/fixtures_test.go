// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package testutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestAddress(t *testing.T) {
	addr1 := TestAddress()
	addr2 := TestAddress()

	assert.NotEmpty(t, addr1)
	assert.NotEmpty(t, addr2)
	assert.NotEqual(t, addr1, addr2, "addresses should be unique")
}

func TestTestImmutables(t *testing.T) {
	imm := TestImmutables()
	assert.NotNil(t, imm)
	assert.NotNil(t, imm.GetImmutablePropertyList())
	assert.Greater(t, len(imm.GetImmutablePropertyList().Get()), 0)
}

func TestTestMutables(t *testing.T) {
	mut := TestMutables()
	assert.NotNil(t, mut)
	assert.NotNil(t, mut.GetMutablePropertyList())
	assert.Greater(t, len(mut.GetMutablePropertyList().Get()), 0)
}

func TestTestClassificationID(t *testing.T) {
	id := TestClassificationID()
	assert.NotNil(t, id)
	assert.NoError(t, id.ValidateBasic())
}

func TestTestIdentityID(t *testing.T) {
	id := TestIdentityID()
	assert.NotNil(t, id)
	assert.NoError(t, id.ValidateBasic())
}
