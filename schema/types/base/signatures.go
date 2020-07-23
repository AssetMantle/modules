package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type signatures struct {
	SignatureList []types.Signature
}

var _ types.Signatures = (*signatures)(nil)

func (signatures signatures) Get(id types.ID) types.Signature {
	for _, signature := range signatures.SignatureList {
		if signature.GetID().Compare(id) == 0 {
			return signature
		}
	}
	return nil
}
func (signatures signatures) GetList() []types.Signature {
	var signatureList []types.Signature
	for _, signature := range signatures.SignatureList {
		signatureList = append(signatureList, signature)
	}
	return signatureList
}
func (signatures signatures) Add(signature types.Signature) types.Signatures {
	signatureList := signatures.GetList()
	for i, oldSignature := range signatureList {
		if oldSignature.GetID().Compare(signature.GetID()) < 0 {
			signatureList = append(append(signatureList[:i], signature), signatureList[i+1:]...)
		}
	}
	return NewSignatures(signatureList)
}
func (signatures signatures) Remove(signature types.Signature) types.Signatures {
	signatureList := signatures.SignatureList
	for i, oldSignature := range signatureList {
		if oldSignature.GetID().Compare(signature.GetID()) == 0 {
			signatureList = append(signatureList[:i], signatureList[i+1:]...)
		}
	}
	return NewSignatures(signatureList)
}
func (signatures signatures) Mutate(signature types.Signature) types.Signatures {
	signatureList := signatures.GetList()
	for i, oldSignature := range signatureList {
		if oldSignature.GetID().Compare(signature.GetID()) == 0 {
			signatureList[i] = signature
		}
	}
	return NewSignatures(signatureList)
}
func NewSignatures(signatureList []types.Signature) types.Signatures {
	return signatures{SignatureList: signatureList}
}
