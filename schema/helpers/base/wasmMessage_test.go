// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWasmMessage(t *testing.T) {
	WasmMessage := wasmMessage{
		Type:       "testMsg",
		RawMessage: json.RawMessage(`{from:tester}`),
	}

	// Type
	require.Equal(t, "testMsg", WasmMessage.GetType())

	// GetRawMessage
	require.Equal(t, json.RawMessage(`{from:tester}`), WasmMessage.GetRawMessage())

	// Prototype
	require.Equal(t, wasmMessage{}, WasmMessagePrototype())
}
