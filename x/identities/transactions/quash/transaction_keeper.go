// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package quash

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"

	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/properties"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	baseTypes "github.com/AssetMantle/schema/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/unbond"
	"github.com/AssetMantle/modules/x/identities/constants"
	"github.com/AssetMantle/modules/x/identities/key"
	"github.com/AssetMantle/modules/x/identities/mappable"
	"github.com/AssetMantle/modules/x/identities/record"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/authorize"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
)

type transactionKeeper struct {
	mapper              helpers.Mapper
	parameterManager    helpers.ParameterManager
	authorizeAuxiliary  helpers.Auxiliary
	supplementAuxiliary helpers.Auxiliary
	unbondAuxiliary     helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	if !transactionKeeper.parameterManager.Fetch(context).GetParameter(constantProperties.QuashEnabledProperty.GetID()).GetMetaProperty().GetData().Get().(data.BooleanData).Get() {
		return nil, errorConstants.NotAuthorized.Wrapf("identity quashing is not enabled")
	}

	fromAddress := message.GetSigners()[0]

	if Mappable := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.FromID)).GetMappable(key.NewKey(message.FromID)); Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("identity with ID %s not found", message.FromID.AsString())
	} else if identity := mappable.GetIdentity(Mappable); !identity.IsProvisioned(fromAddress) {
		return nil, errorConstants.NotAuthorized.Wrapf("address %s is not provisioned for identity with ID %s", fromAddress.String(), message.FromID.AsString())
	}

	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.IdentityID))

	Mappable := identities.GetMappable(key.NewKey(message.IdentityID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("identity with ID %s not found", message.IdentityID.AsString())
	}
	identity := mappable.GetIdentity(Mappable)

	if _, err := transactionKeeper.authorizeAuxiliary.GetKeeper().Help(context, authorize.NewAuxiliaryRequest(identity.GetClassificationID(), message.FromID, constants.CanQuashIdentityPermission)); err != nil {
		return nil, err
	}

	// TODO: Check if identity expired identities must be quashed at the end of block
	if identity.GetExpiry().Compare(baseTypes.NewHeight(sdkTypes.UnwrapSDKContext(context).BlockHeight())) > 0 {
		return nil, errorConstants.NotAuthorized.Wrapf("identity with ID %s is not expired yet", message.IdentityID.AsString())
	}

	bondAmount := sdkTypes.ZeroInt()
	if bondAmountProperty := identity.GetProperty(constantProperties.BondAmountProperty.GetID()); bondAmountProperty == nil || !bondAmountProperty.IsMeta() {
		return nil, errorConstants.MetaDataError.Wrapf("identity with ID %s has no revealed bond amount", message.IdentityID)
	} else {
		bondAmount = bondAmountProperty.Get().(properties.MetaProperty).GetData().Get().(data.NumberData).Get()
	}

	if _, err := transactionKeeper.unbondAuxiliary.GetKeeper().Help(context, unbond.NewAuxiliaryRequest(identity.GetClassificationID(), fromAddress, bondAmount)); err != nil {
		return nil, err
	}

	identities.Remove(record.NewRecord(identity))

	return newTransactionResponse(), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameterManager = mapper, parameterManager

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case authorize.Auxiliary.GetName():
				transactionKeeper.authorizeAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case unbond.Auxiliary.GetName():
				transactionKeeper.unbondAuxiliary = value
			}
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
