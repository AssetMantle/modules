/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mint

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/initialize"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/mint"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper                    helpers.Mapper
	splitsMintAuxiliary       helpers.Auxiliary
	identitiesVerifyAuxiliary helpers.Auxiliary
	metasInitializeAuxiliary  helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if auxiliaryResponse := transactionKeeper.identitiesVerifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}
	var propertyList []types.Property
	for _, property := range message.Properties.GetList() {
		if property.GetFact().IsMeta() {
			if auxiliaryResponse := transactionKeeper.metasInitializeAuxiliary.GetKeeper().Help(context, initialize.NewAuxiliaryRequest(property.GetFact().Get())); auxiliaryResponse != nil {
				return newTransactionResponse(auxiliaryResponse.GetError())
			}
			property = base.NewProperty(property.GetID(), base.MetaFactToFact(property.GetFact()))
		}
		propertyList = append(propertyList, property)
	}
	// TODO segregate immutables for mutables
	mutableProperties := base.NewProperties(propertyList)
	immutableProperties := base.NewProperties(propertyList)
	mutables := base.NewMutables(mutableProperties)
	immutables := base.NewImmutables(immutableProperties)
	assetID := mapper.NewAssetID(message.ClassificationID, immutables.GetHashID())
	assets := mapper.NewAssets(transactionKeeper.mapper, context).Fetch(assetID)
	if assets.Get(assetID) != nil {
		return newTransactionResponse(constants.EntityAlreadyExists)
	}
	if auxiliaryResponse := transactionKeeper.splitsMintAuxiliary.GetKeeper().Help(context, mint.NewAuxiliaryRequest(message.ToID, assetID, sdkTypes.OneDec())); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}
	assets.Add(mapper.NewAsset(assetID, message.Burn, message.Lock, immutables, mutables))
	return newTransactionResponse(nil)
}

func initializeTransactionKeeper(mapper helpers.Mapper, auxiliaries []interface{}) helpers.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case mint.Auxiliary.GetName():
				transactionKeeper.splitsMintAuxiliary = value
			case verify.Auxiliary.GetName():
				transactionKeeper.identitiesVerifyAuxiliary = value
			case initialize.Auxiliary.GetName():
				transactionKeeper.metasInitializeAuxiliary = value
			}
		}
	}
	return transactionKeeper
}
