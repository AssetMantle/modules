package compliance

import (
	"github.com/AssetMantle/modules/helpers"
)

type auxiliaryResponse struct {
	Compliant bool
}

var _ helpers.AuxiliaryResponse = (*auxiliaryResponse)(nil)

func newAuxiliaryResponse(compliant bool) helpers.AuxiliaryResponse {
	return auxiliaryResponse{Compliant: compliant}
}

func GetComplianceFromResponse(response helpers.AuxiliaryResponse) bool {
	return response.(auxiliaryResponse).Compliant
}
