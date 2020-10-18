/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package maintain

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)
	maintainerID := key.NewMaintainerID(auxiliaryRequest.ClassificationID, auxiliaryRequest.IdentityID)
	maintainers := mapper.NewMaintainers(auxiliaryKeeper.mapper, context).Fetch(maintainerID)
	maintainer := maintainers.Get(maintainerID)
	if maintainer == nil {
		return newAuxiliaryResponse(errors.EntityNotFound)
	}
	for _, maintainedProperty := range auxiliaryRequest.MaintainedMutables.Get().GetList() {
		if !maintainer.MaintainsTrait(maintainedProperty.GetID()) {
			return newAuxiliaryResponse(errors.NotAuthorized)
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
