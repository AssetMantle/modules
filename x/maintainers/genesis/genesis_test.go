// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package genesis

import (
	"testing"

	storeTypes "cosmossdk.io/store/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	"github.com/AssetMantle/modules/x/maintainers/mapper"
	"github.com/AssetMantle/modules/x/maintainers/parameters"
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

func TestValidateBasic(t *testing.T) {
	g := Prototype()
	pm := parameters.Prototype()
	assert.NoError(t, g.ValidateBasic(pm))
}

func TestGenesis_Encode_Decode_roundtrip(t *testing.T) {
	g := Prototype()
	codec := baseHelpers.CodecPrototype()
	encoded := g.Encode(codec.GetProtoCodec())
	require.NotNil(t, encoded)
	require.Greater(t, len(encoded), 0)

	decoded := g.Decode(codec.GetProtoCodec(), encoded)
	require.NotNil(t, decoded)
	assert.Equal(t, len(g.GetRecords()), len(decoded.GetRecords()))
}

func TestGenesis_Import_Export_roundtrip(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	paramStoreKey := storeTypes.NewKVStoreKey("testParams")
	ctx := testutil.NewTestContext(t, storeKey, paramStoreKey)

	m := mapper.Prototype().Initialize(storeKey)
	pm := parameters.Prototype().Initialize(paramStoreKey)

	g := Prototype()
	g.Import(ctx, m, pm)

	exported := g.Export(ctx, m, pm)
	assert.NotNil(t, exported)
	assert.Equal(t, len(g.GetRecords()), len(exported.GetRecords()))
}

func TestGenesis_SetRecords_SetParameters(t *testing.T) {
	g := Prototype()
	g = g.SetRecords(g.GetRecords())
	g = g.SetParameters(g.GetParameterList())
	assert.NotNil(t, g)
	assert.NotNil(t, g.GetRecords())
	assert.NotNil(t, g.GetParameterList())
}

func TestGenesis_Initialize(t *testing.T) {
	g := Prototype()
	initialized := g.Initialize(g.GetRecords(), g.GetParameterList())
	assert.NotNil(t, initialized)
	assert.NotNil(t, initialized.GetRecords())
	assert.NotNil(t, initialized.GetParameterList())
}

func TestGenesis_Initialize_empty_records(t *testing.T) {
	g := Prototype()
	initialized := g.Initialize(nil, g.GetParameterList())
	assert.NotNil(t, initialized)
	assert.NotNil(t, initialized.GetParameterList())
	assert.NotNil(t, initialized.GetRecords())
}
