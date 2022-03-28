/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/base64"

	"github.com/tendermint/tendermint/crypto"

	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ types.Signature = (*Signature)(nil)

func (baseSignature Signature) String() string {
	return base64.URLEncoding.EncodeToString(baseSignature.Bytes())
}
func (baseSignature Signature) Bytes() []byte   { return baseSignature.SignatureBytes }
func (baseSignature Signature) GetID() types.ID { return &baseSignature.Id }
func (baseSignature Signature) Verify(pubKey crypto.PubKey, bytes []byte) bool {
	return pubKey.VerifySignature(bytes, baseSignature.Bytes())
}
func (baseSignature Signature) GetValidityHeight() types.Height {
	return &baseSignature.ValidityHeight
}
func (baseSignature Signature) HasExpired(height types.Height) bool {
	return baseSignature.GetValidityHeight().Compare(height) > 0
}

func NewSignature(id types.ID, signatureBytes []byte, validityHeight types.Height) *Signature {
	return &Signature{
		Id:             *NewID(id.String()),
		SignatureBytes: signatureBytes,
		ValidityHeight: *NewHeight(validityHeight.Get()),
	}
}
