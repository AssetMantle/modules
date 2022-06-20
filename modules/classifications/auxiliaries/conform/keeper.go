// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package conform

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.FromID(auxiliaryRequest.ClassificationID))

	classification := classifications.Get(key.FromID(auxiliaryRequest.ClassificationID))
	if classification == nil {
		return newAuxiliaryResponse(constants.EntityNotFound)
	}

	if auxiliaryRequest.ImmutableProperties != nil {
		if len(auxiliaryRequest.ImmutableProperties.GetList()) != len(classification.(mappables.Classification).GetImmutablePropertyList().GetList()) {
			return newAuxiliaryResponse(constants.NotAuthorized)
		}

		for _, immutableProperty := range auxiliaryRequest.ImmutableProperties.GetList() {
			if property := classification.(mappables.Classification).GetImmutablePropertyList().GetProperty(immutableProperty.GetID()); property == nil || property.GetType().Compare(immutableProperty.GetType()) != 0 || property.GetHash().Compare(baseIDs.NewID("")) != 0 && property.GetHash() != immutableProperty.GetHash() {
				return newAuxiliaryResponse(constants.NotAuthorized)
			}
		}
	}

	if auxiliaryRequest.MutableProperties != nil {
		if len(auxiliaryRequest.MutableProperties.GetList()) > len(classification.(mappables.Classification).GetMutablePropertyList().GetList()) {
			return newAuxiliaryResponse(constants.NotAuthorized)
		}

		for _, mutableProperty := range auxiliaryRequest.MutableProperties.GetList() {
			if property := classification.(mappables.Classification).GetMutablePropertyList().GetProperty(mutableProperty.GetID()); property == nil || property.GetType().Compare(mutableProperty.GetType()) != 0 {
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
