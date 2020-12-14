/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package conform

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)
	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.New(auxiliaryRequest.ClassificationID))
	classification := classifications.Get(key.New(auxiliaryRequest.ClassificationID))
	if classification == nil {
		return newAuxiliaryResponse(errors.EntityNotFound)
	}
	if auxiliaryRequest.Immutables != nil {
		if len(auxiliaryRequest.Immutables.Get().GetList()) != len(classification.(mappables.Classification).GetImmutables().Get().GetList()) {
			return newAuxiliaryResponse(errors.NotAuthorized)
		}
		for _, immutableProperty := range auxiliaryRequest.Immutables.Get().GetList() {
			if trait := classification.(mappables.Classification).GetImmutables().Get().Get(immutableProperty.GetID()); trait == nil || trait.GetFact().GetHash() != trait.GetFact().GetType()+constants.DataTypeAndValueSeparator && trait.GetFact().GetHash() != immutableProperty.GetFact().GetHash() {
				return newAuxiliaryResponse(errors.NotAuthorized)
			}
		}
	}
	if auxiliaryRequest.Mutables != nil {
		if len(auxiliaryRequest.Mutables.Get().GetList()) > len(classification.(mappables.Classification).GetMutables().Get().GetList()) {
			return newAuxiliaryResponse(errors.NotAuthorized)
		}
		for _, mutableProperty := range auxiliaryRequest.Mutables.Get().GetList() {
			if classification.(mappables.Classification).GetMutables().Get().Get(mutableProperty.GetID()) == nil {
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
