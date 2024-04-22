// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package conform

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/key"
	"github.com/AssetMantle/modules/x/classifications/mappable"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, AuxiliaryRequest helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest, ok := AuxiliaryRequest.(auxiliaryRequest)
	if !ok {
		return nil, errorConstants.InvalidRequest.Wrapf("invalid request type %T", AuxiliaryRequest)
	}

	if err := auxiliaryRequest.Validate(); err != nil {
		return nil, err
	}

	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(auxiliaryRequest.ClassificationID))

	Mappable := classifications.GetMappable(key.NewKey(auxiliaryRequest.ClassificationID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("classification with ID %s not found", auxiliaryRequest.ClassificationID.AsString())
	}
	classification := mappable.GetClassification(Mappable)

	if !(auxiliaryRequest.Immutables == nil && classification.GetImmutables() == nil) {
		if (auxiliaryRequest.Immutables == nil && classification.GetImmutables() != nil) || (auxiliaryRequest.Immutables != nil && classification.GetImmutables() == nil) {
			return nil, errorConstants.IncorrectFormat.Wrapf("incorrect number of immutables")
		} else {
			if len(auxiliaryRequest.Immutables.GetImmutablePropertyList().Get()) != len(classification.GetImmutables().GetImmutablePropertyList().Get()) {
				return nil, errorConstants.IncorrectFormat.Wrapf("incorrect number of immutables")
			}

			for _, immutableProperty := range classification.GetImmutables().GetImmutablePropertyList().Get() {
				if property := auxiliaryRequest.Immutables.GetImmutablePropertyList().GetProperty(immutableProperty.GetID()); property == nil || immutableProperty.GetDataID().GetHashID().Compare(baseIDs.GenerateHashID()) != 0 && property.GetDataID().GetHashID().Compare(immutableProperty.GetDataID().GetHashID()) != 0 {
					return nil, errorConstants.IncorrectFormat.Wrapf("incorrect immutable %s", immutableProperty.GetID().AsString())
				}
			}
		}
	}

	if !(auxiliaryRequest.Mutables == nil && classification.GetMutables() == nil) {
		if (auxiliaryRequest.Mutables == nil && classification.GetMutables() != nil) || (auxiliaryRequest.Mutables != nil && classification.GetMutables() == nil) {
			return nil, errorConstants.IncorrectFormat.Wrapf("incorrect number of mutables")
		} else {
			if len(auxiliaryRequest.Mutables.GetMutablePropertyList().Get()) != len(classification.GetMutables().GetMutablePropertyList().Get()) {
				return nil, errorConstants.IncorrectFormat.Wrapf("incorrect number of mutables")
			}

			for _, mutableProperty := range classification.GetMutables().GetMutablePropertyList().Get() {
				if property := auxiliaryRequest.Mutables.GetMutablePropertyList().GetProperty(mutableProperty.GetID()); property == nil {
					return nil, errorConstants.IncorrectFormat.Wrapf("incorrect mutable %s", mutableProperty.GetID().AsString())
				}
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
