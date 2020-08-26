/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package define

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)

	if len(auxiliaryRequest.ImmutableTraits.Get().GetList())+len(auxiliaryRequest.MutableTraits.Get().GetList()) > constants.MaxTraitCount {
		return newAuxiliaryResponse(nil, errors.InvalidRequest)
	}

	classificationID := mapper.NewClassificationID(base.NewID(context.ChainID()), auxiliaryRequest.ImmutableTraits, auxiliaryRequest.MutableTraits)
	classifications := mapper.NewClassifications(auxiliaryKeeper.mapper, context).Fetch(classificationID)
	if classifications.Get(classificationID) != nil {
		return newAuxiliaryResponse(classificationID, errors.EntityAlreadyExists)
	}

	classifications = classifications.Add(mapper.NewClassification(classificationID, auxiliaryRequest.ImmutableTraits, auxiliaryRequest.MutableTraits))
	return newAuxiliaryResponse(classificationID, nil)
}

func initializeAuxiliaryKeeper(mapper helpers.Mapper, _ []interface{}) helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{mapper: mapper}
}
