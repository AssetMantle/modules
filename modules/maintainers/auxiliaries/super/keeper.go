// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package super

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/modules/maintainers/internal/mappable"
	"github.com/AssetMantle/modules/schema/helpers"
	baseTypes "github.com/AssetMantle/modules/schema/lists/base"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	maintainerID := key.NewMaintainerID(auxiliaryRequest.ClassificationID, auxiliaryRequest.IdentityID)

	maintainers := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.FromID(maintainerID))
	if maintainers.Get(key.FromID(maintainerID)) != nil {
		return newAuxiliaryResponse(errors.EntityAlreadyExists)
	}

	maintainers.Add(mappable.NewMaintainer(maintainerID, baseTypes.NewPropertyList(), auxiliaryRequest.MutableProperties))

	return newAuxiliaryResponse(nil)
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
