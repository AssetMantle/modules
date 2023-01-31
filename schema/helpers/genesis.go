package helpers

import (
	"context"

	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/gogo/protobuf/proto"
)

type Genesis interface {
	proto.Message

	GetParameterList() []Parameter
	GetMappableList() []Mappable
	Default() Genesis

	Validate() error

	Import(context.Context, Mapper)
	Export(context.Context, Mapper) Genesis

	Encode(sdkCodec.JSONCodec) []byte
	Decode(sdkCodec.JSONCodec, []byte) Genesis

	Initialize(mappableList []Mappable, parameterList []Parameter) Genesis
}
