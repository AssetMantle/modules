package types

type Signatures interface {
	Get(ID) Signature

	SignatureList() []Signature
	Add(Signature) error
	Remove(Signature) error
	Mutate(Signature) error
}
type BaseSignatures struct {
	BaseSignatureList []BaseSignature
}

var _ Signatures = (*BaseSignatures)(nil)

func (baseSignatures BaseSignatures) Get(id ID) Signature {
	for _, signature := range baseSignatures.BaseSignatureList {
		if signature.ID().Compare(id) == 0 {
			return signature
		}
	}
	return nil
}
func (baseSignatures *BaseSignatures) SignatureList() []Signature {

	var signatureList []Signature
	for _, baseSignature := range baseSignatures.BaseSignatureList {
		signatureList = append(signatureList, &baseSignature)
	}
	return signatureList
}
func (baseSignatures *BaseSignatures) Add(signature Signature) error {
	signatureList := baseSignatures.BaseSignatureList
	for i, oldSignature := range signatureList {
		if oldSignature.ID().Compare(signature.ID()) < 0 {
			signatureList = append(append(signatureList[:i], BaseSignature{
				BaseID:             BaseID{signature.ID().String()},
				BaseBytes:          signature.Bytes(),
				ValidityBaseHeight: BaseHeight{signature.ValidityHeight().Count()},
			}), signatureList[i+1:]...)
		}
	}
	return nil
}
func (baseSignatures *BaseSignatures) Remove(signature Signature) error {
	signatureList := baseSignatures.BaseSignatureList
	for i, oldSignature := range signatureList {
		if oldSignature.ID().Compare(signature.ID()) == 0 {
			signatureList = append(signatureList[:i], signatureList[i+1:]...)
		}
	}
	return nil
}
func (baseSignatures *BaseSignatures) Mutate(signature Signature) error {
	signatureList := baseSignatures.BaseSignatureList
	for i, oldSignature := range signatureList {
		if oldSignature.ID().Compare(signature.ID()) == 0 {
			signatureList[i] = BaseSignature{
				BaseID:             BaseID{signature.ID().String()},
				BaseBytes:          signature.Bytes(),
				ValidityBaseHeight: BaseHeight{signature.ValidityHeight().Count()},
			}
		}
	}
	return nil
}
func BaseSignaturesFromInterface(signatures Signatures) BaseSignatures {
	var baseSignatureList []BaseSignature
	for _, signature := range signatures.SignatureList() {
		baseSignatureList = append(baseSignatureList, BaseSignature{BaseID: BaseID{signature.String()}, BaseBytes: signature.Bytes(), ValidityBaseHeight: BaseHeight{signature.ValidityHeight().Count()}})
	}
	return BaseSignatures{BaseSignatureList: baseSignatureList}
}
