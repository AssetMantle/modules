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

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	if Error := transactionKeeper.identitiesVerifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); Error != nil {
		return Error
	}
	var propertyList []types.Property
	for _, property := range message.Properties.GetList() {
		if property.GetFact().IsMeta() {
			if Error := transactionKeeper.metasInitializeAuxiliary.GetKeeper().Help(context, initialize.NewAuxiliaryRequest(property.GetFact().Get())); Error != nil {
				return Error
			}
			propertyList = append(propertyList, base.NewProperty(property.GetID(), base.MetaFactToFact(property.GetFact())))
		} else {
			propertyList = append(propertyList, property)
		}
	}
	// TODO segregate immutables for mutables
	mutableProperties := base.NewProperties(propertyList)
	immutableProperties := base.NewProperties(propertyList)
	mutables := base.NewMutables(mutableProperties, message.MaintainersID)
	immutables := base.NewImmutables(immutableProperties)
	assetID := mapper.NewAssetID(base.NewID(context.ChainID()), mutables.GetMaintainersID(), message.ClassificationID, immutables.GetHashID())
	assets := mapper.NewAssets(transactionKeeper.mapper, context).Fetch(assetID)
	if assets.Get(assetID) != nil {
		return constants.EntityAlreadyExists
	}
	if Error := transactionKeeper.splitsMintAuxiliary.GetKeeper().Help(context, mint.NewAuxiliaryRequest(message.ToID, assetID, sdkTypes.OneDec())); Error != nil {
		return Error
	}
	assets.Add(mapper.NewAsset(assetID, message.Burn, message.Lock, immutables, mutables))
	return nil
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
