// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package verify

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/properties/constants"
)

type auxiliaryKeeper struct {
	mapper              helpers.Mapper
	parameters          helpers.Parameters
	supplementAuxiliary helpers.Auxiliary
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	identity := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(auxiliaryRequest.IdentityID)).Get(key.NewKey(auxiliaryRequest.IdentityID))
	if identity == nil {
		return newAuxiliaryResponse(errorConstants.EntityNotFound)
	}

	metaProperties, err := supplement.GetMetaPropertiesFromResponse(auxiliaryKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(identity.(mappables.Identity).GetAuthentication())))
	if err != nil {
		return newAuxiliaryResponse(err)
	}

	// TODO add test
	if _, found := metaProperties.GetMetaProperty(constants.AuthenticationProperty.GetID()).(lists.DataList).Search(base.NewAccAddressData(auxiliaryRequest.Address)); !found {
		return newAuxiliaryResponse(errorConstants.NotAuthorized)
	}

	return newAuxiliaryResponse(nil)
}

func (auxiliaryKeeper auxiliaryKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	auxiliaryKeeper.mapper, auxiliaryKeeper.parameters = mapper, parameters

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case supplement.Auxiliary.GetName():
				auxiliaryKeeper.supplementAuxiliary = value
			default:
				break
			}
		default:
			panic(errorConstants.UninitializedUsage)
		}
	}

	return auxiliaryKeeper
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
