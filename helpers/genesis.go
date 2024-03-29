package helpers

import (
	"context"

	"github.com/AssetMantle/schema/go/lists"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/gogo/protobuf/proto"
)

type Genesis interface {
	proto.Message

	Default() Genesis

	ValidateBasic(ParameterManager) error

	Import(context.Context, Mapper, ParameterManager)
	Export(context.Context, Mapper, ParameterManager) Genesis

	Encode(sdkCodec.JSONCodec) []byte
	Decode(sdkCodec.JSONCodec, []byte) Genesis

	Initialize([]Record, lists.ParameterList) Genesis
}
