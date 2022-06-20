// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/modules/classifications/internal/mappable"
	"github.com/AssetMantle/modules/schema/helpers"
	constants2 "github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/utilities/property"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	if len(auxiliaryRequest.ImmutableProperties.GetList())+len(auxiliaryRequest.MutableProperties.GetList()) > constants.MaxPropertyCount {
		return newAuxiliaryResponse(nil, constants2.InvalidRequest)
	}

	if property.Duplicate(append(auxiliaryRequest.ImmutableProperties.GetList(), auxiliaryRequest.MutableProperties.GetList()...)) {
		return newAuxiliaryResponse(nil, constants2.InvalidRequest)
	}

	classificationID := key.NewClassificationID(baseIDs.NewID(context.ChainID()), auxiliaryRequest.ImmutableProperties, auxiliaryRequest.MutableProperties)

	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.FromID(classificationID))
	if classifications.Get(key.FromID(classificationID)) != nil {
		return newAuxiliaryResponse(baseIDs.NewID(classificationID.String()), constants2.EntityAlreadyExists)
	}

	classifications.Add(mappable.NewClassification(classificationID, auxiliaryRequest.ImmutableProperties, auxiliaryRequest.MutableProperties))

	return newAuxiliaryResponse(baseIDs.NewID(classificationID.String()), nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
