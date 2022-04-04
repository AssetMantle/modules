// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/modules/classifications/internal/mappable"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/types/base"
	"github.com/AssetMantle/modules/utilities/property"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	if len(auxiliaryRequest.ImmutableProperties.GetList())+len(auxiliaryRequest.MutableProperties.GetList()) > constants.MaxPropertyCount {
		return newAuxiliaryResponse(nil, errors.InvalidRequest)
	}

	if property.Duplicate(append(auxiliaryRequest.ImmutableProperties.GetList(), auxiliaryRequest.MutableProperties.GetList()...)) {
		return newAuxiliaryResponse(nil, errors.InvalidRequest)
	}

	classificationID := key.NewClassificationID(base.NewID(context.ChainID()), auxiliaryRequest.ImmutableProperties, auxiliaryRequest.MutableProperties)

	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.FromID(classificationID))
	if classifications.Get(key.FromID(classificationID)) != nil {
		return newAuxiliaryResponse(base.NewID(classificationID.String()), errors.EntityAlreadyExists)
	}

	classifications.Add(mappable.NewClassification(classificationID, auxiliaryRequest.ImmutableProperties, auxiliaryRequest.MutableProperties))

	return newAuxiliaryResponse(base.NewID(classificationID.String()), nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
