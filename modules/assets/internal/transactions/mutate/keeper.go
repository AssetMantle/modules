// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mutate

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/modules/assets/internal/mappable"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/maintain"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/schema/helpers"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/mappables"
)

type transactionKeeper struct {
	mapper            helpers.Mapper
	verifyAuxiliary   helpers.Auxiliary
	maintainAuxiliary helpers.Auxiliary
	scrubAuxiliary    helpers.Auxiliary
	conformAuxiliary  helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if auxiliaryResponse := transactionKeeper.verifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	assets := transactionKeeper.mapper.NewCollection(context).Fetch(key.FromID(message.AssetID))

	asset := assets.Get(key.FromID(message.AssetID))
	if asset == nil {
		return newTransactionResponse(errors.EntityNotFound)
	}

	mutableMetaProperties, err := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(message.MutableMetaProperties.GetList()...)))
	if err != nil {
		return newTransactionResponse(err)
	}

	mutableProperties := baseLists.NewPropertyList(append(mutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...)

	if auxiliaryResponse := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(asset.(mappables.Asset).GetClassificationID(), nil, mutableProperties)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	if auxiliaryResponse := transactionKeeper.maintainAuxiliary.GetKeeper().Help(context, maintain.NewAuxiliaryRequest(asset.(mappables.Asset).GetClassificationID(), message.FromID, mutableProperties)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	assets.Mutate(mappable.NewAsset(asset.(mappables.Asset).GetID(), asset.(mappables.Asset).GetImmutablePropertyList(), asset.(mappables.Asset).GetImmutablePropertyList().Mutate(mutableProperties.GetList()...)))

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
			case verify.Auxiliary.GetName():
				transactionKeeper.verifyAuxiliary = value
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
