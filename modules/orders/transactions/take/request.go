package take

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
	BaseReq    rest.BaseReq `json:"baseReq"`
	FromID     string       `json:"fromID" valid:"required~required field fromID missing"`
	TakerSplit int64        `json:"takerSplit" valid:"required~required field takerSplit missing"`
	OrderID    string       `json:"orderID" valid:"required~required field orderID missing"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext context.CLIContext) helpers.TransactionRequest {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(constants.FromID),
		cliCommand.ReadInt64(constants.TakerSplit),
		cliCommand.ReadString(constants.OrderID),
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

	return newMessage(
		from,
		base.NewID(transactionRequest.FromID),
		sdkTypes.NewDec(transactionRequest.TakerSplit),
		base.NewID(transactionRequest.OrderID),
	)
}

func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, fromID string, takerSplit int64, orderID string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:    baseReq,
		FromID:     fromID,
		TakerSplit: takerSplit,
		OrderID:    orderID,
	}
}
