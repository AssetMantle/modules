/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type signatures struct {
	SignatureList []types.Signature `json:"signatureList"`
}

var _ types.Signatures = (*signatures)(nil)

func (signatures signatures) Get(id types.ID) types.Signature {
	for _, signature := range signatures.SignatureList {
		if signature.GetID().Equals(id) {
			return signature
		}
	}

	return nil
}
func (signatures signatures) GetList() []types.Signature {
	return signatures.SignatureList
}
func (signatures signatures) Add(signature types.Signature) types.Signatures {
	signatures.SignatureList = append(signatures.SignatureList, signature)
	return signatures
}
func (signatures signatures) Remove(signature types.Signature) types.Signatures {
	signatureList := signatures.SignatureList
	for i, oldSignature := range signatureList {
		if oldSignature.GetID().Equals(signature.GetID()) {
			signatureList = append(signatureList[:i], signatureList[i+1:]...)
		}
	}

	return NewSignatures(signatureList)
}
func (signatures signatures) Mutate(signature types.Signature) types.Signatures {
	signatureList := signatures.GetList()
	for i, oldSignature := range signatureList {
		if oldSignature.GetID().Equals(signature.GetID()) {
			signatureList[i] = signature
		}
	}

	return NewSignatures(signatureList)
}
func NewSignatures(signatureList []types.Signature) types.Signatures {
	return signatures{SignatureList: signatureList}
}
