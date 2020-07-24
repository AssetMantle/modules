package send

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionRequest struct {
	BaseReq   rest.BaseReq `json:"baseReq"`
	ToID      string       `json:"toID" valid:"required~required field toID missing matches(^[A-Za-z]$)~invalid field toID"`
	OwnableID string       `json:"ownableID" valid:"required~required field ownableID missing matches(^[A-Za-z]$)~invalid field ownableID"`
	Split     string       `json:"split" valid:"required~required field split missing matches(^[A-Za-z]$)~invalid field split"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext context.CLIContext) helpers.TransactionRequest {
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
		base.NewID(transactionRequest.ToID),
		base.NewID(transactionRequest.OwnableID),
		split,
	)
}

func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, toID string, ownableID string, split string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:   baseReq,
		ToID:      toID,
		OwnableID: ownableID,
		Split:     split,
	}
}
