package types

type Signatures interface {
	Get(ID) Signature

	Add(Signature) error
	Remove(Signature) error
	Mutate(Signature) error
}

type BaseSignatures struct {
	SignatureList []Signature
}

var _ Signatures = (*BaseSignatures)(nil)

func (baseSignatures BaseSignatures) Get(id ID) Signature {
	for _, signature := range baseSignatures.SignatureList {
		if signature.ID().Compare(id) == 0 {
			return signature
		}
	}
	return nil
}
func (baseSignatures *BaseSignatures) Add(signature Signature) error {
	signatureList := baseSignatures.SignatureList
	for i, oldSignature := range signatureList {
		if oldSignature.ID().Compare(signature.ID()) < 0 {
			signatureList = append(append(signatureList[:i], signature), signatureList[i+1:]...)
		}
	}
	return nil
}
func (baseSignatures *BaseSignatures) Remove(signature Signature) error {
	signatureList := baseSignatures.SignatureList
	for i, oldSignature := range signatureList {
		if oldSignature.ID().Compare(signature.ID()) == 0 {
			signatureList = append(signatureList[:i], signatureList[i+1:]...)
		}
	}
	return nil
}
func (baseSignatures *BaseSignatures) Mutate(signature Signature) error {
	signatureList := baseSignatures.SignatureList
	for i, oldSignature := range signatureList {
		if oldSignature.ID().Compare(signature.ID()) == 0 {
			signatureList[i] = signature
		}
	}
	return nil
}
