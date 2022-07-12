// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package nub

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/define"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/mappable"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
)

type transactionKeeper struct {
	mapper          helpers.Mapper
	defineAuxiliary helpers.Auxiliary
	scrubAuxiliary  helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)

	nubIDProperty := baseProperties.NewMetaProperty(constants.NubIDProperty.GetKey(), baseData.NewIDData(message.NubID))

	immutableProperties, err := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(nubIDProperty)))
	if err != nil {
		return newTransactionResponse(err)
	}

	authenticationProperty := baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), baseData.NewListData(baseLists.NewDataList(baseData.NewAccAddressData(message.From))))

	mutableProperties, err := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(authenticationProperty)))
	if err != nil {
		return newTransactionResponse(err)
	}

	classificationID, err := define.GetClassificationIDFromResponse(transactionKeeper.defineAuxiliary.GetKeeper().Help(context, define.NewAuxiliaryRequest(baseLists.NewPropertyList(constants.NubID), baseLists.NewPropertyList(constants.Authentication))))
	if classificationID == nil && err != nil {
		return newTransactionResponse(err)
	}

	identityID := key.NewIdentityID(classificationID, immutableProperties)

	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.FromID(identityID))
	if identities.Get(key.FromID(identityID)) != nil {
		return newTransactionResponse(errorConstants.EntityAlreadyExists)
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
			panic(errorConstants.UninitializedUsage)
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
