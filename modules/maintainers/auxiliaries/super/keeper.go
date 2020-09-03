/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package super

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)
	maintainerID := mapper.NewMaintainerID(auxiliaryRequest.ClassificationID, auxiliaryRequest.IdentityID)
	maintainers := mapper.NewMaintainers(auxiliaryKeeper.mapper, context).Fetch(maintainerID)
	if maintainers.Get(maintainerID) != nil {
		return newAuxiliaryResponse(errors.EntityAlreadyExists)
	}
	maintainers = maintainers.Add(mapper.NewMaintainer(maintainerID, auxiliaryRequest.MutableTraits, true, true, true))
	return newAuxiliaryResponse(nil)
}

func initializeAuxiliaryKeeper(mapper helpers.Mapper, _ []interface{}) helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{mapper: mapper}
}
