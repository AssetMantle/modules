// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/constants/ids"
	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/modules/assets/internal/mappable"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/verify"
	maintainersVerify "github.com/AssetMantle/modules/modules/maintainers/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/mint"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/types/base"
)

type transactionKeeper struct {
	mapper                     helpers.Mapper
	parameters                 helpers.Parameters
	conformAuxiliary           helpers.Auxiliary
	mintAuxiliary              helpers.Auxiliary
	scrubAuxiliary             helpers.Auxiliary
	verifyAuxiliary            helpers.Auxiliary
	maintainersVerifyAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)

	if auxiliaryResponse := transactionKeeper.maintainersVerifyAuxiliary.GetKeeper().Help(context, maintainersVerify.NewAuxiliaryRequest(message.ClassificationID, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	if auxiliaryResponse := transactionKeeper.verifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	immutableMetaProperties, Error := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(message.ImmutableMetaProperties.GetList()...)))
	if Error != nil {
		return newTransactionResponse(Error)
	}

	immutableProperties := base.NewProperties(append(immutableMetaProperties.GetList(), message.ImmutableProperties.GetList()...)...)

	assetID := key.NewAssetID(message.ClassificationID, immutableProperties)

	assets := transactionKeeper.mapper.NewCollection(context).Fetch(key.FromID(assetID))
	if assets.Get(key.FromID(assetID)) != nil {
		return newTransactionResponse(errors.EntityAlreadyExists)
	}

	mutableMetaProperties, Error := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(message.MutableMetaProperties.GetList()...)))
	if Error != nil {
		return newTransactionResponse(Error)
	}

	mutableProperties := base.NewProperties(append(mutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...)

	if auxiliaryResponse := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(message.ClassificationID, immutableProperties, mutableProperties)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	split := sdkTypes.SmallestDec()

	if metaProperties := base.NewMetaProperties(append(message.ImmutableMetaProperties.GetList(), message.MutableMetaProperties.GetList()...)...); metaProperties.Get(ids.LockProperty) != nil {
		if split, Error = metaProperties.Get(ids.LockProperty).GetData().AsDec(); Error != nil {
			return newTransactionResponse(errors.MetaDataError)
		}
	}

	if auxiliaryResponse := transactionKeeper.mintAuxiliary.GetKeeper().Help(context, mint.NewAuxiliaryRequest(message.ToID, assetID, split)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	assets.Add(mappable.NewAsset(assetID, immutableProperties, mutableProperties))

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case conform.Auxiliary.GetName():
				transactionKeeper.conformAuxiliary = value
			case mint.Auxiliary.GetName():
				transactionKeeper.mintAuxiliary = value
			case scrub.Auxiliary.GetName():
				transactionKeeper.scrubAuxiliary = value
			case verify.Auxiliary.GetName():
				transactionKeeper.verifyAuxiliary = value
			case maintainersVerify.Auxiliary.GetName():
				transactionKeeper.maintainersVerifyAuxiliary = value
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
