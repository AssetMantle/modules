// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/mappables"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	maintainerID := key.NewMaintainerID(auxiliaryRequest.ClassificationID, auxiliaryRequest.FromID)
	maintainers := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.FromID(maintainerID))

	maintainer := maintainers.Get(key.FromID(maintainerID))
	if maintainer == nil {
		return newAuxiliaryResponse(errors.EntityNotFound)
	}

	if !maintainer.(mappables.Maintainer).CanRemoveMaintainer() {
		return newAuxiliaryResponse(errors.NotAuthorized)
	}

	maintainers.Remove(maintainer)

	return newAuxiliaryResponse(nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
