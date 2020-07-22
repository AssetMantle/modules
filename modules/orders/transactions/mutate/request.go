package mutate

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
	BaseReq    rest.BaseReq `json:"baseReq"`
	OrderID    string       `json:"classificationID"`
	Properties string       `json:"properties"`
	Lock       int64        `json:"lock"`
	Burn       int64        `json:"burn"`
}

var _ utility.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand utility.CLICommand, cliContext context.CLIContext) utility.TransactionRequest {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(constants.OrderID),
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

	var propertyList []schema.Property
	for _, property := range properties {
		traitIDAndProperty := strings.Split(property, constants.TraitIDAndPropertySeparator)
		if len(traitIDAndProperty) == 2 && traitIDAndProperty[0] != "" {
			propertyList = append(propertyList, schema.NewProperty(schema.NewID(traitIDAndProperty[0]), schema.NewFact(traitIDAndProperty[1], schema.NewSignatures(nil))))
		}
	}

	return newMessage(
		from,
		schema.NewID(transactionRequest.OrderID),
		schema.NewProperties(propertyList),
		schema.NewHeight(transactionRequest.Lock),
		schema.NewHeight(transactionRequest.Burn),
	)
}

func requestPrototype() utility.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, orderID string, properties string, lock int64, burn int64) utility.TransactionRequest {
	return transactionRequest{
		BaseReq:    baseReq,
		OrderID:    orderID,
		Properties: properties,
		Lock:       lock,
		Burn:       burn,
	}
}
