package compliance

import (
	"cosmossdk.io/math"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/schema/ids"
)

type auxiliaryRequest struct {
	IdentityID      ids.IdentityID
	MinTier         math.Int
	Jurisdiction    string // empty string means any jurisdiction allowed
	RequireSanction bool
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest) Validate() error {
	return nil
}

func NewAuxiliaryRequest(identityID ids.IdentityID, minTier math.Int, jurisdiction string, requireSanction bool) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		IdentityID:      identityID,
		MinTier:         minTier,
		Jurisdiction:    jurisdiction,
		RequireSanction: requireSanction,
	}
}
