package mint

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
	ClassificationID string       `json:"classificationID" valid:"required~required field classificationID missing matches(^[A-Za-z]$)~invalid field classificationID"`
	MaintainersID    string       `json:"maintainersID" valid:"required~required field maintainersID missing matches(^[A-Za-z]$)~invalid field maintainersID"`
	Properties       string       `json:"properties" valid:"required~required field properties missing matches(^[A-Za-z]$)~invalid field properties"`
	Lock             int64        `json:"lock" valid:"required~required field lock missing matches(^[0-9]$)~invalid field lock "`
	Burn             int64        `json:"burn" valid:"required~required field burn missing matches(^[0-9]$)~invalid field burn "`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext context.CLIContext) helpers.TransactionRequest {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(constants.ClassificationID),
		cliCommand.ReadString(constants.MaintainersID),
		cliCommand.ReadString(constants.Properties),
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

	properties := strings.Split(transactionRequest.Properties, constants.PropertiesSeparator)
	if len(properties) > constants.MaxTraitCount {
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
		base.NewID(transactionRequest.MaintainersID),
		base.NewID(transactionRequest.ClassificationID),
		base.NewProperties(propertyList),
		base.NewHeight(transactionRequest.Lock),
		base.NewHeight(transactionRequest.Burn),
	)
}

func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, classificationID string, maintainersID string, properties string, lock int64, burn int64) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:          baseReq,
		ClassificationID: classificationID,
		MaintainersID:    maintainersID,
		Properties:       properties,
		Lock:             lock,
		Burn:             burn,
	}
}
