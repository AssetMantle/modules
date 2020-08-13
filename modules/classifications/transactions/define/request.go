/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package define

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
	"github.com/persistenceOne/persistenceSDK/utilities/request"
)

//TODO add mutable flag to traits
type transactionRequest struct {
	BaseReq             rest.BaseReq `json:"baseReq"`
	FromID              string       `json:"fromID" valid:"required~required field fromID missing"`
	ReadableID          string       `json:"readableId" valid:"required~required field fromID missing"`
	ImmutableMetaTraits string       `json:"immutableMetaTraits" valid:"required~required field immutableMetaTraits missing"`
	ImmutableTraits     string       `json:"immutableTraits" valid:"required~required field immutableTraits missing"`
	MutableMetaTraits   string       `json:"mutableMetaTraits" valid:"required~required field mutableMetaTraits missing"`
	MutableTraits       string       `json:"mutableTraits" valid:"required~required field mutableTraits missing"`
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
		cliCommand.ReadString(constants.ReadableID),
		cliCommand.ReadString(constants.ImmutableMetaTraits),
		cliCommand.ReadString(constants.ImmutableTraits),
		cliCommand.ReadString(constants.MutableMetaTraits),
		cliCommand.ReadString(constants.MutableTraits),
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
		request.ReadImmutableMetaTraits(transactionRequest.ImmutableMetaTraits),
		request.ReadImmutableTraits(transactionRequest.ImmutableTraits),
		request.ReadMutableMetaTraits(transactionRequest.MutableMetaTraits),
		request.ReadMutableTraits(transactionRequest.MutableTraits),
	)
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}
func newTransactionRequest(baseReq rest.BaseReq, fromID string, readableID string, immutableMetaTraits string, immutableTraits string, mutableMetaTraits string, mutableTraits string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:             baseReq,
		FromID:              fromID,
		ReadableID:          readableID,
		ImmutableMetaTraits: immutableMetaTraits,
		ImmutableTraits:     immutableTraits,
		MutableMetaTraits:   mutableMetaTraits,
		MutableTraits:       mutableTraits,
	}
}
