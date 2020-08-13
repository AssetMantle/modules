/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package verify

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/identities/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)
	identities := mapper.NewIdentities(auxiliaryKeeper.mapper, context).Fetch(auxiliaryRequest.IdentityID)
	identity := identities.Get(auxiliaryRequest.IdentityID)
	if identity == nil {
		return newAuxiliaryResponse(constants.EntityNotFound)
	}
	if identity.IsProvisioned(auxiliaryRequest.Address) {
		return newAuxiliaryResponse(nil)
	} else {
		return newAuxiliaryResponse(constants.NotAuthorized)
	}
}

func initializeAuxiliaryKeeper(mapper helpers.Mapper, _ []interface{}) helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{mapper: mapper}
}
