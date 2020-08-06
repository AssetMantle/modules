/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package issue

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"strings"
)

type transactionRequest struct {
	BaseReq          rest.BaseReq `json:"baseReq"`
	To               string       `json:"to" valid:"required~required field to missing matches(^commit[a-z0-9]{39}$)~invalid field to"`
	FromID           string       `json:"fromID" valid:"required~required field fromID missing matches(^[a-z]$)~invalid field fromID "`
	MaintainersID    string       `json:"maintainersID" valid:"required~required field maintainersID missing matches(^[A-Za-z]$)~invalid field maintainersID"`
	ClassificationID string       `json:"classificationID" valid:"required~required field classificationId missing,matches(^[A-Za-z]$)~invalid field classificationID"`
	Properties       string       `json:"properties" valid:"required~required field properties missing matches(^[A-Za-z]$)~invalid field properties"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext context.CLIContext) helpers.TransactionRequest {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(constants.To),
		cliCommand.ReadString(constants.FromID),
		cliCommand.ReadString(constants.MaintainersID),
		cliCommand.ReadString(constants.ClassificationID),
		cliCommand.ReadString(constants.Properties),
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

	to, Error := sdkTypes.AccAddressFromBech32(transactionRequest.To)
	if Error != nil {
		panic(errors.New(fmt.Sprintf("")))
	}

	properties := strings.Split(transactionRequest.Properties, constants.PropertiesSeparator)
	if len(properties) > constants.MaxTraitCount {
		//TODO handle
		panic(errors.New(fmt.Sprintf("")))
	}

	var propertyList []types.Property
	for _, property := range properties {
		traitIDAndProperty := strings.Split(property, constants.TraitIDAndPropertySeparator)
		if len(traitIDAndProperty) == 2 && traitIDAndProperty[0] != "" {
			propertyList = append(propertyList, base.NewProperty(base.NewID(traitIDAndProperty[0]), base.NewFact(traitIDAndProperty[1], base.NewSignatures(nil))))
		}
	}

	return newMessage(
		from,
		to,
		base.NewID(transactionRequest.FromID),
		base.NewID(transactionRequest.MaintainersID),
		base.NewID(transactionRequest.ClassificationID),
		base.NewProperties(propertyList),
	)
}

func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, to string, fromID string, maintainersID string, classificationID string, properties string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:          baseReq,
		To:               to,
		FromID:           fromID,
		MaintainersID:    maintainersID,
		ClassificationID: classificationID,
		Properties:       properties,
	}
}
