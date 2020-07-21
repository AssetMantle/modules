package issue

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
	"strings"
)

type transactionRequest struct {
	BaseReq          rest.BaseReq `json:"baseReq"`
	To               string
	MaintainersID    string
	ClassificationID string
	Properties       string
}

var _ utility.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand utility.CLICommand, cliContext context.CLIContext) utility.TransactionRequest {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(constants.To),
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

	var propertyList []schema.Property
	for _, property := range properties {
		traitIDAndProperty := strings.Split(property, constants.TraitIDAndPropertySeparator)
		if len(traitIDAndProperty) == 2 && traitIDAndProperty[0] != "" {
			propertyList = append(propertyList, schema.NewProperty(schema.NewID(traitIDAndProperty[0]), schema.NewFact(traitIDAndProperty[1], schema.NewSignatures(nil))))
		}
	}

	return newMessage(
		from,
		to,
		schema.NewID(transactionRequest.MaintainersID),
		schema.NewID(transactionRequest.ClassificationID),
		schema.NewProperties(propertyList),
	)
}

func requestPrototype() utility.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, to string, maintainersID string, classificationID string, properties string) utility.TransactionRequest {
	return transactionRequest{
		BaseReq:          baseReq,
		To:               to,
		MaintainersID:    maintainersID,
		ClassificationID: classificationID,
		Properties:       properties,
	}
}
