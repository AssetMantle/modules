package distribute

import (
	"github.com/AssetMantle/modules/helpers"
)

type auxiliaryResponse struct {
	HolderCount     int
	TotalDistributed string
}

var _ helpers.AuxiliaryResponse = (*auxiliaryResponse)(nil)

func newAuxiliaryResponse(holderCount int, totalDistributed string) helpers.AuxiliaryResponse {
	return auxiliaryResponse{HolderCount: holderCount, TotalDistributed: totalDistributed}
}
