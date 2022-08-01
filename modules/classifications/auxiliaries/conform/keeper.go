// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package conform

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(auxiliaryRequest.ClassificationID))

	classification := classifications.Get(key.NewKey(auxiliaryRequest.ClassificationID))
	if classification == nil {
		return newAuxiliaryResponse(constants.EntityNotFound)
	}

	if auxiliaryRequest.Immutables != nil {
		if len(auxiliaryRequest.Immutables.GetImmutablePropertyList().GetList()) != len(classification.(mappables.Classification).GetImmutables().GetImmutablePropertyList().GetList()) {
			return newAuxiliaryResponse(constants.NotAuthorized)
		}

		for _, immutableProperty := range auxiliaryRequest.Immutables.GetImmutablePropertyList().GetList() {
			if property := classification.(mappables.Classification).GetImmutables().GetImmutablePropertyList().GetProperty(immutableProperty.GetID()); property == nil || property.GetType().Compare(immutableProperty.GetType()) != 0 || property.GetDataID().GetHashID().Compare(baseIDs.GenerateHashID()) != 0 && property.GetDataID().GetHashID() != immutableProperty.GetDataID().GetHashID() {
				return newAuxiliaryResponse(constants.NotAuthorized)
			}
		}
	}

	if auxiliaryRequest.Mutables != nil {
		if len(auxiliaryRequest.Mutables.GetMutablePropertyList().GetList()) > len(classification.(mappables.Classification).GetMutables().GetMutablePropertyList().GetList()) {
			return newAuxiliaryResponse(constants.NotAuthorized)
		}

		for _, mutableProperty := range auxiliaryRequest.Mutables.GetMutablePropertyList().GetList() {
			if property := classification.(mappables.Classification).GetMutables().GetMutablePropertyList().GetProperty(mutableProperty.GetID()); property == nil || property.GetType().Compare(mutableProperty.GetType()) != 0 {
				return newAuxiliaryResponse(constants.NotAuthorized)
			}
		}
	}

	return newAuxiliaryResponse(nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
