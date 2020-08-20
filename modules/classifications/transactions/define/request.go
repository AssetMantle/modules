/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package define

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionRequest struct {
	BaseReq             rest.BaseReq `json:"baseReq"`
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
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext context.CLIContext) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(flags.ImmutableMetaTraits),
		cliCommand.ReadString(flags.ImmutableTraits),
		cliCommand.ReadString(flags.MutableMetaTraits),
		cliCommand.ReadString(flags.MutableTraits),
	), nil
}
func (transactionRequest transactionRequest) FromJSON(rawMessage json.RawMessage) (helpers.TransactionRequest, error) {
	if Error := json.Unmarshal(rawMessage, &transactionRequest); Error != nil {
		return nil, Error
	}
	return transactionRequest, nil
}
func (transactionRequest transactionRequest) GetBaseReq() rest.BaseReq {
	return transactionRequest.BaseReq
}
func (transactionRequest transactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	from, Error := sdkTypes.AccAddressFromBech32(transactionRequest.GetBaseReq().From)
	if Error != nil {
		panic(errors.New(fmt.Sprintf("")))
	}
	return newMessage(
		from,
		base.ReadMetaProperties(transactionRequest.ImmutableMetaTraits),
		base.ReadProperties(transactionRequest.ImmutableTraits),
		base.ReadMetaProperties(transactionRequest.MutableMetaTraits),
		base.ReadProperties(transactionRequest.MutableTraits),
	), nil
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}
func newTransactionRequest(baseReq rest.BaseReq, immutableMetaTraits string, immutableTraits string, mutableMetaTraits string, mutableTraits string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:             baseReq,
		ImmutableMetaTraits: immutableMetaTraits,
		ImmutableTraits:     immutableTraits,
		MutableMetaTraits:   mutableMetaTraits,
		MutableTraits:       mutableTraits,
	}
}
