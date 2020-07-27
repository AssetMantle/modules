package base

import (
	"encoding/json"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type wasmMessage struct {
	MsgType string          `json:"msgtype,required"`
	Raw     json.RawMessage `json:"raw,omitempty"`
}

var _ helpers.WasmMessage = (*wasmMessage)(nil)

func (wasmMessage wasmMessage) Type() string                { return wasmMessage.MsgType }
func (wasmMessage wasmMessage) Encode() []byte              { return nil }
func (wasmMessage wasmMessage) Decode() helpers.WasmMessage { return nil }
