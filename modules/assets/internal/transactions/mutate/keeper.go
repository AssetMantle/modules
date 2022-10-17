// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mutate

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/modules/assets/internal/mappable"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/maintain"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/mappables/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	authenticateAuxiliary helpers.Auxiliary
	maintainAuxiliary     helpers.Auxiliary
	scrubAuxiliary        helpers.Auxiliary
	conformAuxiliary      helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if auxiliaryResponse := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	assets := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.AssetID))

	Mappable := assets.Get(key.NewKey(message.AssetID))
	if Mappable == nil {
		return newTransactionResponse(constants.EntityNotFound)
	}
	asset := Mappable.(mappables.Asset)

	mutableMetaProperties, err := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(message.MutableMetaProperties.GetList()...)))
	if err != nil {
		return newTransactionResponse(err)
	}

	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(append(mutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...))

	if auxiliaryResponse := transactionKeeper.maintainAuxiliary.GetKeeper().Help(context, maintain.NewAuxiliaryRequest(asset.GetClassificationID(), message.FromID, mutables)); !auxiliaryResponse.IsSuccessful() {
		if auxiliaryResponse := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(asset.GetClassificationID(), nil, asset.GetMutables().Mutate(mutables.GetMutablePropertyList().GetList()...))); !auxiliaryResponse.IsSuccessful() {
			return newTransactionResponse(auxiliaryResponse.GetError())
		}
	}

	assets.Mutate(mappable.NewMappable(base.NewAsset(asset.GetClassificationID(), asset.GetImmutables(), asset.GetMutables().Mutate(mutables.GetMutablePropertyList().GetList()...))))

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case conform.Auxiliary.GetName():
				transactionKeeper.conformAuxiliary = value
			case maintain.Auxiliary.GetName():
				transactionKeeper.maintainAuxiliary = value
			case scrub.Auxiliary.GetName():
				transactionKeeper.scrubAuxiliary = value
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			}
		default:
			panic(constants.UninitializedUsage)
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
