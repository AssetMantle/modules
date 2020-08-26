/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package make

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionRequest struct {
	BaseReq                 rest.BaseReq `json:"baseReq"`
	FromID                  string       `json:"fromID" valid:"required~required field fromID missing"`
	ClassificationID        string       `json:"classificationID" valid:"required~required field classificationID missing matches(^[A-Za-z]$)~invalid field classificationID"`
	MakerOwnableID          string       `json:"makerOwnableID" valid:"required~required field makerOwnableID missing"`
	TakerOwnableID          string       `json:"takerOwnableID" valid:"required~required field takerOwnableID missing"`
	ExpiresIn               int64        `json:"expiresIn" valid:"required~required field expiresIn missing"`
	MakerOwnableSplit       string       `json:"makerOwnableSplit" valid:"required~required field makerOwnableSplit missing"`
	ImmutableMetaProperties string       `json:"immutableMetaProperties" valid:"required~required field immutableMetaProperties missing matches(^[A-Za-z]$)~invalid field immutableMetaProperties"`
	ImmutableProperties     string       `json:"immutableProperties" valid:"required~required field immutableProperties missing matches(^[A-Za-z]$)~invalid field immutableProperties"`
	MutableMetaProperties   string       `json:"mutableMetaProperties" valid:"required~required field mutableMetaProperties missing matches(^[A-Za-z]$)~invalid field mutableMetaProperties"`
	MutableProperties       string       `json:"mutableProperties" valid:"required~required field mutableProperties missing matches(^[A-Za-z]$)~invalid field mutableProperties"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) Validate() error {
	_, Error := govalidator.ValidateStruct(transactionRequest)
	return Error
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext context.CLIContext) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(flags.FromID),
		cliCommand.ReadString(flags.ClassificationID),
		cliCommand.ReadString(flags.MakerOwnableID),
		cliCommand.ReadString(flags.TakerOwnableID),
		cliCommand.ReadInt64(flags.ExpiresIn),
		cliCommand.ReadString(flags.MakerOwnableSplit),
		cliCommand.ReadString(flags.ImmutableMetaProperties),
		cliCommand.ReadString(flags.ImmutableProperties),
		cliCommand.ReadString(flags.MutableMetaProperties),
		cliCommand.ReadString(flags.MutableProperties),
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
		return nil, Error
	}

	makerOwnableSplit, Error := sdkTypes.NewDecFromStr(transactionRequest.MakerOwnableSplit)
	if Error != nil {
		return nil, Error
	}

	immutableMetaProperties, Error := base.ReadMetaProperties(transactionRequest.ImmutableMetaProperties)
	if Error != nil {
		return nil, Error
	}
	immutableProperties, Error := base.ReadProperties(transactionRequest.ImmutableProperties)
	if Error != nil {
		return nil, Error
	}
	mutableMetaProperties, Error := base.ReadMetaProperties(transactionRequest.MutableMetaProperties)
	if Error != nil {
		return nil, Error
	}
	mutableProperties, Error := base.ReadProperties(transactionRequest.MutableProperties)
	if Error != nil {
		return nil, Error
	}

	return newMessage(
		from,
		base.NewID(transactionRequest.FromID),
		base.NewID(transactionRequest.ClassificationID),
		base.NewID(transactionRequest.MakerOwnableID),
		base.NewID(transactionRequest.TakerOwnableID),
		base.NewHeight(transactionRequest.ExpiresIn),
		makerOwnableSplit,
		immutableMetaProperties,
		immutableProperties,
		mutableMetaProperties,
		mutableProperties,
	), nil
}

func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, fromID string, classificationID string, makerOwnableID string, takerOwnableID string, expiresIn int64, makerOwnableSplit string, immutableMetaProperties string, immutableProperties string, mutableMetaProperties string, mutableProperties string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:                 baseReq,
		FromID:                  fromID,
		ClassificationID:        classificationID,
		MakerOwnableID:          makerOwnableID,
		TakerOwnableID:          takerOwnableID,
		ExpiresIn:               expiresIn,
		MakerOwnableSplit:       makerOwnableSplit,
		ImmutableMetaProperties: immutableMetaProperties,
		ImmutableProperties:     immutableProperties,
		MutableMetaProperties:   mutableMetaProperties,
		MutableProperties:       mutableProperties,
	}
}
