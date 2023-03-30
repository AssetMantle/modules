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

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {

	if _, err := transactionKeeper.maintainersVerifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.ClassificationID, message.FromID)); err != nil {
		return nil, err
	}

	fromAddress, err := types.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); err != nil {
		return nil, err
	}

	immutables := baseQualified.NewImmutables(message.ImmutableMetaProperties.Add(utilities.AnyPropertyListToPropertyList(message.ImmutableProperties.GetList()...)...))

	identityID := baseIDs.NewIdentityID(message.ClassificationID, immutables)

	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(identityID))
	if identities.Get(key.NewKey(identityID)) != nil {
		return nil, errorConstants.EntityAlreadyExists.Wrapf("identity with ID %s already exists", identityID.AsString())
	}

	toAddress, err := types.AccAddressFromBech32(message.To)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	authenticationProperty := baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), baseData.NewListData(baseData.NewAccAddressData(toAddress)))
	mutableMetaProperties := message.MutableMetaProperties.Add(authenticationProperty.ToAnyProperty().Get())

	mutables := baseQualified.NewMutables(mutableMetaProperties.Add(utilities.AnyPropertyListToPropertyList(message.MutableProperties.GetList()...)...))

	if _, err := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(message.ClassificationID, immutables, mutables)); err != nil {
		return nil, err
	}

	if _, err := transactionKeeper.bondAuxiliary.GetKeeper().Help(context, bond.NewAuxiliaryRequest(message.ClassificationID, fromAddress)); err != nil {
		return nil, err
	}

	identities.Add(mappable.NewMappable(base.NewIdentity(message.ClassificationID, immutables, mutables)))

	return newTransactionResponse(identityID.AsString()), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
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
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
