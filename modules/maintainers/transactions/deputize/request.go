/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package deputize

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionRequest struct {
	BaseReq          rest.BaseReq `json:"baseReq"`
	FromID           string       `json:"fromID" valid:"required~required field fromID missing"`
	ToID             string       `json:"toID" valid:"required~required field toID missing"`
	ClassificationID string       `json:"classificationID" valid:"required~required field classificationID missing matches(^[A-Za-z]$)~invalid field classificationID"`
	MaintainedTraits string       `json:"maintainedTraits" valid:"required field maintainedTraits missing"`
	AddMaintainer    bool         `json:"addMaintainer" valid:"required field addMaintainer missing"`
	RemoveMaintainer bool         `json:"removeMaintainer" valid:"required field removeMaintainer missing"`
	MutateMaintainer bool         `json:"mutateMaintainer" valid:"required field mutateMaintainer missing"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) Validate() error {
	_, Error := govalidator.ValidateStruct(transactionRequest)
	return Error
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext context.CLIContext) helpers.TransactionRequest {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(constants.FromID),
		cliCommand.ReadString(constants.ToID),
		cliCommand.ReadString(constants.ClassificationID),
		cliCommand.ReadString(constants.MaintainedTraits),
		cliCommand.ReadBool(constants.AddMaintainer),
		cliCommand.ReadBool(constants.RemoveMaintainer),
		cliCommand.ReadBool(constants.MutateMaintainer),
	)
}
func (transactionRequest transactionRequest) GetBaseReq() rest.BaseReq {
	return transactionRequest.BaseReq
}
func (transactionRequest transactionRequest) MakeMsg() sdkTypes.Msg {
	from, Error := sdkTypes.AccAddressFromBech32(transactionRequest.GetBaseReq().From)
	if Error != nil {
		panic(errors.New(fmt.Sprintf("")))
	}

	return newMessage(
		from,
		base.NewID(transactionRequest.FromID),
		base.NewID(transactionRequest.ToID),
		base.NewID(transactionRequest.ClassificationID),
		base.ReadProperties(transactionRequest.MaintainedTraits),
		transactionRequest.AddMaintainer,
		transactionRequest.RemoveMaintainer,
		transactionRequest.MutateMaintainer,
	)
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, fromID string, toID string, classificationID string, maintainedTraits string, addMaintainer bool, removeMaintainer bool, mutateMaintainer bool) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:          baseReq,
		FromID:           fromID,
		ToID:             toID,
		ClassificationID: classificationID,
		MaintainedTraits: maintainedTraits,
		AddMaintainer:    addMaintainer,
		RemoveMaintainer: removeMaintainer,
		MutateMaintainer: mutateMaintainer,
	}
}
