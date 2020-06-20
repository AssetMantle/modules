package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*Chain)(nil), nil)
	codec.RegisterInterface((*Chains)(nil), nil)
	codec.RegisterInterface((*Classification)(nil), nil)
	codec.RegisterInterface((*Fact)(nil), nil)
	codec.RegisterInterface((*Height)(nil), nil)
	codec.RegisterInterface((*ID)(nil), nil)
	codec.RegisterInterface((*InterNFT)(nil), nil)
	codec.RegisterInterface((*InterNFTs)(nil), nil)
	codec.RegisterInterface((*InterNFTWallet)(nil), nil)
	codec.RegisterInterface((*Maintainer)(nil), nil)
	codec.RegisterInterface((*Maintainers)(nil), nil)
	codec.RegisterInterface((*NFT)(nil), nil)
	codec.RegisterInterface((*NFTWallet)(nil), nil)
	codec.RegisterInterface((*Properties)(nil), nil)
	codec.RegisterInterface((*Property)(nil), nil)
	codec.RegisterInterface((*Query)(nil), nil)
	codec.RegisterInterface((*Request)(nil), nil)
	codec.RegisterInterface((*Share)(nil), nil)
	codec.RegisterInterface((*Signature)(nil), nil)
	codec.RegisterInterface((*Signatures)(nil), nil)
	codec.RegisterInterface((*Trait)(nil), nil)
	codec.RegisterInterface((*Traits)(nil), nil)

	codec.RegisterConcrete(&BaseFact{}, "xprt/fact", nil)
	codec.RegisterConcrete(&BaseHeight{}, "xprt/height", nil)
	codec.RegisterConcrete(&BaseID{}, "xprt/id", nil)
	codec.RegisterConcrete(&BaseProperties{}, "xprt/properties", nil)
	codec.RegisterConcrete(&BaseProperty{}, "xprt/property", nil)
	codec.RegisterConcrete(&BaseSignature{}, "xprt/signature", nil)
	codec.RegisterConcrete(&BaseSignatures{}, "xprt/signatures", nil)
}
