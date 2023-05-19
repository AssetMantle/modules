// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn

import (
	"context"

	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/splits/key"
	"github.com/AssetMantle/modules/x/splits/mappable"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	splitID := baseIDs.NewSplitID(auxiliaryRequest.OwnerID, auxiliaryRequest.OwnableID)
	splits := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(splitID))

	Mappable := splits.GetMappable(key.NewKey(splitID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("split with ID %s not found", splitID.AsString())
	}
	split := mappable.GetSplit(Mappable)

	switch split = split.Send(auxiliaryRequest.Value); {
	case split.GetValue().LT(sdkTypes.ZeroInt()):
		return nil, errorConstants.InvalidRequest.Wrapf("split value cannot be negative")
	case split.GetValue().Equal(sdkTypes.ZeroInt()):
		splits.Remove(mappable.NewMappable(split))
	default:
		splits.Mutate(mappable.NewMappable(split))
	}

	return newAuxiliaryResponse(), nil
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
