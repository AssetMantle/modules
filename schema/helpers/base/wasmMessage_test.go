/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
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
