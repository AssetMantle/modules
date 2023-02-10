// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unbond

import (
	"context"
	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/modules/classifications/internal/mappable"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/ids/constansts"
	"github.com/AssetMantle/modules/schema/properties"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
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

	for _, immutableProperty := range classification.GetImmutables().GetImmutablePropertyList().GetList() {
		if immutableProperty.Get().GetID().Compare(constansts.BondingPropertyID) == 0 {
			coins, err := sdkTypes.ParseCoinsNormalized(immutableProperty.Get().(properties.MetaProperty).GetData().Get().AsString() + dataConstants.Denom)
			if err != nil {
				return newAuxiliaryResponse(err)
			}
			if err := auxiliaryRequest.bankKeeper.SendCoinsFromModuleToAccount(sdkTypes.UnwrapSDKContext(context), auxiliaryRequest.moduleName, auxiliaryRequest.address, coins); err != nil {
				return newAuxiliaryResponse(err)
			}
			return newAuxiliaryResponse(nil)
		}
	}
	return newAuxiliaryResponse(nil)
}

func (auxiliaryKeeper auxiliaryKeeper) FetchCollection(context context.Context, classificationID baseIDs.ClassificationID) helpers.Collection {
	return auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(&classificationID))
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterList, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
