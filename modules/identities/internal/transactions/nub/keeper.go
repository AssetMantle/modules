// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package nub

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/constants/ids"
	"github.com/AssetMantle/modules/constants/properties"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/define"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/mappable"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

type transactionKeeper struct {
	mapper          helpers.Mapper
	defineAuxiliary helpers.Auxiliary
	scrubAuxiliary  helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)

	nubIDProperty := baseTypes.NewMetaProperty(ids.NubIDProperty, baseData.NewIDData(message.NubID))

	immutableProperties, Error := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(nubIDProperty)))
	if Error != nil {
		return newTransactionResponse(Error)
	}

	authenticationProperty := baseTypes.NewMetaProperty(ids.AuthenticationProperty, baseData.NewListData(baseData.NewAccAddressData(message.From)))

	mutableProperties, Error := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(authenticationProperty)))
	if Error != nil {
		return newTransactionResponse(Error)
	}

	classificationID, Error := define.GetClassificationIDFromResponse(transactionKeeper.defineAuxiliary.GetKeeper().Help(context, define.NewAuxiliaryRequest(baseTypes.NewProperties(properties.NubID), baseTypes.NewProperties(properties.Authentication))))
	if classificationID == nil && Error != nil {
		return newTransactionResponse(Error)
	}

	identityID := key.NewIdentityID(classificationID, immutableProperties)

	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.FromID(identityID))
	if identities.Get(key.FromID(identityID)) != nil {
		return newTransactionResponse(errors.EntityAlreadyExists)
	}

	identities.Add(mappable.NewIdentity(identityID, immutableProperties, mutableProperties))

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case define.Auxiliary.GetName():
				transactionKeeper.defineAuxiliary = value
			case scrub.Auxiliary.GetName():
				transactionKeeper.scrubAuxiliary = value
			}
		default:
			panic(errors.UninitializedUsage)
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
