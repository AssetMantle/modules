// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package member

import (
	"context"

	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/key"
	"github.com/AssetMantle/modules/x/classifications/mappable"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(auxiliaryRequest.ClassificationID))

	Mappable := classifications.GetMappable(key.NewKey(auxiliaryRequest.ClassificationID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("classification with ID %s not found", auxiliaryRequest.ClassificationID.AsString())
	}
	classification := mappable.GetClassification(Mappable)

	if auxiliaryRequest.Immutables != nil {
		if len(auxiliaryRequest.Immutables.GetImmutablePropertyList().Get()) > len(classification.GetImmutables().GetImmutablePropertyList().Get()) {
			return nil, errorConstants.IncorrectFormat
		}

		for _, immutableProperty := range auxiliaryRequest.Immutables.GetImmutablePropertyList().Get() {
			if property := classification.GetImmutables().GetImmutablePropertyList().GetProperty(immutableProperty.GetID()); property == nil || immutableProperty.GetDataID().GetHashID().Compare(baseIDs.GenerateHashID()) != 0 && property.GetDataID().GetHashID().Compare(immutableProperty.GetDataID().GetHashID()) != 0 {
				return nil, errorConstants.IncorrectFormat
			}
		}
	}

	if auxiliaryRequest.Mutables != nil {
		if len(auxiliaryRequest.Mutables.GetMutablePropertyList().Get()) > len(classification.GetMutables().GetMutablePropertyList().Get()) {
			return nil, errorConstants.IncorrectFormat
		}

		for _, mutableProperty := range auxiliaryRequest.Mutables.GetMutablePropertyList().Get() {
			if property := classification.GetMutables().GetMutablePropertyList().GetProperty(mutableProperty.GetID()); property == nil {
				return nil, errorConstants.IncorrectFormat
			}
		}
	}

	return newAuxiliaryResponse(), nil
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
