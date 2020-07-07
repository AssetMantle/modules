package mint

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/modules/assets/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"strings"
)

type transactionRequest struct {
	BaseReq          rest.BaseReq `json:"baseReq"`
	ClassificationID string       `json:"classificationID"`
	MaintainersID    string       `json:"maintainersID"`
	Properties       string       `json:"properties"`
	Lock             int64        `json:"lock"`
	Burn             int64        `json:"burn"`
}

var _ types.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand types.CLICommand, cliContext context.CLIContext) types.TransactionRequest {
	return NewTransactionRequest(
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
			propertyList = append(propertyList, types.NewProperty(types.NewID(traitIDAndProperty[0]), types.NewFact(traitIDAndProperty[1], types.NewSignatures(nil))))
		}
	}

	return NewMessage(
		from,
		types.NewID(transactionRequest.GetBaseReq().ChainID),
		types.NewID(transactionRequest.MaintainersID),
		types.NewID(transactionRequest.ClassificationID),
		types.NewProperties(propertyList),
		types.NewHeight(transactionRequest.Lock),
		types.NewHeight(transactionRequest.Burn),
	)
}

func requestPrototype() types.TransactionRequest {
	return transactionRequest{}
}

func NewTransactionRequest(baseReq rest.BaseReq, classificationID string, maintainersID string, properties string, lock int64, burn int64) types.TransactionRequest {
	return transactionRequest{
		BaseReq:          baseReq,
		ClassificationID: classificationID,
		MaintainersID:    maintainersID,
		Properties:       properties,
		Lock:             lock,
		Burn:             burn,
	}
}
