// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/types"
)

type signatures struct {
	// TODO replace with List
	SignatureList []types.Signature `json:"signatureList"`
}

var _ lists.Signatures = (*signatures)(nil)

func (signatures signatures) Get(id ids.ID) types.Signature {
	for _, signature := range signatures.SignatureList {
		if signature.GetID().Compare(id) == 0 {
			return signature
		}
	}

	return nil
}
func (signatures signatures) GetList() []types.Signature {
	return signatures.SignatureList
}
func (signatures signatures) Add(signature types.Signature) lists.Signatures {
	signatures.SignatureList = append(signatures.SignatureList, signature)
	return signatures
}
func (signatures signatures) Remove(signature types.Signature) lists.Signatures {
	signatureList := signatures.SignatureList
	for i, oldSignature := range signatureList {
		if oldSignature.GetID().Compare(signature.GetID()) == 0 {
			signatureList = append(signatureList[:i], signatureList[i+1:]...)
		}
	}

	return NewSignatures(signatureList)
}
func (signatures signatures) Mutate(signature types.Signature) lists.Signatures {
	signatureList := signatures.GetList()
	for i, oldSignature := range signatureList {
		if oldSignature.GetID().Compare(signature.GetID()) == 0 {
			signatureList[i] = signature
		}
	}

	return NewSignatures(signatureList)
}

func NewSignatures(signatureList []types.Signature) lists.Signatures {
	return signatures{SignatureList: signatureList}
}
