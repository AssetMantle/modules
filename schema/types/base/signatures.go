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
	//return signatures.SignatureList
	panic("Implement me")
}
func (signatures Signatures) Add(signature types.Signature) types.Signatures {
	newSignature := Signature{
		Id:             *NewID(signature.GetID().String()),
		SignatureBytes: signature.Bytes(),
		ValidityHeight: *NewHeight(signature.GetValidityHeight().Get())}
	signatures.SignatureList = append(signatures.SignatureList, newSignature)
	return &signatures
}
func (signatures Signatures) Remove(signature types.Signature) types.Signatures {
	//signatureList := signatures.SignatureList
	//for i, oldSignature := range signatureList {
	//	if oldSignature.GetID().Compare(signature.GetID()) == 0 {
	//		signatureList = append(signatureList[:i], signatureList[i+1:]...)
	//	}
	//}
	//
	//return NewSignatures(signatureList)
	panic("Implement me")
}
func (signatures Signatures) Mutate(signature types.Signature) types.Signatures {
	//signatureList := signatures.GetList()
	//for i, oldSignature := range signatureList {
	//	if oldSignature.GetID().Compare(signature.GetID()) == 0 {
	//		signatureList[i] = signature
	//	}
	//}
	//
	//return NewSignatures(signatureList)
	panic("Implement me")
}
func NewSignatures(signatureList []Signature) *Signatures {
	return &Signatures{SignatureList: signatureList}
}
