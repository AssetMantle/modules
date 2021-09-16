/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ types.Signatures = (*Signatures)(nil)

func (signatures Signatures) Get(id types.ID) types.Signature {
	for _, signature := range signatures.SignatureList {
		if signature.GetID().Compare(id) == 0 {
			return &signature
		}
	}

	return nil
}
func (signatures Signatures) GetList() []types.Signature {
	newSignatureList := make([]types.Signature, len(signatures.SignatureList))
	for i, element := range signatures.SignatureList {
		newSignatureList[i] = &element
	}
	return newSignatureList
}
func (signatures Signatures) Add(signature types.Signature) types.Signatures {
	signatures.SignatureList = append(signatures.SignatureList, *NewSignature(signature.GetID(), signature.Bytes(), signature.GetValidityHeight()))
	return &signatures
}
func (signatures Signatures) Remove(signature types.Signature) types.Signatures {
	signatureList := signatures.GetList()
	for i, oldSignature := range signatureList {
		if oldSignature.GetID().Compare(signature.GetID()) == 0 {
			signatureList = append(signatureList[:i], signatureList[i+1:]...)
		}
	}

	return NewSignatures(signatureList)
}
func (signatures Signatures) Mutate(signature types.Signature) types.Signatures {
	signatureList := signatures.GetList()
	for i, oldSignature := range signatureList {
		if oldSignature.GetID().Compare(signature.GetID()) == 0 {
			signatureList[i] = signature
		}
	}

	return NewSignatures(signatureList)
}
func NewSignatures(signatureList []types.Signature) *Signatures {
	newSignatureList := make([]Signature, len(signatureList))
	for i, element := range signatureList {
		newSignatureList[i] = *NewSignature(element.GetID(), element.Bytes(), element.GetValidityHeight())
	}
	return &Signatures{
		SignatureList: newSignatureList,
	}
}
