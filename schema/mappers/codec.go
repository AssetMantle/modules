package mappers

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*Chains)(nil), nil)
	codec.RegisterInterface((*InterIdentities)(nil), nil)
	codec.RegisterInterface((*InterNFTs)(nil), nil)
	codec.RegisterInterface((*Maintainers)(nil), nil)
	codec.RegisterInterface((*Orders)(nil), nil)
	codec.RegisterInterface((*Splits)(nil), nil)
}
