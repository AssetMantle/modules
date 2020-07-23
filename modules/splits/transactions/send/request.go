package send

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type transactionRequest struct {
	BaseReq   rest.BaseReq `json:"baseReq"`
	ToID      string       `json:"toID" valid:"required~Enter the ToID,matches(^[A-Za-z]$)~ToID is Invalid"`
	OwnableID string       `json:"ownableID" valid:"required~Enter the OwnableID,matches(^[A-Za-z]$)~OwnableID is Invalid"`
	Split     string       `json:"split" valid:"required~Enter the Split,matches(^[A-Za-z]$)~Split is Invalid"`
}

var _ utility.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand utility.CLICommand, cliContext context.CLIContext) utility.TransactionRequest {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(constants.ToID),
		cliCommand.ReadString(constants.OwnableID),
		cliCommand.ReadString(constants.Split),
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
	split, Error := sdkTypes.NewDecFromStr(transactionRequest.Split)
	if Error != nil {
		panic(errors.New(fmt.Sprintf("")))
	}

	return newMessage(
		from,
		schema.NewID(transactionRequest.ToID),
		schema.NewID(transactionRequest.OwnableID),
		split,
	)
}

func requestPrototype() utility.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, toID string, ownableID string, split string) utility.TransactionRequest {
	return transactionRequest{
		BaseReq:   baseReq,
		ToID:      toID,
		OwnableID: ownableID,
		Split:     split,
	}
}
