package base

import (
	"encoding/base64"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/tendermint/tendermint/crypto"
)

type signature struct {
	ID             types.ID
	SignatureBytes []byte
	ValidityHeight types.Height
}

var _ types.Signature = (*signature)(nil)

func (baseSignature signature) String() string {
	return base64.URLEncoding.EncodeToString(baseSignature.Bytes())
}
func (baseSignature signature) Bytes() []byte   { return baseSignature.SignatureBytes }
func (baseSignature signature) GetID() types.ID { return baseSignature.ID }
func (baseSignature signature) Verify(pubKey crypto.PubKey, bytes []byte) bool {
	return pubKey.VerifyBytes(bytes, baseSignature.Bytes())
}
func (baseSignature signature) GetValidityHeight() types.Height {
	return baseSignature.GetValidityHeight()
}
func (baseSignature signature) HasExpired(height types.Height) bool {
	return baseSignature.GetValidityHeight().IsGreaterThan(height)
}

func NewSignature(id types.ID, signatureBytes []byte, validityHeight types.Height) types.Signature {
	return signature{
		ID:             id,
		SignatureBytes: signatureBytes,
		ValidityHeight: validityHeight,
	}
}
