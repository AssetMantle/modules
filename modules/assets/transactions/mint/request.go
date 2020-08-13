/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mint

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

type transactionRequest struct {
	BaseReq          rest.BaseReq `json:"baseReq"`
	FromID           string       `json:"fromID" valid:"required~required field fromID missing"`
	ToID             string       `json:"toID" valid:"required~required field toID missing"`
	ClassificationID string       `json:"classificationID" valid:"required~required field classificationID missing matches(^[A-Za-z]$)~invalid field classificationID"`
	MaintainersID    string       `json:"maintainersID" valid:"required~required field maintainersID missing matches(^[A-Za-z]$)~invalid field maintainersID"`
	Properties       string       `json:"properties" valid:"required~required field properties missing matches(^[A-Za-z]$)~invalid field properties"`
	MetaProperties   string       `json:"metaProperties" valid:"required~required field metaProperties missing matches(^[A-Za-z]$)~invalid field metaProperties"`
	Lock             int64        `json:"lock" valid:"required~required field lock missing matches(^[0-9]$)~invalid field lock "`
	Burn             int64        `json:"burn" valid:"required~required field burn missing matches(^[0-9]$)~invalid field burn "`
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
		cliCommand.ReadString(constants.MaintainersID),
		cliCommand.ReadString(constants.Properties),
		cliCommand.ReadString(constants.MetaProperties),
		cliCommand.ReadInt64(constants.Lock),
		cliCommand.ReadInt64(constants.Burn),
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

	properties := base.NewProperties(append(request.ReadProperties(transactionRequest.Properties), request.ReadMetaProperties(transactionRequest.MetaProperties)...))
	if len(properties.GetList()) > constants.MaxTraitCount {
		panic(errors.New(fmt.Sprintf("")))
	}

	return newMessage(
		from,
		base.NewID(transactionRequest.FromID),
		base.NewID(transactionRequest.ToID),
		base.NewID(transactionRequest.MaintainersID),
		base.NewID(transactionRequest.ClassificationID),
		properties,
		base.NewHeight(transactionRequest.Lock),
		base.NewHeight(transactionRequest.Burn),
	)
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}
func newTransactionRequest(baseReq rest.BaseReq, fromID string, toID string, classificationID string, maintainersID string, properties string, metaProperties string, lock int64, burn int64) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:          baseReq,
		FromID:           fromID,
		ToID:             toID,
		ClassificationID: classificationID,
		MaintainersID:    maintainersID,
		Properties:       properties,
		MetaProperties:   metaProperties,
		Lock:             lock,
		Burn:             burn,
	}
}
