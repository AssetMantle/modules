// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package member

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/types"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(auxiliaryRequest.ClassificationID))

	Mappable := classifications.Get(key.NewKey(auxiliaryRequest.ClassificationID))
	if Mappable == nil {
		return newAuxiliaryResponse(errorConstants.EntityNotFound)
	}
	classification := Mappable.(types.Classification)

	if auxiliaryRequest.Immutables != nil {
		if len(auxiliaryRequest.Immutables.GetImmutablePropertyList().GetList()) > len(classification.GetImmutables().GetImmutablePropertyList().GetList()) {
			return newAuxiliaryResponse(errorConstants.IncorrectFormat)
		}

		for _, immutableProperty := range auxiliaryRequest.Immutables.GetImmutablePropertyList().GetList() {
			if property := classification.GetImmutables().GetImmutablePropertyList().GetProperty(immutableProperty.GetID()); property == nil || immutableProperty.GetDataID().GetHashID().Compare(baseIDs.GenerateHashID()) != 0 && property.GetDataID().GetHashID().Compare(immutableProperty.GetDataID().GetHashID()) != 0 {
				return newAuxiliaryResponse(errorConstants.IncorrectFormat)
			}
		}
	}

	if auxiliaryRequest.Mutables != nil {
		if len(auxiliaryRequest.Mutables.GetMutablePropertyList().GetList()) > len(classification.GetMutables().GetMutablePropertyList().GetList()) {
			return newAuxiliaryResponse(errorConstants.IncorrectFormat)
		}

		for _, mutableProperty := range auxiliaryRequest.Mutables.GetMutablePropertyList().GetList() {
			if property := classification.GetMutables().GetMutablePropertyList().GetProperty(mutableProperty.GetID()); property == nil {
				return newAuxiliaryResponse(errorConstants.IncorrectFormat)
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
