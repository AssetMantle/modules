// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/properties/constants"
	"github.com/AssetMantle/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/define"
	"github.com/AssetMantle/modules/x/identities/key"
	"github.com/AssetMantle/modules/x/identities/mappable"
	"github.com/AssetMantle/modules/x/identities/utilities"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/super"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
)

type transactionKeeper struct {
	mapper              helpers.Mapper
	defineAuxiliary     helpers.Auxiliary
	superAuxiliary      helpers.Auxiliary
	supplementAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}
func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	fromAddress, err := types.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	Mappable := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.FromID)).GetMappable(key.NewKey(message.FromID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("identity with ID %s not found", message.FromID.AsString())
	}
	identity := mappable.GetIdentity(Mappable)

	if !identity.IsProvisioned(fromAddress) {
		return nil, errorConstants.NotAuthorized.Wrapf("address %s is not provisioned by identity with ID %s", message.From, message.FromID.AsString())
	}

	immutables := base.NewImmutables(message.ImmutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(message.ImmutableProperties.Get()...)...))

	mutables := base.NewMutables(
		message.MutableMetaProperties.Add(
			baseLists.AnyPropertiesToProperties(
				message.MutableProperties.Add(
					constants.AuthenticationProperty,
				).Get()...,
			)...,
		),
	)

	auxiliaryResponse, err := transactionKeeper.defineAuxiliary.GetKeeper().Help(context, define.NewAuxiliaryRequest(fromAddress, immutables, mutables))
	if err != nil {
		return nil, err
	}
	classificationID := define.GetClassificationIDFromResponse(auxiliaryResponse)

	if _, err := transactionKeeper.superAuxiliary.GetKeeper().Help(context, super.NewAuxiliaryRequest(classificationID, message.FromID, base.NewMutables(mutables.GetMutablePropertyList().Remove(constants.AuthenticationProperty)), utilities.SetModulePermissions(true, true)...)); err != nil {
		return nil, err
	}

	return newTransactionResponse(classificationID), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case define.Auxiliary.GetName():
				transactionKeeper.defineAuxiliary = value
			case super.Auxiliary.GetName():
				transactionKeeper.superAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			}
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
