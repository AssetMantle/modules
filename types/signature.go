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
	ValidityHeight() Height
	HasExpired(Height) bool
}

type BaseSignature struct {
	BaseID             BaseID
	BaseBytes          []byte
	ValidityBaseHeight BaseHeight
}

var _ Signature = (*BaseSignature)(nil)

func (baseSignature BaseSignature) String() string {
	return base64.URLEncoding.EncodeToString(baseSignature.Bytes())
}
func (baseSignature BaseSignature) Bytes() []byte { return baseSignature.BaseBytes }
func (baseSignature BaseSignature) ID() ID        { return baseSignature.BaseID }
func (baseSignature BaseSignature) Verify(pubKey crypto.PubKey, bytes []byte) bool {
	return pubKey.VerifyBytes(bytes, baseSignature.Bytes())
}
func (baseSignature BaseSignature) ValidityHeight() Height { return baseSignature.ValidityBaseHeight }
func (baseSignature BaseSignature) HasExpired(height Height) bool {
	return baseSignature.ValidityBaseHeight.IsGraterThat(height)
}
