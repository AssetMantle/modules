// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/modules/classifications/internal/mappable"
	"github.com/AssetMantle/modules/modules/classifications/internal/module"
	"github.com/AssetMantle/modules/schema/documents/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/utilities/property"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	if len(auxiliaryRequest.Immutables.GetImmutablePropertyList().GetList())+len(auxiliaryRequest.Mutables.GetMutablePropertyList().GetList()) > module.MaxPropertyCount {
		return newAuxiliaryResponse(nil, errorConstants.InvalidRequest)
	}

	if property.Duplicate(append(auxiliaryRequest.Immutables.GetImmutablePropertyList().GetList(), auxiliaryRequest.Mutables.GetMutablePropertyList().GetList()...)) {
		return newAuxiliaryResponse(nil, errorConstants.InvalidRequest)
	}

	classificationID := baseIDs.NewClassificationID(auxiliaryRequest.Immutables, auxiliaryRequest.Mutables)

	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(classificationID))
	if classifications.Get(key.NewKey(classificationID)) != nil {
		return newAuxiliaryResponse(classificationID, errorConstants.EntityAlreadyExists)
	}

	classifications.Add(mappable.NewMappable(base.NewClassification(auxiliaryRequest.Immutables, auxiliaryRequest.Mutables)))

	return newAuxiliaryResponse(classificationID, nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
