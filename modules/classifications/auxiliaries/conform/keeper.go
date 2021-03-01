/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package conform

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
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
		return newAuxiliaryResponse(errors.EntityNotFound)
	}

	if auxiliaryRequest.ImmutableProperties != nil {
		if len(auxiliaryRequest.ImmutableProperties.GetList()) != len(classification.(mappables.Classification).GetImmutableProperties().GetList()) {
			return newAuxiliaryResponse(errors.NotAuthorized)
		}

		for _, immutableProperty := range auxiliaryRequest.ImmutableProperties.GetList() {
			if property := classification.(mappables.Classification).GetImmutableProperties().Get(immutableProperty.GetID()); property == nil || !property.GetFact().GetTypeID().Equals(immutableProperty.GetFact().GetTypeID()) || !property.GetFact().GetHashID().Equals(base.NewID("")) && property.GetFact().GetHashID() != immutableProperty.GetFact().GetHashID() {
				return newAuxiliaryResponse(errors.NotAuthorized)
			}
		}
	}

	if auxiliaryRequest.MutableProperties != nil {
		if len(auxiliaryRequest.MutableProperties.GetList()) > len(classification.(mappables.Classification).GetMutableProperties().GetList()) {
			return newAuxiliaryResponse(errors.NotAuthorized)
		}

		for _, mutableProperty := range auxiliaryRequest.MutableProperties.GetList() {
			if property := classification.(mappables.Classification).GetMutableProperties().Get(mutableProperty.GetID()); property == nil || !property.GetFact().GetTypeID().Equals(mutableProperty.GetFact().GetTypeID()) {
				return newAuxiliaryResponse(errors.NotAuthorized)
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
