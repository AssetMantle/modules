// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mutate

import (
	"context"

	"github.com/AssetMantle/modules/x/identities/key"
	"github.com/AssetMantle/modules/x/identities/mappable"

	"github.com/AssetMantle/schema/go/documents/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/properties/utilities"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/member"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/maintain"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	authenticateAuxiliary helpers.Auxiliary
	maintainAuxiliary     helpers.Auxiliary
	memberAuxiliary       helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	fromAddress, err := sdkTypes.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); err != nil {
		return nil, err
	}

	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.IdentityID))

	Mappable := identities.Get(key.NewKey(message.IdentityID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("identity with ID %s not found", message.IdentityID.AsString())
	}
	identity := mappable.GetIdentity(Mappable)

	mutables := baseQualified.NewMutables(message.MutableMetaProperties.Add(utilities.AnyPropertyListToPropertyList(message.MutableProperties.GetList()...)...))

	if _, err := transactionKeeper.memberAuxiliary.GetKeeper().Help(context, member.NewAuxiliaryRequest(identity.GetClassificationID(), nil, mutables)); err != nil {
		return nil, err
	}

	if _, err := transactionKeeper.maintainAuxiliary.GetKeeper().Help(context, maintain.NewAuxiliaryRequest(identity.GetClassificationID(), message.FromID, mutables)); err != nil {
		return nil, err
	}

	identities.Mutate(mappable.NewMappable(base.NewIdentity(identity.GetClassificationID(), identity.GetImmutables(), baseQualified.NewMutables(identity.GetMutables().GetMutablePropertyList().Mutate(utilities.AnyPropertyListToPropertyList(mutables.GetMutablePropertyList().GetList()...)...)))))

	return newTransactionResponse(), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case member.Auxiliary.GetName():
				transactionKeeper.memberAuxiliary = value
			case maintain.Auxiliary.GetName():
				transactionKeeper.maintainAuxiliary = value
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			}
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
