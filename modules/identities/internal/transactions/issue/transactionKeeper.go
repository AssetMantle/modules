// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package issue

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/bond"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/mappable"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/verify"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/properties/utilities"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

type transactionKeeper struct {
	mapper                     helpers.Mapper
	authenticateAuxiliary      helpers.Auxiliary
	conformAuxiliary           helpers.Auxiliary
	bondAuxiliary              helpers.Auxiliary
	maintainersVerifyAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) helpers.TransactionResponse {
	_, err := transactionKeeper.Handle(context, message.(*Message))
	return newTransactionResponse(err)
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*Response, error) {

	if auxiliaryResponse := transactionKeeper.maintainersVerifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.ClassificationID, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	fromAddress, err := types.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	if auxiliaryResponse := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(utilities.AnyPropertyListToPropertyList(append(message.ImmutableMetaProperties.GetList(), message.ImmutableProperties.GetList()...)...)...))

	identityID := baseIDs.NewIdentityID(message.ClassificationID, immutables)

	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(identityID))
	if identities.Get(key.NewKey(identityID)) != nil {
		return nil, errorConstants.EntityAlreadyExists
	}

	toAddress, err := types.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	authenticationProperty := baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), baseData.NewListData(baseData.NewAccAddressData(toAddress)))
	mutableMetaProperties := baseLists.NewPropertyList(utilities.AnyPropertyListToPropertyList(append(message.MutableMetaProperties.GetList(), authenticationProperty.ToAnyProperty())...)...)

	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(utilities.AnyPropertyListToPropertyList(append(mutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...)...))

	if auxiliaryResponse := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(message.ClassificationID, immutables, mutables)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	if auxiliaryResponse := transactionKeeper.bondAuxiliary.GetKeeper().Help(context, bond.NewAuxiliaryRequest(message.ClassificationID, fromAddress)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	identities.Add(mappable.NewMappable(base.NewIdentity(message.ClassificationID, immutables, mutables)))

	return &Response{}, nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterList, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			case conform.Auxiliary.GetName():
				transactionKeeper.conformAuxiliary = value
			case bond.Auxiliary.GetName():
				transactionKeeper.bondAuxiliary = value
			case verify.Auxiliary.GetName():
				transactionKeeper.maintainersVerifyAuxiliary = value
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
