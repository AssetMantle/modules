package types

import (
	"encoding/base64"
	"github.com/tendermint/tendermint/crypto"
)

type Signature interface {
	String() string
	Bytes() []byte

	ID() ID

	Verify(crypto.PubKey, []byte) bool
	HasExpired(Height) bool
}

type BaseSignature struct {
	SignatureID    ID
	SignatureBytes []byte
	ValidityHeight Height
}

var _ Signature = (*BaseSignature)(nil)

func (baseSignature BaseSignature) String() string {
	return base64.URLEncoding.EncodeToString(baseSignature.Bytes())
}
func (baseSignature BaseSignature) Bytes() []byte { return baseSignature.SignatureBytes }
func (baseSignature BaseSignature) ID() ID        { return baseSignature.SignatureID }
func (baseSignature BaseSignature) Verify(pubKey crypto.PubKey, bytes []byte) bool {
	return pubKey.VerifyBytes(bytes, baseSignature.Bytes())
}
func (baseSignature BaseSignature) HasExpired(height Height) bool {
	return baseSignature.ValidityHeight.IsGraterThat(height)
}
