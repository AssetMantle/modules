// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package send

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/assets/key"
	"github.com/AssetMantle/modules/x/assets/mappable"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/properties"
	propertyConstants "github.com/AssetMantle/schema/properties/constants"
	baseTypes "github.com/AssetMantle/schema/types/base"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameterManager      helpers.ParameterManager
	authenticateAuxiliary helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	transferAuxiliary     helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

// Handle is a method of the transactionKeeper struct. It processes the transaction message passed to it.
//
// Parameters:
// - context (context.Context): Used for process synchronization and carrying deadlines, among other things.
// - message (*Message): The transaction message being processed. It should contain details such as the source, destination, and the asset involved.
//
// Return values:
// - *TransactionResponse: A response object detailing the result of the transaction.
// - error: In case of an error, this will contain the error message.
//
// The Handle method performs the following steps:
//
// 1. It authenticates the transaction request using the authenticateAuxiliary getter from the transactionKeeper object. If this fails, it returns the error encountered.
//
// 2. It extracts the value from the message as an integer. If there's an error in this operation, it returns the error encountered.
//
// 3. It fetches the asset from the mapper attached to the transactionKeeper object using the asset ID from the message.
// - If the asset does not exist, it returns an EntityNotFound error.
//
// 4. It checks the asset's lock height property.
// - If the property exists and is not a meta property, it makes a call to supplementAuxiliary's Help method with the lock height property.
// - If the returned lock height property is a meta property, it updates the lock height.
// - If the property isn't a meta property, it returns a MetaDataError error.
//
// 5. If the lock height is greater than the current context height, it returns a NotAuthorized error.
//
// 6. It attempts to perform a transfer via the transferAuxiliary's Help method.
// - If an error occurs during the transfer, it returns the error.
//
// 7. If the process completes successfully, it returns a new transaction response object.
//
// Note: The errorConstants and propertyConstants are used for error handling and property checking respectively.
func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(message)); err != nil {
		return nil, err
	}

	value, err := message.GetValueAsInt()
	if err != nil {
		return nil, err
	}

	assets := transactionKeeper.mapper.NewCollection(context)

	Mappable := assets.Fetch(key.NewKey(message.AssetID)).GetMappable(key.NewKey(message.AssetID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("asset with ID %s not found", message.AssetID.AsString())
	}
	asset := mappable.GetAsset(Mappable)

	lockHeight := asset.GetLockHeight()

	if lockHeightProperty := asset.GetProperty(propertyConstants.LockHeightProperty.GetID()); lockHeightProperty != nil && !lockHeightProperty.IsMeta() {
		auxiliaryResponse, err := transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(lockHeightProperty))
		if err != nil {
			return nil, err
		}

		if lockHeightProperty = supplement.GetMetaPropertiesFromResponse(auxiliaryResponse).GetProperty(propertyConstants.LockHeightProperty.GetID()); lockHeightProperty != nil && lockHeightProperty.IsMeta() {
			lockHeight = lockHeightProperty.Get().(properties.MetaProperty).GetData().Get().(data.HeightData).Get()
		} else {
			return nil, errorConstants.MetaDataError.Wrapf("lock height property is not revealed")
		}
	}

	if lockHeight.Compare(baseTypes.CurrentHeight(context)) > 0 {
		return nil, errorConstants.NotAuthorized.Wrapf("transfer is not allowed until height %d", lockHeight.Get())
	}

	if _, err := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(message.GetFromIdentityID(), message.ToID, message.AssetID, value)); err != nil {
		return nil, err
	}

	return newTransactionResponse(), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameterManager = mapper, parameterManager

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case transfer.Auxiliary.GetName():
				transactionKeeper.transferAuxiliary = value
			}
		}
	}

	helpers.PanicOnUninitializedKeeperFields(transactionKeeper)
	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
