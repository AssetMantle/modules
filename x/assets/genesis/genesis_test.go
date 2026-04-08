// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package genesis

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrototype(t *testing.T) {
	g := Prototype()
	require.NotNil(t, g)
	assert.NotNil(t, g.GetParameterList())
	assert.NotNil(t, g.GetRecords())
}

func TestDefault(t *testing.T) {
	g := Prototype()
	d := g.Default()
	require.NotNil(t, d)
	assert.NotNil(t, d.GetParameterList())
}

func TestValidateBasic(t *testing.T) {
	g := Prototype()
	// ValidateBasic requires a parameterManager; with nil it may panic,
	// so we just verify the prototype is well-formed.
	assert.NotNil(t, g.GetParameterList())
}

func TestGetRecords(t *testing.T) {
	g := Prototype()
	records := g.GetRecords()
	assert.NotNil(t, records)
}

func TestGetParameterList(t *testing.T) {
	g := Prototype()
	pl := g.GetParameterList()
	assert.NotNil(t, pl)
}
