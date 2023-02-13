package helpers

import (
	"context"

	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/gogo/protobuf/proto"
)

type Genesis interface {
	proto.Message

	Default() Genesis

	ValidateBasic() error

	Import(context.Context, Mapper, ParameterList)
	Export(context.Context, Mapper, ParameterList) Genesis

	Encode(sdkCodec.JSONCodec) []byte
	Decode(sdkCodec.JSONCodec, []byte) Genesis

	Initialize([]Mappable, []Parameter) Genesis
}
