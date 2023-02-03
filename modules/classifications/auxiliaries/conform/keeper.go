// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package conform

import (
	"context"
	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/modules/classifications/internal/mappable"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(auxiliaryRequest.ClassificationID))

	Mappable := classifications.Get(key.NewKey(auxiliaryRequest.ClassificationID))
	if Mappable == nil {
		return newAuxiliaryResponse(errorConstants.EntityNotFound)
	}
	classification := mappable.GetClassification(Mappable)

	if auxiliaryRequest.Immutables != nil {
		//if len(auxiliaryRequest.Immutables.GetImmutablePropertyList().GetList()) != len(classification.GetImmutables().GetImmutablePropertyList().GetList()) {
		//	return newAuxiliaryResponse(errorConstants.IncorrectFormat)
		//}

		for _, immutableProperty := range classification.GetImmutables().GetImmutablePropertyList().GetList() {
			if immutableProperty.Get().IsMeta() && immutableProperty.Get().GetID().AsString() == "BondingAmount.S" {
				continue
			}
			if property := auxiliaryRequest.Immutables.GetImmutablePropertyList().GetProperty(immutableProperty.GetID()); property == nil || immutableProperty.GetDataID().GetHashID().Compare(baseIDs.GenerateHashID()) != 0 && property.GetDataID().GetHashID().Compare(immutableProperty.GetDataID().GetHashID()) != 0 {
				return newAuxiliaryResponse(errorConstants.IncorrectFormat)
			}
		}
	}

	if auxiliaryRequest.Mutables != nil {
		if len(auxiliaryRequest.Mutables.GetMutablePropertyList().GetList()) != len(classification.GetMutables().GetMutablePropertyList().GetList()) {
			return newAuxiliaryResponse(errorConstants.IncorrectFormat)
		}

		for _, mutableProperty := range classification.GetMutables().GetMutablePropertyList().GetList() {
			if property := auxiliaryRequest.Mutables.GetMutablePropertyList().GetProperty(mutableProperty.GetID()); property == nil {
				return newAuxiliaryResponse(errorConstants.IncorrectFormat)
			}
		}
	}

	return newAuxiliaryResponse(nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterList, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
