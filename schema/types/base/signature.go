// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/base64"

	"github.com/tendermint/tendermint/crypto"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/types"
)

type signature struct {
	ID             ids.ID       `json:"id"`
	SignatureBytes []byte       `json:"signatureBytes"`
	ValidityHeight types.Height `json:"validityHeight"`
}

var _ types.Signature = (*signature)(nil)

func (baseSignature signature) String() string {
	return base64.URLEncoding.EncodeToString(baseSignature.Bytes())
}
func (baseSignature signature) Bytes() []byte { return baseSignature.SignatureBytes }
func (baseSignature signature) GetID() ids.ID { return baseSignature.ID }
func (baseSignature signature) Verify(pubKey crypto.PubKey, bytes []byte) bool {
	return pubKey.VerifyBytes(bytes, baseSignature.Bytes())
}
func (baseSignature signature) GetValidityHeight() types.Height {
	return baseSignature.ValidityHeight
}
func (baseSignature signature) HasExpired(height types.Height) bool {
	return baseSignature.GetValidityHeight().Compare(height) > 0
}

func NewSignature(id ids.ID, signatureBytes []byte, validityHeight types.Height) types.Signature {
	return signature{
		ID:             id,
		SignatureBytes: signatureBytes,
		ValidityHeight: validityHeight,
	}
}
