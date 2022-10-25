// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maintain

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/schema/errors/constants"
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
	maintainerID := baseIDs.NewMaintainerID(auxiliaryRequest.ClassificationID, auxiliaryRequest.IdentityID)
	maintainers := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(maintainerID))

	maintainer := maintainers.Get(key.NewKey(maintainerID))
	if maintainer == nil {
		return newAuxiliaryResponse(constants.EntityNotFound)
	}

	for _, maintainedProperty := range auxiliaryRequest.MaintainedMutables.GetMutablePropertyList().GetList() {
		if !maintainer.(types.Maintainer).MaintainsProperty(maintainedProperty.GetID()) {
			return newAuxiliaryResponse(constants.NotAuthorized)
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
