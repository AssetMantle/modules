// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package name

import (
	"context"

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
	"github.com/AssetMantle/modules/x/identities/key"
	"github.com/AssetMantle/modules/x/identities/record"
)

type transactionKeeper struct {
	mapper helpers.Mapper
}

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

	identityID := baseIDs.NewIdentityID(base.PrototypeNubIdentity().GetClassificationID(), immutables)

	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(identityID))
	if identities.GetMappable(key.NewKey(identityID)) != nil {
		return nil, errorConstants.EntityAlreadyExists.Wrapf("identity with ID %s already exists", identityID.AsString())
	}

	identities.Add(record.NewRecord(base.NewIdentity(NubIdentityClassificationID, immutables, baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), baseData.NewListData(baseData.NewAccAddressData(address))))))))

	return newTransactionResponse(identityID), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
