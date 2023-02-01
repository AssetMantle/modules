// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package send

import (
	"context"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/splits/internal/utilities"
	"github.com/AssetMantle/modules/schema/helpers"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameters            helpers.ParameterList
	authenticateAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) helpers.TransactionResponse {
	_, err := transactionKeeper.Handle(context, message.(*Message))
	return newTransactionResponse(err)
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*Response, error) {

	fromAddress, err := sdkTypes.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	if auxiliaryResponse := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	splits := transactionKeeper.mapper.NewCollection(context)

	if _, err := utilities.SubtractSplits(splits, message.FromID, message.OwnableID, message.Value); err != nil {
		return nil, err
	}

	if _, err := utilities.AddSplits(splits, message.ToID, message.OwnableID, message.Value); err != nil {
		return nil, err
	}

	return &Response{}, nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.ParameterList, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			default:
				break
			}
		}
	}

	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
