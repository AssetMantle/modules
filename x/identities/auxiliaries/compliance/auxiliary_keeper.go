package compliance

import (
	"fmt"
	"context"

	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/documents/base"
	baseTypes "github.com/AssetMantle/schema/types/base"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/identities/key"
	"github.com/AssetMantle/modules/x/identities/mappable"
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

	identityMappable := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(request.IdentityID)).GetMappable(key.NewKey(request.IdentityID))
	if identityMappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("identity with ID %s not found", request.IdentityID.AsString())
	}

	identity := base.NewIdentityFromDocument(mappable.GetIdentity(identityMappable))

	// Check compliance tier
	if identity.GetComplianceTier().LT(request.MinTier) {
		return nil, errorConstants.NotAuthorized.Wrapf(
			"identity %s has compliance tier %s, minimum required is %s",
			request.IdentityID.AsString(),
			identity.GetComplianceTier().String(),
			request.MinTier.String(),
		)
	}

	// Check jurisdiction if specified
	if request.Jurisdiction != "" && identity.GetJurisdiction() != request.Jurisdiction {
		return nil, errorConstants.NotAuthorized.Wrapf(
			"identity %s has jurisdiction %s, required jurisdiction is %s",
			request.IdentityID.AsString(),
			identity.GetJurisdiction(),
			request.Jurisdiction,
		)
	}

	// Check sanctions clearance
	if request.RequireSanction && !identity.IsSanctionsCleared() {
		return nil, errorConstants.NotAuthorized.Wrapf(
			"identity %s has not passed sanctions screening",
			request.IdentityID.AsString(),
		)
	}

	// Check accreditation expiry
	currentHeight := baseTypes.CurrentHeight(context)
	expiryHeight := identity.GetAccreditationExpiry()
	if expiryHeight.Compare(baseTypes.NewHeight(0)) > 0 && expiryHeight.Compare(currentHeight) <= 0 {
		return nil, errorConstants.NotAuthorized.Wrapf(
			"identity %s accreditation expired at height %s, current height is %s",
			request.IdentityID.AsString(),
			fmt.Sprint(expiryHeight.Get()),
			fmt.Sprint(currentHeight.Get()),
		)
	}

	// Check identity expiry
	if identity.GetExpiry().Compare(baseTypes.NewHeight(0)) > 0 && identity.GetExpiry().Compare(currentHeight) <= 0 {
		return nil, errorConstants.NotAuthorized.Wrapf(
			"identity %s expired at height %s",
			request.IdentityID.AsString(),
			fmt.Sprint(identity.GetExpiry().Get()),
		)
	}


	return newAuxiliaryResponse(true), nil
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
