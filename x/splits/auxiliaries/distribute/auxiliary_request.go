package distribute

import (
	"cosmossdk.io/math"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/schema/ids"
)

type auxiliaryRequest struct {
	AssetID          ids.AssetID    // the asset whose holders receive the distribution
	DistributionID   ids.AssetID    // the asset being distributed (e.g., USDC split ID)
	FromID           ids.IdentityID // identity funding the distribution
	TotalAmount      math.Int       // total amount to distribute
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest) Validate() error {
	return nil
}

func NewAuxiliaryRequest(assetID ids.AssetID, distributionID ids.AssetID, fromID ids.IdentityID, totalAmount math.Int) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		AssetID:        assetID,
		DistributionID: distributionID,
		FromID:         fromID,
		TotalAmount:    totalAmount,
	}
}
