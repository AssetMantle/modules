package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*InterNFT)(nil), nil)
	codec.RegisterInterface((*InterNFTs)(nil), nil)
	codec.RegisterInterface((*Classification)(nil), nil)
	codec.RegisterInterface((*Fact)(nil), nil)
	codec.RegisterInterface((*Height)(nil), nil)
	codec.RegisterInterface((*ID)(nil), nil)
	codec.RegisterInterface((*Maintainer)(nil), nil)
	codec.RegisterInterface((*Maintainers)(nil), nil)
	codec.RegisterInterface((*Properties)(nil), nil)
	codec.RegisterInterface((*Property)(nil), nil)
	codec.RegisterInterface((*Share)(nil), nil)
	codec.RegisterInterface((*Signature)(nil), nil)
	codec.RegisterInterface((*Signatures)(nil), nil)
	codec.RegisterInterface((*Trait)(nil), nil)
	codec.RegisterInterface((*Traits)(nil), nil)
}
