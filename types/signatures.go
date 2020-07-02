package types

type Signatures interface {
	Get(ID) Signature

	GetList() []Signature

	Add(Signature) Signatures
	Remove(Signature) Signatures
	Mutate(Signature) Signatures
}
type signatures struct {
	SignatureList []Signature
}

var _ Signatures = (*signatures)(nil)

func (signatures signatures) Get(id ID) Signature {
	for _, signature := range signatures.SignatureList {
		if signature.GetID().Compare(id) == 0 {
			return signature
		}
	}
	return nil
}
func (signatures signatures) GetList() []Signature {
	var signatureList []Signature
	for _, signature := range signatures.SignatureList {
		signatureList = append(signatureList, signature)
	}
	return signatureList
}
func (signatures signatures) Add(signature Signature) Signatures {
	signatureList := signatures.GetList()
	for i, oldSignature := range signatureList {
		if oldSignature.GetID().Compare(signature.GetID()) < 0 {
			signatureList = append(append(signatureList[:i], signature), signatureList[i+1:]...)
		}
	}
	return NewSignatures(signatureList)
}
func (signatures signatures) Remove(signature Signature) Signatures {
	signatureList := signatures.SignatureList
	for i, oldSignature := range signatureList {
		if oldSignature.GetID().Compare(signature.GetID()) == 0 {
			signatureList = append(signatureList[:i], signatureList[i+1:]...)
		}
	}
	return NewSignatures(signatureList)
}
func (signatures signatures) Mutate(signature Signature) Signatures {
	signatureList := signatures.GetList()
	for i, oldSignature := range signatureList {
		if oldSignature.GetID().Compare(signature.GetID()) == 0 {
			signatureList[i] = signature
		}
	}
	return NewSignatures(signatureList)
}
func NewSignatures(signatureList []Signature) Signatures {
	return signatures{SignatureList: signatureList}
}
