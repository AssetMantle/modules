package distribute

import (
	"context"

	"cosmossdk.io/math"
	"github.com/AssetMantle/modules/helpers"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/splits/key"
	"github.com/AssetMantle/modules/x/splits/mappable"
	"github.com/AssetMantle/modules/x/splits/utilities"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
)

type auxiliaryKeeper struct {
	mapper           helpers.Mapper
	parameterManager helpers.ParameterManager
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, auxiliaryRequestI helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	request, ok := auxiliaryRequestI.(auxiliaryRequest)
	if !ok {
		return nil, errorConstants.InvalidRequest.Wrapf("invalid request type %T", auxiliaryRequestI)
	}

	if request.TotalAmount.LTE(math.ZeroInt()) {
		return nil, errorConstants.InvalidParameter.Wrapf("distribution amount must be positive")
	}

	splits := auxiliaryKeeper.mapper.NewCollection(context)

	// Calculate total supply of the asset being distributed TO holders of
	totalSupply := math.ZeroInt()
	type holder struct {
		ownerID ids.IdentityID
		balance math.Int
	}
	var holders []holder

	// Iterate all splits for the target asset to build holder list and total supply
	// Key prefix is the asset ID bytes, so this iterates all owners of this asset
	splits.Iterate(key.NewKey(baseIDs.NewSplitID(request.AssetID, baseIDs.PrototypeIdentityID())), func(record helpers.Record) bool {
		split := mappable.GetSplit(record.GetMappable())
		recordKey := record.GetKey().(*key.Key)
		ownerID := recordKey.SplitID.GetOwnerID()
		if split.GetValue().GT(math.ZeroInt()) {
			totalSupply = totalSupply.Add(split.GetValue())
			holders = append(holders, holder{
				ownerID: ownerID,
				balance: split.GetValue(),
			})
		}
		return false
	})

	if totalSupply.IsZero() {
		return nil, errorConstants.EntityNotFound.Wrapf("no holders found for asset %s", request.AssetID.AsString())
	}

	// Subtract total from funder
	if _, err := utilities.SubtractSplits(splits, request.FromID, request.DistributionID, request.TotalAmount); err != nil {
		return nil, errorConstants.InsufficientBalance.Wrapf("funder does not have enough of distribution asset: %s", err.Error())
	}

	// Distribute pro-rata to each holder
	distributed := math.ZeroInt()
	for i, h := range holders {
		var share math.Int
		if i == len(holders)-1 {
			// Last holder gets remainder to avoid rounding dust
			share = request.TotalAmount.Sub(distributed)
		} else {
			// Pro-rata: share = totalAmount * holderBalance / totalSupply
			share = request.TotalAmount.Mul(h.balance).Quo(totalSupply)
		}

		if share.GT(math.ZeroInt()) {
			if _, err := utilities.AddSplits(splits, h.ownerID, request.DistributionID, share); err != nil {
				return nil, err
			}
			distributed = distributed.Add(share)
		}
	}

	return newAuxiliaryResponse(len(holders), distributed.String()), nil
}

func (auxiliaryKeeper auxiliaryKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	auxiliaryKeeper.mapper = mapper
	auxiliaryKeeper.parameterManager = parameterManager

	helpers.PanicOnUninitializedKeeperFields(auxiliaryKeeper)
	return auxiliaryKeeper
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
