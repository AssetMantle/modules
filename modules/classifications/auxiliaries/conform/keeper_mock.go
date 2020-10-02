/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package conform

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type auxiliaryKeeperMock struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeperMock)(nil)

func (auxiliaryKeeper auxiliaryKeeperMock) Help(_ sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)
	if auxiliaryRequest.Mutables.Get().Get(base.NewID("conformError")) != nil {
		return newAuxiliaryResponse(errors.MockError)
	}
	return newAuxiliaryResponse(nil)
}

func initializeAuxiliaryKeeperMock(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.AuxiliaryKeeper {
	return auxiliaryKeeperMock{mapper: mapper}
}
