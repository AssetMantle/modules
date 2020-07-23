package mint

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
	"strings"
)

type transactionRequest struct {
	BaseReq          rest.BaseReq `json:"baseReq"`
	ClassificationID string       `json:"classificationID" valid:"required~Enter the ClassificationID,matches(^[A-Za-z]$)~ClassificationID is Invalid"`
	MaintainersID    string       `json:"maintainersID" valid:"required~Enter the MaintainersID,matches(^[A-Za-z]$)~MaintainersID is Invalid"`
	Properties       string       `json:"properties" valid:"required~Enter the Properties,matches(^[A-Za-z]$)~Properties is Invalid"`
	Lock             int64        `json:"lock" valid:"required~Enter the Lock,matches(^[0-9]$)~Lock is Invalid"`
	Burn             int64        `json:"burn" valid:"required~Enter the Burn,matches(^[0-9]$)~Burn is Invalid"`
}

var _ utilities.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand utilities.CLICommand, cliContext context.CLIContext) utilities.TransactionRequest {
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

func requestPrototype() utilities.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, classificationID string, maintainersID string, properties string, lock int64, burn int64) utilities.TransactionRequest {
	return transactionRequest{
		BaseReq:          baseReq,
		ClassificationID: classificationID,
		MaintainersID:    maintainersID,
		Properties:       properties,
		Lock:             lock,
		Burn:             burn,
	}
}
