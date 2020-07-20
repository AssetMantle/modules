package schema

import (
	"encoding/base64"
	"github.com/tendermint/tendermint/crypto"
)

type Signature interface {
	String() string
	Bytes() []byte

	GetID() ID

	Verify(crypto.PubKey, []byte) bool
	GetValidityHeight() Height
	HasExpired(Height) bool
}

type signature struct {
	ID             ID
	SignatureBytes []byte
	ValidityHeight Height
}

var _ Signature = (*signature)(nil)

func (baseSignature signature) String() string {
	return base64.URLEncoding.EncodeToString(baseSignature.Bytes())
}
func (baseSignature signature) Bytes() []byte { return baseSignature.SignatureBytes }
func (baseSignature signature) GetID() ID     { return baseSignature.ID }
func (baseSignature signature) Verify(pubKey crypto.PubKey, bytes []byte) bool {
	return pubKey.VerifyBytes(bytes, baseSignature.Bytes())
}
func (baseSignature signature) GetValidityHeight() Height { return baseSignature.GetValidityHeight() }
func (baseSignature signature) HasExpired(height Height) bool {
	return baseSignature.GetValidityHeight().IsGreaterThan(height)
}

func NewSignature(id ID, signatureBytes []byte, validityHeight Height) Signature {
	return signature{
		ID:             id,
		SignatureBytes: signatureBytes,
		ValidityHeight: validityHeight,
	}
}
