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
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type auxiliaryKeeper struct {
	mapper         helpers.Mapper
	scrubAuxiliary helpers.Auxiliary
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)

	scrubImmutableMetaTraitsAuxiliaryResponse, Error := scrub.ValidateResponse(auxiliaryKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(auxiliaryRequest.ImmutableMetaTraits.GetMetaPropertyList()...)))
	if Error != nil {
		return newAuxiliaryResponse(Error)
	}
	immutableTraits := base.NewImmutables(base.NewProperties(append(scrubImmutableMetaTraitsAuxiliaryResponse.Properties.GetList(), auxiliaryRequest.ImmutableTraits.GetList()...)))

	scrubMutableMetaTraitsAuxiliaryResponse, Error := scrub.ValidateResponse(auxiliaryKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(auxiliaryRequest.MutableMetaTraits.GetMetaPropertyList()...)))
	if Error != nil {
		return newAuxiliaryResponse(Error)
	}
	mutableTraits := base.NewMutables(base.NewProperties(append(scrubMutableMetaTraitsAuxiliaryResponse.Properties.GetList(), auxiliaryRequest.MutableTraits.GetList()...)))

	if len(immutableTraits.Get().GetList())+len(mutableTraits.Get().GetList()) > constants.MaxTraitCount {
		return newAuxiliaryResponse(errors.NotAuthorized)
	}

	classificationID := mapper.NewClassificationID(base.NewID(context.ChainID()), immutableTraits, mutableTraits)
	classifications := mapper.NewClassifications(auxiliaryKeeper.mapper, context).Fetch(classificationID)
	if classifications.Get(classificationID) != nil {
		return newAuxiliaryResponse(errors.EntityAlreadyExists)
	}

	classifications = classifications.Add(mapper.NewClassification(classificationID, immutableTraits, mutableTraits))
	return newAuxiliaryResponse(nil)
}

func initializeAuxiliaryKeeper(mapper helpers.Mapper, auxiliaries []interface{}) helpers.AuxiliaryKeeper {
	transactionKeeper := auxiliaryKeeper{mapper: mapper}
	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case scrub.Auxiliary.GetName():
				transactionKeeper.scrubAuxiliary = value
			}
		}
	}
	return transactionKeeper
}
