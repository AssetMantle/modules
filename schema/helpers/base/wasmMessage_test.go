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
