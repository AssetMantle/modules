/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"

	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type wasmMessage struct {
	Type       string          `json:"msgtype"`
	RawMessage json.RawMessage `json:"raw,omitempty"`
}

var _ helpers.WasmMessage = (*wasmMessage)(nil)

func (wasmMessage wasmMessage) GetType() string                { return wasmMessage.Type }
func (wasmMessage wasmMessage) GetRawMessage() json.RawMessage { return wasmMessage.RawMessage }

var WasmMessagePrototype helpers.WasmMessagePrototype = func() helpers.WasmMessage {
	return wasmMessage{}
}
