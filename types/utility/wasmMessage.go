package utility

import "encoding/json"

type WasmMessage interface {
	Type() string
	Encode() []byte
	Decode() WasmMessage
}
type wasmMessage struct {
	MsgType string          `json:"msgtype,required"`
	Raw     json.RawMessage `json:"raw,omitempty"`
}

var _ WasmMessage = (*wasmMessage)(nil)

func (wasmMessage wasmMessage) Type() string        { return wasmMessage.MsgType }
func (wasmMessage wasmMessage) Encode() []byte      { return nil }
func (wasmMessage wasmMessage) Decode() WasmMessage { return nil }
