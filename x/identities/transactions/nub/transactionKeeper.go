// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package nub

import (
	"context"

	"github.com/AssetMantle/modules/x/identities/key"
	"github.com/AssetMantle/modules/x/identities/mappable"

	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/define"
)

type transactionKeeper struct {
	mapper          helpers.Mapper
	defineAuxiliary helpers.Auxiliary
}

// TODO move to proper package
var NubClassificationID = baseIDs.NewClassificationID(baseQualified.NewImmutables(baseLists.NewPropertyList(constants.NubIDProperty)), baseQualified.NewMutables(baseLists.NewPropertyList(constants.AuthenticationProperty)))

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	address, err := types.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constants.NubIDProperty.GetKey(), baseData.NewIDData(message.NubID))))

	// TODO ***** add nub classificationID to genesis

	identityID := baseIDs.NewIdentityID(NubClassificationID, immutables)

	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(identityID))
	if identities.Get(key.NewKey(identityID)) != nil {
		return nil, errorConstants.EntityAlreadyExists.Wrapf("identity with ID %s already exists", identityID.AsString())
	}

	identities.Add(mappable.NewMappable(base.NewIdentity(NubClassificationID, immutables, baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), baseData.NewListData(baseData.NewAccAddressData(address))))))))

	return newTransactionResponse(identityID.AsString()), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case define.Auxiliary.GetName():
				transactionKeeper.defineAuxiliary = value
			}
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
