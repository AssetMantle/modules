/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package define

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/mappable"
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

	classificationID := key.NewClassificationID(base.NewID(context.ChainID()), auxiliaryRequest.ImmutableTraits, auxiliaryRequest.MutableTraits)
	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.New(classificationID))
	if classifications.Get(key.New(classificationID)) != nil {
		return newAuxiliaryResponse(base.NewID(classificationID.String()), errors.EntityAlreadyExists)
	}

	classifications = classifications.Add(mappable.NewClassification(classificationID, auxiliaryRequest.ImmutableTraits, auxiliaryRequest.MutableTraits))
	return newAuxiliaryResponse(base.NewID(classificationID.String()), nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
