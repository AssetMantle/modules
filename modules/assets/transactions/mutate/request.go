/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mutate

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
	BaseReq               rest.BaseReq `json:"baseReq"`
	FromID                string       `json:"fromID" valid:"required~required field fromID missing"`
	AssetID               string       `json:"assetID" valid:"required~required field assetID missing matches(^[A-Za-z]$)~invalid field assetID"`
	MutableMetaProperties string       `json:"mutableMetaProperties" valid:"required~required field mutableMetaProperties missing matches(^[A-Za-z]$)~invalid field mutableMetaProperties"`
	MutableProperties     string       `json:"mutableProperties" valid:"required~required field mutableProperties missing matches(^[A-Za-z]$)~invalid field mutableProperties"`
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
		cliCommand.ReadString(constants.AssetID),
		cliCommand.ReadString(constants.MutableMetaProperties),
		cliCommand.ReadString(constants.MutableProperties),
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
		base.NewID(transactionRequest.AssetID),
		base.ReadMetaProperties(transactionRequest.MutableMetaProperties),
		base.ReadProperties(transactionRequest.MutableProperties),
	)
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, fromID string, assetID string, mutableMetaProperties string, mutableProperties string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:               baseReq,
		FromID:                fromID,
		AssetID:               assetID,
		MutableMetaProperties: mutableMetaProperties,
		MutableProperties:     mutableProperties,
	}
}
