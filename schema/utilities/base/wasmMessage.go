package base

import (
	"encoding/json"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type wasmMessage struct {
	MsgType string          `json:"msgtype,required"`
	Raw     json.RawMessage `json:"raw,omitempty"`
}

var _ utilities.WasmMessage = (*wasmMessage)(nil)

func (wasmMessage wasmMessage) Type() string                  { return wasmMessage.MsgType }
func (wasmMessage wasmMessage) Encode() []byte                { return nil }
func (wasmMessage wasmMessage) Decode() utilities.WasmMessage { return nil }
